package challenge

const (
	// UnknownInvalidChallengeReason indicates that a reason was undefined.
	UnknownInvalidChallengeReason InvalidChallengeReason = iota

	// NoPlayToChallengeReason indicates that there is no play available to
	// challenge.
	NoPlayToChallengeReason

	// PlayAlreadyChallengedReason indicates that a play has already been
	// challenged.
	PlayAlreadyChallengedReason
)

// InvalidChallengeReason indicates the reason for an InvalidChallengeError.
type InvalidChallengeReason int

// GoString returns the Go syntax representation of the reason, or
// UnknownInvalidChallengeReason if it is not a valid reason.
func (r InvalidChallengeReason) GoString() string {
	switch r {
	case NoPlayToChallengeReason:
		return "NoPlayToChallengeReason"
	case PlayAlreadyChallengedReason:
		return "PlayAlreadyChallengedReason"
	default:
		return "UnknownInvalidChallengeReason"
	}
}

// String returns the textual representation of the reason, or "Unknown" if
// it is not a valid reason.
func (r InvalidChallengeReason) String() string {
	switch r {
	case NoPlayToChallengeReason:
		return "NoPlayToChallenge"
	case PlayAlreadyChallengedReason:
		return "PlayAlreadyChallenged"
	default:
		return "Unknown"
	}
}
