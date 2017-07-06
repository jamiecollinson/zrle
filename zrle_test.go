package zrle

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestReadsHeader(t *testing.T) {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, 123)

	buf := &bytes.Buffer{}
	buf.Write(bs)

	length, _ := getLength(buf)

	if length != 123 {
		t.Errorf("expected 123, got: %v", length)
	}
}
