package scrubble

// RackValidator represents a function which validates whether a rack holds the
// necessary tiles for making a play, which tiles would be used from the rack,
// and what the remaining rack would contain.
type RackValidator func(rack Rack, toPlay []Tile) (used []Tile, remaining []Tile, err error)
