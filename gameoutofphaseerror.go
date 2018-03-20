package scrubble

import "fmt"

// GameOutOfPhaseError indicates that an operation required a game to be in a
// specific phase when it was in a different phase.
type GameOutOfPhaseError struct {
	Required GamePhase
	Current  GamePhase
}

func (e GameOutOfPhaseError) Error() string {
	return fmt.Sprintf("%#v", e)
}
