package scrubble

import "testing"

func TestNextGamePhase(t *testing.T) {
	game := &Game{}

	t.Run("allows the game to continue when the rack can be replenished with at least one tile", func(t *testing.T) {
		seat := &Seat{
			Rack: []Tile{
				{'A', 1},
			},
		}

		next := NextGamePhase(seat, 123, game)

		if next != MainPhase {
			t.Errorf("Expected that the game should continue but got %#v", next)
		}
	})

	t.Run("ends the game when the rack is emptied and can't be replenished from the bag", func(t *testing.T) {
		seat := &Seat{}

		next := NextGamePhase(seat, 123, game)

		if next != EndPhase {
			t.Errorf("Expected that the game should end but got %#v", next)
		}
	})
}
