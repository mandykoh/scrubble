package scrubble

import "github.com/mandykoh/scrubble/tile"

// RackValidator represents a function which validates whether a rack holds the
// necessary tiles for making a play, which tiles would be used from the rack,
// and what the remaining rack would contain.
type RackValidator func(rack tile.Rack, toPlay []tile.Tile) (used, remaining []tile.Tile, err error)
