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

// BagWithStandardEnglishTiles returns a Bag containing tiles corresponding to
// a standard English tile and letter distribution.
func BagWithStandardEnglishTiles() Bag {
	return BagWithDistribution([]TileDistribution{
		{Tile{' ', 0}, 2},
		{Tile{'E', 1}, 12},
		{Tile{'A', 1}, 9},
		{Tile{'I', 1}, 9},
		{Tile{'O', 1}, 8},
		{Tile{'N', 1}, 6},
		{Tile{'R', 1}, 6},
		{Tile{'T', 1}, 6},
		{Tile{'L', 1}, 4},
		{Tile{'S', 1}, 4},
		{Tile{'U', 1}, 4},
		{Tile{'D', 2}, 4},
		{Tile{'G', 2}, 3},
		{Tile{'B', 3}, 2},
		{Tile{'C', 3}, 2},
		{Tile{'M', 3}, 2},
		{Tile{'P', 3}, 2},
		{Tile{'F', 4}, 2},
		{Tile{'H', 4}, 2},
		{Tile{'V', 4}, 2},
		{Tile{'W', 4}, 2},
		{Tile{'Y', 4}, 2},
		{Tile{'K', 5}, 1},
		{Tile{'J', 8}, 1},
		{Tile{'X', 8}, 1},
		{Tile{'Q', 10}, 1},
		{Tile{'Z', 10}, 1},
	})
}
