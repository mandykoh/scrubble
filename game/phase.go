package game

// Phase represents the phase that a Game is in.
type Phase int

const (
	SetupPhase Phase = iota
	MainPhase
	EndPhase
	UnknownPhase
)

// GoString returns the Go syntax representation of the game phase, or
// UnknownPhase if it is not a valid phase.
func (p Phase) GoString() string {
	switch p {
	case SetupPhase:
		return "SetupPhase"
	case MainPhase:
		return "MainPhase"
	case EndPhase:
		return "EndPhase"
	default:
		return "UnknownPhase"
	}
}

// String returns the textual representation of the game phase, or "Unknown" if
// it is not a valid phase.
func (p Phase) String() string {
	switch p {
	case SetupPhase:
		return "Setup"
	case MainPhase:
		return "Main"
	case EndPhase:
		return "End"
	default:
		return "Unknown"
	}
}
