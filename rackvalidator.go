package scrubble

// RackValidator represents a function which validates whether a rack holds the
// necessary tiles for making a play, and what the remaining rack would contain.
type RackValidator func(rack Rack, toPlay []Tile) (remaining []Tile, err error)
