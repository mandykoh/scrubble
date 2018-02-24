package scrubble

import "math/rand"

type Bag []Tile

type TileDistribution struct {
	Tile  Tile
	Count int
}

func BagFromDistribution(dist []TileDistribution) (bag Bag) {

	for _, d := range dist {
		for i := 0; i < d.Count; i++ {
			bag = append(bag, d.Tile)
		}
	}

	rand.Shuffle(len(bag), func(i, j int) {
		bag[i], bag[j] = bag[j], bag[i]
	})

	return
}
