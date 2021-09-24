package testutils

import (
	"fmt"
	"os"

	"github.com/cilium/ebpf/internal"
)

func init() {
	// Increase the memlock for all tests unconditionally. It's a great source of
	// weird bugs, since different distros have different default limits.
	if err := internal.RemoveMemlockRlimit(); err != nil {
		fmt.Fprintln(os.Stderr, "WARNING: Failed to adjust rlimit, tests may fail")
	}
}
