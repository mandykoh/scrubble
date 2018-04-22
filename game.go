package scrubble

import (
	"math/rand"
)

const GameMinPlayers = 1
const ChallengeFailPenaltyPoints = 5

// Game represents the rules and simulation for a single game. The zero-value of
// a Game is a game in the SetupPhase with no players.
type Game struct {
	Phase            GamePhase
	Seats            []Seat
	Bag              Bag
	Board            Board
	CurrentSeatIndex int
	Rules            Rules
	History          History
}

// NewGame returns an initialised game in the SetupPhase with no players.
func NewGame(bag Bag, board Board) *Game {
	return &Game{
		Bag:   bag,
		Board: board,
	}
}

// NewGameWithDefaults returns an initialised game in the SetupPhase with no
// players, with a default bag and board layout.
func NewGameWithDefaults() *Game {
	return NewGame(BagWithStandardEnglishTiles(), BoardWithStandardLayout())
}

// AddPlayer adds a seat for a new player to the game.
//
// If the game is not in the Setup phase, GameOutOfPhaseError is returned.
func (g *Game) AddPlayer() (seat *Seat, err error) {
	return seat, g.requirePhase(SetupPhase, func() error {
		g.Seats = append(g.Seats, Seat{})
		seat = &g.Seats[len(g.Seats)-1]
		return nil
	})
}

// Challenge challenges the last turn's play. A challenge succeeds if any of the
// words formed in the play are invalid according to the dictionary in use,
// which then causes the play to be withdrawn and play to proceed (with the
// challenged player effectively losing their turn). If all words are found to
// be valid, the challenge fails and the challenger is penalised.
//
// The supplied random number generator is used to reshuffle drawn tiles back
// into the bag upon a successful challenge.
//
// If a challenge is not allowed, an InvalidChallengeError is returned with the
// reason. Otherwise, whether the challenge succeeded or failed is returned.
func (g *Game) Challenge(challengerSeatIndex int, r *rand.Rand) (success bool, err error) {
	if len(g.History) == 0 {
		return false, InvalidChallengeError{NoPlayToChallengeReason}
	}
	if challengerSeatIndex < 0 || challengerSeatIndex >= len(g.Seats) {
		return false, InvalidChallengeError{InvalidChallengerReason}
	}

	play := g.History.Last()
	switch play.Type {
	case ChallengeFailHistoryEntryType, ChallengeSuccessHistoryEntryType:
		return false, InvalidChallengeError{PlayAlreadyChallengedReason}

	case PlayHistoryEntryType:
		break

	default:
		return false, InvalidChallengeError{NoPlayToChallengeReason}
	}

	success = g.Rules.IsChallengeSuccessful(play.WordsFormed)
	if success {
		challenged := g.prevSeat()
		challenged.Rack.Remove(play.TilesDrawn...)
		challenged.Rack = append(challenged.Rack, play.TilesSpent...)
		challenged.Score -= play.Score

		for _, p := range play.TilesPlayed {
			g.Board.Position(p.Coord).Tile = nil
		}

		g.Bag = append(g.Bag, play.TilesDrawn...)
		g.Bag.Shuffle(r)

		g.History.AppendChallengeSuccess(challengerSeatIndex)
		g.Phase = MainPhase

	} else {
		challenger := &g.Seats[challengerSeatIndex]
		challenger.Score -= ChallengeFailPenaltyPoints
		g.History.AppendChallengeFail(challengerSeatIndex)
	}

	return success, nil
}

// CurrentSeat returns the seat for the player whose turn it currently is.
func (g *Game) CurrentSeat() *Seat {
	return &g.Seats[g.CurrentSeatIndex]
}

// ExchangeTiles exchanges tiles from the current player's rack with tiles from
// the bag, ending the turn.
//
// The supplied random number generator is used to reshuffle the bag after
// replacing the exchanged tiles.
//
// If the game is not in the Main phase, GameOutOfPhaseError is returned.
//
// If the current player doesn't have the required tiles to exchange, an
// InsufficientTilesError is returned.
//
// If an attempt is made to exchange zero tiles, or there are fewer than
// MaxRackTiles in the bag, tile exchange is illegal and an
// InvalidTileExchangeError is returned.
func (g *Game) ExchangeTiles(tiles []Tile, r *rand.Rand) error {
	return g.requirePhase(MainPhase, func() error {
		if len(tiles) == 0 {
			return InvalidTileExchangeError{NoTilesExchangedReason}
		}
		if len(g.Bag) < MaxRackTiles {
			return InvalidTileExchangeError{InsufficientTilesInBagReason}
		}

		seat := g.CurrentSeat()

		used, remaining, err := g.Rules.ValidateTilesFromRack(seat.Rack, tiles)
		if err != nil {
			return err
		}

		seat.Rack = remaining
		for i := 0; i < len(used); i++ {
			seat.Rack = append(seat.Rack, g.Bag.DrawTile())
		}

		g.Bag = append(g.Bag, used...)
		g.Bag.Shuffle(r)

		g.endTurn(0, used, nil, nil)

		return nil
	})
}

// Pass forfeits the current player's turn.
//
// If the game is not in the Main phase, GameOutOfPhaseError is returned.
func (g *Game) Pass() error {
	return g.requirePhase(MainPhase, func() error {
		g.endTurn(0, nil, nil, nil)
		return nil
	})
}

// Play attempts to place tiles from the current player's rack on the board. On
// success, the words formed by the play are returned, the game is updated, and
// play moves to the next player in turn.
//
// Wildcard tiles (tiles with a zero-point tile value) can be played by passing
// a TilePlacement with the letter replaced by any desired letter (but keeping
// the point value at zero). The wildcard will be correctly deducted from the
// player's rack.
//
// If the game is not in the Main phase, GameOutOfPhaseError is returned.
//
// If the current player doesn't have the tiles required to make the play, an
// InsufficientTilesError is returned.
//
// If the tile placement is illegal, an InvalidTilePlacementError is returned.
//
// If any formed words are invalid, an InvalidWordError is returned.
func (g *Game) Play(placements TilePlacements) (playedWords []PlayedWord, err error) {
	return playedWords, g.requirePhase(MainPhase, func() error {
		seat := g.CurrentSeat()

		used, remaining, err := g.Rules.ValidateTilesFromRack(seat.Rack, placements.Tiles())
		if err != nil {
			return err
		}

		err = g.Rules.ValidatePlacements(placements, &g.Board)
		if err != nil {
			return err
		}

		var score int
		score, playedWords, err = g.Rules.ScoreWords(placements, &g.Board)
		if err != nil {
			return err
		}

		seat.Rack = remaining
		g.Board.placeTiles(placements)
		g.endTurn(score, used, placements, playedWords)

		return nil
	})
}

// RemovePlayer removes the seat at the specified index. If no such seat exists,
// this has no effect.
//
// If the game is not in the Setup phase, GameOutOfPhaseError is returned.
func (g *Game) RemovePlayer(seatIndex int) error {
	return g.requirePhase(SetupPhase, func() error {
		if seatIndex >= 0 && seatIndex < len(g.Seats) {
			g.Seats = append(g.Seats[:seatIndex], g.Seats[seatIndex+1:]...)
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

func (g *Game) endTurn(score int, tilesSpent []Tile, tilesPlayed TilePlacements, wordsFormed []PlayedWord) {
	seat := g.CurrentSeat()
	seat.Score += score
	tilesDrawn := seat.Rack.FillFromBag(&g.Bag)

	if len(tilesPlayed) > 0 {
		g.History.AppendPlay(g.CurrentSeatIndex, score, tilesSpent, tilesPlayed, tilesDrawn, wordsFormed)
	} else if len(tilesSpent) > 0 {
		g.History.AppendExchange(g.CurrentSeatIndex, tilesSpent, tilesDrawn)
	} else {
		g.History.AppendPass(g.CurrentSeatIndex)
	}

	g.CurrentSeatIndex = g.nextSeatIndex()
	g.Phase = g.Rules.NextGamePhase(g)

	if g.Phase == EndPhase {
		if len(tilesPlayed) > 0 {
			playOutBonus := 0
			for i, s := range g.Seats {
				if i == g.prevSeatIndex() {
					continue
				}
				for _, t := range s.Rack {
					playOutBonus += t.Points
				}
			}
			playOutBonus *= 2

			seat.Score += playOutBonus
			g.History.Last().Score += playOutBonus

		} else {
			for i := range g.Seats {
				s := &g.Seats[i]
				gameEndPenalty := 0
				for _, t := range s.Rack {
					gameEndPenalty += t.Points
				}
				s.Score -= gameEndPenalty
			}
		}
	}
}

func (g *Game) nextSeatIndex() int {
	return (g.CurrentSeatIndex + 1) % len(g.Seats)
}

func (g *Game) prevSeat() *Seat {
	return &g.Seats[g.prevSeatIndex()]
}

func (g *Game) prevSeatIndex() int {
	return (g.CurrentSeatIndex + (len(g.Seats) - 1)) % len(g.Seats)
}

func (g *Game) requirePhase(phase GamePhase, action func() error) error {
	if g.Phase != phase {
		return GameOutOfPhaseError{phase, g.Phase}
	}

	return action()
}
