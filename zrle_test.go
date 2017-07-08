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
	buf := &bytes.Buffer{}

	// RAW
	buf.WriteByte(0)
	subencoding, err := getSubencoding(buf)
	if err != nil {
		t.Error(err)
	}
	if subencoding != raw {
		t.Errorf("expected %v, got %v", raw, subencoding)
	}

	// solid
	buf.Reset()
	buf.WriteByte(1)
	subencoding, err = getSubencoding(buf)
	if err != nil {
		t.Error(err)
	}
	if subencoding != solid {
		t.Errorf("expected %v, got %v", solid, subencoding)
	}

	// solid
	buf.Reset()
	buf.WriteByte(1)
	subencoding, err = getSubencoding(buf)
	if err != nil {
		t.Error(err)
	}
	if subencoding != solid {
		t.Errorf("expected %v, got %v", solid, subencoding)
	}

	// packedPalette
	buf.Reset()
	buf.WriteByte(5)
	subencoding, err = getSubencoding(buf)
	if err != nil {
		t.Error(err)
	}
	if subencoding != packedPalette {
		t.Errorf("expected %v, got %v", packedPalette, subencoding)
	}

	// 17-127 invalid
	buf.Reset()
	buf.WriteByte(20)
	subencoding, err = getSubencoding(buf)
	if err != nil {
		t.Error(err)
	}
	if subencoding != invalid {
		t.Errorf("expected %v, got %v", invalid, subencoding)
	}

	// RLE
	buf.Reset()
	buf.WriteByte(128)
	subencoding, err = getSubencoding(buf)
	if err != nil {
		t.Error(err)
	}
	if subencoding != rle {
		t.Errorf("expected %v, got %v", rle, subencoding)
	}

	// 129 invalid
	buf.Reset()
	buf.WriteByte(129)
	subencoding, err = getSubencoding(buf)
	if err != nil {
		t.Error(err)
	}
	if subencoding != invalid {
		t.Errorf("expected %v, got %v", invalid, subencoding)
	}

	// PRLE
	buf.Reset()
	buf.WriteByte(130)
	subencoding, err = getSubencoding(buf)
	if err != nil {
		t.Error(err)
	}
	if subencoding != prle {
		t.Errorf("expected %v, got %v", prle, subencoding)
	}
}
