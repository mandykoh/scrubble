package scrubble

import "strings"

// ScoreWords determines the scoring from a set of proposed tile placements.
// This assumes that the tiles are being placed in valid positions according to
// the game rules.
//
// If a score cannot be determined because not all formed words are valid, an
// InvalidWordError is returned containing the invalid words.
//
// Otherwise, the total score is returned along with the words that would be
// formed on the board should the tiles be placed.
func ScoreWords(placements TilePlacements, board *Board) (score int, words []PlayedWord, err error) {
	var wordSpans []CoordRange
	findSpans(Coord.West, Coord.East, placements, &wordSpans, board)
	findSpans(Coord.North, Coord.South, placements, &wordSpans, board)

	var invalidWordSpans []CoordRange
	findUnspanned(placements, wordSpans, &invalidWordSpans)

	if len(invalidWordSpans) > 0 {
		_, invalidWords := spansToPlayedWords(invalidWordSpans, placements, board)
		return 0, nil, InvalidWordError{invalidWords}
	}

	score, words = spansToPlayedWords(wordSpans, placements, board)

	return
}

func findSpans(growMinDir, growMaxDir func(Coord) Coord, placements TilePlacements, results *[]CoordRange, board *Board) {
	unspanned := append(TilePlacements{}, placements...)

	for p := unspanned.takeLast(); p != nil; p = unspanned.takeLast() {
		span := CoordRange{p.Coord, p.Coord}
		growSpan(&span.Min, growMinDir, &unspanned, board)
		growSpan(&span.Max, growMaxDir, &unspanned, board)

		if span.Min.Row != span.Max.Row || span.Min.Column != span.Max.Column {
			*results = append(*results, span)
		}
	}
}

func findUnspanned(placements TilePlacements, wordSpans []CoordRange, result *[]CoordRange) {
	for _, p := range placements {
		inSpan := false

		for _, s := range wordSpans {
			if s.Includes(p.Coord) {
				inSpan = true
				break
			}
		}

		if !inSpan {
			*result = append(*result, CoordRange{p.Coord, p.Coord})
		}
	}
}

func growSpan(growCoord *Coord, growDir func(Coord) Coord, unspanned *TilePlacements, board *Board) {
	for {
		c := growDir(*growCoord)
		pos := board.Position(c)
		if pos == nil {
			break
		}

		if pos.Tile != nil {
			*growCoord = c
		} else if placement := unspanned.take(c); placement != nil {
			*growCoord = c
		} else {
			break
		}
	}
}

func spansToPlayedWords(wordSpans []CoordRange, placements TilePlacements, board *Board) (totalScore int, words []PlayedWord) {
	for _, s := range wordSpans {
		var playedWord = PlayedWord{CoordRange: s}
		var word strings.Builder

		s.EachCoord(func(c Coord) error {
			var tile *Tile

			position := board.Position(c)
			if position.Tile != nil {
				tile = position.Tile
				playedWord.Score += tile.Points
			} else {
				tile = &placements.Find(c).Tile
				playedWord.Score += position.Type.ModifyTileScore(*tile)
			}

			word.WriteRune(tile.Letter)

			return nil
		})

		totalScore += playedWord.Score

		playedWord.Word = word.String()
		words = append(words, playedWord)
	}

	return
}
