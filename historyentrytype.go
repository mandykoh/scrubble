package scrubble

const (
	// UnknownHistoryEntryType indicates that the type of a history entry was indeterminate.
	UnknownHistoryEntryType HistoryEntryType = iota

	// PlayHistoryEntryType indicates that a history entry represents a play.
	PlayHistoryEntryType

	// PassHistoryEntryType indicates that a history entry represents a pass.
	PassHistoryEntryType

	// ExchangeTilesHistoryEntryType indicates that a history entry represents a tile exchange.
	ExchangeTilesHistoryEntryType

	// ChallengeFailHistoryEntryType indicates that a history entry represents an unsuccessful challenge.
	ChallengeFailHistoryEntryType

	// ChallengeSuccessHistoryEntryType indicates that a history entry represents a successful challenge.
	ChallengeSuccessHistoryEntryType
)

// HistoryEntryType represents a type of history entry.
type HistoryEntryType int

// GoString returns the Go syntax representation of the history entry type, or
// UnknownHistoryEntryType if it is not a valid type.
func (t HistoryEntryType) GoString() string {
	switch t {
	case PlayHistoryEntryType:
		return "PlayHistoryEntryType"
	case PassHistoryEntryType:
		return "PassHistoryEntryType"
	case ExchangeTilesHistoryEntryType:
		return "ExchangeTilesHistoryEntryType"
	case ChallengeFailHistoryEntryType:
		return "ChallengeFailHistoryEntryType"
	case ChallengeSuccessHistoryEntryType:
		return "ChallengeSuccessHistoryEntryType"
	default:
		return "UnknownHistoryEntryType"
	}
}

// String returns the textual representation of the history entry type, or
// "Unknown" if it is not a valid type.
func (t HistoryEntryType) String() string {
	switch t {
	case PlayHistoryEntryType:
		return "Play"
	case PassHistoryEntryType:
		return "Pass"
	case ExchangeTilesHistoryEntryType:
		return "ExchangeTiles"
	case ChallengeFailHistoryEntryType:
		return "ChallengeFail"
	case ChallengeSuccessHistoryEntryType:
		return "ChallengeSuccess"
	default:
		return "Unknown"
	}
}
