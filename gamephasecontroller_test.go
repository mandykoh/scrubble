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
			History: History{{0, 123, nil, nil, nil}},
		}

		next := NextGamePhase(game)

		if next != MainPhase {
			t.Errorf("Expected that the game should continue but got %#v next", next)
		}
	})

	t.Run("ends the game when the rack was emptied and couldn't be replenished from the bag", func(t *testing.T) {
		game := &Game{
			Seats: []Seat{
				{Rack: []Tile{}},
			},
			History: History{{0, 123, nil, nil, nil}},
		}

		next := NextGamePhase(game)

		if next != EndPhase {
			t.Errorf("Expected that the game should end but got %#v next", next)
		}
	})

	t.Run("ends the game after six consecutive scoreless turns", func(t *testing.T) {
		game := &Game{
			Seats: []Seat{
				{Rack: []Tile{{'A', 1}}},
				{Rack: []Tile{{'B', 2}}},
			},
			History: History{
				{0, 0, nil, nil, nil},
				{1, 0, nil, nil, nil},
				{0, 0, nil, nil, nil},
				{1, 0, nil, nil, nil},
				{0, 0, nil, nil, nil},
			},
		}

		next := NextGamePhase(game)

		if next == EndPhase {
			t.Errorf("Expected that the game should still continue for one turn but got %#v next", next)
		}

		game.History.AppendPlay(1, 0, nil, nil, nil)
		next = NextGamePhase(game)

		if next != EndPhase {
			t.Errorf("Expected that the game should end but got %#v next", next)
		}
	})
}
