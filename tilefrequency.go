package scrubble

// TileFrequency is a count of how many of a particular Tile should appear.
// This is used to specify how frequently a tile should appear in a bag.
type TileFrequency struct {
	Tile  Tile
	Count int
}
