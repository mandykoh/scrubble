package scrubble

// GamePhase represents the phase that a Game is in.
type GamePhase int

const (
	SetupPhase GamePhase = iota
	MainPhase
	EndPhase
)

// String returns the textual representation of the game phase, or "Unknown" if
// it is not a valid phase.
func (p GamePhase) String() string {
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
