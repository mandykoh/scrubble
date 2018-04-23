package scrubble

import (
	"fmt"

	"github.com/mandykoh/scrubble/tile"
)

// InsufficientTilesError indicates that a play called for tiles which the
// player doesn't possess in their rack.
type InsufficientTilesError struct {
	Missing []tile.Tile
}

func (e InsufficientTilesError) Error() string {
	return fmt.Sprintf("%#v", e)
}
