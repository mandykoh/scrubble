package scrubble

import "github.com/mandykoh/scrubble/positiontype"

// BoardPosition represents a single position on a Board, which may be occupied
// by a tile and may have a special property.
type BoardPosition struct {
	Type positiontype.Interface
	Tile *Tile
}
