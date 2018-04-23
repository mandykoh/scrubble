package tile

import "testing"

func TestValidateTilesFromRack(t *testing.T) {

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

	t.Run("returns missing tiles when the rack has insufficient tiles for the play", func(t *testing.T) {
		r := Rack{
			{'A', 1},
			{'B', 1},
			{'O', 1},
			{'M', 1},
		}

		used, remaining, err := ValidateFromRack(r, []Tile{
			{'B', 1},
			{'O', 1},
			{'O', 1},
			{'M', 1},
			{'S', 1},
		})

		switch e := err.(type) {

		case InsufficientTilesError:
			expectTiles(t, "missing", e.Missing,
				Tile{'O', 1},
				Tile{'S', 1},
			)

		default:
			t.Errorf("Expected an InsufficientTilesError but got %v", err)
		}

		expectRackContains(t, used, 'B', 'O', 'M')
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

		used, remaining, err := ValidateFromRack(r, []Tile{
			{'B', 1},
			{'O', 1},
			{'O', 1},
			{'M', 1},
		})

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			expectRackContains(t, used, 'B', 'O', 'O', 'M')
			expectRackContains(t, remaining, 'A')
			expectRackContains(t, r, 'A', 'O', 'M', 'B', 'O')
		}
	})

	t.Run("treats zero-point tiles as wildcard tiles", func(t *testing.T) {
		r := Rack{
			{'A', 1},
			{' ', 0},
			{'M', 1},
			{'B', 1},
			{'O', 1},
		}

		used, remaining, err := ValidateFromRack(r, []Tile{
			{'B', 1},
			{'O', 1},
			{'O', 0},
			{'M', 1},
		})

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			expectRackContains(t, used, 'B', 'O', ' ', 'M')
			expectRackContains(t, remaining, 'A')
		}
	})
}
