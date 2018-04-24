package board

import (
	"github.com/mandykoh/scrubble/tile"
)

// Position represents a single position on a Board, which may be occupied
// by a tile and may have a special property.
type Position struct {
	Type PositionType
	Tile *tile.Tile
}
