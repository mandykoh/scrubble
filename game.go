package scrubble

import (
	"math/rand"
)

const GameMinPlayers = 1

// Game represents the rules and simulation for a single game. The zero-value of
// a Game is a game in the SetupPhase with no players.
type Game struct {
	Phase            GamePhase
	Seats            []Seat
	Bag              Bag
	Board            Board
	CurrentSeatIndex int
	Rules            Rules
}

// AddPlayer adds a seat for a new player to the game.
//
// If the game is not in the Setup phase, GameOutOfPhaseError is returned.
func (g *Game) AddPlayer(p *Player) error {
	return g.requirePhase(SetupPhase, func() error {
		g.Seats = append(g.Seats, Seat{OccupiedBy: p})
		return nil
	})
}

// Play attempts to place tiles from the current player's rack on the board.
//
// If the game is not in the Main phase, GameOutOfPhaseError is returned.
//
// If the current player doesn't have the tiles required to make the play, an
// InsufficientTilesError is returned.
//
// If the tile placement is illegal, an InvalidTilePlacementError is returned.
func (g *Game) Play(placements TilePlacements) error {
	return g.requirePhase(MainPhase, func() error {
		remaining, err := g.Rules.ValidateTilesFromRack(g.currentSeat().Rack, placements)
		if err != nil {
			return err
		}

		err = g.Rules.ValidatePlacements(placements, &g.Board)
		if err != nil {
			return err
		}

		g.Board.placeTiles(placements)

		seat := g.currentSeat()
		seat.Rack = remaining
		seat.Rack.FillFromBag(&g.Bag)

		g.CurrentSeatIndex = (g.CurrentSeatIndex + 1) % len(g.Seats)

		return nil
	})
}

// RemovePlayer removes the seat occupied by the specified player. If no such
// seat exists, this has no effect.
//
// If the game is not in the Setup phase, GameOutOfPhaseError is returned.
func (g *Game) RemovePlayer(p *Player) error {
	return g.requirePhase(SetupPhase, func() error {
		for i, s := range g.Seats {
			if s.OccupiedBy == p {
				g.Seats = append(g.Seats[:i], g.Seats[i+1:]...)
				break
			}
		}
		return nil
	})
}

// Start begins the game by shuffling the bag, picking a random seat for the
// first turn, filling all players' tile racks from the bag, and moving the game
// into the MainPhase.
//
// The supplied random number generator is used to determine the bag shuffling
// and the starting player.
//
// If the game has no players, NotEnoughPlayersError is returned.
//
// If the game is not in the Setup phase, GameOutOfPhaseError is returned.
func (g *Game) Start(r *rand.Rand) error {
	return g.requirePhase(SetupPhase, func() error {

		if len(g.Seats) < GameMinPlayers {
			return NotEnoughPlayersError{GameMinPlayers, len(g.Seats)}
		}

		g.CurrentSeatIndex = r.Intn(len(g.Seats))
		g.Bag.Shuffle(r)

		for i := range g.Seats {
			g.Seats[i].Rack.FillFromBag(&g.Bag)
		}

		g.Phase = MainPhase
		return nil
	})
}

func (g *Game) currentSeat() *Seat {
	return &g.Seats[g.CurrentSeatIndex]
}

func (g *Game) requirePhase(phase GamePhase, action func() error) error {
	if g.Phase != phase {
		return GameOutOfPhaseError{phase, g.Phase}
	}

	return action()
}
