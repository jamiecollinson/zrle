package zrle

import (
	"io"
	"encoding/binary"
)

func getLength(buf io.Reader) (uint32, error) {
	p := make([]byte, 4)
	_, err := buf.Read(p)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(p), nil
}
