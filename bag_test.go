package scrubble

import (
	"testing"
)

func TestBag(t *testing.T) {

	tileCount := func(t Tile, b Bag) (count int) {
		for _, tile := range b {
			if tile == t {
				count++
			}
		}

		return
	}

	t.Run("BagWithDistribution()", func(t *testing.T) {

		t.Run("creates bag with correct distribution of tiles", func(t *testing.T) {
			dist := []TileDistribution{
				{Tile{'A', 1}, 9},
				{Tile{'B', 3}, 2},
				{Tile{'C', 3}, 2},
				{Tile{'D', 2}, 4},
			}
			bag := BagWithDistribution(dist)

			if actual, expected := len(bag), 17; actual != expected {
				t.Fatalf("Expected bag of %d tiles but got %d", expected, actual)
			}

			for _, d := range dist {
				if actual, expected := tileCount(d.Tile, bag), d.Count; actual != expected {
					t.Errorf("Expected %d of tile %c(%d) but found %d", expected, d.Tile.Letter, d.Tile.Points, actual)
				}
			}
		})

		t.Run("creates bag with deterministic ordering", func(t *testing.T) {
			dist := []TileDistribution{
				{Tile{'A', 1}, 2},
				{Tile{'B', 3}, 2},
				{Tile{'C', 3}, 2},
				{Tile{'D', 2}, 2},
			}
			bag := BagWithDistribution(dist)

			tileNum := 0

			for _, d := range dist {
				for i := 0; i < d.Count; i++ {
					if actual, expected := bag[tileNum], d.Tile; actual != expected {
						t.Errorf("Expected tile %c(%d) in position %d but found %c(%d) instead", expected.Letter, expected.Points, tileNum, actual.Letter, actual.Points)
					}
					tileNum++
				}
			}
		})
	})

	t.Run("BagWithStandardEnglishTiles()", func(t *testing.T) {

		t.Run("creates a bag with correct distribution of tiles", func(t *testing.T) {
			expectedDist := []TileDistribution{
				{Tile{' ', 0}, 2},
				{Tile{'A', 1}, 9},
				{Tile{'B', 3}, 2},
				{Tile{'C', 3}, 2},
				{Tile{'D', 2}, 4},
				{Tile{'E', 1}, 12},
				{Tile{'F', 4}, 2},
				{Tile{'G', 2}, 3},
				{Tile{'H', 4}, 2},
				{Tile{'I', 1}, 9},
				{Tile{'J', 8}, 1},
				{Tile{'K', 5}, 1},
				{Tile{'L', 1}, 4},
				{Tile{'M', 3}, 2},
				{Tile{'N', 1}, 6},
				{Tile{'O', 1}, 8},
				{Tile{'P', 3}, 2},
				{Tile{'Q', 10}, 1},
				{Tile{'R', 1}, 6},
				{Tile{'S', 1}, 4},
				{Tile{'T', 1}, 6},
				{Tile{'U', 1}, 4},
				{Tile{'V', 4}, 2},
				{Tile{'W', 4}, 2},
				{Tile{'X', 8}, 1},
				{Tile{'Y', 4}, 2},
				{Tile{'Z', 10}, 1},
			}

			bag := BagWithStandardEnglishTiles()

			if actual, expected := len(bag), 100; actual != expected {
				t.Fatalf("Expected bag of %d tiles but got %d", expected, actual)
			}

			for _, d := range expectedDist {
				if actual, expected := tileCount(d.Tile, bag), d.Count; actual != expected {
					t.Errorf("Expected %d of tile %c(%d) but found %d", expected, d.Tile.Letter, d.Tile.Points, actual)
				}
			}
		})
	})
}
