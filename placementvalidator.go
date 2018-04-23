package scrubble

import "github.com/mandykoh/scrubble/play"

// PlacementValidator represents a function which validates whether the
// positions of tiles being placed on a board are valid.
type PlacementValidator func(placements play.Tiles, board *Board) error
