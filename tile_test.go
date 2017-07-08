package zrle

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCreateSingleTile(t *testing.T) {
	tiles := createTiles(64, 64)
	if len(tiles) != 1 {
		t.Errorf("expected 1 tile, got %d", len(tiles))
	}
	t1 := tiles[0]
	if t1.width != 64 || t1.height != 64 {
		t.Errorf("expected height and width 64, got %v", t1)
	}
}

func TestCreateRowOfTiles(t *testing.T) {
	tiles := createTiles(128, 64)
	if len(tiles) != 2 {
		t.Fatalf("expected 2 tiles, got %d", len(tiles))
	}
	t1 := tiles[0]
	if t1.width != 64 || t1.height != 64 {
		t.Errorf("expected height and width 64, got %v", t1)
	}
	t2 := tiles[1]
	if t2.width != 64 || t2.height != 64 {
		t.Errorf("expected height and width 64, got %v", t1)
	}
}

func TestCreateColumnOfTiles(t *testing.T) {
	tiles := createTiles(64, 128)
	if len(tiles) != 2 {
		t.Fatalf("expected 2 tiles, got %d", len(tiles))
	}
	t1 := tiles[0]
	if t1.width != 64 || t1.height != 64 {
		t.Errorf("expected height and width 64, got %v", t1)
	}
	t2 := tiles[1]
	if t2.width != 64 || t2.height != 64 {
		t.Errorf("expected height and width 64, got %v", t1)
	}
}

func TestCreateGridOfTiles(t *testing.T) {
	tiles := createTiles(128, 128)
	if len(tiles) != 4 {
		t.Fatalf("expected 4 tiles, got %d", len(tiles))
	}
}

func TestCreateIrregularTile(t *testing.T) {
	tiles := createTiles(30, 40)
	t1 := tiles[0]
	if t1.width != 30 || t1.height != 40 {
		t.Errorf("expected width 30 and height 40, got %v", t1)
	}
}

func TestCreateIrregularGridOfTiles(t *testing.T) {
	tiles := createTiles(80, 70)
	if len(tiles) != 4 {
		t.Fatalf("expected 4 tiles, got %d", len(tiles))
	}
	t1 := tiles[0]
	if t1.width != 64 || t1.height != 64 {
		t.Errorf("expected width 64 and height 64, got %v", t1)
	}
	t2 := tiles[1]
	if t2.width != 16 || t2.height != 64 {
		t.Errorf("expected width 16 and height 64, got %v", t2)
	}
	t3 := tiles[2]
	if t3.width != 64 || t3.height != 6 {
		t.Errorf("expected width 64 and height 6, got %v", t3)
	}
	t4 := tiles[3]
	if t4.width != 16 || t4.height != 6 {
		t.Errorf("expected width 16 and height 6, got %v", t4)
	}
}

func TestRawTileRead(t *testing.T) {
	buf := &bytes.Buffer{}
	bs := []byte{1, 2, 3, 4}
	buf.Write(bs)

	rawtile := RawTile{tile{width: 2, height: 1, bytesPerCPixel: 2}}
	n, err := rawtile.Read(buf)
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
