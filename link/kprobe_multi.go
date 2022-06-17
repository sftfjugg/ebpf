package link

import (
	"errors"
	"fmt"
	"runtime"
	"unsafe"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/asm"
	"github.com/cilium/ebpf/internal"
	"github.com/cilium/ebpf/internal/sys"
	"github.com/cilium/ebpf/internal/unix"
)

// KprobeMultiOptions defines additional parameters that will be used
// when opening a KprobeMulti Link.
//
// Symbols and Addresses are mutually exclusive.
type KprobeMultiOptions struct {
	// Symbols is an array of kernel symbols to attach the ebpf program to.
	Symbols []string
	// Addresses is an array of kernel symbols' addresses.
	Addresses []uint64
	// Cookies is an array of arbitrary values that can be fetched from an eBPF program
	// via `bpf_get_attach_cookie()`.
	//
	// If set, its length should be equal to the length of Symbols.
	//
	// Cookies will be assigned to Symbols based on their ordering.
	Cookies []uint64

	// Internal field. Only used for retprobes.
	flags uint32
}

// KprobeMulti attaches the given eBPF program to the entry point of a set of
// kernel symbols via the fprobe API, making it faster than attaching N Kprobes.
//
// The fprobe API limits the attach point to the function entry or return.
func KprobeMulti(prog *ebpf.Program, opts *KprobeMultiOptions) (Link, error) {
	return kprobeMulti(prog, opts)
}

// KretprobeMulti attaches the given eBPF program to the return point of a set of
// kernel symbols via the fprobe API, making it faster than attaching N Kprobes.
func KretprobeMulti(prog *ebpf.Program, opts *KprobeMultiOptions) (Link, error) {
	opts.flags = 1 << 0
	return kprobeMulti(prog, opts)
}

func kprobeMulti(prog *ebpf.Program, opts *KprobeMultiOptions) (Link, error) {
	if err := haveBPFLinkKprobeMulti(); err != nil {
		return nil, fmt.Errorf("kprobe.multi: %w", err)
	}

	if opts == nil {
		return nil, errors.New("kprobe.multi: missing options")
	}

	symbolsCnt := uint32(len(opts.Symbols))
	addressesCnt := uint32(len(opts.Addresses))

	if symbolsCnt > 0 && addressesCnt > 0 {
		return nil, errors.New("kprobe.multi: symbols and addresses are mutually exclusive")
	}
	if symbolsCnt == 0 && addressesCnt == 0 {
		return nil, errors.New("kprobe.multi: missing symbols/addresses array")
	}

	cookiesCnt := uint32(len(opts.Cookies))
	attr := &sys.LinkCreateKprobeMultiAttr{
		ProgFd:           uint32(prog.FD()),
		AttachType:       sys.BPF_TRACE_KPROBE_MULTI,
		KprobeMultiFlags: opts.flags,
	}

	syms := make([]uint64, 0)
	if symbolsCnt > addressesCnt {
		if cookiesCnt > 0 && cookiesCnt != symbolsCnt {
			return nil, errors.New("kprobe.multi: invalid cookies array length")
		}

		for _, s := range opts.Symbols {
			sptr, err := unsafeStringPtr(s)
			if err != nil {
				return nil, fmt.Errorf("kprobe.multi: %w", err)
			}
			syms = append(syms, uint64(uintptr(sptr)))
		}

		attr.Cnt = symbolsCnt
		attr.Syms = sys.NewPointer(unsafe.Pointer(&syms[0]))
	} else {
		if cookiesCnt > 0 && cookiesCnt != addressesCnt {
			return nil, errors.New("kprobe.multi: invalid cookies array length")
		}

		attr.Cnt = addressesCnt
		attr.Addrs = sys.NewPointer(unsafe.Pointer(&opts.Addresses[0]))
	}

	if cookiesCnt > 0 {
		attr.Cookies = sys.NewPointer(unsafe.Pointer(&opts.Cookies[0]))
	}

	fd, err := sys.LinkCreateKprobeMulti(attr)
	if err != nil {
		return nil, fmt.Errorf("kprobe.multi: link_create: %w", err)
	}

	runtime.KeepAlive(syms)
	runtime.KeepAlive(opts)

	return &kprobeMultiLink{RawLink{fd, ""}}, nil
}

type kprobeMultiLink struct {
	RawLink
}

var _ Link = (*kprobeMultiLink)(nil)

func (kml *kprobeMultiLink) Update(prog *ebpf.Program) error {
	return fmt.Errorf("kprobe.multi: link update: %w", ErrNotSupported)
}

func (kml *kprobeMultiLink) Pin(string) error {
	return fmt.Errorf("kprobe.multi: link pin: %w", ErrNotSupported)
}

func (kml *kprobeMultiLink) Unpin() error {
	return fmt.Errorf("kprobe.multi: link unpin: %w", ErrNotSupported)
}

// Probe BPF kprobe multi link.
//
// patchwork.kernel.org/project/netdevbpf/list/?series=623878&state=*
//
// TODO(matt): do we need this feature check?
var haveBPFLinkKprobeMulti = internal.FeatureTest("bpf_link_kprobe_multi", "5.18", func() error {
	prog, err := ebpf.NewProgram(&ebpf.ProgramSpec{
		Name: "probe_bpf_kprobe_multi_link",
		Type: ebpf.Kprobe,
		Instructions: asm.Instructions{
			asm.Mov.Imm(asm.R0, 0),
			asm.Return(),
		},
		AttachType: ebpf.AttachTraceKprobeMulti,
		License:    "MIT",
	})
	if err != nil {
		return err
	}
	defer prog.Close()

	sp, err := unsafeStringPtr("vprintk")
	if err != nil {
		return err
	}

	syms := []uint64{uint64(uintptr(sp))}
	_, err = sys.LinkCreateKprobeMulti(&sys.LinkCreateKprobeMultiAttr{
		ProgFd:     uint32(prog.FD()),
		AttachType: sys.BPF_TRACE_KPROBE_MULTI,
		Cnt:        1,
		Syms:       sys.NewPointer(unsafe.Pointer(&syms[0])),
	})
	if errors.Is(err, unix.EINVAL) {
		return internal.ErrNotSupported
	}
	runtime.KeepAlive(sp)

	return err
})
