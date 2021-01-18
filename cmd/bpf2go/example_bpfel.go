// Code generated by bpf2go; DO NOT EDIT.
// +build 386 amd64 amd64p32 arm arm64 mipsle mips64le mips64p32le ppc64le riscv64

package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type exampleSpecs struct {
	ProgramFilter *ebpf.ProgramSpec `ebpf:"filter"`
}

func newExampleSpecs() (*exampleSpecs, error) {
	reader := bytes.NewReader(_ExampleBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load example: %w", err)
	}

	specs := new(exampleSpecs)
	if err := spec.LoadAndAssign(specs, nil); err != nil {
		return nil, fmt.Errorf("can't assign example: %w", err)
	}

	return specs, nil
}

func (s *exampleSpecs) CollectionSpec() *ebpf.CollectionSpec {
	return &ebpf.CollectionSpec{
		Programs: map[string]*ebpf.ProgramSpec{
			"filter": s.ProgramFilter,
		},
		Maps: map[string]*ebpf.MapSpec{},
	}
}

func (s *exampleSpecs) Load(opts *ebpf.CollectionOptions) (*exampleObjects, error) {
	var objs exampleObjects
	if err := s.CollectionSpec().LoadAndAssign(&objs, opts); err != nil {
		return nil, err
	}
	return &objs, nil
}

func (s *exampleSpecs) Copy() *exampleSpecs {
	return &exampleSpecs{
		ProgramFilter: s.ProgramFilter.Copy(),
	}
}

type exampleObjects struct {
	ProgramFilter *ebpf.Program `ebpf:"filter"`
}

func (o *exampleObjects) Close() error {
	for _, closer := range []io.Closer{
		o.ProgramFilter,
	} {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
var _ExampleBytes = []byte("\x7f\x45\x4c\x46\x02\x01\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\xf7\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xb0\x06\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x40\x00\x00\x00\x00\x00\x40\x00\x14\x00\x01\x00\xb7\x00\x00\x00\x00\x00\x00\x00\x95\x00\x00\x00\x00\x00\x00\x00\x4d\x49\x54\x00\x00\x74\x65\x73\x74\x64\x61\x74\x61\x2f\x6d\x69\x6e\x69\x6d\x61\x6c\x2e\x63\x00\x2e\x00\x5f\x5f\x6c\x69\x63\x65\x6e\x73\x65\x00\x63\x68\x61\x72\x00\x5f\x5f\x41\x52\x52\x41\x59\x5f\x53\x49\x5a\x45\x5f\x54\x59\x50\x45\x5f\x5f\x00\x66\x69\x6c\x74\x65\x72\x00\x69\x6e\x74\x00\x01\x11\x01\x25\x0e\x13\x05\x03\x0e\x10\x17\x1b\x0e\x11\x01\x12\x06\x00\x00\x02\x34\x00\x03\x0e\x49\x13\x3f\x19\x3a\x0b\x3b\x0b\x02\x18\x00\x00\x03\x01\x01\x49\x13\x00\x00\x04\x21\x00\x49\x13\x37\x0b\x00\x00\x05\x24\x00\x03\x0e\x3e\x0b\x0b\x0b\x00\x00\x06\x24\x00\x03\x0e\x0b\x0b\x3e\x0b\x00\x00\x07\x2e\x00\x11\x01\x12\x06\x40\x18\x03\x0e\x3a\x0b\x3b\x0b\x49\x13\x3f\x19\x00\x00\x00\x76\x00\x00\x00\x04\x00\x00\x00\x00\x00\x08\x01\x00\x00\x00\x00\x0c\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x10\x00\x00\x00\x02\x00\x00\x00\x00\x3f\x00\x00\x00\x01\x03\x09\x03\x00\x00\x00\x00\x00\x00\x00\x00\x03\x4b\x00\x00\x00\x04\x52\x00\x00\x00\x04\x00\x05\x00\x00\x00\x00\x06\x01\x06\x00\x00\x00\x00\x08\x07\x07\x00\x00\x00\x00\x00\x00\x00\x00\x10\x00\x00\x00\x01\x5a\x00\x00\x00\x00\x01\x05\x72\x00\x00\x00\x05\x00\x00\x00\x00\x05\x04\x00\x00\x9f\xeb\x01\x00\x18\x00\x00\x00\x00\x00\x00\x00\x88\x00\x00\x00\x88\x00\x00\x00\x5e\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0d\x02\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x01\x04\x00\x00\x00\x20\x00\x00\x01\x05\x00\x00\x00\x00\x00\x00\x0c\x01\x00\x00\x00\x33\x00\x00\x00\x00\x00\x00\x01\x01\x00\x00\x00\x08\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x03\x00\x00\x00\x00\x04\x00\x00\x00\x06\x00\x00\x00\x04\x00\x00\x00\x38\x00\x00\x00\x00\x00\x00\x01\x04\x00\x00\x00\x20\x00\x00\x00\x4c\x00\x00\x00\x00\x00\x00\x0e\x05\x00\x00\x00\x01\x00\x00\x00\x56\x00\x00\x00\x01\x00\x00\x0f\x00\x00\x00\x00\x07\x00\x00\x00\x00\x00\x00\x00\x04\x00\x00\x00\x00\x69\x6e\x74\x00\x66\x69\x6c\x74\x65\x72\x00\x73\x6f\x63\x6b\x65\x74\x00\x2e\x2f\x74\x65\x73\x74\x64\x61\x74\x61\x2f\x6d\x69\x6e\x69\x6d\x61\x6c\x2e\x63\x00\x09\x72\x65\x74\x75\x72\x6e\x20\x30\x3b\x00\x63\x68\x61\x72\x00\x5f\x5f\x41\x52\x52\x41\x59\x5f\x53\x49\x5a\x45\x5f\x54\x59\x50\x45\x5f\x5f\x00\x5f\x5f\x6c\x69\x63\x65\x6e\x73\x65\x00\x6c\x69\x63\x65\x6e\x73\x65\x00\x9f\xeb\x01\x00\x28\x00\x00\x00\x00\x00\x00\x00\x14\x00\x00\x00\x14\x00\x00\x00\x1c\x00\x00\x00\x30\x00\x00\x00\x00\x00\x00\x00\x30\x00\x00\x00\x00\x00\x00\x00\x08\x00\x00\x00\x0c\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x03\x00\x00\x00\x10\x00\x00\x00\x0c\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x13\x00\x00\x00\x28\x00\x00\x00\x02\x18\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0c\x00\x00\x00\xff\xff\xff\xff\x04\x00\x08\x00\x08\x7c\x0b\x00\x14\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x10\x00\x00\x00\x00\x00\x00\x00\x45\x00\x00\x00\x04\x00\x2a\x00\x00\x00\x08\x01\x01\xfb\x0e\x0d\x00\x01\x01\x01\x01\x00\x00\x00\x01\x00\x00\x01\x74\x65\x73\x74\x64\x61\x74\x61\x00\x00\x6d\x69\x6e\x69\x6d\x61\x6c\x2e\x63\x00\x01\x00\x00\x00\x00\x09\x02\x00\x00\x00\x00\x00\x00\x00\x00\x16\x05\x02\x0a\x13\x02\x02\x00\x01\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x93\x00\x00\x00\x04\x00\xf1\xff\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x05\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x05\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x05\x00\x14\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x05\x00\x16\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x05\x00\x20\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x05\x00\x25\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x05\x00\x39\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x05\x00\x40\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\x00\x03\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\x00\x06\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\x00\x0e\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\x00\x10\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x68\x00\x00\x00\x11\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x34\x00\x00\x00\x12\x00\x03\x00\x00\x00\x00\x00\x00\x00\x00\x00\x10\x00\x00\x00\x00\x00\x00\x00\x06\x00\x00\x00\x00\x00\x00\x00\x0a\x00\x00\x00\x0b\x00\x00\x00\x0c\x00\x00\x00\x00\x00\x00\x00\x0a\x00\x00\x00\x02\x00\x00\x00\x12\x00\x00\x00\x00\x00\x00\x00\x0a\x00\x00\x00\x03\x00\x00\x00\x16\x00\x00\x00\x00\x00\x00\x00\x0a\x00\x00\x00\x0d\x00\x00\x00\x1a\x00\x00\x00\x00\x00\x00\x00\x0a\x00\x00\x00\x04\x00\x00\x00\x1e\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x0a\x00\x00\x00\x2b\x00\x00\x00\x00\x00\x00\x00\x0a\x00\x00\x00\x05\x00\x00\x00\x37\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x0e\x00\x00\x00\x4c\x00\x00\x00\x00\x00\x00\x00\x0a\x00\x00\x00\x06\x00\x00\x00\x53\x00\x00\x00\x00\x00\x00\x00\x0a\x00\x00\x00\x07\x00\x00\x00\x5a\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x0a\x00\x00\x00\x68\x00\x00\x00\x00\x00\x00\x00\x0a\x00\x00\x00\x08\x00\x00\x00\x73\x00\x00\x00\x00\x00\x00\x00\x0a\x00\x00\x00\x09\x00\x00\x00\x98\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0e\x00\x00\x00\x34\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0a\x00\x00\x00\x48\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0a\x00\x00\x00\x14\x00\x00\x00\x00\x00\x00\x00\x0a\x00\x00\x00\x0c\x00\x00\x00\x18\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x0a\x00\x00\x00\x37\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x0a\x00\x00\x00\x0f\x0e\x00\x2e\x64\x65\x62\x75\x67\x5f\x61\x62\x62\x72\x65\x76\x00\x2e\x74\x65\x78\x74\x00\x2e\x72\x65\x6c\x2e\x42\x54\x46\x2e\x65\x78\x74\x00\x73\x6f\x63\x6b\x65\x74\x00\x2e\x64\x65\x62\x75\x67\x5f\x73\x74\x72\x00\x66\x69\x6c\x74\x65\x72\x00\x2e\x64\x65\x62\x75\x67\x5f\x6d\x61\x63\x69\x6e\x66\x6f\x00\x2e\x72\x65\x6c\x2e\x64\x65\x62\x75\x67\x5f\x69\x6e\x66\x6f\x00\x2e\x6c\x6c\x76\x6d\x5f\x61\x64\x64\x72\x73\x69\x67\x00\x5f\x5f\x6c\x69\x63\x65\x6e\x73\x65\x00\x2e\x72\x65\x6c\x2e\x64\x65\x62\x75\x67\x5f\x6c\x69\x6e\x65\x00\x2e\x72\x65\x6c\x2e\x64\x65\x62\x75\x67\x5f\x66\x72\x61\x6d\x65\x00\x6d\x69\x6e\x69\x6d\x61\x6c\x2e\x63\x00\x2e\x73\x74\x72\x74\x61\x62\x00\x2e\x73\x79\x6d\x74\x61\x62\x00\x2e\x72\x65\x6c\x2e\x42\x54\x46\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x9d\x00\x00\x00\x03\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xfa\x05\x00\x00\x00\x00\x00\x00\xb6\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0f\x00\x00\x00\x01\x00\x00\x00\x06\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x40\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x22\x00\x00\x00\x01\x00\x00\x00\x06\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x40\x00\x00\x00\x00\x00\x00\x00\x10\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x08\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x6a\x00\x00\x00\x01\x00\x00\x00\x03\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x50\x00\x00\x00\x00\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x29\x00\x00\x00\x01\x00\x00\x00\x30\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x54\x00\x00\x00\x00\x00\x00\x00\x44\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x98\x00\x00\x00\x00\x00\x00\x00\x60\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x4e\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf8\x00\x00\x00\x00\x00\x00\x00\x7a\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x4a\x00\x00\x00\x09\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xc8\x04\x00\x00\x00\x00\x00\x00\xd0\x00\x00\x00\x00\x00\x00\x00\x13\x00\x00\x00\x07\x00\x00\x00\x08\x00\x00\x00\x00\x00\x00\x00\x10\x00\x00\x00\x00\x00\x00\x00\x3b\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x72\x01\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xb1\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x73\x01\x00\x00\x00\x00\x00\x00\xfe\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xad\x00\x00\x00\x09\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x98\x05\x00\x00\x00\x00\x00\x00\x10\x00\x00\x00\x00\x00\x00\x00\x13\x00\x00\x00\x0a\x00\x00\x00\x08\x00\x00\x00\x00\x00\x00\x00\x10\x00\x00\x00\x00\x00\x00\x00\x19\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x71\x02\x00\x00\x00\x00\x00\x00\x58\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x15\x00\x00\x00\x09\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa8\x05\x00\x00\x00\x00\x00\x00\x20\x00\x00\x00\x00\x00\x00\x00\x13\x00\x00\x00\x0c\x00\x00\x00\x08\x00\x00\x00\x00\x00\x00\x00\x10\x00\x00\x00\x00\x00\x00\x00\x86\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xd0\x02\x00\x00\x00\x00\x00\x00\x28\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x08\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x82\x00\x00\x00\x09\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xc8\x05\x00\x00\x00\x00\x00\x00\x20\x00\x00\x00\x00\x00\x00\x00\x13\x00\x00\x00\x0e\x00\x00\x00\x08\x00\x00\x00\x00\x00\x00\x00\x10\x00\x00\x00\x00\x00\x00\x00\x76\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf8\x02\x00\x00\x00\x00\x00\x00\x49\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x72\x00\x00\x00\x09\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xe8\x05\x00\x00\x00\x00\x00\x00\x10\x00\x00\x00\x00\x00\x00\x00\x13\x00\x00\x00\x10\x00\x00\x00\x08\x00\x00\x00\x00\x00\x00\x00\x10\x00\x00\x00\x00\x00\x00\x00\x5a\x00\x00\x00\x03\x4c\xff\x6f\x00\x00\x00\x80\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf8\x05\x00\x00\x00\x00\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x13\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa5\x00\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x48\x03\x00\x00\x00\x00\x00\x00\x80\x01\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x0e\x00\x00\x00\x08\x00\x00\x00\x00\x00\x00\x00\x18\x00\x00\x00\x00\x00\x00\x00")
