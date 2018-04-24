package board

import (
	"testing"
)

func TestLayout(t *testing.T) {

	__, st, _, _, _, _ := AllPositionTypes()

	t.Run(".WidestRow()", func(t *testing.T) {

		t.Run("returns the number of columns in the widest row", func(t *testing.T) {
			l := Layout{}

			if actual, expected := l.WidestRow(), 0; actual != expected {
				t.Errorf("Expected empty layout to be %d columns wide but was %d", expected, actual)
			}

			l = Layout{
				{},
			}

			if actual, expected := l.WidestRow(), 0; actual != expected {
				t.Errorf("Expected degenerate layout with one row to be %d columns wide but was %d", expected, actual)
			}

			l = Layout{
				{__},
			}

			if actual, expected := l.WidestRow(), 1; actual != expected {
				t.Errorf("Expected layout to be %d columns wide but was %d", expected, actual)
			}

			l = Layout{
				{__},
				{},
			}

			if actual, expected := l.WidestRow(), 1; actual != expected {
				t.Errorf("Expected layout to be %d columns wide but was %d", expected, actual)
			}

			l = Layout{
				{__},
				{__, __, st, __},
			}

			if actual, expected := l.WidestRow(), 4; actual != expected {
				t.Errorf("Expected layout to be %d columns wide but was %d", expected, actual)
			}

			l = Layout{
				{__},
				{__, __, st, __},
				{__, __},
			}

			if actual, expected := l.WidestRow(), 4; actual != expected {
				t.Errorf("Expected layout to be %d columns wide but was %d", expected, actual)
			}
		})
	})
}
