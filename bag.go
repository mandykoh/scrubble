package scrubble

// Bag represents a bag of Tiles to be used during a game.
type Bag []Tile

// TileDistribution is a count of how many of a particular Tile should appear.
// This is used to create a Bag containing a particular distribution of tiles.
type TileDistribution struct {
	Tile  Tile
	Count int
}

// BagWithDistribution returns a Bag containing tiles according to the specified
// distribution.
func BagWithDistribution(dist []TileDistribution) (bag Bag) {

	for _, d := range dist {
		for i := 0; i < d.Count; i++ {
			bag = append(bag, d.Tile)
		}
	}

	return
}
