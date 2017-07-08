package zrle

import ()

const (
	TILE_WIDTH  int = 64
	TILE_HEIGHT     = 64
)

// tile represents a subencoded tile
type tile struct {
	width    int
	height   int
	encoding subencodingType
	bytes    []byte
}

func createTiles(width int, height int) (tiles []tile) {
	for height > 0 {
		rowWidth := width

		// If row is shorter than TILE_HEIGHT adjust
		rowHeight := TILE_HEIGHT
		height -= rowHeight
		if height < 0 {
			rowHeight += height
		}

		for rowWidth > 0 {

			// If tile is narrower than TILE_WIDTH adjust
			tileWidth := TILE_WIDTH
			rowWidth -= tileWidth
			if rowWidth < 0 {
				tileWidth += rowWidth
			}
			newTile := tile{width: tileWidth, height: rowHeight}
			tiles = append(tiles, newTile)
		}
	}
	return
}
