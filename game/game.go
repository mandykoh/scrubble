package game

import (
	"math/rand"

	"github.com/mandykoh/scrubble/board"
	"github.com/mandykoh/scrubble/exchange"
	"github.com/mandykoh/scrubble/history"
	"github.com/mandykoh/scrubble/play"
	"github.com/mandykoh/scrubble/seat"
	"github.com/mandykoh/scrubble/tile"
)

// MinPlayers is the minimum number of players required before a game can be
// started.
const MinPlayers = 1

// ChallengeFailPenaltyPoints is the number of points deducted for a failed
// challenge.
const ChallengeFailPenaltyPoints = 5

// Game represents the rules and simulation for a single game. The zero-value of
// a Game is a game in the SetupPhase with no players.
type Game struct {
	Phase            Phase
	Seats            []seat.Seat
	Bag              tile.Bag
	Board            board.Board
	CurrentSeatIndex int
	Rules            Rules
	History          history.History
}

// New returns an initialised game in the SetupPhase with no players.
func New(bag tile.Bag, board board.Board) *Game {
	return &Game{
		Bag:   bag,
		Board: board,
	}
}

// NewWithDefaults returns an initialised game in the SetupPhase with no
// players, with a default bag and board layout.
func NewWithDefaults() *Game {
	return New(tile.BagWithStandardEnglishTiles(), board.WithStandardLayout())
}

// AddPlayer adds a seat for a new player to the game.
//
// If the game is not in the Setup phase, GameOutOfPhaseError is returned.
func (g *Game) AddPlayer() (s *seat.Seat, err error) {
	return s, g.requirePhase(SetupPhase, func() error {
		g.Seats = append(g.Seats, seat.Seat{})
		s = &g.Seats[len(g.Seats)-1]
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
	var lastPlay *history.Entry
	if len(g.History) > 0 {
		lastPlay = g.History.Last()
	}

	success, err = g.Rules.ValidateChallenge(lastPlay)
	if err != nil {
		return
	}

	if success {
		challenged := g.prevSeat()
		challenged.Rack.Remove(lastPlay.TilesDrawn...)
		challenged.Rack = append(challenged.Rack, lastPlay.TilesSpent...)
		challenged.Score -= lastPlay.Score

		for _, p := range lastPlay.TilesPlayed {
			g.Board.Position(p.Coord).Tile = nil
		}

		g.Bag = append(g.Bag, lastPlay.TilesDrawn...)
		g.Bag.Shuffle(r)

		g.History.AppendChallengeSuccess(challengerSeatIndex)
		g.Phase = MainPhase

	} else {
		challenger := &g.Seats[challengerSeatIndex]
		challenger.Score -= ChallengeFailPenaltyPoints
		g.History.AppendChallengeFail(challengerSeatIndex)
	}

	return
}

// CurrentSeat returns the seat for the player whose turn it currently is.
func (g *Game) CurrentSeat() *seat.Seat {
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
func (g *Game) ExchangeTiles(tiles []tile.Tile, r *rand.Rand) error {
	return g.requirePhase(MainPhase, func() error {
		if len(tiles) == 0 {
			return exchange.InvalidTileExchangeError{Reason: exchange.NoTilesExchangedReason}
		}
		if len(g.Bag) < tile.MaxRackTiles {
			return exchange.InvalidTileExchangeError{Reason: exchange.InsufficientTilesInBagReason}
		}

		s := g.CurrentSeat()

		used, remaining, err := g.Rules.ValidateTilesFromRack(s.Rack, tiles)
		if err != nil {
			return err
		}

		s.Rack = remaining
		for i := 0; i < len(used); i++ {
			s.Rack = append(s.Rack, g.Bag.DrawTile())
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
func (g *Game) Play(placements play.Tiles) (playedWords []play.Word, err error) {
	return playedWords, g.requirePhase(MainPhase, func() error {
		s := g.CurrentSeat()

		used, remaining, err := g.Rules.ValidateTilesFromRack(s.Rack, placements.Tiles())
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

		s.Rack = remaining
		placements.Place(&g.Board)
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

		if len(g.Seats) < MinPlayers {
			return NotEnoughPlayersError{MinPlayers, len(g.Seats)}
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

func (g *Game) endTurn(score int, tilesSpent []tile.Tile, tilesPlayed play.Tiles, wordsFormed []play.Word) {
	s := g.CurrentSeat()
	s.Score += score
	tilesDrawn := s.Rack.FillFromBag(&g.Bag)

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
		endGameScores := g.Rules.ScoreEndGame(g.History.Last(), g.Seats)
		for i, score := range endGameScores {
			s := &g.Seats[i]
			s.Score += score
			if i == g.prevSeatIndex() {
				g.History.Last().Score += score
			}
		}
	}
}

func (g *Game) nextSeatIndex() int {
	return (g.CurrentSeatIndex + 1) % len(g.Seats)
}

func (g *Game) prevSeat() *seat.Seat {
	return &g.Seats[g.prevSeatIndex()]
}

func (g *Game) prevSeatIndex() int {
	return (g.CurrentSeatIndex + (len(g.Seats) - 1)) % len(g.Seats)
}

func (g *Game) requirePhase(phase Phase, action func() error) error {
	if g.Phase != phase {
		return OutOfPhaseError{phase, g.Phase}
	}

	return action()
}
