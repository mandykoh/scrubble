package scrubble

import "testing"

func TestBag(t *testing.T) {

	tileCount := func(t Tile, b Bag) (count int) {
		for _, tile := range b {
			if tile == t {
				count++
			}
		}

		return
	}

	t.Run("BagFromDistribution()", func(t *testing.T) {

		t.Run("creates bag with correct distribution of tiles", func(t *testing.T) {
			dist := []TileDistribution{
				{Tile{'A', 1}, 9},
				{Tile{'B', 3}, 2},
				{Tile{'C', 3}, 2},
				{Tile{'D', 2}, 4},
			}
			bag := BagFromDistribution(dist)

			if actual, expected := len(bag), 17; actual != expected {
				t.Fatalf("Expected bag of %d tiles but got %d", expected, actual)
			}

			for _, d := range dist {
				if actual, expected := tileCount(d.Tile, bag), d.Count; actual != expected {
					t.Errorf("Expected %d of tile %c(%d) but found %d", expected, d.Tile.Letter, d.Tile.Points, actual)
				}
			}
		})

		t.Run("creates bag with tiles shuffled", func(t *testing.T) {
			dist := []TileDistribution{
				{Tile{'A', 1}, 100},
				{Tile{'B', 3}, 100},
			}
			bag := BagFromDistribution(dist)

			consecutive := 0
			for _, t := range bag {
				if t != dist[0].Tile {
					break
				}
				consecutive++
			}

			if consecutive == dist[0].Count {
				t.Errorf("Got %d consecutive %c(%d) tiles but expected tiles to be shuffled", consecutive, dist[0].Tile.Letter, dist[0].Tile.Points)
			}
		})
	})
}
