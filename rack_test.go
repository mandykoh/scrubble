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

			drawn := r.FillFromBag(&b)

			if actual, expected := len(r), MaxRackTiles; actual != expected {
				t.Errorf("Expected filled rack to contain %d tiles but found %d", expected, actual)
			}
			if actual, expected := len(b), 9-MaxRackTiles; actual != expected {
				t.Errorf("Expected bag to contain %d tiles but found %d", expected, actual)
			}

			expectRackContains(t, r, 'C', 'C', 'C', 'B', 'B', 'B', 'A')
			expectRackContains(t, drawn, 'C', 'C', 'C', 'B', 'B', 'B', 'A')
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

			drawn := r.FillFromBag(&b)

			if actual, expected := len(r), MaxRackTiles; actual != expected {
				t.Errorf("Expected filled rack to contain %d tiles but found %d", expected, actual)
			}
			if actual, expected := len(b), 5-(MaxRackTiles-4); actual != expected {
				t.Errorf("Expected bag to contain %d tiles but found %d", expected, actual)
			}

			expectRackContains(t, r, 'F', 'G', 'H', 'I', 'E', 'D', 'C')
			expectRackContains(t, drawn, 'E', 'D', 'C')
		})

		t.Run("moves all tiles from the bag when not enough to reach MaxRackTiles", func(t *testing.T) {
			var r Rack

			b := BagWithDistribution(TileDistribution{
				{Tile{'A', 1}, 1},
				{Tile{'B', 1}, 1},
				{Tile{'C', 1}, 1},
			})

			drawn := r.FillFromBag(&b)

			if actual, expected := len(r), 3; actual != expected {
				t.Errorf("Expected filled rack to contain %d tiles but found %d", expected, actual)
			}
			if actual, expected := len(b), 0; actual != expected {
				t.Errorf("Expected bag to be empty but found %d tiles", actual)
			}

			expectRackContains(t, r, 'C', 'B', 'A')
			expectRackContains(t, drawn, 'C', 'B', 'A')
		})
	})

	t.Run(".Remove()", func(t *testing.T) {

		t.Run("removes the specified tiles", func(t *testing.T) {
			r := Rack{
				{'F', 1},
				{'G', 1},
				{'H', 1},
				{'I', 1},
			}

			r.Remove(Tile{'F', 1}, Tile{'H', 1})

			expectRackContains(t, r, 'G', 'I')
		})

		t.Run("ignores nonexistent tiles", func(t *testing.T) {
			r := Rack{
				{'F', 1},
				{'G', 1},
				{'H', 1},
			}

			r.Remove(Tile{'G', 1}, Tile{'X', 1})

			expectRackContains(t, r, 'F', 'H')
		})
	})
}
