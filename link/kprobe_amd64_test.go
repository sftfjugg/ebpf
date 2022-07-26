//go:build amd64
// +build amd64

package link

import (
	"errors"
	"os"
	"testing"

	"github.com/cilium/ebpf"
)

func TestKprobeOffset(t *testing.T) {
	prog := mustLoadProgram(t, ebpf.Kprobe, 0, "")

	// This test only works on amd64 because arches like arm64 don't allow
	// probes on unaligned addresses and return EINVAL instead of EILSEQ,
	// see arch_prepare_kprobe().
	for i := uint64(2); i < 10; i++ {
		k, err := Kprobe("inet6_release", prog, &KprobeOptions{Offset: i})
		if err != nil {
			continue
		}
		k.Close()

		_, err = Kprobe("inet6_release", prog, &KprobeOptions{Offset: i - 1})
		if !errors.Is(err, os.ErrNotExist) {
			t.Fatalf("expected 'os.ErrNotExist', got: '%v'", err)
		}
		return
	}

	t.Fatal("no valid offsets found")
}
