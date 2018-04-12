package scrubble

import (
	"testing"
)

func TestRules(t *testing.T) {
	rules := Rules{}
	board := BoardWithStandardLayout()

	t.Run("zero-value", func(t *testing.T) {

		t.Run("can score words", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Expected ScoreWords to succeed without panic but got %v", r)
				}
			}()

			rules.ScoreWords(TilePlacements{}, &board)
		})

		t.Run("can validate placements", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Expected ValidatePlacements to succeed without panic but got %v", r)
				}
			}()

			rules.ValidatePlacements(TilePlacements{}, &board)
		})

		t.Run("can validate rack", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Expected ValidateTilesFromRack to succeed without panic but got %v", r)
				}
			}()

			rules.ValidateTilesFromRack(Rack{}, TilePlacements{})
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
			overriddenRules.ScoreWords(TilePlacements{
				{Tile{'C', 1}, Coord{0, 0}},
				{Tile{'A', 1}, Coord{1, 0}},
				{Tile{'T', 1}, Coord{2, 0}},
			}, &board)

			if actual, expected := dictionaryCalled, 1; actual != expected {
				t.Errorf("Expected overridden dictionary to be called once but got %d invocations", actual)
			}
		})

		t.Run("leaves the original rules unmodified", func(t *testing.T) {
			if actual := rules.dictionary; actual != nil {
				t.Errorf("Expected original dictionary to be unmodified but wasn't")
			}
		})
	})

	t.Run(".WithPlacementValidator()", func(t *testing.T) {
		validatorCalled := 0
		validator := func(TilePlacements, *Board) error {
			validatorCalled++
			return nil
		}

		overriddenRules := Rules{}.WithPlacementValidator(validator)

		t.Run("sets the validator to use for placement validation", func(t *testing.T) {
			overriddenRules.ValidatePlacements(TilePlacements{}, &board)

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
		validator := func(Rack, TilePlacements) (Rack, error) {
			validatorCalled++
			return nil, nil
		}

		overriddenRules := Rules{}.WithRackValidator(validator)

		t.Run("sets the validator to use for rack validation", func(t *testing.T) {
			overriddenRules.ValidateTilesFromRack(Rack{}, TilePlacements{})

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
		scorer := func(TilePlacements, *Board, Dictionary) (int, []PlayedWord, error) {
			scorerCalled++
			return 0, nil, nil
		}

		overriddenRules := Rules{}.WithWordScorer(scorer)

		t.Run("sets the word scorer to use", func(t *testing.T) {
			overriddenRules.ScoreWords(TilePlacements{}, &board)

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
