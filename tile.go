package zrle

import (
	"io"
)

const (
	TILE_WIDTH  int = 64
	TILE_HEIGHT     = 64
)

type cpixel []byte

type tile struct {
	width, height, bytesPerCPixel int
	pixels                        []cpixel
}

type RawTile struct {
	tile
}

func (t *RawTile) Read(buf io.Reader) (int, error) {
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

func createTiles(width int, height int) (tiles []tile) {
	for height > 0 {
		rowWidth := width

		// If row is shorter than TILE_HEIGHT adjust
		rowHeight := TILE_HEIGHT
		if height < rowHeight {
			rowHeight = height
		}
		height -= rowHeight

		for rowWidth > 0 {

			// If tile is narrower than TILE_WIDTH adjust
			tileWidth := TILE_WIDTH
			if rowWidth < tileWidth {
				tileWidth = rowWidth
			}
			rowWidth -= tileWidth

			newTile := tile{width: tileWidth, height: rowHeight}
			tiles = append(tiles, newTile)
		}
	}
	return
}
