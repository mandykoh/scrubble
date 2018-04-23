package scrubble

import (
	"testing"

	"github.com/mandykoh/scrubble/tile"
)

func expectTilePlacements(t *testing.T, placements TilePlacements, expected ...TilePlacement) {
	t.Helper()

	if actual, expectedLen := len(placements), len(expected); actual != expectedLen {
		t.Errorf("Expected there to be %d tiles placed but found %d", expectedLen, actual)

	} else {
		for i, e := range expected {
			if placements[i] != e {
				t.Errorf("Expected tile placement %v in position %d but found %v instead", e, i, placements[i])
			}
		}
	}
}

func TestTilePlacements(t *testing.T) {

	t.Run(".Bounds()", func(t *testing.T) {

		t.Run("returns the correct bounds of the placements", func(t *testing.T) {
			placements := TilePlacements{
				{tile.Tile{Letter: 'A', Points: 1}, Coord{3, 0}},
				{tile.Tile{Letter: 'B', Points: 1}, Coord{3, 1}},
				{tile.Tile{Letter: 'C', Points: 1}, Coord{2, 7}},
				{tile.Tile{Letter: 'D', Points: 1}, Coord{5, 6}},
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
			{tile.Tile{Letter: 'Z', Points: 1}, Coord{2, 5}},
			{tile.Tile{Letter: 'A', Points: 1}, Coord{3, 7}},
			{tile.Tile{Letter: 'B', Points: 1}, Coord{3, 2}},
			{tile.Tile{Letter: 'C', Points: 1}, Coord{3, 7}},
			{tile.Tile{Letter: 'D', Points: 1}, Coord{1, 7}},
		}

		t.Run("returns the first matching placement", func(t *testing.T) {
			p := placements.Find(Coord{3, 7})

			if p == nil {
				t.Errorf("Expected to find a placement but got nil")
			} else {
				if actual, expected := p.Tile, (tile.Tile{Letter: 'A', Points: 1}); actual != expected {
					t.Errorf("Expected to find tile %v but instead found %v", expected, actual)
				}
			}
		})

		t.Run("returns nil if none match", func(t *testing.T) {
			p := placements.Find(Coord{3, 99})

			if p != nil {
				t.Errorf("Expected not to find a placement but got one for tile %v at position %d,%d", p.Tile, p.Row, p.Column)
			}
		})
	})

	t.Run(".Tiles()", func(t *testing.T) {

		placements := TilePlacements{
			{tile.Tile{Letter: 'Z', Points: 1}, Coord{2, 5}},
			{tile.Tile{Letter: 'A', Points: 1}, Coord{3, 7}},
			{tile.Tile{Letter: 'B', Points: 1}, Coord{3, 2}},
			{tile.Tile{Letter: 'C', Points: 1}, Coord{3, 7}},
			{tile.Tile{Letter: 'D', Points: 1}, Coord{1, 7}},
		}

		t.Run("returns just the tiles being placed", func(t *testing.T) {
			tiles := placements.Tiles()

			if actual, expected := len(tiles), len(placements); actual != expected {
				t.Errorf("Expected %d tiles but got %d", expected, actual)
			} else {
				for i := 0; i < len(tiles); i++ {
					if actual, expected := tiles[i], placements[i].Tile; actual != expected {
						t.Errorf("Expected to find tile %v but instead found %v", expected, actual)
					}
				}
			}
		})

		t.Run("returns nil if none match", func(t *testing.T) {
			p := placements.Find(Coord{3, 99})

			if p != nil {
				t.Errorf("Expected not to find a placement but got one for tile %v at position %d,%d", p.Tile, p.Row, p.Column)
			}
		})
	})

	t.Run(".take()", func(t *testing.T) {

		t.Run("removes and returns the specified placement", func(t *testing.T) {
			placements := TilePlacements{
				{tile.Tile{Letter: 'A', Points: 1}, Coord{3, 0}},
				{tile.Tile{Letter: 'B', Points: 1}, Coord{3, 1}},
				{tile.Tile{Letter: 'C', Points: 1}, Coord{2, 7}},
				{tile.Tile{Letter: 'D', Points: 1}, Coord{5, 6}},
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
				{tile.Tile{Letter: 'A', Points: 1}, Coord{3, 0}},
				{tile.Tile{Letter: 'B', Points: 1}, Coord{3, 1}},
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
				{tile.Tile{Letter: 'A', Points: 1}, Coord{3, 0}},
				{tile.Tile{Letter: 'B', Points: 1}, Coord{3, 1}},
				{tile.Tile{Letter: 'C', Points: 1}, Coord{2, 7}},
				{tile.Tile{Letter: 'D', Points: 1}, Coord{5, 6}},
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
