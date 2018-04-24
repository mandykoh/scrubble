package scoring

import (
	"strings"

	"github.com/mandykoh/scrubble/board"
	"github.com/mandykoh/scrubble/coord"
	"github.com/mandykoh/scrubble/dict"
	"github.com/mandykoh/scrubble/history"
	"github.com/mandykoh/scrubble/play"
	"github.com/mandykoh/scrubble/seat"
	"github.com/mandykoh/scrubble/tile"
)

// MaxRackTilesBonus is the number of bonus points awarded for playing all the
// tiles on a full rack in one turn.
const MaxRackTilesBonus = 50

// ScoreEndGame determines the final scores to be added to each player's total
// after the last play of the game is made.
func ScoreEndGame(lastPlay *history.Entry, seats []seat.Seat) (finalScores []int) {
	finalScores = make([]int, len(seats))

	if lastPlay.Type == history.PlayEntryType {
		playOutBonus := 0
		for i, s := range seats {
			if i == lastPlay.SeatIndex {
				continue
			}
			for _, t := range s.Rack {
				playOutBonus += t.Points
			}
		}
		playOutBonus *= 2
		finalScores[lastPlay.SeatIndex] = playOutBonus

	} else {
		for i := range seats {
			s := &seats[i]
			gameEndPenalty := 0
			for _, t := range s.Rack {
				gameEndPenalty += t.Points
			}
			finalScores[i] = -gameEndPenalty
		}
	}

	return
}

// ScoreWords determines the scoring from a set of proposed tile placements.
// This assumes that the tiles are being placed in valid positions according to
// the game rules. This implements standard scoring rules.
//
// If a score cannot be determined because not all formed words are valid, an
// InvalidWordError is returned containing the invalid words.
//
// Otherwise, the total score is returned along with the words that would be
// formed on the board should the tiles be placed.
func ScoreWords(placements play.Tiles, board *board.Board, isWordValid dict.Dictionary) (score int, words []play.Word, err error) {
	var wordSpans []coord.Range
	findSpans(coord.Coord.West, coord.Coord.East, placements, &wordSpans, board)
	findSpans(coord.Coord.North, coord.Coord.South, placements, &wordSpans, board)

	score, words = spansToPlayedWords(wordSpans, placements, board)

	var singleTileSpans []coord.Range
	findUnspanned(placements, wordSpans, &singleTileSpans)
	_, invalidWords := spansToPlayedWords(singleTileSpans, placements, board)

	for _, w := range words {
		if !isWordValid(w.Word) {
			invalidWords = append(invalidWords, w)
		}
	}

	if len(invalidWords) > 0 {
		return 0, nil, play.InvalidWordError{Words: invalidWords}
	}

	if len(placements) >= tile.MaxRackTiles {
		score += MaxRackTilesBonus
	}

	return
}

func findSpans(growMinDir, growMaxDir func(coord.Coord) coord.Coord, placements play.Tiles, results *[]coord.Range, board *board.Board) {
	unspanned := append(play.Tiles{}, placements...)

	for p := unspanned.TakeLast(); p != nil; p = unspanned.TakeLast() {
		span := coord.Range{Min: p.Coord, Max: p.Coord}
		growSpan(&span.Min, growMinDir, &unspanned, board)
		growSpan(&span.Max, growMaxDir, &unspanned, board)

		if span.Min.Row != span.Max.Row || span.Min.Column != span.Max.Column {
			*results = append(*results, span)
		}
	}
}

func findUnspanned(placements play.Tiles, wordSpans []coord.Range, result *[]coord.Range) {
	for _, p := range placements {
		inSpan := false

		for _, s := range wordSpans {
			if s.Includes(p.Coord) {
				inSpan = true
				break
			}
		}

		if !inSpan {
			*result = append(*result, coord.Range{Min: p.Coord, Max: p.Coord})
		}
	}
}

func growSpan(growCoord *coord.Coord, growDir func(coord.Coord) coord.Coord, unspanned *play.Tiles, board *board.Board) {
	for {
		c := growDir(*growCoord)
		pos := board.Position(c)
		if pos == nil {
			break
		}

		if pos.Tile != nil {
			*growCoord = c
		} else if placement := unspanned.Take(c); placement != nil {
			*growCoord = c
		} else {
			break
		}
	}
}

func spansToPlayedWords(wordSpans []coord.Range, placements play.Tiles, b *board.Board) (totalScore int, words []play.Word) {
	for _, s := range wordSpans {
		var playedWord = play.Word{Range: s}
		var word strings.Builder
		var wordScoreModifiers []board.PositionType

		s.Each(func(c coord.Coord) error {
			var t *tile.Tile

			position := b.Position(c)
			if position.Tile != nil {
				t = position.Tile
				playedWord.Score += t.Points
			} else {
				t = &placements.Find(c).Tile
				playedWord.Score += position.Type.ModifyTileScore(t.Points)
				wordScoreModifiers = append(wordScoreModifiers, position.Type)
			}

			word.WriteRune(t.Letter)

			return nil
		})

		for _, m := range wordScoreModifiers {
			playedWord.Score = m.ModifyWordScore(playedWord.Score)
		}

		totalScore += playedWord.Score

		playedWord.Word = word.String()
		words = append(words, playedWord)
	}

	return
}
