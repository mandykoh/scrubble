package scrubble

// Rules represent the rules used by the game to check and validate various
// conditions for legality. The zero-value Rules uses default game play rules
// with a permissive dictionary that allows any word.
type Rules struct {
	dictionary         Dictionary
	placementValidator PlacementValidator
	rackValidator      RackValidator
	wordScorer         WordScorer
}

func (r *Rules) ScoreWords(placements TilePlacements, board *Board) (score int, words []PlayedWord, err error) {
	dictionary := r.dictionary
	if dictionary == nil {
		dictionary = func(string) bool { return true }
	}

	wordScorer := r.wordScorer
	if wordScorer == nil {
		wordScorer = ScoreWords
	}
	return wordScorer(placements, board, dictionary)
}

func (r *Rules) ValidatePlacements(placements TilePlacements, board *Board) error {
	placementValidator := r.placementValidator
	if placementValidator == nil {
		placementValidator = ValidatePlacements
	}
	return placementValidator(placements, board)
}

func (r *Rules) ValidateTilesFromRack(rack Rack, placements TilePlacements) (remaining Rack, err error) {
	rackValidator := r.rackValidator
	if rackValidator == nil {
		rackValidator = ValidateTilesFromRack
	}
	return rackValidator(rack, placements)
}

func (r *Rules) WithDictionary(dict Dictionary) Rules {
	rules := *r
	rules.dictionary = dict
	return rules
}

func (r *Rules) WithPlacementValidator(validator PlacementValidator) Rules {
	rules := *r
	rules.placementValidator = validator
	return rules
}

func (r *Rules) WithRackValidator(validator RackValidator) Rules {
	rules := *r
	rules.rackValidator = validator
	return rules
}

func (r *Rules) WithWordScorer(scorer WordScorer) Rules {
	rules := *r
	rules.wordScorer = scorer
	return rules
}
