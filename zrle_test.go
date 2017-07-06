package zrle

import (
	"bytes"
	"encoding/binary"
	"testing"
	"compress/zlib"
)

func TestReadsLength(t *testing.T) {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, 123)

	buf := &bytes.Buffer{}
	buf.Write(bs)

	length, _ := getLength(buf)

	if length != 123 {
		t.Errorf("expected 123, got: %v", length)
	}
}

func TestDecodesZlib(t *testing.T) {
	bs := []byte("testtesttest")

	buf := &bytes.Buffer{}
	w := zlib.NewWriter(buf)
	w.Write(bs)
	
	length := uint32(buf.Len())

	decoded, err := decode(buf, length)

	// no error should be raised
	if err != nil {
		t.Error(err)
	}

	// buf should now be empty
	if buf.Len() != 0 {
		t.Errorf("buffer should be empty but has length: %d", buf.Len())
	}

	// decoded should be original
	if !bytes.Equal(decoded, bs) {
		t.Errorf("expected %v, got %v", bs, decoded)
	}
}
