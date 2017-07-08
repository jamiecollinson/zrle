package zrle

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
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

func TestSubencodingDispatch(t *testing.T) {
	// RAW
	subencoding := getSubencoding(0)
	if subencoding != raw {
		t.Errorf("expected %v, got %v", raw, subencoding)
	}

	// solid
	subencoding = getSubencoding(1)
	if subencoding != solid {
		t.Errorf("expected %v, got %v", solid, subencoding)
	}

	// packedPalette
	subencoding = getSubencoding(5)
	if subencoding != packedPalette {
		t.Errorf("expected %v, got %v", packedPalette, subencoding)
	}

	// 17-127 invalid
	subencoding = getSubencoding(20)
	if subencoding != invalid {
		t.Errorf("expected %v, got %v", invalid, subencoding)
	}

	// RLE
	subencoding = getSubencoding(128)
	if subencoding != rle {
		t.Errorf("expected %v, got %v", rle, subencoding)
	}

	// 129 invalid
	subencoding = getSubencoding(129)
	if subencoding != invalid {
		t.Errorf("expected %v, got %v", invalid, subencoding)
	}

	// PRLE
	subencoding = getSubencoding(130)
	if subencoding != prle {
		t.Errorf("expected %v, got %v", prle, subencoding)
	}
}
