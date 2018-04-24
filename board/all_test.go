package board

import (
	"testing"
)

func TestAllPositionTypes(t *testing.T) {

	__, st, dl, dw, tl, tw := AllPositionTypes()

	t.Run("returns the correct set of board position types", func(t *testing.T) {
		cases := []struct {
			actual, expected PositionType
		}{
			{__, normalInstance},
			{st, startInstance},
			{dl, doubleLetterScoreInstance},
			{dw, doubleWordScoreInstance},
			{tl, tripleLetterScoreInstance},
			{tw, tripleWordScoreInstance},
		}

		for _, c := range cases {
			if c.actual != c.expected {
				t.Errorf("Expected '%s' position type but got '%s' instead", c.expected.Name(), c.actual.Name())
			}
		}
	})
}
