package scrubble

import (
	"errors"
	"testing"
)

func TestCoordRange(t *testing.T) {

	t.Run(".EachCoord()", func(t *testing.T) {
		b := CoordRange{Coord{2, 4}, Coord{4, 6}}

		t.Run("visits each coordinate in the range", func(t *testing.T) {
			expectedCoords := []Coord{
				{2, 4}, {2, 5}, {2, 6},
				{3, 4}, {3, 5}, {3, 6},
				{4, 4}, {4, 5}, {4, 6},
			}

			offset := 0

			b.EachCoord(func(c Coord) error {
				if actual, expected := c, expectedCoords[offset]; actual != expected {
					t.Errorf("Expected to visit coordinate %v but got %v", expected, actual)
				}
				offset++
				return nil
			})

			if offset < len(expectedCoords) {
				t.Errorf("Expected to visit %d coordinates but only got %d", len(expectedCoords), offset)
			}
		})

		t.Run("stops when an error is returned and returns the error", func(t *testing.T) {
			expectedError := errors.New("some error")
			offset := 0

			err := b.EachCoord(func(c Coord) error {
				offset++
				if offset >= 3 {
					return expectedError
				}
				return nil
			})

			if actual, expected := offset, 3; actual != expected {
				t.Errorf("Expected to visit %d coordinates but got %d", expected, actual)
			}
			if err != expectedError {
				t.Errorf("Expected to get error %v but was %v", expectedError, err)
			}
		})
	})

	t.Run(".Include()", func(t *testing.T) {

		t.Run("returns expanded bounds for an outside point", func(t *testing.T) {
			b := CoordRange{Coord{2, 4}, Coord{3, 5}}

			included := b.Include(Coord{1, 7})

			if actual, expected := included, (CoordRange{Coord{1, 4}, Coord{3, 7}}); actual != expected {
				t.Errorf("Expected expanded bounds to be %v but got %v", expected, actual)
			}
		})

		t.Run("returns same bounds for an already included point", func(t *testing.T) {
			b := CoordRange{Coord{2, 4}, Coord{3, 7}}

			included := b.Include(Coord{2, 5})

			if actual, expected := included, b; actual != expected {
				t.Errorf("Expected bounds to not have changed but was %v", actual)
			}
		})
	})
}
