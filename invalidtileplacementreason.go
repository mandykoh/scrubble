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

	// PlacementNotLinearReason indicates that a play was attempted such that
	// the tiles were not all in a straight (horizontal or vertical) line.
	PlacementNotLinearReason

	// PlacementNotContiguousReason indicates that a play was attempted such
	// that gaps would be created in the tiles.
	PlacementNotContiguousReason
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
	case PlacementNotLinearReason:
		return "PlacementNotLinearReason"
	case PlacementNotContiguousReason:
		return "PlacementNotContiguousReason"
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
	case PlacementNotLinearReason:
		return "PlacementNotLinear"
	case PlacementNotContiguousReason:
		return "PlacementNotContiguous"
	default:
		return "Unknown"
	}
}
