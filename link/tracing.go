package link

import (
	"errors"
	"fmt"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/asm"
	"github.com/cilium/ebpf/btf"
	"github.com/cilium/ebpf/internal"
	"github.com/cilium/ebpf/internal/sys"
	"github.com/cilium/ebpf/internal/unix"
)

type tracing struct {
	RawLink
}

func (f *tracing) Update(new *ebpf.Program) error {
	return fmt.Errorf("tracing update: %w", ErrNotSupported)
}

// AttachFreplace attaches the given eBPF program to the function it replaces.
//
// The program and name can either be provided at link time, or can be provided
// at program load time. If they were provided at load time, they should be nil
// and empty respectively here, as they will be ignored by the kernel.
// Examples:
//
//	AttachFreplace(dispatcher, "function", replacement)
//	AttachFreplace(nil, "", replacement)
func AttachFreplace(targetProg *ebpf.Program, name string, prog *ebpf.Program) (Link, error) {
	if (name == "") != (targetProg == nil) {
		return nil, fmt.Errorf("must provide both or neither of name and targetProg: %w", errInvalidInput)
	}
	if prog == nil {
		return nil, fmt.Errorf("prog cannot be nil: %w", errInvalidInput)
	}
	if prog.Type() != ebpf.Extension {
		return nil, fmt.Errorf("eBPF program type %s is not an Extension: %w", prog.Type(), errInvalidInput)
	}

	var (
		target int
		typeID btf.TypeID
	)
	if targetProg != nil {
		btfHandle, err := targetProg.Handle()
		if err != nil {
			return nil, err
		}
		defer btfHandle.Close()

		spec, err := btfHandle.Spec(nil)
		if err != nil {
			return nil, err
		}

		var function *btf.Func
		if err := spec.TypeByName(name, &function); err != nil {
			return nil, err
		}

		target = targetProg.FD()
		typeID, err = spec.TypeID(function)
		if err != nil {
			return nil, err
		}
	}

	link, err := AttachRawLink(RawLinkOptions{
		Target:  target,
		Program: prog,
		Attach:  ebpf.AttachNone,
		BTF:     typeID,
	})
	if errors.Is(err, sys.ENOTSUPP) {
		// This may be returned by bpf_tracing_prog_attach via bpf_arch_text_poke.
		return nil, fmt.Errorf("create raw tracepoint: %w", ErrNotSupported)
	}
	if err != nil {
		return nil, err
	}

	return &tracing{*link}, nil
}

type TracingOptions struct {
	// Program must be of type Tracing with attach type
	// AttachTraceFEntry/AttachTraceFExit/AttachModifyReturn or
	// AttachTraceRawTp.
	Program *ebpf.Program
	// Arbitrary value that can be fetched from an eBPF program
	// via `bpf_get_attach_cookie()`.
	Cookie uint64
}

type LSMOptions struct {
	// Program must be of type LSM with attach type
	// AttachLSMMac.
	Program *ebpf.Program
	// Arbitrary value that can be fetched from an eBPF program
	// via `bpf_get_attach_cookie()`.
	Cookie uint64
}

// attachBTFID links all BPF program types (Tracing/LSM) that they attach to a btf_id.
func attachBTFID(program *ebpf.Program, cookie uint64) (Link, error) {
	if program.FD() < 0 {
		return nil, fmt.Errorf("invalid program %w", sys.ErrClosedFd)
	}

	var fd *sys.FD
	if err := haveBPFLinkTracing(); err == nil {
		// Use tracing bpf link
		fd, err = sys.LinkCreateTracing(&sys.LinkCreateTracingAttr{
			ProgFd:     uint32(program.FD()),
			AttachType: sys.AttachType(program.ExpectedAttachType()),
			Cookie:     cookie,
		})
		if err != nil {
			return nil, fmt.Errorf("create tracing link: %w", err)
		}
	} else {
		// Fallback to RawTracepointOpen
		if cookie > 0 {
			return nil, fmt.Errorf("attach BTF ID with cookie: %w", ErrNotSupported)
		}

		fd, err = sys.RawTracepointOpen(&sys.RawTracepointOpenAttr{
			ProgFd: uint32(program.FD()),
		})
		if errors.Is(err, sys.ENOTSUPP) {
			// This may be returned by bpf_tracing_prog_attach via bpf_arch_text_poke.
			return nil, fmt.Errorf("create raw tracepoint: %w", ErrNotSupported)
		}
		if err != nil {
			return nil, fmt.Errorf("create raw tracepoint: %w", err)
		}
	}

	raw := RawLink{fd: fd}
	info, err := raw.Info()
	if err != nil {
		raw.Close()
		return nil, err
	}

	if info.Type == RawTracepointType {
		// Sadness upon sadness: a Tracing program with AttachRawTp returns
		// a raw_tracepoint link. Other types return a tracing link.
		return &rawTracepoint{raw}, nil
	}
	return &tracing{raw}, nil
}

// AttachTracing links a tracing (fentry/fexit/fmod_ret) BPF program or
// a BTF-powered raw tracepoint (tp_btf) BPF Program to a BPF hook defined
// in kernel modules.
//
// TODO(matt): BPF_TRACE_ITER, BPF_LSM_CGROUP ?
func AttachTracing(opts TracingOptions) (Link, error) {
	if t := opts.Program.Type(); t != ebpf.Tracing {
		return nil, fmt.Errorf("invalid program type %s, expected Tracing", t)
	}

	switch opts.Program.ExpectedAttachType() {
	case ebpf.AttachTraceFEntry:
	case ebpf.AttachTraceFExit:
	case ebpf.AttachTraceRawTp:
	case ebpf.AttachModifyReturn:
	default:
		return nil, fmt.Errorf("invalid program expected attach type: %s", opts.Program.ExpectedAttachType())
	}

	return attachBTFID(opts.Program, opts.Cookie)
}

// AttachLSM links a Linux security module (LSM) BPF Program to a BPF
// hook defined in kernel modules.
func AttachLSM(opts LSMOptions) (Link, error) {
	if t := opts.Program.Type(); t != ebpf.LSM {
		return nil, fmt.Errorf("invalid program type %s, expected LSM", t)
	}

	if opts.Program.ExpectedAttachType() != ebpf.AttachLSMMac {
		return nil, fmt.Errorf("invalid program expected attach type: %s", opts.Program.ExpectedAttachType())
	}

	return attachBTFID(opts.Program, opts.Cookie)
}

// https://github.com/torvalds/linux/commit/2fcc82411e74e5e6aba336561cf56fb899bfae4e
var haveBPFLinkTracing = internal.FeatureTest("bpf_link_tracing", "5.19", func() error {
	prog, err := ebpf.NewProgram(&ebpf.ProgramSpec{
		Name:         "probe_tracing_link",
		Type:         ebpf.Tracing,
		AttachType:   ebpf.AttachTraceFEntry,
		Instructions: asm.Instructions{asm.Mov.Imm(asm.R0, 0), asm.Return()},
		AttachTo:     "vprintk",
		License:      "MIT",
	})
	if errors.Is(err, unix.E2BIG) || errors.Is(err, unix.EINVAL) {
		return internal.ErrNotSupported
	}
	if err != nil {
		return err
	}
	defer prog.Close()

	fd, err := sys.LinkCreateTracing(&sys.LinkCreateTracingAttr{
		ProgFd:     uint32(prog.FD()),
		AttachType: sys.BPF_TRACE_FENTRY,
		Cookie:     1,
	})
	if errors.Is(err, unix.EINVAL) {
		return internal.ErrNotSupported
	}
	if err != nil {
		return err
	}
	fd.Close()

	return nil
})
