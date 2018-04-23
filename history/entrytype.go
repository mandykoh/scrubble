package history

const (
	// UnknownEntryType indicates that the type of a history entry was indeterminate.
	UnknownEntryType EntryType = iota

	// PlayEntryType indicates that a history entry represents a play.
	PlayEntryType

	// PassEntryType indicates that a history entry represents a pass.
	PassEntryType

	// ExchangeTilesEntryType indicates that a history entry represents a tile exchange.
	ExchangeTilesEntryType

	// ChallengeFailEntryType indicates that a history entry represents an unsuccessful challenge.
	ChallengeFailEntryType

	// ChallengeSuccessEntryType indicates that a history entry represents a successful challenge.
	ChallengeSuccessEntryType
)

// EntryType represents a type of history entry.
type EntryType int

// GoString returns the Go syntax representation of the history entry type, or
// UnknownEntryType if it is not a valid type.
func (t EntryType) GoString() string {
	switch t {
	case PlayEntryType:
		return "PlayEntryType"
	case PassEntryType:
		return "PassEntryType"
	case ExchangeTilesEntryType:
		return "ExchangeTilesEntryType"
	case ChallengeFailEntryType:
		return "ChallengeFailEntryType"
	case ChallengeSuccessEntryType:
		return "ChallengeSuccessEntryType"
	default:
		return "UnknownEntryType"
	}
}

// String returns the textual representation of the history entry type, or
// "Unknown" if it is not a valid type.
func (t EntryType) String() string {
	switch t {
	case PlayEntryType:
		return "Play"
	case PassEntryType:
		return "Pass"
	case ExchangeTilesEntryType:
		return "ExchangeTiles"
	case ChallengeFailEntryType:
		return "ChallengeFail"
	case ChallengeSuccessEntryType:
		return "ChallengeSuccess"
	default:
		return "Unknown"
	}
}
