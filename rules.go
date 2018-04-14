package scrubble

// Rules is an immutable struct representing the rules used by the game to check
// and validate various conditions for legality. The zero-value Rules uses
// default game play rules with a default English dictionary of words.
type Rules struct {
	dictionary          Dictionary
	gamePhaseController GamePhaseController
	placementValidator  PlacementValidator
	rackValidator       RackValidator
	wordScorer          WordScorer
}

// NextGamePhase determines the next game phase given the last turn and the
// game's current state. Unless overridden by WithGamePhaseController, this uses
// the default implementation provided by the NextGamePhase function.
func (r *Rules) NextGamePhase(lastSeat *Seat, lastScore int, game *Game) GamePhase {
	nextGamePhase := r.gamePhaseController
	if nextGamePhase == nil {
		nextGamePhase = NextGamePhase
	}
	return nextGamePhase(lastSeat, lastScore, game)
}

// ScoreWords determines the scoring from a set of proposed tile placements.
// This assumes that the tiles are being placed in valid positions according to
// placement validation. Unless overridden by WithWordScorer, this uses the
// default implementation provided by the ScoreWords function, with the current
// dictionary.
//
// If a score cannot be determined because not all formed words are valid, an
// InvalidWordError is returned containing the invalid words.
//
// Otherwise, the total score is returned along with the words that would be
// formed on the board should the tiles be placed.
func (r *Rules) ScoreWords(placements TilePlacements, board *Board) (score int, words []PlayedWord, err error) {
	dictionary := r.dictionary
	if dictionary == nil {
		dictionary = DefaultEnglishDictionary
	}

	wordScorer := r.wordScorer
	if wordScorer == nil {
		wordScorer = ScoreWords
	}
	return wordScorer(placements, board, dictionary)
}

// ValidatePlacements checks the intended placement of tiles on a board for
// legality. Unless overridden by WithPlacementValidator, this uses the default
// implementation provided by the ValidatePlacements function.
//
// If any violations are detected, InvalidTilePlacementError is returned with
// the reason indicating the violation.
//
// Otherwise, nil is returned, indicating that it would be safe to place the
// given tiles on the board (word validity not withstanding).
func (r *Rules) ValidatePlacements(placements TilePlacements, board *Board) error {
	placementValidator := r.placementValidator
	if placementValidator == nil {
		placementValidator = ValidatePlacements
	}
	return placementValidator(placements, board)
}

// ValidateTilesFromRack checks the intended usage of tiles from a rack for
// legality. Unless overridden by WithRackValidator, this uses the default
// implementation provided by the ValidateTilesFromRack function.
//
// An InsufficientTilesError is returned with the missing tiles in the Missing
// field if the rack doesn't contain sufficient tiles to make the play.
//
// Otherwise, the remainder (after the placed tiles have been removed from the
// rack) is returned with no error, indicating that it would be safe to update
// the rack for placement.
func (r *Rules) ValidateTilesFromRack(rack Rack, placements TilePlacements) (remaining Rack, err error) {
	rackValidator := r.rackValidator
	if rackValidator == nil {
		rackValidator = ValidateTilesFromRack
	}
	return rackValidator(rack, placements)
}

// WithDictionary returns a copy of these Rules which uses the specified
// dictionary for word validation.
func (r Rules) WithDictionary(dict Dictionary) Rules {
	r.dictionary = dict
	return r
}

// WithGamePhaseController returns a copy of these Rules which uses the
// specified function for determining the progression of the game, and the
// conditions under which the game ends.
func (r Rules) WithGamePhaseController(controller GamePhaseController) Rules {
	r.gamePhaseController = controller
	return r
}

// WithPlacementValidator returns a copy of these Rules which uses the specified
// function for tile placement validation.
func (r Rules) WithPlacementValidator(validator PlacementValidator) Rules {
	r.placementValidator = validator
	return r
}

// WithRackValidator returns a copy of these Rules which uses the specified
// function for tile rack validation.
func (r Rules) WithRackValidator(validator RackValidator) Rules {
	r.rackValidator = validator
	return r
}

// WithWordScorer returns a copy of these Rules which uses the specified word
// scorer for computing formed words and their scores.
func (r Rules) WithWordScorer(scorer WordScorer) Rules {
	r.wordScorer = scorer
	return r
}
