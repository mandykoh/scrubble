package scrubble

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ExampleBagWithDistribution() {

	// Creates a Bag with 9 x A tiles, 2 x B tiles, 2 x C tiles, and 4 x D tiles
	bag := BagWithDistribution(TileDistribution{
		{Tile{'A', 1}, 9},
		{Tile{'B', 3}, 2},
		{Tile{'C', 3}, 2},
		{Tile{'D', 2}, 4},
	})

	// Output: Number of tiles in bag: 17
	fmt.Println("Number of tiles in bag:", len(bag))
}

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
			dist := TileDistribution{
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
					t.Errorf("Expected %d of tile %v but found %d", expected, d.Tile, actual)
				}
			}
		})

		t.Run("creates bag with deterministic ordering", func(t *testing.T) {
			dist := TileDistribution{
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
						t.Errorf("Expected tile %v in position %d but found %v instead", expected, tileNum, actual)
					}
					tileNum++
				}
			}
		})

		t.Run("allocates exact capacity for requested tiles", func(t *testing.T) {
			dist := TileDistribution{
				{Tile{'A', 1}, 3},
				{Tile{'B', 2}, 3},
				{Tile{'C', 3}, 3},
			}
			bag := BagWithDistribution(dist)

			if actual, expected := cap(bag), 9; actual != expected {
				t.Errorf("Expected bag of exactly %d capacity but was %d instead", expected, actual)
			}
		})
	})

	t.Run("BagWithStandardEnglishTiles()", func(t *testing.T) {

		t.Run("creates a bag with correct distribution of tiles", func(t *testing.T) {
			expectedDist := TileDistribution{
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
					t.Errorf("Expected %d of tile %v but found %d", expected, d.Tile, actual)
				}
			}
		})
	})

	t.Run(".DrawTile()", func(t *testing.T) {

		t.Run("removes and returns tiles in last to first order", func(t *testing.T) {
			dist := TileDistribution{
				{Tile{'A', 1}, 1},
				{Tile{'B', 2}, 1},
				{Tile{'C', 3}, 1},
			}
			bag := BagWithDistribution(dist)

			if actual, expected := bag.DrawTile(), dist[2].Tile; actual != expected {
				t.Errorf("Expected first tile drawn to be %v but got %v", expected, actual)
			}
			if actual, expected := len(bag), 2; actual != expected {
				t.Errorf("Expected %d tiles remaining in bag but found %d", expected, actual)
			}
			if actual, expected := bag.DrawTile(), dist[1].Tile; actual != expected {
				t.Errorf("Expected second tile drawn to be %v but got %v", expected, actual)
			}
			if actual, expected := len(bag), 1; actual != expected {
				t.Errorf("Expected %d tiles remaining in bag but found %d", expected, actual)
			}
			if actual, expected := bag.DrawTile(), dist[0].Tile; actual != expected {
				t.Errorf("Expected third tile drawn to be %v but got %v", expected, actual)
			}
			if actual, expected := len(bag), 0; actual != expected {
				t.Errorf("Expected no tiles remaining in bag but found %d", actual)
			}
		})

		t.Run("panics when no tiles are left, leaving the bag unchanged", func(t *testing.T) {
			var bag Bag

			defer func() {
				if recovered := recover(); recovered == nil {
					t.Errorf("Expected a panic but nothing happened")
				}
			}()

			bag.DrawTile()
		})
	})

	t.Run(".Shuffle()", func(t *testing.T) {

		t.Run("randomises the order of the tiles using the specified random generator", func(t *testing.T) {

			tiles := []Tile{
				{'A', 1},
				{'B', 2},
				{'C', 3},
				{'D', 4},
			}

			bag := make(Bag, len(tiles))
			copy(bag, tiles)

			seed := time.Now().UnixNano()

			r1 := rand.New(rand.NewSource(seed))
			bag.Shuffle(r1)

			r2 := rand.New(rand.NewSource(seed))
			r2.Shuffle(len(tiles), func(i, j int) {
				tiles[i], tiles[j] = tiles[j], tiles[i]
			})

			expectTiles(t, "shuffled", bag, tiles...)
		})
	})
}
