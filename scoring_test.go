package scrubble

import "testing"

func TestScoreWords(t *testing.T) {

	setupBoard := func() *Board {
		board := BoardWithStandardLayout()
		return &board
	}

	t.Run("returns an error for single-letter words", func(t *testing.T) {
		board := setupBoard()

		_, _, err := ScoreWords(TilePlacements{
			{Tile{'A', 2}, Coord{7, 7}},
		}, board)

		if err == nil {
			t.Errorf("Expected an error for single-letter word but got nil")
		} else {
			if actual, expected := err, (InvalidWordError{SingleLetterWordDisallowedReason}); actual != expected {
				t.Errorf("Expected error %#v but got %#v", expected, actual)
			}
		}
	})

	t.Run("counts entire horizontal word placed on starting position", func(t *testing.T) {
		board := setupBoard()

		score, wordSpans, err := ScoreWords(TilePlacements{
			{Tile{'D', 2}, Coord{7, 6}},
			{Tile{'O', 1}, Coord{7, 7}},
			{Tile{'G', 2}, Coord{7, 8}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 5; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(wordSpans), 1; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := wordSpans[0].Min, (Coord{7, 6}); actual != expected {
					t.Errorf("Expected word to begin at %v but was %v", expected, actual)
				}
				if actual, expected := wordSpans[0].Max, (Coord{7, 8}); actual != expected {
					t.Errorf("Expected word to end at %v but was %v", expected, actual)
				}
			}
		}
	})

	t.Run("counts entire vertical word placed on starting position", func(t *testing.T) {
		board := setupBoard()

		score, wordSpans, err := ScoreWords(TilePlacements{
			{Tile{'D', 2}, Coord{6, 7}},
			{Tile{'O', 1}, Coord{7, 7}},
			{Tile{'G', 2}, Coord{8, 7}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 5; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(wordSpans), 1; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := wordSpans[0].Min, (Coord{6, 7}); actual != expected {
					t.Errorf("Expected word to begin at %v but was %v", expected, actual)
				}
				if actual, expected := wordSpans[0].Max, (Coord{8, 7}); actual != expected {
					t.Errorf("Expected word to end at %v but was %v", expected, actual)
				}
			}
		}
	})

	t.Run("counts hooked words connected to the main word", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{5, 4}).Tile = &Tile{'D', 2}
		board.Position(Coord{6, 4}).Tile = &Tile{'O', 1}
		board.Position(Coord{7, 4}).Tile = &Tile{'G', 2}

		score, wordSpans, err := ScoreWords(TilePlacements{
			{Tile{'S', 2}, Coord{8, 4}},
			{Tile{'O', 2}, Coord{8, 5}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 11; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(wordSpans), 2; actual != expected {
				t.Errorf("Expected two words formed but found %d", actual)
			} else {
				if actual, expected := wordSpans[0].Min, (Coord{8, 4}); actual != expected {
					t.Errorf("Expected first word to begin at %v but was %v", expected, actual)
				}
				if actual, expected := wordSpans[0].Max, (Coord{8, 5}); actual != expected {
					t.Errorf("Expected first word to end at %v but was %v", expected, actual)
				}
				if actual, expected := wordSpans[1].Min, (Coord{5, 4}); actual != expected {
					t.Errorf("Expected second word to begin at %v but was %v", expected, actual)
				}
				if actual, expected := wordSpans[1].Max, (Coord{8, 4}); actual != expected {
					t.Errorf("Expected second word to end at %v but was %v", expected, actual)
				}
			}
		}
	})

	t.Run("does not count connected words which were unmodified by the play", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{5, 4}).Tile = &Tile{'D', 2}
		board.Position(Coord{6, 4}).Tile = &Tile{'O', 1}
		board.Position(Coord{7, 4}).Tile = &Tile{'G', 2}

		score, wordSpans, err := ScoreWords(TilePlacements{
			{Tile{'G', 2}, Coord{6, 3}},
			{Tile{'D', 2}, Coord{6, 5}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 5; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(wordSpans), 1; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := wordSpans[0].Min, (Coord{6, 3}); actual != expected {
					t.Errorf("Expected word to begin at %v but was %v", expected, actual)
				}
				if actual, expected := wordSpans[0].Max, (Coord{6, 5}); actual != expected {
					t.Errorf("Expected word to end at %v but was %v", expected, actual)
				}
			}
		}
	})
}
