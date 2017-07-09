package zrle

import (
	"bytes"
	"reflect"
	"testing"
)

func TestSubencodingDispatch(t *testing.T) {
	// RAW
	subencoding, err := getSubencoding(0)
	if err != nil {
		t.Error(err)
	}
	if subencoding.SubType() != raw {
		t.Errorf("expected %v, got %v", raw, subencoding)
	}

	// solid
	subencoding, err = getSubencoding(1)
	if err != nil {
		t.Error(err)
	}
	if subencoding.SubType() != solid {
		t.Errorf("expected %v, got %v", solid, subencoding)
	}

	// packedPalette
	subencoding, err = getSubencoding(5)
	if err != nil {
		t.Error(err)
	}
	if subencoding.SubType() != packedPalette {
		t.Errorf("expected %v, got %v", packedPalette, subencoding)
	}

	// 17-127 invalid
	_, err = getSubencoding(20)
	if err == nil {
		t.Errorf("expected error, got none")
	}

	// RLE
	subencoding, err = getSubencoding(128)
	if err != nil {
		t.Error(err)
	}
	if subencoding.SubType() != rle {
		t.Errorf("expected %v, got %v", rle, subencoding)
	}

	// 129 invalid
	_, err = getSubencoding(129)
	if err == nil {
		t.Errorf("expected error, got none")
	}

	// PRLE
	subencoding, err = getSubencoding(130)
	if err != nil {
		t.Error(err)
	}
	if subencoding.SubType() != prle {
		t.Errorf("expected %v, got %v", prle, subencoding)
	}
}

func TestRawEncodingRead(t *testing.T) {
	buf := &bytes.Buffer{}
	bs := []byte{1, 2, 3, 4}
	buf.Write(bs)

	rawtile := &tile{width: 2, height: 1, bytesPerCPixel: 2}
	n, err := rawEncoding{}.Read(buf, rawtile)
	if err != nil {
		t.Fatal(err)
	}

	// expect 4 bytes read
	if n != 4 {
		t.Errorf("Expected length 4, got %d", n)
	}
	if buf.Len() != 0 {
		t.Errorf("Expected 0 bytes remaining, got %d", buf.Len())
	}

	// expect pixels to be parsed
	if len(rawtile.pixels) != 2 {
		t.Errorf("expected 2 pixels, got %d", len(rawtile.pixels))
	}
	if !reflect.DeepEqual(rawtile.pixels[0], cpixel([]byte{1, 2})) {
		t.Errorf("pixel doesn't match: %v", rawtile.pixels[0])
	}
	if !reflect.DeepEqual(rawtile.pixels[1], cpixel([]byte{3, 4})) {
		t.Errorf("pixel doesn't match: %v", rawtile.pixels[1])
	}
}
