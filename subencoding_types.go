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

func getSubencoding(b byte) (encoding subencodingType) {
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
	return
}
