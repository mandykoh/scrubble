package scrubble

import "github.com/mandykoh/scrubble/tile"

// TilePlacement represents the placement of a single tile on a board.
type TilePlacement struct {
	Tile tile.Tile
	Coord
}
