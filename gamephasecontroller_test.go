package scrubble

import (
	"testing"

	"github.com/mandykoh/scrubble/tile"
)

func TestNextGamePhase(t *testing.T) {

	t.Run("allows the game to continue when the rack could be replenished with at least one tile", func(t *testing.T) {
		game := &Game{
			Seats: []Seat{
				{
					Rack: tile.Rack{
						{'A', 1},
					},
				},
			},
			History: History{{Type: UnknownHistoryEntryType, Score: 123}},
		}

		next := NextGamePhase(game)

		if next != MainPhase {
			t.Errorf("Expected that the game should continue but got %#v next", next)
		}
	})

	t.Run("ends the game when the rack was emptied and couldn't be replenished from the bag", func(t *testing.T) {
		game := &Game{
			Seats: []Seat{
				{Rack: tile.Rack{}},
			},
			History: History{{Type: UnknownHistoryEntryType, Score: 123}},
		}

		next := NextGamePhase(game)

		if next != EndPhase {
			t.Errorf("Expected that the game should end but got %#v next", next)
		}
	})

	t.Run("ends the game after six consecutive scoreless turns", func(t *testing.T) {
		game := &Game{
			Seats: []Seat{
				{Rack: tile.Rack{{'A', 1}}},
				{Rack: tile.Rack{{'B', 2}}},
			},
			History: History{
				{Type: PassHistoryEntryType},
				{Type: PassHistoryEntryType, SeatIndex: 1},
				{Type: PassHistoryEntryType},
				{Type: PassHistoryEntryType, SeatIndex: 1},
				{Type: PassHistoryEntryType},
			},
		}

		next := NextGamePhase(game)

		if next == EndPhase {
			t.Errorf("Expected that the game should still continue for one turn but got %#v next", next)
		}

		game.History.AppendPass(1)
		next = NextGamePhase(game)

		if next != EndPhase {
			t.Errorf("Expected that the game should end but got %#v next", next)
		}
	})

	t.Run("treats successfully challenged plays as scoreless turns", func(t *testing.T) {
		game := &Game{
			Seats: []Seat{
				{Rack: tile.Rack{{'A', 1}}},
				{Rack: tile.Rack{{'B', 2}}},
			},
			History: History{
				{Type: PlayHistoryEntryType, Score: 123},
				{Type: ChallengeSuccessHistoryEntryType, SeatIndex: 1},
				{Type: PlayHistoryEntryType, SeatIndex: 1, Score: 234},
				{Type: ChallengeSuccessHistoryEntryType},
				{Type: PassHistoryEntryType},
				{Type: PassHistoryEntryType, SeatIndex: 1},
				{Type: PassHistoryEntryType},
			},
		}

		next := NextGamePhase(game)

		if next == EndPhase {
			t.Errorf("Expected that the game should still continue for one turn but got %#v next", next)
		}

		game.History.AppendPass(1)
		next = NextGamePhase(game)

		if next != EndPhase {
			t.Errorf("Expected that the game should end but got %#v next", next)
		}
	})
}
