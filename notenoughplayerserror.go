package scrubble

import "fmt"

// NotEnoughPlayersError indicates that a game doesn't have enough players to
// be started.
type NotEnoughPlayersError struct {
	Required int
	Current  int
}

func (e NotEnoughPlayersError) Error() string {
	return fmt.Sprintf("%#v", e)
}
