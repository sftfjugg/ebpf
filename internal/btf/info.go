package btf

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/cilium/ebpf/internal"
)

type BTFInfo struct {
	BTF       *Spec
	id        TypeID
	Name      string
	KernelBTF bool
}

func newBTFInfoFromFd(fd *internal.FD) (*BTFInfo, error) {
	// We invoke the syscall once with a empty BTF and name buffers to get size
	// information to allocate buffers. Then we invoke it a second time with
	// buffers to receive the data.
	info, err := bpfGetBTFInfoByFD(fd, nil, nil)
	if err != nil {
		return nil, err
	}

	btfBuffer := make([]byte, info.btfSize)
	nameBuffer := make([]byte, info.nameLen)
	info, err = bpfGetBTFInfoByFD(fd, btfBuffer, nameBuffer)
	if err != nil {
		return nil, err
	}

	spec, err := loadNakedSpec(bytes.NewReader(btfBuffer), internal.NativeEndian, nil, nil)
	if err != nil {
		return nil, err
	}

	return &BTFInfo{
		BTF:       spec,
		id:        TypeID(info.id),
		Name:      internal.CString(nameBuffer),
		KernelBTF: info.kernelBTF != 0,
	}, nil
}

func scanFdInfo(fd *internal.FD, fields map[string]interface{}) error {
	raw, err := fd.Value()
	if err != nil {
		return err
	}

	fh, err := os.Open(fmt.Sprintf("/proc/self/fdinfo/%d", raw))
	if err != nil {
		return err
	}
	defer fh.Close()

	if err := scanFdInfoReader(fh, fields); err != nil {
		return fmt.Errorf("%s: %w", fh.Name(), err)
	}
	return nil
}

var errMissingFields = errors.New("missing fields")

func scanFdInfoReader(r io.Reader, fields map[string]interface{}) error {
	var (
		scanner = bufio.NewScanner(r)
		scanned int
	)

	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "\t", 2)
		if len(parts) != 2 {
			continue
		}

		name := strings.TrimSuffix(parts[0], ":")
		field, ok := fields[string(name)]
		if !ok {
			continue
		}

		if n, err := fmt.Sscanln(parts[1], field); err != nil || n != 1 {
			return fmt.Errorf("can't parse field %s: %v", name, err)
		}

		scanned++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if scanned != len(fields) {
		return errMissingFields
	}

	return nil
}
