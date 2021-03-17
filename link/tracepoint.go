package link

import (
	"errors"

	"github.com/cilium/ebpf"
)

// Tracepoint returns a trace event for the given tracepoint group and name.
// See /sys/kernel/debug/tracing/events to find available tracepoints. The
// top-level directory is group, the event's subdirectory is name. Example:
//
//	Tracepoint("syscalls", "sys_enter_fork")
//
// Multiple calls with the same group and name are unnecessary, as they can
// be re-used to attach an arbitrary amount of eBPF programs.
func Tracepoint(group, name string) (*TraceEvent, error) {

	if group == "" || name == "" {
		return nil, errors.New("Tracepoint group and name cannot be empty")
	}

	tid, err := getTraceEventID(group, name)
	if err != nil {
		return nil, err
	}

	return &TraceEvent{
		tracefsID: tid,
		group:     group,
		name:      name,
		progType:  ebpf.TracePoint,
		finalizer: nil,
	}, nil
}
