package zrle

type subType uint8

const (
	raw subType = iota
	solid
	packedPalette
	rle
	prle
	invalid
)
