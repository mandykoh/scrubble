package scrubble

const (
	// UnknownInvalidTilePlacementReason indicates that a reason was undefined.
	UnknownInvalidTilePlacementReason InvalidTilePlacementReason = iota

	// NoTilesPlacedReason indicates that a play attempted to place zero tiles.
	NoTilesPlacedReason
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
	default:
		return "Unknown"
	}
}
