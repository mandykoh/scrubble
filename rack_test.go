package scrubble

import "testing"

func TestRack(t *testing.T) {

	expectRackContains := func(t *testing.T, r Rack, letters ...rune) {
		if actual, expected := len(r), len(letters); actual != expected {
			t.Fatalf("Expected rack to contain %d tiles but found %d", expected, actual)
		}

		for i, expected := range letters {
			if actual := r[i].Letter; actual != expected {
				t.Errorf("Expected letter '%c' on the rack but found '%c' instead", expected, actual)
			}
		}
	}

	t.Run(".FillFromBag()", func(t *testing.T) {

		t.Run("moves MaxRackTiles from the bag to an empty rack", func(t *testing.T) {
			var r Rack

			b := BagWithDistribution(TileDistribution{
				{Tile{'A', 1}, 3},
				{Tile{'B', 1}, 3},
				{Tile{'C', 1}, 3},
			})

			r.FillFromBag(&b)

			if actual, expected := len(r), MaxRackTiles; actual != expected {
				t.Errorf("Expected filled rack to contain %d tiles but found %d", expected, actual)
			}
			if actual, expected := len(b), 9-MaxRackTiles; actual != expected {
				t.Errorf("Expected bag to contain %d tiles but found %d", expected, actual)
			}

			expectRackContains(t, r, 'C', 'C', 'C', 'B', 'B', 'B', 'A')
		})

		t.Run("moves enough from the bag to a partially filled rack to reach MaxRackTiles", func(t *testing.T) {
			r := Rack{
				{'F', 1},
				{'G', 1},
				{'H', 1},
				{'I', 1},
			}

			b := BagWithDistribution(TileDistribution{
				{Tile{'A', 1}, 1},
				{Tile{'B', 1}, 1},
				{Tile{'C', 1}, 1},
				{Tile{'D', 1}, 1},
				{Tile{'E', 1}, 1},
			})

			r.FillFromBag(&b)

			if actual, expected := len(r), MaxRackTiles; actual != expected {
				t.Errorf("Expected filled rack to contain %d tiles but found %d", expected, actual)
			}
			if actual, expected := len(b), 5-(MaxRackTiles-4); actual != expected {
				t.Errorf("Expected bag to contain %d tiles but found %d", expected, actual)
			}

			expectRackContains(t, r, 'F', 'G', 'H', 'I', 'E', 'D', 'C')
		})

		t.Run("moves all tiles from the bag when not enough to reach MaxRackTiles", func(t *testing.T) {
			var r Rack

			b := BagWithDistribution(TileDistribution{
				{Tile{'A', 1}, 1},
				{Tile{'B', 1}, 1},
				{Tile{'C', 1}, 1},
			})

			r.FillFromBag(&b)

			if actual, expected := len(r), 3; actual != expected {
				t.Errorf("Expected filled rack to contain %d tiles but found %d", expected, actual)
			}
			if actual, expected := len(b), 0; actual != expected {
				t.Errorf("Expected bag to be empty but found %d tiles", actual)
			}
		})
	})

	t.Run(".tryPlayTiles()", func(t *testing.T) {

		t.Run("returns missing tiles when the rack has insufficient tiles for the play", func(t *testing.T) {
			r := Rack{
				{'A', 1},
				{'B', 1},
				{'O', 1},
				{'M', 1},
			}

			remaining, missing := r.tryPlayTiles(TilePlacements{
				{Tile{'B', 1}, Coord{0, 0}},
				{Tile{'O', 1}, Coord{0, 1}},
				{Tile{'O', 1}, Coord{0, 2}},
				{Tile{'M', 1}, Coord{0, 3}},
				{Tile{'S', 1}, Coord{0, 3}},
			})

			if actual, expected := len(missing), 2; actual != expected {
				t.Errorf("Expected %d missing tile but found %d", expected, actual)
			} else {
				if actual, expected := missing[0], (Tile{'O', 1}); actual != expected {
					t.Errorf("Expected missing tile to be %c(%d) but was %c(%d)", expected.Letter, expected.Points, actual.Letter, actual.Points)
				}
				if actual, expected := missing[1], (Tile{'S', 1}); actual != expected {
					t.Errorf("Expected missing tile to be %c(%d) but was %c(%d)", expected.Letter, expected.Points, actual.Letter, actual.Points)
				}
			}

			expectRackContains(t, remaining, 'A')
			expectRackContains(t, r, 'A', 'B', 'O', 'M')
		})

		t.Run("returns no missing tiles and the remainder if successful", func(t *testing.T) {
			r := Rack{
				{'A', 1},
				{'O', 1},
				{'M', 1},
				{'B', 1},
				{'O', 1},
			}

			remaining, missing := r.tryPlayTiles(TilePlacements{
				{Tile{'B', 1}, Coord{0, 0}},
				{Tile{'O', 1}, Coord{0, 1}},
				{Tile{'O', 1}, Coord{0, 2}},
				{Tile{'M', 1}, Coord{0, 3}},
			})

			if actual, expected := len(missing), 0; actual != expected {
				t.Errorf("Expected no missing tiles but found %d", missing)
			}

			expectRackContains(t, remaining, 'A')
			expectRackContains(t, r, 'A', 'O', 'M', 'B', 'O')
		})
	})
}
