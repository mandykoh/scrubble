package scrubble

import (
	"github.com/mandykoh/scrubble/positiontype"
	"github.com/mandykoh/scrubble/tile"
)

// BoardPosition represents a single position on a Board, which may be occupied
// by a tile and may have a special property.
type BoardPosition struct {
	Type positiontype.Interface
	Tile *tile.Tile
}
