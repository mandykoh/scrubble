package scrubble

// BoardPosition represents a single position on a Board, which may be occupied
// by a tile and may have a special property.
type BoardPosition struct {
	Type PositionType
	Tile *Tile
}
