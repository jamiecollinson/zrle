package zrle

import (
	"errors"
	"io"
)

type subencoding interface {
	SubType() subType
}

type rawEncoding struct{}

func (rawEncoding) SubType() subType { return raw }

func (rawEncoding) Read(buf io.Reader, t *tile) (int, error) {
	bytesToRead := t.width * t.height * t.bytesPerCPixel
	bytesRead := 0
	for bytesToRead > 0 {
		pixel := make(cpixel, t.bytesPerCPixel)
		n, err := buf.Read(pixel)
		bytesRead += n
		if err != nil {
			return bytesRead, err
		}
		t.pixels = append(t.pixels, pixel)
		bytesToRead -= len(pixel)
	}
	return bytesRead, nil
}

type solidEncoding struct{}

func (solidEncoding) SubType() subType { return solid }

type packedPaletteEncoding struct{}

func (packedPaletteEncoding) SubType() subType { return packedPalette }

type rleEncoding struct{}

func (rleEncoding) SubType() subType { return rle }

type prleEncoding struct{}

func (prleEncoding) SubType() subType { return prle }

func getSubencoding(b byte) (encoding subencoding, err error) {
	switch {
	case b == 0:
		encoding = rawEncoding{}
	case b == 1:
		encoding = solidEncoding{}
	case b <= 16:
		encoding = packedPaletteEncoding{}
	case b == 128:
		encoding = rleEncoding{}
	case b >= 130:
		encoding = prleEncoding{}
	default:
		err = errors.New("No valid encoding")
	}
	return
}
