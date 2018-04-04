package scrubble

import "testing"

func TestTilePlacements(t *testing.T) {

	t.Run(".Bounds()", func(t *testing.T) {

		t.Run("returns the correct bounds of the placements", func(t *testing.T) {
			placements := TilePlacements{
				{Tile{'A', 1}, 3, 0},
				{Tile{'B', 1}, 3, 1},
				{Tile{'C', 1}, 2, 7},
				{Tile{'D', 1}, 5, 6},
			}

			minRow, minCol, maxRow, maxCol := placements.Bounds()

			if expected := 2; minRow != expected {
				t.Errorf("Expected minimum row to be %d but was %d", expected, minRow)
			}
			if expected := 0; minCol != expected {
				t.Errorf("Expected minimum column to be %d but was %d", expected, minCol)
			}
			if expected := 5; maxRow != expected {
				t.Errorf("Expected maximum row to be %d but was %d", expected, maxRow)
			}
			if expected := 7; maxCol != expected {
				t.Errorf("Expected maximum column to be %d but was %d", expected, maxCol)
			}
		})
	})
}
