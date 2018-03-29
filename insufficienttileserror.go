package scrubble

import "fmt"

// InsufficientTilesError indicates that a play called for tiles which the
// player doesn't possess in their rack.
type InsufficientTilesError struct {
	Missing []Tile
}

func (e InsufficientTilesError) Error() string {
	return fmt.Sprintf("%#v", e)
}
