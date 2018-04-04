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

	t.Run(".Find()", func(t *testing.T) {

		placements := TilePlacements{
			{Tile{'Z', 1}, 2, 5},
			{Tile{'A', 1}, 3, 7},
			{Tile{'B', 1}, 3, 2},
			{Tile{'C', 1}, 3, 7},
			{Tile{'D', 1}, 1, 7},
		}

		t.Run("returns the first matching placement", func(t *testing.T) {
			p := placements.Find(3, 7)

			if p == nil {
				t.Errorf("Expected to find a placement but got nil")
			} else {
				if actual, expected := p.Tile, (Tile{'A', 1}); actual != expected {
					t.Errorf("Expected to find tile %c(%d) but instead found %c(%d)", expected.Letter, expected.Points, actual.Letter, actual.Points)
				}
			}
		})

		t.Run("returns nil if none match", func(t *testing.T) {
			p := placements.Find(3, 99)

			if p != nil {
				t.Errorf("Expected not to find a placement but got one for tile %c(%d) at position %d,%d", p.Tile.Letter, p.Tile.Points, p.Row, p.Column)
			}
		})
	})
}
