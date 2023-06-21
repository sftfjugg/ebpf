package sysenc

import (
	"unsafe"

	"github.com/cilium/ebpf/internal/sys"
)

type Buffer struct {
	ptr unsafe.Pointer
	// Size of the buffer. 0 if created from UnsafeBuffer or when using
	// zero-copy unmarshaling.
	size int
}

func newBuffer(buf []byte) Buffer {
	if len(buf) == 0 {
		return Buffer{}
	}
	return Buffer{unsafe.Pointer(&buf[0]), len(buf)}
}

// UnsafeBuffer constructs a Buffer for zero-copy unmarshaling.
//
// [Pointer] is the only valid method to call on such a Buffer.
// Use [SyscallBuffer] instead if possible.
func UnsafeBuffer(ptr unsafe.Pointer) Buffer {
	return Buffer{ptr, 0}
}

// SyscallOutput prepares a Buffer for a syscall to write into.
//
// The buffer may point at the underlying memory of dst, in which case [Unmarshal]
// becomes a no-op.
//
// The contents of the buffer are undefined and may be non-zero.
func SyscallOutput(dst any, size int) Buffer {
	if dstBuf := unsafeBackingMemory(dst); len(dstBuf) == size {
		buf := newBuffer(dstBuf)
		buf.size = 0
		return buf
	}

	return newBuffer(make([]byte, size))
}

// Copy the contents into dst.
//
// Returns the number of copied bytes.
func (b Buffer) Copy(dst []byte) int {
	return copy(dst, b.unsafeBytes())
}

// Pointer returns the location where a syscall should write.
func (b Buffer) Pointer() sys.Pointer {
	// NB: This deliberately ignores b.layout.valid() to support zero-copy
	// marshaling / unmarshaling using unsafe.Pointer.
	return sys.NewPointer(b.ptr)
}

// Unmarshal the buffer into the provided value.
//
// This is a no-op on a zero buffer.
func (b Buffer) Unmarshal(data any) error {
	if b.size == 0 {
		return nil
	}

	return Unmarshal(data, b.unsafeBytes())
}

func (b Buffer) unsafeBytes() []byte {
	return unsafe.Slice((*byte)(b.ptr), b.size)
}