package zrle

import (
	"testing"
)

func TestSubencodingDispatch(t *testing.T) {
	// RAW
	subencoding := getSubencoding(0)
	if subencoding.SubType() != raw {
		t.Errorf("expected %v, got %v", raw, subencoding)
	}

	// solid
	subencoding = getSubencoding(1)
	if subencoding.SubType() != solid {
		t.Errorf("expected %v, got %v", solid, subencoding)
	}

	// packedPalette
	subencoding = getSubencoding(5)
	if subencoding.SubType() != packedPalette {
		t.Errorf("expected %v, got %v", packedPalette, subencoding)
	}

	// 17-127 invalid
	subencoding = getSubencoding(20)
	if subencoding.SubType() != invalid {
		t.Errorf("expected %v, got %v", invalid, subencoding)
	}

	// RLE
	subencoding = getSubencoding(128)
	if subencoding.SubType() != rle {
		t.Errorf("expected %v, got %v", rle, subencoding)
	}

	// 129 invalid
	subencoding = getSubencoding(129)
	if subencoding.SubType() != invalid {
		t.Errorf("expected %v, got %v", invalid, subencoding)
	}

	// PRLE
	subencoding = getSubencoding(130)
	if subencoding.SubType() != prle {
		t.Errorf("expected %v, got %v", prle, subencoding)
	}
}
