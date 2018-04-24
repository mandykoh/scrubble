package game

import (
	"testing"

	"github.com/mandykoh/scrubble/history"
	"github.com/mandykoh/scrubble/seat"
	"github.com/mandykoh/scrubble/tile"
)

func TestNextPhase(t *testing.T) {

	t.Run("allows the game to continue when the rack could be replenished with at least one tile", func(t *testing.T) {
		game := &Game{
			Seats: []seat.Seat{
				{
					Rack: tile.Rack{
						{'A', 1},
					},
				},
			},
			History: history.History{{Type: history.UnknownEntryType, Score: 123}},
		}

		next := NextPhase(game)

		if next != MainPhase {
			t.Errorf("Expected that the game should continue but got %#v next", next)
		}
	})

	t.Run("ends the game when the rack was emptied and couldn't be replenished from the bag", func(t *testing.T) {
		game := &Game{
			Seats: []seat.Seat{
				{Rack: tile.Rack{}},
			},
			History: history.History{{Type: history.UnknownEntryType, Score: 123}},
		}

		next := NextPhase(game)

		if next != EndPhase {
			t.Errorf("Expected that the game should end but got %#v next", next)
		}
	})

	t.Run("ends the game after six consecutive scoreless turns", func(t *testing.T) {
		game := &Game{
			Seats: []seat.Seat{
				{Rack: tile.Rack{{'A', 1}}},
				{Rack: tile.Rack{{'B', 2}}},
			},
			History: history.History{
				{Type: history.PassEntryType},
				{Type: history.PassEntryType, SeatIndex: 1},
				{Type: history.PassEntryType},
				{Type: history.PassEntryType, SeatIndex: 1},
				{Type: history.PassEntryType},
			},
		}

		next := NextPhase(game)

		if next == EndPhase {
			t.Errorf("Expected that the game should still continue for one turn but got %#v next", next)
		}

		game.History.AppendPass(1)
		next = NextPhase(game)

		if next != EndPhase {
			t.Errorf("Expected that the game should end but got %#v next", next)
		}
	})

	t.Run("treats successfully challenged plays as scoreless turns", func(t *testing.T) {
		game := &Game{
			Seats: []seat.Seat{
				{Rack: tile.Rack{{'A', 1}}},
				{Rack: tile.Rack{{'B', 2}}},
			},
			History: history.History{
				{Type: history.PlayEntryType, Score: 123},
				{Type: history.ChallengeSuccessEntryType, SeatIndex: 1},
				{Type: history.PlayEntryType, SeatIndex: 1, Score: 234},
				{Type: history.ChallengeSuccessEntryType},
				{Type: history.PassEntryType},
				{Type: history.PassEntryType, SeatIndex: 1},
				{Type: history.PassEntryType},
			},
		}

		next := NextPhase(game)

		if next == EndPhase {
			t.Errorf("Expected that the game should still continue for one turn but got %#v next", next)
		}

		game.History.AppendPass(1)
		next = NextPhase(game)

		if next != EndPhase {
			t.Errorf("Expected that the game should end but got %#v next", next)
		}
	})
}
