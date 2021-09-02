// Code generated by bpf2go; DO NOT EDIT.
// +build 386 amd64 amd64p32 arm arm64 mips64le mips64p32le mipsle ppc64le riscv64

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// LoadHandler returns the embedded CollectionSpec for Handler.
func LoadHandler() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_HandlerBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load Handler: %w", err)
	}

	return spec, err
}

// LoadHandlerObjects loads Handler and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//     *HandlerObjects
//     *HandlerPrograms
//     *HandlerMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadHandlerObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadHandler()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// HandlerSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type HandlerSpecs struct {
	HandlerProgramSpecs
	HandlerMapSpecs
}

// HandlerSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type HandlerProgramSpecs struct {
	MmPageAlloc *ebpf.ProgramSpec `ebpf:"mm_page_alloc"`
}

// HandlerMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type HandlerMapSpecs struct {
	CountingMap *ebpf.MapSpec `ebpf:"counting_map"`
}

// HandlerObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadHandlerObjects or ebpf.CollectionSpec.LoadAndAssign.
type HandlerObjects struct {
	HandlerPrograms
	HandlerMaps
}

func (o *HandlerObjects) Close() error {
	return _HandlerClose(
		&o.HandlerPrograms,
		&o.HandlerMaps,
	)
}

// HandlerMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadHandlerObjects or ebpf.CollectionSpec.LoadAndAssign.
type HandlerMaps struct {
	CountingMap *ebpf.Map `ebpf:"counting_map"`
}

func (m *HandlerMaps) Close() error {
	return _HandlerClose(
		m.CountingMap,
	)
}

// HandlerPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadHandlerObjects or ebpf.CollectionSpec.LoadAndAssign.
type HandlerPrograms struct {
	MmPageAlloc *ebpf.Program `ebpf:"mm_page_alloc"`
}

func (p *HandlerPrograms) Close() error {
	return _HandlerClose(
		p.MmPageAlloc,
	)
}

func _HandlerClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//go:embed handler_bpfel.o
var _HandlerBytes []byte
