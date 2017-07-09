package zrle

import ()

const (
	TILE_WIDTH  int = 64
	TILE_HEIGHT     = 64
)

type tile struct {
	width, height, bytesPerCPixel int
	pixels                        []cpixel
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
