package scrubble

import (
	"testing"
)

func TestBoardLayout(t *testing.T) {

	t.Run("LayoutPositionTypes()", func(t *testing.T) {

		t.Run("returns the correct set of board position types", func(t *testing.T) {
			__, st, dl, dw, tl, tw := LayoutPositionTypes()

			cases := []struct {
				actual, expected PositionType
			}{
				{__, normalPositionTypeInstance},
				{st, startPositionTypeInstance},
				{dl, doubleLetterScorePositionTypeInstance},
				{dw, doubleWordScorePositionTypeInstance},
				{tl, tripleLetterScorePositionTypeInstance},
				{tw, tripleWordScorePositionTypeInstance},
			}

			for _, c := range cases {
				if c.actual != c.expected {
					t.Errorf("Expected '%s' position type but got '%s' instead", c.expected.Name(), c.actual.Name())
				}
			}
		})
	})

	t.Run(".widestRow()", func(t *testing.T) {

		t.Run("returns the number of columns in the widest row", func(t *testing.T) {
			__, st, _, _, _, _ := LayoutPositionTypes()

			l := BoardLayout{}

			if actual, expected := l.widestRow(), 0; actual != expected {
				t.Errorf("Expected empty layout to be %d columns wide but was %d", expected, actual)
			}

			l = BoardLayout{
				{},
			}

			if actual, expected := l.widestRow(), 0; actual != expected {
				t.Errorf("Expected degenerate layout with one row to be %d columns wide but was %d", expected, actual)
			}

			l = BoardLayout{
				{__},
			}

			if actual, expected := l.widestRow(), 1; actual != expected {
				t.Errorf("Expected layout to be %d columns wide but was %d", expected, actual)
			}

			l = BoardLayout{
				{__},
				{},
			}

			if actual, expected := l.widestRow(), 1; actual != expected {
				t.Errorf("Expected layout to be %d columns wide but was %d", expected, actual)
			}

			l = BoardLayout{
				{__},
				{__, __, st, __},
			}

			if actual, expected := l.widestRow(), 4; actual != expected {
				t.Errorf("Expected layout to be %d columns wide but was %d", expected, actual)
			}

			l = BoardLayout{
				{__},
				{__, __, st, __},
				{__, __},
			}

			if actual, expected := l.widestRow(), 4; actual != expected {
				t.Errorf("Expected layout to be %d columns wide but was %d", expected, actual)
			}
		})
	})
}
