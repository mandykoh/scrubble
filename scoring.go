package scrubble

// ScoreWords determines the total score from a set of proposed tile placements.
// This assumes that the tiles are being placed in valid positions according to
// the game rules.
//
// If a score cannot be determined (eg because not all formed words are valid),
// an error is returned.
func ScoreWords(placements TilePlacements, board *Board) (score int, err error) {
	var wordSpans []CoordRange
	findSpans(Coord.West, Coord.East, placements, &wordSpans, board)
	findSpans(Coord.North, Coord.South, placements, &wordSpans, board)

	if len(wordSpans) == 0 && len(placements) > 0 {
		return 0, InvalidWordError{SingleLetterWordDisallowedReason}
	}

	for _, s := range wordSpans {
		s.EachCoord(func(c Coord) error {
			position := board.Position(c)

			if position.Tile != nil {
				score += position.Tile.Points
			} else {
				placement := placements.Find(c)
				score += placement.Tile.Points
			}

			return nil
		})
	}

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
