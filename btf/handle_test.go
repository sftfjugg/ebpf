package btf_test

import (
	"testing"
	"time"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/asm"
	"github.com/cilium/ebpf/btf"
	"github.com/cilium/ebpf/internal/testutils"
)

func TestNewHandleFromID(t *testing.T) {
	const vmlinux = btf.ID(1)

	// See https://github.com/torvalds/linux/commit/5329722057d41aebc31e391907a501feaa42f7d9
	testutils.SkipOnOldKernel(t, "5.11", "vmlinux BTF ID")

	// We need to call into the verifier at least once to ensure that
	// vmlinux BTF has been loaded.
	prog, err := ebpf.NewProgram(&ebpf.ProgramSpec{
		Type:    ebpf.SocketFilter,
		License: "MIT",
		Instructions: asm.Instructions{
			asm.Mov.Imm(asm.R0, 0),
			asm.Return(),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer prog.Close()

	time.Sleep(100 * time.Millisecond)

	h, err := btf.NewHandleFromID(vmlinux)
	if err != nil {
		t.Fatal(err)
	}
	h.Close()
}
