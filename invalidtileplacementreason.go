package scrubble

const (
	// UnknownInvalidTilePlacementReason indicates that a reason was undefined.
	UnknownInvalidTilePlacementReason InvalidTilePlacementReason = iota

	// NoTilesPlacedReason indicates that a play attempted to place zero tiles.
	NoTilesPlacedReason

	// PositionOccupiedReason indicates that a play attempted to place a tile in
	// an already occupied position.
	PositionOccupiedReason

	// PlacementOutOfBoundsReason indicates that a play attempted to place a
	// tile outside the valid bounds of the board.
	PlacementOutOfBoundsReason
)

// InvalidTilePlacementReason indicates the reason for an
// InvalidTilePlacementError.
type InvalidTilePlacementReason int

// GoString returns the Go syntax representation of the reason, or
// UnknownInvalidTilePlacementReason if it is not a valid reason.
func (r InvalidTilePlacementReason) GoString() string {
	switch r {
	case NoTilesPlacedReason:
		return "NoTilesPlacedReason"
	case PositionOccupiedReason:
		return "PositionOccupiedReason"
	case PlacementOutOfBoundsReason:
		return "PlacementOutOfBoundsReason"
	default:
		return "UnknownInvalidTilePlacementReason"
	}
}

// String returns the textual representation of the reason, or "Unknown" if
// it is not a valid reason.
func (r InvalidTilePlacementReason) String() string {
	switch r {
	case NoTilesPlacedReason:
		return "NoTilesPlaced"
	case PositionOccupiedReason:
		return "PositionOccupied"
	case PlacementOutOfBoundsReason:
		return "PlacementOutOfBounds"
	default:
		return "Unknown"
	}
}
