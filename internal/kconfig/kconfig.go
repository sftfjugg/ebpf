package kconfig

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cilium/ebpf/btf"
	"github.com/cilium/ebpf/internal"
)

// ParseKconfig parses the kconfig file which path is given at parameter.
// This function can read gzip'ed kconfig file is the corresponding flag is set
// to true.
// All the CONFIG_* set will be put in the returned map as key with their
// corresponding value as map value.
// If the kconfig file is not valid, error will be returned.
func ParseKconfig(path string, gziped bool) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	if gziped {
		zr, err := gzip.NewReader(f)
		if err != nil {
			return nil, err
		}
		defer zr.Close()

		s = bufio.NewScanner(zr)
	}

	ret := make(map[string]string)

	for s.Scan() {
		line := s.Text()
		err = processKconfigLine(line, ret, false)
		if err != nil {
			return nil, fmt.Errorf("cannot parse line %q: %w", line, err)
		}
	}

	return ret, nil
}

// PatchKconfig adds the CONFIG_* options which are given in the patch to the
// kconfig map given as parameter.
// If a CONFIG_* already exists in the map and is also present in the patch,
// the final value will be that of the patch.
func PatchKconfig(kconfig map[string]string, patch string) error {
	sr := strings.NewReader(patch)
	s := bufio.NewScanner(sr)

	for s.Scan() {
		line := s.Text()
		err := processKconfigLine(line, kconfig, true)
		if err != nil {
			return fmt.Errorf("cannot parse line %q: %w", line, err)
		}
	}

	return nil
}

// Golang translation of libbpf bpf_object__process_kconfig_line():
// https://github.com/libbpf/libbpf/blob/fbd60dbff51c870f5e80a17c4f2fd639eb80af90/src/libbpf.c#L1874
// It does the same checks but does not put the data inside the BPF map.
func processKconfigLine(line string, m map[string]string, overwrite bool) error {
	// Ignore empty lines and "# CONFIG_* is not set".
	if !strings.HasPrefix(line, "CONFIG_") {
		return nil
	}

	key, value, found := strings.Cut(line, "=")
	if !found {
		return fmt.Errorf("line %q does not contain separator '='", line)
	}

	value = strings.Trim(value, "\n")
	if len(value) == 0 {
		return fmt.Errorf("line %q has no value", line)
	}

	_, ok := m[key]
	if !ok || overwrite {
		m[key] = value
	}

	return nil
}

// PutKconfigValue translates the value given as parameter depending on the BTF
// type, the translated value is then written to the byte array.
func PutKconfigValue(data []byte, typ btf.Type, value string) error {
	typ = btf.UnderlyingType(typ)

	switch value {
	case "y", "n", "m":
		return putKconfigValueTri(data, typ, value)
	default:
		if strings.HasPrefix(value, `"`) {
			return putKconfigValueString(data, typ, value)
		}
		return putKconfigValueNumber(data, typ, value)
	}
}

// Golang translation of libbpf_tristate enum:
// https://github.com/libbpf/libbpf/blob/fbd60dbff51c870f5e80a17c4f2fd639eb80af90/src/bpf_helpers.h#L169
type TriState int

const (
	TriNo     TriState = 0
	TriYes    TriState = 1
	TriModule TriState = 2
)

func putKconfigValueTri(data []byte, typ btf.Type, value string) error {
	switch v := typ.(type) {
	case *btf.Int:
		if v.Encoding != btf.Bool {
			return fmt.Errorf("cannot add tri value, expected btf.Bool, got: %v", v.Encoding)
		}

		if v.Size != 1 {
			return fmt.Errorf("cannot add tri value, expected size of 1 byte, got: %d", v.Size)
		}

		switch value {
		case "y":
			data[0] = 1
		case "n":
			data[0] = 0
		default:
			return fmt.Errorf("cannot use %q for btf.Bool", value)
		}
	case *btf.Enum:
		if v.Name != "libbpf_tristate" {
			return fmt.Errorf("cannot use enum %q, only libbpf_tristate is supported", v.Name)
		}

		var tri TriState
		switch value {
		case "y":
			tri = TriYes
		case "m":
			tri = TriModule
		case "n":
			tri = TriNo
		default:
			return fmt.Errorf("value %q is not support for libbpf_tristate", value)
		}

		internal.NativeEndian.PutUint64(data, uint64(tri))
	default:
		return fmt.Errorf("cannot add number value, expected btf.Int or btf.Enum, got: %T", v)
	}

	return nil
}

func putKconfigValueString(data []byte, typ btf.Type, value string) error {
	array, ok := typ.(*btf.Array)
	if !ok {
		return fmt.Errorf("cannot add string value, expected btf.Array, got %T", array)
	}

	contentType, ok := array.Type.(*btf.Int)
	if !ok {
		return fmt.Errorf("cannot add string value, expected array of btf.Int, got %T", contentType)
	}

	// Treat unsigned int8 as char.
	if contentType.Encoding != btf.Char && contentType.Encoding != btf.Unsigned {
		return fmt.Errorf("cannot add string value, expected array of btf.Char, got array of: %v", contentType.Encoding)
	}

	if contentType.Size != 1 {
		return fmt.Errorf("cannot add string value, expected array of btf.Char of size 1, got array of btf.Char of size: %v", contentType.Size)
	}

	if !strings.HasPrefix(value, `"`) || !strings.HasSuffix(value, `"`) {
		return fmt.Errorf(`value %q must start and finish with '"'`, value)
	}

	str := strings.Trim(value, `"`)

	// We need to trim string if the bpf array is smaller.
	if uint32(len(str)) >= array.Nelems {
		str = str[:array.Nelems]
	}

	// Write the string content to .kconfig.
	copy(data, str)

	return nil
}

func putKconfigValueNumber(data []byte, typ btf.Type, value string) error {
	integer, ok := typ.(*btf.Int)
	if !ok {
		return fmt.Errorf("cannot add number value, expected *btf.Int, got: %T", integer)
	}

	size := integer.Size
	sizeInBits := size * 8

	var n int
	var err error
	if integer.Encoding == btf.Signed {
		parsed, e := strconv.ParseInt(value, 0, int(sizeInBits))

		n = int(parsed)
		err = e
	} else {
		parsed, e := strconv.ParseUint(value, 0, int(sizeInBits))

		n = int(parsed)
		err = e
	}

	if err != nil {
		return fmt.Errorf("cannot parse value: %w", err)
	}

	switch size {
	case 1:
		data[0] = byte(n)
	case 2:
		internal.NativeEndian.PutUint16(data, uint16(n))
	case 4:
		internal.NativeEndian.PutUint32(data, uint32(n))
	case 8:
		internal.NativeEndian.PutUint64(data, uint64(n))
	default:
		return fmt.Errorf("size (%d) is not valid, expected: 1, 2, 4 or 8", size)
	}

	return nil
}
