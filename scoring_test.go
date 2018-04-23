package scrubble

import (
	"testing"

	"github.com/mandykoh/scrubble/tile"
)

func TestScoreWords(t *testing.T) {
	setupBoard := func() *Board {
		board := BoardWithStandardLayout()
		return &board
	}

	dictionary := func(word string) (valid bool) {
		return true
	}

	expectFormedWords := func(t *testing.T, actualWords []PlayedWord, expectedWords ...PlayedWord) {
		if actual, expected := len(actualWords), len(expectedWords); actual != expected {
			t.Errorf("Expected %d word(s) formed but found %d", expected, actual)
		} else {
			for i := range expectedWords {
				if actual, expected := actualWords[i], expectedWords[i]; actual != expected {
					t.Errorf("Expected formed word %v but found %v", expected, actual)
				}
			}
		}
	}

	t.Run("returns an error for single-letter words", func(t *testing.T) {
		board := setupBoard()

		_, _, err := ScoreWords(TilePlacements{
			{tile.Make('A', 2), Coord{7, 7}},
		}, board, dictionary)

		if err == nil {
			t.Errorf("Expected an error for single-letter word but got nil")
		} else {
			switch e := err.(type) {
			case InvalidWordError:
				expectFormedWords(t, e.Words, PlayedWord{"A", 4, CoordRange{Coord{7, 7}, Coord{7, 7}}})

			default:
				t.Errorf("Expected InvalidWordError but got %v", e)
			}
		}
	})

	t.Run("checks each word against dictionary for validity", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{0, 3}).Tile = &tile.Tile{Letter: 'G', Points: 2}
		board.Position(Coord{2, 3}).Tile = &tile.Tile{Letter: 'D', Points: 2}

		var wordsLookedUp []string

		dictionary := func(word string) bool {
			wordsLookedUp = append(wordsLookedUp, word)
			return false
		}

		_, _, err := ScoreWords(TilePlacements{
			{tile.Make('D', 2), Coord{1, 2}},
			{tile.Make('O', 1), Coord{1, 3}},
			{tile.Make('G', 2), Coord{1, 4}},
		}, board, dictionary)

		if err == nil {
			t.Errorf("Expected an error but succeeded")
		} else {
			switch e := err.(type) {
			case InvalidWordError:
				if actual, expected := len(wordsLookedUp), 2; actual != expected {
					t.Errorf("Expected %d words to be validated against dictionary but only got %d", expected, actual)
				} else {
					if actual, expected := wordsLookedUp[0], "DOG"; actual != expected {
						t.Errorf("Expected '%s' to be validated against dictionary but got '%s'", expected, actual)
					}
					if actual, expected := wordsLookedUp[1], "GOD"; actual != expected {
						t.Errorf("Expected '%s' to be validated against dictionary but got '%s'", expected, actual)
					}
				}
				expectFormedWords(t, e.Words,
					PlayedWord{"DOG", 5, CoordRange{Coord{1, 2}, Coord{1, 4}}},
					PlayedWord{"GOD", 5, CoordRange{Coord{0, 3}, Coord{2, 3}}})

			default:
				t.Errorf("Expected InvalidWordError but got %v", e)
			}
		}
	})

	t.Run("counts entire horizontal word", func(t *testing.T) {
		board := setupBoard()

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('D', 2), Coord{1, 2}},
			{tile.Make('O', 1), Coord{1, 3}},
			{tile.Make('G', 2), Coord{1, 4}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 5; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words, PlayedWord{"DOG", 5, CoordRange{Coord{1, 2}, Coord{1, 4}}})
		}
	})

	t.Run("counts entire horizontal word connected to existing tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{2, 3}).Tile = &tile.Tile{Letter: 'D', Points: 2}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('O', 1), Coord{2, 4}},
			{tile.Make('G', 2), Coord{2, 5}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 5; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words, PlayedWord{"DOG", 5, CoordRange{Coord{2, 3}, Coord{2, 5}}})
		}
	})

	t.Run("counts entire vertical word", func(t *testing.T) {
		board := setupBoard()

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('D', 2), Coord{2, 1}},
			{tile.Make('O', 1), Coord{3, 1}},
			{tile.Make('G', 2), Coord{4, 1}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 5; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words, PlayedWord{"DOG", 5, CoordRange{Coord{2, 1}, Coord{4, 1}}})
		}
	})

	t.Run("counts entire vertical word connected to existing tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{1, 4}).Tile = &tile.Tile{Letter: 'D', Points: 2}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('O', 1), Coord{2, 4}},
			{tile.Make('G', 2), Coord{3, 4}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 5; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words, PlayedWord{"DOG", 5, CoordRange{Coord{1, 4}, Coord{3, 4}}})
		}
	})

	t.Run("counts hooked words connected to the main word", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{5, 4}).Tile = &tile.Tile{Letter: 'D', Points: 2}
		board.Position(Coord{6, 4}).Tile = &tile.Tile{Letter: 'O', Points: 1}
		board.Position(Coord{7, 4}).Tile = &tile.Tile{Letter: 'G', Points: 2}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('S', 2), Coord{8, 4}},
			{tile.Make('O', 2), Coord{8, 5}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 11; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words,
				PlayedWord{"SO", 4, CoordRange{Coord{8, 4}, Coord{8, 5}}},
				PlayedWord{"DOGS", 7, CoordRange{Coord{5, 4}, Coord{8, 4}}})
		}
	})

	t.Run("does not count connected words which were unmodified by the play", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{5, 4}).Tile = &tile.Tile{Letter: 'D', Points: 2}
		board.Position(Coord{6, 4}).Tile = &tile.Tile{Letter: 'O', Points: 1}
		board.Position(Coord{7, 4}).Tile = &tile.Tile{Letter: 'G', Points: 2}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('G', 2), Coord{6, 3}},
			{tile.Make('D', 2), Coord{6, 5}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 5; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words, PlayedWord{"GOD", 5, CoordRange{Coord{6, 3}, Coord{6, 5}}})
		}
	})

	t.Run("awards double-letter score under a newly placed tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{2, 4}).Tile = &tile.Tile{Letter: 'D', Points: 2}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('O', 1), Coord{2, 5}},
			{tile.Make('G', 2), Coord{2, 6}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 7; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words, PlayedWord{"DOG", 7, CoordRange{Coord{2, 4}, Coord{2, 6}}})
		}
	})

	t.Run("awards double-letter score for each word formed", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{2, 4}).Tile = &tile.Tile{Letter: 'D', Points: 2}
		board.Position(Coord{3, 6}).Tile = &tile.Tile{Letter: 'O', Points: 1}
		board.Position(Coord{4, 6}).Tile = &tile.Tile{Letter: 'D', Points: 2}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('O', 1), Coord{2, 5}},
			{tile.Make('G', 2), Coord{2, 6}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 14; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words,
				PlayedWord{"DOG", 7, CoordRange{Coord{2, 4}, Coord{2, 6}}},
				PlayedWord{"GOD", 7, CoordRange{Coord{2, 6}, Coord{4, 6}}})
		}
	})

	t.Run("does not award double-letter score under an existing tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{2, 8}).Tile = &tile.Tile{Letter: 'D', Points: 2}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('O', 1), Coord{2, 9}},
			{tile.Make('G', 2), Coord{2, 10}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 5; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words, PlayedWord{"DOG", 5, CoordRange{Coord{2, 8}, Coord{2, 10}}})
		}
	})

	t.Run("awards triple-letter score under a newly placed tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{1, 3}).Tile = &tile.Tile{Letter: 'D', Points: 2}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('O', 1), Coord{1, 4}},
			{tile.Make('G', 2), Coord{1, 5}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 9; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words, PlayedWord{"DOG", 9, CoordRange{Coord{1, 3}, Coord{1, 5}}})
		}
	})

	t.Run("awards triple-letter score for each word formed", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{1, 3}).Tile = &tile.Tile{Letter: 'D', Points: 2}
		board.Position(Coord{2, 5}).Tile = &tile.Tile{Letter: 'O', Points: 1}
		board.Position(Coord{3, 5}).Tile = &tile.Tile{Letter: 'D', Points: 2}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('O', 1), Coord{1, 4}},
			{tile.Make('G', 2), Coord{1, 5}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 18; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words,
				PlayedWord{"DOG", 9, CoordRange{Coord{1, 3}, Coord{1, 5}}},
				PlayedWord{"GOD", 9, CoordRange{Coord{1, 5}, Coord{3, 5}}})
		}
	})

	t.Run("does not award triple-letter score under an existing tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{1, 9}).Tile = &tile.Tile{Letter: 'D', Points: 2}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('O', 1), Coord{1, 10}},
			{tile.Make('G', 2), Coord{1, 11}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 5; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words, PlayedWord{"DOG", 5, CoordRange{Coord{1, 9}, Coord{1, 11}}})
		}
	})

	t.Run("awards double-word score under a newly placed tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{3, 1}).Tile = &tile.Tile{Letter: 'D', Points: 2}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('O', 1), Coord{3, 2}},
			{tile.Make('G', 2), Coord{3, 3}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 10; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words, PlayedWord{"DOG", 10, CoordRange{Coord{3, 1}, Coord{3, 3}}})
		}
	})

	t.Run("awards double-word score for the start position", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{7, 5}).Tile = &tile.Tile{Letter: 'D', Points: 2}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('O', 1), Coord{7, 6}},
			{tile.Make('G', 2), Coord{7, 7}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 10; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words, PlayedWord{"DOG", 10, CoordRange{Coord{7, 5}, Coord{7, 7}}})
		}
	})

	t.Run("awards double-word score for each word formed", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{3, 1}).Tile = &tile.Tile{Letter: 'D', Points: 2}
		board.Position(Coord{4, 3}).Tile = &tile.Tile{Letter: 'O', Points: 1}
		board.Position(Coord{5, 3}).Tile = &tile.Tile{Letter: 'D', Points: 2}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('O', 1), Coord{3, 2}},
			{tile.Make('G', 2), Coord{3, 3}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 20; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words,
				PlayedWord{"DOG", 10, CoordRange{Coord{3, 1}, Coord{3, 3}}},
				PlayedWord{"GOD", 10, CoordRange{Coord{3, 3}, Coord{5, 3}}})
		}
	})

	t.Run("does not award double-word score under an existing tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{1, 1}).Tile = &tile.Tile{Letter: 'D', Points: 2}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('O', 1), Coord{1, 2}},
			{tile.Make('G', 2), Coord{1, 3}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 5; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words, PlayedWord{"DOG", 5, CoordRange{Coord{1, 1}, Coord{1, 3}}})
		}
	})

	t.Run("awards double-word score only for the word its under", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{1, 1}).Tile = &tile.Tile{Letter: 'D', Points: 2}
		board.Position(Coord{1, 2}).Tile = &tile.Tile{Letter: 'O', Points: 1}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('G', 2), Coord{1, 3}},
			{tile.Make('O', 1), Coord{2, 3}},
			{tile.Make('D', 2), Coord{3, 3}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 15; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words,
				PlayedWord{"DOG", 5, CoordRange{Coord{1, 1}, Coord{1, 3}}},
				PlayedWord{"GOD", 10, CoordRange{Coord{1, 3}, Coord{3, 3}}})
		}
	})

	t.Run("awards triple-word score under a newly placed tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{0, 12}).Tile = &tile.Tile{Letter: 'D', Points: 2}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('O', 1), Coord{0, 13}},
			{tile.Make('G', 2), Coord{0, 14}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 15; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words, PlayedWord{"DOG", 15, CoordRange{Coord{0, 12}, Coord{0, 14}}})
		}
	})

	t.Run("awards triple-word score for each word formed", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{0, 12}).Tile = &tile.Tile{Letter: 'D', Points: 2}
		board.Position(Coord{1, 14}).Tile = &tile.Tile{Letter: 'O', Points: 1}
		board.Position(Coord{2, 14}).Tile = &tile.Tile{Letter: 'D', Points: 2}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('O', 1), Coord{0, 13}},
			{tile.Make('G', 2), Coord{0, 14}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 30; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words,
				PlayedWord{"DOG", 15, CoordRange{Coord{0, 12}, Coord{0, 14}}},
				PlayedWord{"GOD", 15, CoordRange{Coord{0, 14}, Coord{2, 14}}})
		}
	})

	t.Run("does not award triple-word score under an existing tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{0, 0}).Tile = &tile.Tile{Letter: 'D', Points: 2}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('O', 1), Coord{0, 1}},
			{tile.Make('G', 2), Coord{0, 2}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 5; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words, PlayedWord{"DOG", 5, CoordRange{Coord{0, 0}, Coord{0, 2}}})
		}
	})

	t.Run("awards triple-word score only for the word its under", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{13, 6}).Tile = &tile.Tile{Letter: 'D', Points: 2}
		board.Position(Coord{13, 8}).Tile = &tile.Tile{Letter: 'G', Points: 2}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('G', 2), Coord{12, 7}},
			{tile.Make('O', 1), Coord{13, 7}},
			{tile.Make('D', 2), Coord{14, 7}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 20; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words,
				PlayedWord{"DOG", 5, CoordRange{Coord{13, 6}, Coord{13, 8}}},
				PlayedWord{"GOD", 15, CoordRange{Coord{12, 7}, Coord{14, 7}}})
		}
	})

	t.Run("awards an extra point bonus if a full rack's worth of tiles is played", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{2, 6}).Tile = &tile.Tile{Letter: 'P', Points: 3}
		board.Position(Coord{2, 8}).Tile = &tile.Tile{Letter: 'A', Points: 1}
		board.Position(Coord{3, 3}).Tile = &tile.Tile{Letter: 'E', Points: 1}
		board.Position(Coord{4, 3}).Tile = &tile.Tile{Letter: 'L', Points: 1}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('E', 1), Coord{2, 3}},
			{tile.Make('L', 1), Coord{2, 4}},
			{tile.Make('E', 1), Coord{2, 5}},
			{tile.Make('H', 4), Coord{2, 7}},
			{tile.Make('N', 1), Coord{2, 9}},
			{tile.Make('T', 1), Coord{2, 10}},
			{tile.Make('S', 1), Coord{2, 11}},
		}, board, dictionary)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 14+3+MaxRackTilesBonus; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words,
				PlayedWord{"ELEPHANTS", 14, CoordRange{Coord{2, 3}, Coord{2, 11}}},
				PlayedWord{"EEL", 3, CoordRange{Coord{2, 3}, Coord{4, 3}}})
		}
	})

	t.Run("awards stacked word score bonuses", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{0, 8}).Tile = &tile.Tile{Letter: 'S', Points: 1}

		score, words, err := ScoreWords(TilePlacements{
			{tile.Make('E', 1), Coord{0, 0}},
			{tile.Make('L', 1), Coord{0, 1}},
			{tile.Make('E', 1), Coord{0, 2}},
			{tile.Make('P', 3), Coord{0, 3}},
			{tile.Make('H', 4), Coord{0, 4}},
			{tile.Make('A', 1), Coord{0, 5}},
			{tile.Make('N', 1), Coord{0, 6}},
			{tile.Make('T', 1), Coord{0, 7}},
		}, board, dictionary)

		expectedWordScore := 3 * 3 * (1 + 1 + 1 + 2*3 + 4 + 1 + 1 + 1 + 1)
		expectedTotalScore := expectedWordScore + MaxRackTilesBonus

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, expectedTotalScore; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}
			expectFormedWords(t, words, PlayedWord{"ELEPHANTS", expectedWordScore, CoordRange{Coord{0, 0}, Coord{0, 8}}})
		}
	})
}
