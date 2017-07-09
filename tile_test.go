package zrle

import (
	"reflect"
	"testing"
)

func TestCreateSingleTile(t *testing.T) {
	tiles := createTiles(64, 64)
	if len(tiles) != 1 {
		t.Errorf("expected 1 tile, got %d", len(tiles))
	}
	t1 := tiles[0]
	expected := tile{x: 0, y: 0, width: 64, height: 64}
	if !reflect.DeepEqual(t1, expected) {
		t.Errorf("expected %v, got %v", expected, t1)
	}
}

func TestCreateRowOfTiles(t *testing.T) {
	tiles := createTiles(128, 64)
	if len(tiles) != 2 {
		t.Fatalf("expected 2 tiles, got %d", len(tiles))
	}
	t1 := tiles[0]
	expected := tile{x: 0, y: 0, width: 64, height: 64}
	if !reflect.DeepEqual(t1, expected) {
		t.Errorf("expected %v, got %v", expected, t1)
	}
	t2 := tiles[1]
	expected = tile{x: 64, y: 0, width: 64, height: 64}
	if !reflect.DeepEqual(t2, expected) {
		t.Errorf("expected %v, got %v", expected, t2)
	}
}

func TestCreateColumnOfTiles(t *testing.T) {
	tiles := createTiles(64, 128)
	if len(tiles) != 2 {
		t.Fatalf("expected 2 tiles, got %d", len(tiles))
	}
	t1 := tiles[0]
	expected := tile{x: 0, y: 0, width: 64, height: 64}
	if !reflect.DeepEqual(t1, expected) {
		t.Errorf("expected %v, got %v", expected, t1)
	}
	t2 := tiles[1]
	expected = tile{x: 0, y: 64, width: 64, height: 64}
	if !reflect.DeepEqual(t2, expected) {
		t.Errorf("expected %v, got %v", expected, t2)
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
	expected := tile{x: 0, y: 0, width: 64, height: 64}
	if !reflect.DeepEqual(t1, expected) {
		t.Errorf("expected %v, got %v", expected, t1)
	}
	t2 := tiles[1]
	expected = tile{x: 64, y: 0, width: 16, height: 64}
	if !reflect.DeepEqual(t2, expected) {
		t.Errorf("expected %v, got %v", expected, t2)
	}
	t3 := tiles[2]
	expected = tile{x: 0, y: 64, width: 64, height: 6}
	if !reflect.DeepEqual(t3, expected) {
		t.Errorf("expected %v, got %v", expected, t3)
	}
	t4 := tiles[3]
	expected = tile{x: 64, y: 64, width: 16, height: 6}
	if !reflect.DeepEqual(t4, expected) {
		t.Errorf("expected %v, got %v", expected, t4)
	}
}

func TestTileToPixelGrid_GridCase(t *testing.T) {
	tile := tile{
		width:  2,
		height: 2,
		pixels: []cpixel{
			cpixel{0}, cpixel{1}, cpixel{2}, cpixel{3},
		},
	}
	pixels := tile.toPixelGrid()
	expected := [][]cpixel{
		[]cpixel{cpixel{0}, cpixel{1}},
		[]cpixel{cpixel{2}, cpixel{3}},
	}
	if !reflect.DeepEqual(expected, pixels) {
		t.Errorf("expected %v, got %v", expected, pixels)
	}
}

func TestTileToPixelGrid_ColumnCase(t *testing.T) {
	tile := tile{
		width:  4,
		height: 1,
		pixels: []cpixel{
			cpixel{0}, cpixel{1}, cpixel{2}, cpixel{3},
		},
	}
	pixels := tile.toPixelGrid()
	expected := [][]cpixel{
		[]cpixel{cpixel{0}, cpixel{1}, cpixel{2}, cpixel{3}},
	}
	if !reflect.DeepEqual(expected, pixels) {
		t.Errorf("expected %v, got %v", expected, pixels)
	}
}

func TestTilesToPixels(t *testing.T) {
	tiles := []tile{
		tile{x: 0, y: 0, width: 1, height: 1, pixels: []cpixel{cpixel{0}}},
		tile{x: 1, y: 0, width: 2, height: 1, pixels: []cpixel{cpixel{1}, cpixel{2}}},
	}
	pixels := tilesToPixels(3, 1, tiles)
	expected := [][]cpixel{[]cpixel{cpixel{0}, cpixel{1}, cpixel{2}}}
	if !reflect.DeepEqual(expected, pixels) {
		t.Errorf("expected %v, got %v", expected, pixels)
	}
}
