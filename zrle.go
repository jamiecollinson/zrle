package zrle

import (
	"compress/zlib"
	"encoding/binary"
	"io"
)

func getLength(buf io.Reader) (uint32, error) {
	p := make([]byte, 4)
	_, err := buf.Read(p)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(p), nil
}

func decode(buf io.Reader, length uint32) ([]byte, error) {
	p := make([]byte, length)
	r, err := zlib.NewReader(buf)
	if err != nil {
		return []byte{}, err
	}
	_, err = r.Read(p)
	return p, err
}

// subEncoding takes a tile and returns an array of cpixels
type subencoding interface {
	decode() []cpixel
	scheme() subencodingType
}

// cpixel represents a compressed pixel
type cpixel []byte

// tile represents a subencoded tile
type tile struct {
	width    int
	height   int
	encoding subencoding
}

func getSubencoding(buf io.Reader) (subencodingType, error) {
	p := make([]byte, 1)
	_, err := buf.Read(p)
	if err != nil {
		return invalid, err
	}
	b := p[0]

	var encoding subencodingType
	switch {
	case b == 0:
		encoding = raw
	case b == 1:
		encoding = solid
	case b <= 16:
		encoding = packedPalette
	case b == 128:
		encoding = rle
	case b >= 130:
		encoding = prle
	default:
		encoding = invalid
	}

	return encoding, nil
}
