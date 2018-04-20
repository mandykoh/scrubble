package scrubble

import "testing"

func TestRack(t *testing.T) {

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

			expectTiles(t, "racked", r,
				Tile{'C', 1},
				Tile{'C', 1},
				Tile{'C', 1},
				Tile{'B', 1},
				Tile{'B', 1},
				Tile{'B', 1},
				Tile{'A', 1})

			expectTiles(t, "drawn", drawn,
				Tile{'C', 1},
				Tile{'C', 1},
				Tile{'C', 1},
				Tile{'B', 1},
				Tile{'B', 1},
				Tile{'B', 1},
				Tile{'A', 1})
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

			expectTiles(t, "racked", r,
				Tile{'F', 1},
				Tile{'G', 1},
				Tile{'H', 1},
				Tile{'I', 1},
				Tile{'E', 1},
				Tile{'D', 1},
				Tile{'C', 1})

			expectTiles(t, "drawn", drawn,
				Tile{'E', 1},
				Tile{'D', 1},
				Tile{'C', 1})
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

			expectTiles(t, "racked", r,
				Tile{'C', 1},
				Tile{'B', 1},
				Tile{'A', 1})

			expectTiles(t, "drawn", drawn,
				Tile{'C', 1},
				Tile{'B', 1},
				Tile{'A', 1})
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

			expectTiles(t, "racked", r,
				Tile{'G', 1},
				Tile{'I', 1},
			)
		})

		t.Run("ignores nonexistent tiles", func(t *testing.T) {
			r := Rack{
				{'F', 1},
				{'G', 1},
				{'H', 1},
			}

			r.Remove(Tile{'G', 1}, Tile{'X', 1})

			expectTiles(t, "racked", r,
				Tile{'F', 1},
				Tile{'H', 1},
			)
		})
	})
}
