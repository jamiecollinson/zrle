package zrle

type subencoding interface {
	SubType() subType
	String() string
}

type rawEncoding struct{}

func (e rawEncoding) SubType() subType { return raw }
func (e rawEncoding) String() string   { return "Raw" }

type solidEncoding struct{}

func (e solidEncoding) SubType() subType { return solid }
func (e solidEncoding) String() string   { return "Solid" }

type packedPaletteEncoding struct{}

func (e packedPaletteEncoding) SubType() subType { return packedPalette }
func (e packedPaletteEncoding) String() string   { return "PackedPalette" }

type rleEncoding struct{}

func (e rleEncoding) SubType() subType { return rle }
func (e rleEncoding) String() string   { return "RLE" }

type prleEncoding struct{}

func (e prleEncoding) SubType() subType { return prle }
func (e prleEncoding) String() string   { return "PRLE" }

type nullEncoding struct{}

func (e nullEncoding) SubType() subType { return invalid }
func (e nullEncoding) String() string   { return "INVALID" }

func getSubencoding(b byte) (encoding subencoding) {
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
		encoding = nullEncoding{}
	}
	return
}
