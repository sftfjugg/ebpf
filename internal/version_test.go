package internal

import (
	"testing"
)

func TestVersion(t *testing.T) {
	a, err := NewVersion("1.2")
	if err != nil {
		t.Fatal(err)
	}

	b, err := NewVersion("2.2.1")
	if err != nil {
		t.Fatal(err)
	}

	if !a.Less(b) {
		t.Error("A should be less than B")
	}

	if b.Less(a) {
		t.Error("B shouldn't be less than A")
	}

	v200 := Version{2, 0, 0}
	if !a.Less(v200) {
		t.Error("1.2.1 should not be less than 2.0.0")
	}

	if v200.Less(a) {
		t.Error("2.0.0 should not be less than 1.2.1")
	}

	// 256.256.256 overflows all three uint8s of the kernel version and should
	// result in a zero version.
	if v := (Version{256, 256, 256}); v.Kernel() != 0 {
		t.Error("256.256.256 should result in a zero kernel version")
	}
}
