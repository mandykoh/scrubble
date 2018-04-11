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
						Score:      4,
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

	t.Run("counts entire horizontal word", func(t *testing.T) {
		board := setupBoard()

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'D', 2}, Coord{1, 2}},
			{Tile{'O', 1}, Coord{1, 3}},
			{Tile{'G', 2}, Coord{1, 4}},
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
				if actual, expected := words[0].Min, (Coord{1, 2}); actual != expected {
					t.Errorf("Expected word to begin at %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Max, (Coord{1, 4}); actual != expected {
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

	t.Run("counts entire vertical word", func(t *testing.T) {
		board := setupBoard()

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'D', 2}, Coord{2, 1}},
			{Tile{'O', 1}, Coord{3, 1}},
			{Tile{'G', 2}, Coord{4, 1}},
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
				if actual, expected := words[0].Min, (Coord{2, 1}); actual != expected {
					t.Errorf("Expected word to begin at %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Max, (Coord{4, 1}); actual != expected {
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

	t.Run("awards double-word score under a newly placed tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{3, 1}).Tile = &Tile{'D', 2}

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'O', 1}, Coord{3, 2}},
			{Tile{'G', 2}, Coord{3, 3}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 10; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(words), 1; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "DOG"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, 10; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
			}
		}
	})

	t.Run("awards double-word score for the start position", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{7, 5}).Tile = &Tile{'D', 2}

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'O', 1}, Coord{7, 6}},
			{Tile{'G', 2}, Coord{7, 7}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 10; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(words), 1; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "DOG"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, 10; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
			}
		}
	})

	t.Run("awards double-word score for each word formed", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{3, 1}).Tile = &Tile{'D', 2}
		board.Position(Coord{4, 3}).Tile = &Tile{'O', 1}
		board.Position(Coord{5, 3}).Tile = &Tile{'D', 2}

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'O', 1}, Coord{3, 2}},
			{Tile{'G', 2}, Coord{3, 3}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 20; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(words), 2; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "DOG"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, 10; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
				if actual, expected := words[1].Word, "GOD"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[1].Score, 10; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
			}
		}
	})

	t.Run("does not award double-word score under an existing tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{1, 1}).Tile = &Tile{'D', 2}

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'O', 1}, Coord{1, 2}},
			{Tile{'G', 2}, Coord{1, 3}},
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

	t.Run("awards double-word score only for the word its under", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{1, 1}).Tile = &Tile{'D', 2}
		board.Position(Coord{1, 2}).Tile = &Tile{'O', 1}

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'G', 2}, Coord{1, 3}},
			{Tile{'O', 1}, Coord{2, 3}},
			{Tile{'D', 2}, Coord{3, 3}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 15; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(words), 2; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "DOG"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, 5; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
				if actual, expected := words[1].Word, "GOD"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[1].Score, 10; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
			}
		}
	})

	t.Run("awards triple-word score under a newly placed tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{0, 12}).Tile = &Tile{'D', 2}

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'O', 1}, Coord{0, 13}},
			{Tile{'G', 2}, Coord{0, 14}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 15; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(words), 1; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "DOG"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, 15; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
			}
		}
	})

	t.Run("awards triple-word score for each word formed", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{0, 12}).Tile = &Tile{'D', 2}
		board.Position(Coord{1, 14}).Tile = &Tile{'O', 1}
		board.Position(Coord{2, 14}).Tile = &Tile{'D', 2}

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'O', 1}, Coord{0, 13}},
			{Tile{'G', 2}, Coord{0, 14}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 30; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(words), 2; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "DOG"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, 15; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
				if actual, expected := words[1].Word, "GOD"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[1].Score, 15; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
			}
		}
	})

	t.Run("does not award triple-word score under an existing tile", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{0, 0}).Tile = &Tile{'D', 2}

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'O', 1}, Coord{0, 1}},
			{Tile{'G', 2}, Coord{0, 2}},
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

	t.Run("awards triple-word score only for the word its under", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{13, 6}).Tile = &Tile{'D', 2}
		board.Position(Coord{13, 8}).Tile = &Tile{'G', 2}

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'G', 2}, Coord{12, 7}},
			{Tile{'O', 1}, Coord{13, 7}},
			{Tile{'D', 2}, Coord{14, 7}},
		}, board)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, 20; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(words), 2; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "DOG"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, 5; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
				if actual, expected := words[1].Word, "GOD"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[1].Score, 15; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
			}
		}
	})

	t.Run("awards stacked word score bonuses", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{0, 8}).Tile = &Tile{'S', 1}

		score, words, err := ScoreWords(TilePlacements{
			{Tile{'E', 1}, Coord{0, 0}},
			{Tile{'L', 1}, Coord{0, 1}},
			{Tile{'E', 1}, Coord{0, 2}},
			{Tile{'P', 3}, Coord{0, 3}},
			{Tile{'H', 4}, Coord{0, 4}},
			{Tile{'A', 1}, Coord{0, 5}},
			{Tile{'N', 1}, Coord{0, 6}},
			{Tile{'T', 1}, Coord{0, 7}},
		}, board)

		expectedScore := 3 * 3 * (1 + 1 + 1 + 2*3 + 4 + 1 + 1 + 1 + 1)

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			if actual, expected := score, expectedScore; actual != expected {
				t.Errorf("Expected a total score of %d but got %d", expected, actual)
			}

			if actual, expected := len(words), 1; actual != expected {
				t.Errorf("Expected one word formed but found %d", actual)
			} else {
				if actual, expected := words[0].Word, "ELEPHANTS"; actual != expected {
					t.Errorf("Expected formed word to be %v but was %v", expected, actual)
				}
				if actual, expected := words[0].Score, expectedScore; actual != expected {
					t.Errorf("Expected formed word to be worth %d points but was %d", expected, actual)
				}
			}
		}
	})
}
