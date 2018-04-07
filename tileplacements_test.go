package scrubble

import "testing"

func TestTilePlacements(t *testing.T) {

	t.Run(".Bounds()", func(t *testing.T) {

		t.Run("returns the correct bounds of the placements", func(t *testing.T) {
			placements := TilePlacements{
				{Tile{'A', 1}, Coord{3, 0}},
				{Tile{'B', 1}, Coord{3, 1}},
				{Tile{'C', 1}, Coord{2, 7}},
				{Tile{'D', 1}, Coord{5, 6}},
			}

			bounds := placements.Bounds()

			if expected := 2; bounds.Min.Row != expected {
				t.Errorf("Expected minimum row to be %d but was %d", expected, bounds.Min.Row)
			}
			if expected := 0; bounds.Min.Column != expected {
				t.Errorf("Expected minimum column to be %d but was %d", expected, bounds.Min.Column)
			}
			if expected := 5; bounds.Max.Row != expected {
				t.Errorf("Expected maximum row to be %d but was %d", expected, bounds.Max.Row)
			}
			if expected := 7; bounds.Max.Column != expected {
				t.Errorf("Expected maximum column to be %d but was %d", expected, bounds.Max.Column)
			}
		})
	})

	t.Run(".Find()", func(t *testing.T) {

		placements := TilePlacements{
			{Tile{'Z', 1}, Coord{2, 5}},
			{Tile{'A', 1}, Coord{3, 7}},
			{Tile{'B', 1}, Coord{3, 2}},
			{Tile{'C', 1}, Coord{3, 7}},
			{Tile{'D', 1}, Coord{1, 7}},
		}

		t.Run("returns the first matching placement", func(t *testing.T) {
			p := placements.Find(Coord{3, 7})

			if p == nil {
				t.Errorf("Expected to find a placement but got nil")
			} else {
				if actual, expected := p.Tile, (Tile{'A', 1}); actual != expected {
					t.Errorf("Expected to find tile %c(%d) but instead found %c(%d)", expected.Letter, expected.Points, actual.Letter, actual.Points)
				}
			}
		})

		t.Run("returns nil if none match", func(t *testing.T) {
			p := placements.Find(Coord{3, 99})

			if p != nil {
				t.Errorf("Expected not to find a placement but got one for tile %c(%d) at position %d,%d", p.Tile.Letter, p.Tile.Points, p.Row, p.Column)
			}
		})
	})

	t.Run(".take()", func(t *testing.T) {

		t.Run("removes and returns the specified placement", func(t *testing.T) {
			placements := TilePlacements{
				{Tile{'A', 1}, Coord{3, 0}},
				{Tile{'B', 1}, Coord{3, 1}},
				{Tile{'C', 1}, Coord{2, 7}},
				{Tile{'D', 1}, Coord{5, 6}},
			}

			p := placements.take(Coord{2, 7})

			if p == nil {
				t.Errorf("Expected a valid placement but was nil")
			} else {
				if actual, expected := p.Tile.Letter, 'C'; actual != expected {
					t.Errorf("Expected to get placement for tile %c first but got %c", expected, actual)
				}
				if actual, expected := len(placements), 3; actual != expected {
					t.Errorf("Expected %d placements to remain but found %d", expected, actual)
				}
			}

			p = placements.take(Coord{5, 6})

			if p == nil {
				t.Errorf("Expected a valid placement but was nil")
			} else {
				if actual, expected := p.Tile.Letter, 'D'; actual != expected {
					t.Errorf("Expected to get placement for tile %c first but got %c", expected, actual)
				}
				if actual, expected := len(placements), 2; actual != expected {
					t.Errorf("Expected %d placements to remain but found %d", expected, actual)
				}
			}
		})

		t.Run("returns nil when placement is not found", func(t *testing.T) {
			placements := TilePlacements{
				{Tile{'A', 1}, Coord{3, 0}},
				{Tile{'B', 1}, Coord{3, 1}},
			}

			p := placements.take(Coord{3, 8})

			if p != nil {
				t.Errorf("Expected nil but still got a placement %v", p)
			}
		})
	})

	t.Run(".takeLast()", func(t *testing.T) {

		t.Run("removes and returns the last placement", func(t *testing.T) {
			placements := TilePlacements{
				{Tile{'A', 1}, Coord{3, 0}},
				{Tile{'B', 1}, Coord{3, 1}},
				{Tile{'C', 1}, Coord{2, 7}},
				{Tile{'D', 1}, Coord{5, 6}},
			}

			p := placements.takeLast()

			if p == nil {
				t.Errorf("Expected a valid placement but got nil")
			} else {
				if actual, expected := p.Tile.Letter, 'D'; actual != expected {
					t.Errorf("Expected to get placement for tile %c first but got %c", expected, actual)
				}
				if actual, expected := len(placements), 3; actual != expected {
					t.Errorf("Expected %d placements to remain but found %d", expected, actual)
				}
			}

			p = placements.takeLast()

			if p == nil {
				t.Errorf("Expected a valid placement but got nil")
			} else {
				if actual, expected := p.Tile.Letter, 'C'; actual != expected {
					t.Errorf("Expected to get placement for tile %c second but got %c", expected, actual)
				}
				if actual, expected := len(placements), 2; actual != expected {
					t.Errorf("Expected %d placements to remain but found %d", expected, actual)
				}
			}
		})

		t.Run("returns nil when no placements are left", func(t *testing.T) {
			placements := TilePlacements{}

			p := placements.takeLast()

			if p != nil {
				t.Errorf("Expected nil but still got a placement %v", p)
			}
		})
	})
}
