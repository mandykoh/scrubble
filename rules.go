package scrubble

// Rules represent the rules used by the game to check and validate various
// conditions for legality. The zero-value Rules uses default game play rules.
type Rules struct {
	ScoreWordsFunc            func(placements TilePlacements, board *Board) (score int, wordSpans []CoordRange, err error)
	ValidatePlacementsFunc    func(placements TilePlacements, board *Board) error
	ValidateTilesFromRackFunc func(rack Rack, placements TilePlacements) (remaining Rack, err error)
}

func (r *Rules) ScoreWords(placements TilePlacements, board *Board) (score int, wordSpans []CoordRange, err error) {
	if r.ScoreWordsFunc == nil {
		return ScoreWords(placements, board)
	}
	return r.ScoreWordsFunc(placements, board)
}

func (r *Rules) ValidatePlacements(placements TilePlacements, board *Board) error {
	if r.ValidatePlacementsFunc == nil {
		return ValidatePlacements(placements, board)
	}
	return r.ValidatePlacementsFunc(placements, board)
}

func (r *Rules) ValidateTilesFromRack(rack Rack, placements TilePlacements) (remaining Rack, err error) {
	if r.ValidateTilesFromRackFunc == nil {
		return ValidateTilesFromRack(rack, placements)
	}
	return r.ValidateTilesFromRackFunc(rack, placements)
}
