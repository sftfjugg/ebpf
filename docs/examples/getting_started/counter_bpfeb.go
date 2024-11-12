// Code generated by bpf2go; DO NOT EDIT.
//go:build mips || mips64 || ppc64 || s390x

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// loadCounter returns the embedded CollectionSpec for counter.
func loadCounter() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_CounterBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load counter: %w", err)
	}

	return spec, err
}

// loadCounterObjects loads counter and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*counterObjects
//	*counterPrograms
//	*counterMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadCounterObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadCounter()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// counterSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type counterSpecs struct {
	counterProgramSpecs
	counterMapSpecs
	counterVariableSpecs
}

// counterProgramSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type counterProgramSpecs struct {
	CountPackets *ebpf.ProgramSpec `ebpf:"count_packets"`
}

// counterMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type counterMapSpecs struct {
	PktCount *ebpf.MapSpec `ebpf:"pkt_count"`
}

// counterVariableSpecs contains variables before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type counterVariableSpecs struct {
}

// counterObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadCounterObjects or ebpf.CollectionSpec.LoadAndAssign.
type counterObjects struct {
	counterPrograms
	counterMaps
	counterVariables
}

func (o *counterObjects) Close() error {
	return _CounterClose(
		&o.counterPrograms,
		&o.counterMaps,
	)
}

// counterMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadCounterObjects or ebpf.CollectionSpec.LoadAndAssign.
type counterMaps struct {
	PktCount *ebpf.Map `ebpf:"pkt_count"`
}

func (m *counterMaps) Close() error {
	return _CounterClose(
		m.PktCount,
	)
}

// counterVariables contains all variables after they have been loaded into the kernel.
//
// It can be passed to loadCounterObjects or ebpf.CollectionSpec.LoadAndAssign.
type counterVariables struct {
}

// counterPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadCounterObjects or ebpf.CollectionSpec.LoadAndAssign.
type counterPrograms struct {
	CountPackets *ebpf.Program `ebpf:"count_packets"`
}

func (p *counterPrograms) Close() error {
	return _CounterClose(
		p.CountPackets,
	)
}

func _CounterClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed counter_bpfeb.o
var _CounterBytes []byte
