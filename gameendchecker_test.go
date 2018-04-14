package scrubble

import "testing"

func TestGameEnds(t *testing.T) {
	game := &Game{}

	t.Run("allows the game to continue when the rack can be replenished with at least one tile", func(t *testing.T) {
		seat := &Seat{
			Rack: []Tile{
				{'A', 1},
			},
		}

		shouldEnd := GameEnds(seat, 123, game)

		if shouldEnd {
			t.Errorf("Expected that the game should continue but it would end")
		}
	})

	t.Run("ends the game when the rack is emptied and can't be replenished from the bag", func(t *testing.T) {
		seat := &Seat{}

		shouldEnd := GameEnds(seat, 123, game)

		if !shouldEnd {
			t.Errorf("Expected that the game should end but it would continue")
		}
	})
}
