package zrle

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"io"
	"reflect"
	"testing"
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
	n, err := w.Write(bs)
	if err != nil {
		t.Error(err)
	}
	w.Flush()

	encoded_length := uint32(n)

	decoded, err := decode(buf, encoded_length)

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

func TestZRLERead(t *testing.T) {
	enc := ZRLEEncoding{
		getLength: func(buf io.Reader) (uint32, error) { return 1, nil },
		decode: func(buf io.Reader, length uint32) ([]byte, error) {
			p := make([]byte, length)
			_, err := buf.Read(p)
			if err != nil {
				return p, err
			}
			return p, nil
		},
	}
	buf := bytes.NewBuffer([]byte{1})
	enc.Read(buf)
	if !reflect.DeepEqual(enc.data, []byte{1}) {
		t.Errorf("expected %v, got %v", []byte{1}, enc.data)
	}
}
