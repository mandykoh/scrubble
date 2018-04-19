package scrubble

const (
	// UnknownInvalidTileExchangeReason indicates that a reason was undefined.
	UnknownInvalidTileExchangeReason InvalidTileExchangeReason = iota

	// NoTilesExchangedReason indicates that an exchange of zero tiles was attempted.
	NoTilesExchangedReason

	// InsufficientTilesInBagReason indicates that an exchange was attempted to
	// when the bag did not contain enough tiles.
	InsufficientTilesInBagReason
)

// InvalidTileExchangeReason indicates the reason for an
// InvalidTileExchangeError.
type InvalidTileExchangeReason int

// GoString returns the Go syntax representation of the reason, or
// UnknownInvalidTileExchangeReason if it is not a valid reason.
func (r InvalidTileExchangeReason) GoString() string {
	switch r {
	case NoTilesExchangedReason:
		return "NoTilesExchangedReason"
	case InsufficientTilesInBagReason:
		return "InsufficientTilesInBagReason"
	default:
		return "UnknownInvalidTileExchangeReason"
	}
}

// String returns the textual representation of the reason, or "Unknown" if
// it is not a valid reason.
func (r InvalidTileExchangeReason) String() string {
	switch r {
	case NoTilesExchangedReason:
		return "NoTilesExchanged"
	case InsufficientTilesInBagReason:
		return "InsufficientTilesInBag"
	default:
		return "Unknown"
	}
}
