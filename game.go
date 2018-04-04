package scrubble

import "math/rand"

const GameMinPlayers = 1

// Game represents the rules and simulation for a single game. The zero-value of
// a Game is a game in the SetupPhase with no players.
type Game struct {
	Phase            GamePhase
	Seats            []Seat
	Bag              Bag
	Board            Board
	CurrentSeatIndex int
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
		if len(placements) == 0 {
			return InvalidTilePlacementError{NoTilesPlacedReason}
		}

		remaining, missing := g.currentSeat().Rack.tryPlayTiles(placements)
		if len(missing) != 0 {
			return InsufficientTilesError{missing}
		}

		err := g.validateTilePositions(placements)
		if err != nil {
			return err
		}

		g.Board.placeTiles(placements)
		g.currentSeat().Rack = remaining

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

func (g *Game) validateAvailableAndContiguous(placements TilePlacements) error {
	minRow, minCol, maxRow, maxCol := placements.Bounds()

	if minRow != maxRow && minCol != maxCol {
		return InvalidTilePlacementError{PlacementNotLinearReason}
	}

	placed := 0

	for r := minRow; r <= maxRow; r++ {
		for c := minCol; c <= maxCol; c++ {
			placement := placements.Find(r, c)
			position := g.Board.Position(r, c)

			if position == nil {
				return InvalidTilePlacementError{PlacementOutOfBoundsReason}
			}

			if placement != nil {
				if position.Tile != nil {
					return InvalidTilePlacementError{PositionOccupiedReason}
				}
				placed++

			} else if position.Tile == nil {
				return InvalidTilePlacementError{PlacementNotContiguousReason}
			}
		}
	}

	if placed != len(placements) {
		return InvalidTilePlacementError{PlacementOverlapReason}
	}

	return nil
}

func (g *Game) validateConnected(placements TilePlacements) error {
	return nil
}

func (g *Game) validateTilePositions(placements TilePlacements) error {
	err := g.validateAvailableAndContiguous(placements)
	if err == nil {
		err = g.validateConnected(placements)
	}
	return err
}
