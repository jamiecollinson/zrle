package zrle

import ()

const (
	TILE_WIDTH  int = 64
	TILE_HEIGHT     = 64
)

type tileConfig struct {
	width, height int
}

type tile struct {
	x, y, width, height, bytesPerCPixel int
	pixels                              []cpixel
}

func (t tile) toPixelGrid() [][]cpixel {
	pixels := make([][]cpixel, t.height)
	for i := range pixels {
		pixels[i] = make([]cpixel, t.width)
	}

	i, j := 0, 0
	for _, pixel := range t.pixels {
		if j >= t.width {
			i++
			j = 0
		}
		pixels[i][j] = pixel
		j++
	}
	return pixels
}

func createTiles(width int, height int) (tiles []tile) {
	x, y := 0, 0
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

			newTile := tile{x: x, y: y, width: tileWidth, height: rowHeight}
			tiles = append(tiles, newTile)

			x += tileWidth
		}
		x = 0
		y += rowHeight
	}
	return
}

func tilesToPixels(width int, height int, tiles []tile) [][]cpixel {
	pixels := make([][]cpixel, height)
	for i := range pixels {
		pixels[i] = make([]cpixel, width)
	}
	for _, tile := range tiles {
		tilePixels := tile.toPixelGrid()
		for i, tileRow := range tilePixels {
			for j, pixel := range tileRow {
				pixels[tile.y+i][tile.x+j] = pixel
			}
		}
	}
	return pixels
}
