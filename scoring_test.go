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
			switch e := err.(type) {
			case InvalidWordError:
				if actual, expected := len(e.Words), 1; actual != expected {
					t.Errorf("Expected one word to be marked as invalid but found %d", actual)
				} else {
					expected := PlayedWord{
						Word:       "A",
						Score:      2,
						CoordRange: CoordRange{Coord{7, 7}, Coord{7, 7}},
					}
					if actual := e.Words[0]; actual != expected {
						t.Errorf("Expected invalid played word %#v but was %#v", expected, actual)
					}
				}

			default:
				t.Errorf("Expected InvalidWordError but got %v", e)
			}
		}
	})

	t.Run("counts entire horizontal word placed on starting position", func(t *testing.T) {
		board := setupBoard()

		score, words, err := ScoreWords(TilePlacements{
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

			if actual, expected := len(words), 1; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "DOG"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, 5; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
				if actual, expected := words[0].Min, (Coord{7, 6}); actual != expected {
					t.Errorf("Expected word to begin at %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Max, (Coord{7, 8}); actual != expected {
					t.Errorf("Expected word to end at %v but was %v", expected, actual)
				}
			}
		}
	})

	t.Run("counts entire horizontal word connected to existing tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{2, 3}).Tile = &Tile{'D', 2}

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'O', 1}, Coord{2, 4}},
			{Tile{'G', 2}, Coord{2, 5}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 5; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(words), 1; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "DOG"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, 5; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
				if actual, expected := words[0].Min, (Coord{2, 3}); actual != expected {
					t.Errorf("Expected word to begin at %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Max, (Coord{2, 5}); actual != expected {
					t.Errorf("Expected word to end at %v but was %v", expected, actual)
				}
			}
		}
	})

	t.Run("counts entire vertical word placed on starting position", func(t *testing.T) {
		board := setupBoard()

		score, words, err := ScoreWords(TilePlacements{
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

			if actual, expected := len(words), 1; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "DOG"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, 5; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
				if actual, expected := words[0].Min, (Coord{6, 7}); actual != expected {
					t.Errorf("Expected word to begin at %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Max, (Coord{8, 7}); actual != expected {
					t.Errorf("Expected word to end at %v but was %v", expected, actual)
				}
			}
		}
	})

	t.Run("counts entire vertical word connected to existing tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{1, 4}).Tile = &Tile{'D', 2}

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'O', 1}, Coord{2, 4}},
			{Tile{'G', 2}, Coord{3, 4}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 5; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(words), 1; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "DOG"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, 5; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
				if actual, expected := words[0].Min, (Coord{1, 4}); actual != expected {
					t.Errorf("Expected word to begin at %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Max, (Coord{3, 4}); actual != expected {
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

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'S', 2}, Coord{8, 4}},
			{Tile{'O', 2}, Coord{8, 5}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 11; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(words), 2; actual != expected {
				t.Errorf("Expected two words formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "SO"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, 4; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
				if actual, expected := words[0].Min, (Coord{8, 4}); actual != expected {
					t.Errorf("Expected first word to begin at %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Max, (Coord{8, 5}); actual != expected {
					t.Errorf("Expected first word to end at %v but was %v", expected, actual)
				}

				if actual, expected := words[1].Word, "DOGS"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[1].Score, 7; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
				if actual, expected := words[1].Min, (Coord{5, 4}); actual != expected {
					t.Errorf("Expected second word to begin at %v but was %v", expected, actual)
				}
				if actual, expected := words[1].Max, (Coord{8, 4}); actual != expected {
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

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'G', 2}, Coord{6, 3}},
			{Tile{'D', 2}, Coord{6, 5}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 5; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(words), 1; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "GOD"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, 5; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
				if actual, expected := words[0].Min, (Coord{6, 3}); actual != expected {
					t.Errorf("Expected word to begin at %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Max, (Coord{6, 5}); actual != expected {
					t.Errorf("Expected word to end at %v but was %v", expected, actual)
				}
			}
		}
	})

	t.Run("awards double-letter score under a newly placed tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{2, 4}).Tile = &Tile{'D', 2}

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'O', 1}, Coord{2, 5}},
			{Tile{'G', 2}, Coord{2, 6}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 7; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(words), 1; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "DOG"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, 7; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
			}
		}
	})

	t.Run("awards double-letter score for each word formed", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{2, 4}).Tile = &Tile{'D', 2}
		board.Position(Coord{3, 6}).Tile = &Tile{'O', 1}
		board.Position(Coord{4, 6}).Tile = &Tile{'D', 2}

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'O', 1}, Coord{2, 5}},
			{Tile{'G', 2}, Coord{2, 6}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 14; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(words), 2; actual != expected {
				t.Errorf("Expected two words formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "DOG"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, 7; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
				if actual, expected := words[1].Word, "GOD"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[1].Score, 7; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
			}
		}
	})

	t.Run("does not award double-letter score under an existing tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{2, 8}).Tile = &Tile{'D', 2}

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'O', 1}, Coord{2, 9}},
			{Tile{'G', 2}, Coord{2, 10}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 5; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(words), 1; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "DOG"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, 5; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
			}
		}
	})

	t.Run("awards triple-letter score under a newly placed tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{1, 3}).Tile = &Tile{'D', 2}

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'O', 1}, Coord{1, 4}},
			{Tile{'G', 2}, Coord{1, 5}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 9; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(words), 1; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "DOG"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, 9; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
			}
		}
	})

	t.Run("awards triple-letter score for each word formed", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{1, 3}).Tile = &Tile{'D', 2}
		board.Position(Coord{2, 5}).Tile = &Tile{'O', 1}
		board.Position(Coord{3, 5}).Tile = &Tile{'D', 2}

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'O', 1}, Coord{1, 4}},
			{Tile{'G', 2}, Coord{1, 5}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 18; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(words), 2; actual != expected {
				t.Errorf("Expected two words formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "DOG"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, 9; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
				if actual, expected := words[1].Word, "GOD"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[1].Score, 9; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
			}
		}
	})

	t.Run("does not award triple-letter score under an existing tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{1, 9}).Tile = &Tile{'D', 2}

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'O', 1}, Coord{1, 10}},
			{Tile{'G', 2}, Coord{1, 11}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 5; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(words), 1; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "DOG"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, 5; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
			}
		}
	})
}
