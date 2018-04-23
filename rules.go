package scrubble

import (
	"github.com/mandykoh/scrubble/dict"
	"github.com/mandykoh/scrubble/play"
	"github.com/mandykoh/scrubble/tile"
)

// Rules is an immutable struct representing the rules used by the game to check
// and validate various conditions for legality. The zero-value Rules uses
// default game play rules with a default English dictionary of words, without
// automatic word validation (words are only validated against the dictionary
// when a play is challenged, rather than automatically upon word scoring).
type Rules struct {
	dictionary          dict.Dictionary
	gamePhaseController GamePhaseController
	placementValidator  PlacementValidator
	rackValidator       RackValidator
	challengeValidator  ChallengeValidator
	wordScorer          WordScorer
	useDictForScoring   bool
}

// IsChallengeSuccessful determines if a challenge to a play is successful.
// Unless overridden by WithChallengeValidator, this uses the default
// implementation provided by the IsChallengeSuccessful function.
func (r *Rules) IsChallengeSuccessful(formedWords []play.Word) bool {
	dictionary := r.dictionary
	if dictionary == nil {
		dictionary = dict.DefaultEnglish
	}

	isChallengeSuccessful := r.challengeValidator
	if isChallengeSuccessful == nil {
		isChallengeSuccessful = IsChallengeSuccessful
	}
	return isChallengeSuccessful(formedWords, dictionary)
}

// NextGamePhase determines the next game phase given the game's current state.
// Unless overridden by WithGamePhaseController, this uses the default
// implementation provided by the NextGamePhase function.
//
// This is called at the end of each turn to determine the phase of the game.
func (r *Rules) NextGamePhase(game *Game) GamePhase {
	nextGamePhase := r.gamePhaseController
	if nextGamePhase == nil {
		nextGamePhase = NextGamePhase
	}
	return nextGamePhase(game)
}

// ScoreWords determines the scoring from a set of proposed tile placements.
// This assumes that the tiles are being placed in valid positions according to
// placement validation. Unless overridden by WithWordScorer, this uses the
// default implementation provided by the ScoreWords function. If
// WithDictionaryForScoring is set to true, words are validated against the
// current dictionary.
//
// If a score cannot be determined because not all formed words are valid, an
// InvalidWordError is returned containing the invalid words.
//
// Otherwise, the total score is returned along with the words that would be
// formed on the board should the tiles be placed.
func (r *Rules) ScoreWords(placements play.Tiles, board *Board) (score int, words []play.Word, err error) {
	dictionary := r.dictionary
	if !r.useDictForScoring {
		dictionary = func(string) bool { return true }
	} else if dictionary == nil {
		dictionary = dict.DefaultEnglish
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
func (r *Rules) ValidatePlacements(placements play.Tiles, board *Board) error {
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
func (r *Rules) ValidateTilesFromRack(rack tile.Rack, toPlay []tile.Tile) (used, remaining []tile.Tile, err error) {
	rackValidator := r.rackValidator
	if rackValidator == nil {
		rackValidator = ValidateTilesFromRack
	}
	return rackValidator(rack, toPlay)
}

// WithChallengeValidator returns a copy of these Rules which uses the
// specified function for determining the success or failure of challenges.
func (r Rules) WithChallengeValidator(validator ChallengeValidator) Rules {
	r.challengeValidator = validator
	return r
}

// WithDictionary returns a copy of these Rules which uses the specified
// dictionary for word validation.
func (r Rules) WithDictionary(dict dict.Dictionary) Rules {
	r.dictionary = dict
	return r
}

// WithDictionaryForScoring returns a copy of these Rules which optionally uses
// the current dictionary for word scoring. The default is to only use the
// dictionary when a play is challenged. Setting this to true will check all
// words against the dictionary and cause an InvalidWordError on scoring if any
// aren't valid.
func (r Rules) WithDictionaryForScoring(use bool) Rules {
	r.useDictForScoring = use
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
