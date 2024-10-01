package ebpf

import (
	"io"
	"os"
	"runtime"
	"testing"
	"unsafe"

	"github.com/go-quicktest/qt"

	"github.com/cilium/ebpf/internal/sys"
	"github.com/cilium/ebpf/internal/testutils"
)

func mustMmapableArray(tb testing.TB, extraFlags uint32) *Map {
	tb.Helper()

	m, err := NewMap(&MapSpec{
		Name:       "ebpf_mmap",
		Type:       Array,
		KeySize:    4,
		ValueSize:  8,
		MaxEntries: 8,
		Flags:      sys.BPF_F_MMAPABLE | extraFlags,
	})
	testutils.SkipIfNotSupported(tb, err)
	qt.Assert(tb, qt.IsNil(err))
	tb.Cleanup(func() {
		m.Close()
	})
	return m
}

func TestMemory(t *testing.T) {
	mm, err := mustMmapableArray(t, 0).Memory()
	qt.Assert(t, qt.IsNil(err))

	// The mapping is always at least one page long, and the Map created here fits
	// in a single page.
	qt.Assert(t, qt.Equals(mm.Size(), os.Getpagesize()))

	// No BPF_F_RDONLY_PROG flag, so the Memory should be read-write.
	qt.Assert(t, qt.IsFalse(mm.ReadOnly()))

	want := []byte{1, 2, 3, 4, 4, 3, 2, 1}
	w := io.NewOffsetWriter(mm, 16)
	n, err := w.Write(want)
	qt.Assert(t, qt.IsNil(err))
	qt.Assert(t, qt.Equals(n, 8))

	r := io.NewSectionReader(mm, 16, int64(len(want)))
	got := make([]byte, len(want))
	n, err = r.Read(got)
	qt.Assert(t, qt.IsNil(err))
	qt.Assert(t, qt.Equals(n, len(want)))
}

func TestMemoryBounds(t *testing.T) {
	mm, err := mustMmapableArray(t, 0).Memory()
	qt.Assert(t, qt.IsNil(err))

	size := uint64(mm.Size())
	end := size - 1

	qt.Assert(t, qt.IsTrue(mm.bounds(0, 0)))
	qt.Assert(t, qt.IsTrue(mm.bounds(end, 0)))
	qt.Assert(t, qt.IsTrue(mm.bounds(end-8, 8)))
	qt.Assert(t, qt.IsTrue(mm.bounds(0, end)))

	qt.Assert(t, qt.IsFalse(mm.bounds(end-8, 9)))
	qt.Assert(t, qt.IsFalse(mm.bounds(end, 1)))
	qt.Assert(t, qt.IsFalse(mm.bounds(0, size)))
}

func TestMemoryReadOnly(t *testing.T) {
	rd, err := mustMmapableArray(t, sys.BPF_F_RDONLY_PROG).Memory()
	qt.Assert(t, qt.IsNil(err))

	// BPF_F_RDONLY_PROG flag, so the Memory should be read-only.
	qt.Assert(t, qt.IsTrue(rd.ReadOnly()))

	// Frozen maps can't be mapped rw either.
	frozen := mustMmapableArray(t, 0)
	qt.Assert(t, qt.IsNil(frozen.Freeze()))
	fz, err := frozen.Memory()
	qt.Assert(t, qt.IsNil(err))
	qt.Assert(t, qt.IsTrue(fz.ReadOnly()))
}

func TestMemoryUnmap(t *testing.T) {
	mm, err := mustMmapableArray(t, 0).Memory()
	qt.Assert(t, qt.IsNil(err))

	// Avoid unmap running twice.
	runtime.SetFinalizer(unsafe.SliceData(mm.b), nil)

	// unmap panics if the operation fails.
	unmap(mm.Size())(unsafe.SliceData(mm.b))
}

func TestMemoryPointer(t *testing.T) {
	mm, err := mustMmapableArray(t, 0).Memory()
	qt.Assert(t, qt.IsNil(err))

	// Requesting an unaligned value should fail.
	_, err = MemoryPointer[uint32](mm, 7)
	qt.Assert(t, qt.IsNotNil(err))

	u32, err := MemoryPointer[uint32](mm, 12)
	qt.Assert(t, qt.IsNil(err))

	*u32 = 0xf00d
	qt.Assert(t, qt.Equals(*u32, 0xf00d))

	_, err = MemoryPointer[*uint32](mm, 0)
	qt.Assert(t, qt.ErrorIs(err, ErrInvalidType))
}

func testReadOnly(tb testing.TB, mm *Memory) {
	tb.Helper()

	_, err := mm.WriteAt([]byte{1}, 0)
	qt.Assert(tb, qt.ErrorIs(err, ErrReadOnly))

	_, err = MemoryPointer[uint32](mm, 0)
	qt.Assert(tb, qt.ErrorIs(err, ErrReadOnly))
}

func TestMemoryReadonly(t *testing.T) {
	m := mustMmapableArray(t, sys.BPF_F_RDONLY_PROG)

	mm, err := m.Memory()
	qt.Assert(t, qt.IsNil(err))

	testReadOnly(t, mm)
}

func TestMemoryFrozen(t *testing.T) {
	m := mustMmapableArray(t, 0)
	qt.Assert(t, qt.IsNil(m.Freeze()))

	mm, err := m.Memory()
	qt.Assert(t, qt.IsNil(err))

	testReadOnly(t, mm)
}
