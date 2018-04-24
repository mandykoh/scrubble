package game

import "fmt"

// OutOfPhaseError indicates that an operation required a game to be in a
// specific phase when it was in a different phase.
type OutOfPhaseError struct {
	Required Phase
	Current  Phase
}

func (e OutOfPhaseError) Error() string {
	return fmt.Sprintf("%#v", e)
}
