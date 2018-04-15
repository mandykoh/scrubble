package scrubble

import "testing"

func TestNextGamePhase(t *testing.T) {

	t.Run("allows the game to continue when the rack could be replenished with at least one tile", func(t *testing.T) {
		game := &Game{
			Seats: []Seat{
				{
					Rack: []Tile{
						{'A', 1},
					},
				},
			},
			History: History{{0, 123, TilePlacements{}, []PlayedWord{}}},
		}

		next := NextGamePhase(game)

		if next != MainPhase {
			t.Errorf("Expected that the game should continue but got %#v", next)
		}
	})

	t.Run("ends the game when the rack was emptied and couldn't be replenished from the bag", func(t *testing.T) {
		game := &Game{
			Seats: []Seat{
				{Rack: []Tile{}},
			},
			History: History{{0, 123, TilePlacements{}, []PlayedWord{}}},
		}

		next := NextGamePhase(game)

		if next != EndPhase {
			t.Errorf("Expected that the game should end but got %#v", next)
		}
	})
}
