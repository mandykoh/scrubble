package scrubble

import "testing"

func TestBoardLayout(t *testing.T) {

	t.Run(".BeginRow()", func(t *testing.T) {
		t.Run("starts a new, empty row in the layout", func(t *testing.T) {
			l := BoardLayout{}

			if actual, expected := len(l), 0; actual != expected {
				t.Errorf("Expected %d rows in the layout position but found %d", expected, actual)
			}

			l = l.BeginRow()

			if actual, expected := len(l), 1; actual != expected {
				t.Errorf("Expected %d rows in the layout position but found %d", expected, actual)
			}

			l = l.Em().BeginRow()

			if actual, expected := len(l), 2; actual != expected {
				t.Errorf("Expected %d rows in the layout position but found %d", expected, actual)
			}
			if actual, expected := len(l[0]), 1; actual != expected {
				t.Errorf("Expected first row to have one column but found %d columns", actual)
			}
			if actual, expected := len(l[1]), 0; actual != expected {
				t.Errorf("Expected second row to be empty but found %d columns", actual)
			}
		})
	})

	t.Run(".DL()", func(t *testing.T) {
		t.Run("adds a double-letter score position to the layout", func(t *testing.T) {
			l := BoardLayout{}.BeginRow().DL()

			if actual, expected := l[0][0], DoubleLetterScorePositionType; actual != expected {
				t.Errorf("Expected a '%s' layout position but found '%s'", expected.Name(), actual.Name())
			}
		})
	})

	t.Run(".DW()", func(t *testing.T) {
		t.Run("adds a double-word score position to the layout", func(t *testing.T) {
			l := BoardLayout{}.BeginRow().DW()

			if actual, expected := l[0][0], DoubleWordScorePositionType; actual != expected {
				t.Errorf("Expected a '%s' layout position but found '%s'", expected.Name(), actual.Name())
			}
		})
	})

	t.Run(".Em()", func(t *testing.T) {
		t.Run("adds a normal, non-special position to the layout", func(t *testing.T) {
			l := BoardLayout{}.BeginRow().Em()

			if actual, expected := l[0][0], NormalPositionType; actual != expected {
				t.Errorf("Expected a '%s' layout position but found '%s'", expected.Name(), actual.Name())
			}
		})
	})

	t.Run(".St()", func(t *testing.T) {
		t.Run("adds a starting position to the layout", func(t *testing.T) {
			l := BoardLayout{}.BeginRow().St()

			if actual, expected := l[0][0], StartPositionType; actual != expected {
				t.Errorf("Expected a '%s' layout position but found '%s'", expected.Name(), actual.Name())
			}
		})
	})

	t.Run(".TL()", func(t *testing.T) {
		t.Run("adds a triple-letter score position to the layout", func(t *testing.T) {
			l := BoardLayout{}.BeginRow().TL()

			if actual, expected := l[0][0], TripleLetterScorePositionType; actual != expected {
				t.Errorf("Expected a '%s' layout position but found '%s'", expected.Name(), actual.Name())
			}
		})
	})

	t.Run(".TW()", func(t *testing.T) {
		t.Run("adds a triple-word score position to the layout", func(t *testing.T) {
			l := BoardLayout{}.BeginRow().TW()

			if actual, expected := l[0][0], TripleWordScorePositionType; actual != expected {
				t.Errorf("Expected a '%s' layout position but found '%s'", expected.Name(), actual.Name())
			}
		})
	})

	t.Run(".widestRow()", func(t *testing.T) {

		t.Run("returns the number of columns in the widest row", func(t *testing.T) {
			l := BoardLayout{}

			if actual, expected := l.widestRow(), 0; actual != expected {
				t.Errorf("Expected empty layout to be %d columns wide but was %d", expected, actual)
			}

			l = l.BeginRow()

			if actual, expected := l.widestRow(), 0; actual != expected {
				t.Errorf("Expected degenerate layout with one row to be %d columns wide but was %d", expected, actual)
			}

			l = l.Em()

			if actual, expected := l.widestRow(), 1; actual != expected {
				t.Errorf("Expected layout to be %d columns wide but was %d", expected, actual)
			}

			l = l.BeginRow()

			if actual, expected := l.widestRow(), 1; actual != expected {
				t.Errorf("Expected layout to be %d columns wide but was %d", expected, actual)
			}

			l = l.Em().Em().St().Em()

			if actual, expected := l.widestRow(), 4; actual != expected {
				t.Errorf("Expected layout to be %d columns wide but was %d", expected, actual)
			}

			l = l.BeginRow().Em().Em()

			if actual, expected := l.widestRow(), 4; actual != expected {
				t.Errorf("Expected layout to be %d columns wide but was %d", expected, actual)
			}
		})
	})
}
