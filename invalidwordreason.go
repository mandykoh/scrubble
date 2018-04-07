package scrubble

const (
	// UnknownInvalidWordReason indicates that a reason was undefined.
	UnknownInvalidWordReason InvalidWordReason = iota

	// SingleLetterWordDisallowedReason indicates that a play attempted to form
	// a single-letter word, which is not allowed.
	SingleLetterWordDisallowedReason
)

// InvalidWordReason indicates the reason for an InvalidWordError.
type InvalidWordReason int

// GoString returns the Go syntax representation of the reason, or
// UnknownInvalidWordReason if it is not a valid reason.
func (r InvalidWordReason) GoString() string {
	switch r {
	case SingleLetterWordDisallowedReason:
		return "SingleLetterWordDisallowedReason"
	default:
		return "UnknownInvalidWordReason"
	}
}

// String returns the textual representation of the reason, or "Unknown" if
// it is not a valid reason.
func (r InvalidWordReason) String() string {
	switch r {
	case SingleLetterWordDisallowedReason:
		return "SingleLetterWordDisallowed"
	default:
		return "Unknown"
	}
}
