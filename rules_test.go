package scrubble

import (
	"testing"

	"github.com/mandykoh/scrubble/coord"
	"github.com/mandykoh/scrubble/dict"
	"github.com/mandykoh/scrubble/play"
	"github.com/mandykoh/scrubble/tile"
)

func TestRules(t *testing.T) {
	rules := Rules{}
	board := BoardWithStandardLayout()

	t.Run("zero-value", func(t *testing.T) {

		t.Run("allows any word for scoring", func(t *testing.T) {
			_, _, err := rules.ScoreWords(play.Tiles{
				{tile.Make('A', 1), coord.Make(0, 0)},
				{tile.Make('A', 1), coord.Make(1, 0)},
				{tile.Make('R', 1), coord.Make(2, 0)},
				{tile.Make('D', 1), coord.Make(3, 0)},
				{tile.Make('V', 1), coord.Make(4, 0)},
				{tile.Make('A', 1), coord.Make(5, 0)},
				{tile.Make('R', 1), coord.Make(6, 0)},
				{tile.Make('K', 1), coord.Make(7, 0)},
			}, &board)

			if err != nil {
				t.Errorf("Expected success but got error %v", err)
			}

			_, _, err = rules.ScoreWords(play.Tiles{
				{tile.Make('V', 1), coord.Make(0, 0)},
				{tile.Make('X', 1), coord.Make(1, 0)},
				{tile.Make('T', 1), coord.Make(2, 0)},
				{tile.Make('Q', 1), coord.Make(3, 0)},
				{tile.Make('R', 1), coord.Make(4, 0)},
				{tile.Make('P', 1), coord.Make(5, 0)},
			}, &board)

			if err != nil {
				t.Errorf("Expected success but got error %v", err)
			}
		})

		t.Run("can check for next phase", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Expected NextGamePhase to succeed without panic but got %v", r)
				}
			}()

			rules.NextGamePhase(&Game{
				Seats: []Seat{
					{},
				},
				History: History{
					{Type: UnknownHistoryEntryType, Score: 123},
				},
			})
		})

		t.Run("can score words", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Expected ScoreWords to succeed without panic but got %v", r)
				}
			}()

			rules.ScoreWords(play.Tiles{}, &board)
		})

		t.Run("can validate challenges", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Expected IsChallengeSuccessful to succeed without panic but got %v", r)
				}
			}()

			rules.IsChallengeSuccessful([]play.Word{})
		})

		t.Run("can validate placements", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Expected ValidatePlacements to succeed without panic but got %v", r)
				}
			}()

			rules.ValidatePlacements(play.Tiles{}, &board)
		})

		t.Run("can validate rack", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Expected ValidateTilesFromRack to succeed without panic but got %v", r)
				}
			}()

			rules.ValidateTilesFromRack(tile.Rack{}, []tile.Tile{})
		})
	})

	t.Run(".WithChallengeValidator()", func(t *testing.T) {
		validatorCalled := 0
		validator := func([]play.Word, dict.Dictionary) bool {
			validatorCalled++
			return false
		}

		overriddenRules := Rules{}.WithChallengeValidator(validator)

		t.Run("sets the validator to use for challenge validation", func(t *testing.T) {
			overriddenRules.IsChallengeSuccessful([]play.Word{})

			if actual, expected := validatorCalled, 1; actual != expected {
				t.Errorf("Expected overridden validator to be called once but got %d invocations", actual)
			}
		})

		t.Run("leaves the original rules unmodified", func(t *testing.T) {
			if actual := rules.challengeValidator; actual != nil {
				t.Errorf("Expected original challenge validator to be unmodified but wasn't")
			}
		})
	})

	t.Run(".WithDictionary()", func(t *testing.T) {
		dictionaryCalled := 0
		dictionary := func(string) bool {
			dictionaryCalled++
			return true
		}

		overriddenRules := Rules{}.WithDictionary(dictionary)

		t.Run("sets the dictionary to use for word scoring", func(t *testing.T) {
			dictionaryCalled = 0

			r := overriddenRules.WithDictionaryForScoring(true)
			r.ScoreWords(play.Tiles{
				{tile.Make('C', 1), coord.Make(0, 0)},
				{tile.Make('A', 1), coord.Make(1, 0)},
				{tile.Make('T', 1), coord.Make(2, 0)},
			}, &board)

			if actual, expected := dictionaryCalled, 1; actual != expected {
				t.Errorf("Expected overridden dictionary to be called once but got %d invocations", actual)
			}
		})

		t.Run("continues to allow any word if the dictionary for word scoring is disabled", func(t *testing.T) {
			dictionaryCalled = 0

			r := overriddenRules.WithDictionaryForScoring(false)
			_, _, err := r.ScoreWords(play.Tiles{
				{tile.Make('D', 1), coord.Make(0, 0)},
				{tile.Make('J', 1), coord.Make(1, 0)},
				{tile.Make('K', 1), coord.Make(2, 0)},
			}, &board)

			if err != nil {
				t.Errorf("Expected success but got error %v", err)
			}

			if actual, expected := dictionaryCalled, 0; actual != expected {
				t.Errorf("Expected overridden dictionary to remain unused but got %d invocations", actual)
			}
		})

		t.Run("leaves the original rules unmodified", func(t *testing.T) {
			if actual := rules.dictionary; actual != nil {
				t.Errorf("Expected original dictionary to be unmodified but wasn't")
			}
		})
	})

	t.Run(".WithGamePhaseController()", func(t *testing.T) {
		controllerCalled := 0
		controller := func(*Game) GamePhase {
			controllerCalled++
			return MainPhase
		}

		overriddenRules := Rules{}.WithGamePhaseController(controller)

		t.Run("sets the function to use for game phase progression", func(t *testing.T) {
			overriddenRules.NextGamePhase(nil)

			if actual, expected := controllerCalled, 1; actual != expected {
				t.Errorf("Expected overridden end-of-game checker to be called once but got %d invocations", actual)
			}
		})

		t.Run("leaves the original rules unmodified", func(t *testing.T) {
			if actual := rules.gamePhaseController; actual != nil {
				t.Errorf("Expected original game phase controller  to be unmodified but wasn't")
			}
		})
	})

	t.Run(".WithPlacementValidator()", func(t *testing.T) {
		validatorCalled := 0
		validator := func(play.Tiles, *Board) error {
			validatorCalled++
			return nil
		}

		overriddenRules := Rules{}.WithPlacementValidator(validator)

		t.Run("sets the validator to use for placement validation", func(t *testing.T) {
			overriddenRules.ValidatePlacements(play.Tiles{}, &board)

			if actual, expected := validatorCalled, 1; actual != expected {
				t.Errorf("Expected overridden validator to be called once but got %d invocations", actual)
			}
		})

		t.Run("leaves the original rules unmodified", func(t *testing.T) {
			if actual := rules.placementValidator; actual != nil {
				t.Errorf("Expected original placement validator to be unmodified but wasn't")
			}
		})
	})

	t.Run(".WithRackValidator()", func(t *testing.T) {
		validatorCalled := 0
		validator := func(tile.Rack, []tile.Tile) ([]tile.Tile, []tile.Tile, error) {
			validatorCalled++
			return nil, nil, nil
		}

		overriddenRules := Rules{}.WithRackValidator(validator)

		t.Run("sets the validator to use for rack validation", func(t *testing.T) {
			overriddenRules.ValidateTilesFromRack(tile.Rack{}, []tile.Tile{})

			if actual, expected := validatorCalled, 1; actual != expected {
				t.Errorf("Expected overridden validator to be called once but got %d invocations", actual)
			}
		})

		t.Run("leaves the original rules unmodified", func(t *testing.T) {
			if actual := rules.rackValidator; actual != nil {
				t.Errorf("Expected original rack validator to be unmodified but wasn't")
			}
		})
	})

	t.Run(".WithWordScorer()", func(t *testing.T) {
		scorerCalled := 0
		scorer := func(play.Tiles, *Board, dict.Dictionary) (int, []play.Word, error) {
			scorerCalled++
			return 0, nil, nil
		}

		overriddenRules := Rules{}.WithWordScorer(scorer)

		t.Run("sets the word scorer to use", func(t *testing.T) {
			overriddenRules.ScoreWords(play.Tiles{}, &board)

			if actual, expected := scorerCalled, 1; actual != expected {
				t.Errorf("Expected overridden word scorer to be called once but got %d invocations", actual)
			}
		})

		t.Run("leaves the original rules unmodified", func(t *testing.T) {
			if actual := rules.wordScorer; actual != nil {
				t.Errorf("Expected original word scorer to be unmodified but wasn't")
			}
		})
	})
}
