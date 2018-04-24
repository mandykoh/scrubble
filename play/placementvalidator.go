package play

import (
	"github.com/mandykoh/scrubble/board"
)

// PlacementValidator represents a function which validates whether the
// positions of tiles being placed on a board are valid.
type PlacementValidator func(placements Tiles, board *board.Board) error
