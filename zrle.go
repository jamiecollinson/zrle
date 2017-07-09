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

type ZRLEEncoding struct {
	width, height int
	data          []byte
	getLength     func(io.Reader) (uint32, error)
	decode        func(io.Reader, uint32) ([]byte, error)
}

func NewZRLEEncoding(width, height int) *ZRLEEncoding {
	return &ZRLEEncoding{
		width:     width,
		height:    height,
		getLength: getLength,
		decode:    decode,
	}
}

func (e *ZRLEEncoding) Read(buf io.Reader) (err error) {
	n, err := e.getLength(buf)
	if err != nil {
		return err
	}
	e.data, err = e.decode(buf, n)
	return
}
