package zrle

type subencodingType uint8

const (
	raw subencodingType = iota
	solid
	packedPalette
	rle
	prle
	invalid
)
