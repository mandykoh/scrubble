package scrubble

import (
	"fmt"
	"testing"
)

func ExampleBoardWithLayout() {
	__, st, dl, dw, tl, tw := BoardPositionTypes()

	board := BoardWithLayout(BoardLayout{
		{tw, __, __, dl, __, __, __, tw, __, __, __, dl, __, __, tw},
		{__, dw, __, __, __, tl, __, __, __, tl, __, __, __, dw, __},
		{__, __, dw, __, __, __, dl, __, dl, __, __, __, dw, __, __},
		{dl, __, __, dw, __, __, __, dl, __, __, __, dw, __, __, dl},
		{__, __, __, __, dw, __, __, __, __, __, dw, __, __, __, __},
		{__, tl, __, __, __, tl, __, __, __, tl, __, __, __, tl, __},
		{__, __, dl, __, __, __, dl, __, dl, __, __, __, dl, __, __},
		{tw, __, __, dl, __, __, __, st, __, __, __, dl, __, __, tw},
		{__, __, dl, __, __, __, dl, __, dl, __, __, __, dl, __, __},
		{__, tl, __, __, __, tl, __, __, __, tl, __, __, __, tl, __},
		{__, __, __, __, dw, __, __, __, __, __, dw, __, __, __, __},
		{dl, __, __, dw, __, __, __, dl, __, __, __, dw, __, __, dl},
		{__, __, dw, __, __, __, dl, __, dl, __, __, __, dw, __, __},
		{__, dw, __, __, __, tl, __, __, __, tl, __, __, __, dw, __},
		{tw, __, __, dl, __, __, __, tw, __, __, __, dl, __, __, tw},
	})

	fmt.Printf("The board: %v", board)
}

func TestBoard(t *testing.T) {

	__, st, dl, dw, tl, tw := BoardPositionTypes()

	expectEmptyBoardWithLayout := func(t *testing.T, b Board, layout BoardLayout) {
		rows := len(layout)
		columns := layout.widestRow()

		if actual, expected := b.Rows, rows; actual != expected {
			t.Errorf("Expected board to have %d rows but found %d instead", expected, actual)
		}
		if actual, expected := b.Columns, columns; actual != expected {
			t.Errorf("Expected board to have %d columns but found %d instead", expected, actual)
		}

		if actual, expected := len(b.Positions), rows*columns; actual != expected {
			t.Fatalf("Expected a total of %d positions on the board but found %d instead", expected, actual)
		}

		for row, lRow := range layout {
			posType := PositionType(nil)

			for col := 0; col < columns; col++ {
				if col < len(lRow) {
					posType = lRow[col]
				}

				pos := b.Position(Coord{row, col})

				if tile := pos.Tile; tile != nil {
					t.Errorf("Expected position %d,%d to be empty but found tile %v", row, col, tile)
				}

				if actual, expected := pos.Type, posType; actual != expected {
					t.Errorf("Expected position %d,%d to be of '%s' type but was '%s'", row, col, expected.Name(), actual.Name())
				}
			}
		}
	}

	t.Run("BoardWithLayout()", func(t *testing.T) {

		t.Run("creates an empty board with the specified layout", func(t *testing.T) {
			layout := BoardLayout{
				{__, __, __, __, __, __, __},
				{__, __, __, st, __, __, __},
				{__, __, __, __, __, __, __},
			}

			board := BoardWithLayout(layout)

			expectEmptyBoardWithLayout(t, board, layout)
		})

		t.Run("always creates a rectangular board by filling out with empties to match the longest column", func(t *testing.T) {
			board := BoardWithLayout(BoardLayout{
				{__, __, __, __, __, __, __},
				{__, __, __, st},
				{},
			})

			expectEmptyBoardWithLayout(t, board, BoardLayout{
				{__, __, __, __, __, __, __},
				{__, __, __, st, __, __, __},
				{__, __, __, __, __, __, __},
			})
		})
	})

	t.Run("BoardWithStandardLayout()", func(t *testing.T) {

		t.Run("creates an empty board with a standardised layout", func(t *testing.T) {
			board := BoardWithStandardLayout()

			expectEmptyBoardWithLayout(t, board, BoardLayout{
				{tw, __, __, dl, __, __, __, tw, __, __, __, dl, __, __, tw},
				{__, dw, __, __, __, tl, __, __, __, tl, __, __, __, dw, __},
				{__, __, dw, __, __, __, dl, __, dl, __, __, __, dw, __, __},
				{dl, __, __, dw, __, __, __, dl, __, __, __, dw, __, __, dl},
				{__, __, __, __, dw, __, __, __, __, __, dw, __, __, __, __},
				{__, tl, __, __, __, tl, __, __, __, tl, __, __, __, tl, __},
				{__, __, dl, __, __, __, dl, __, dl, __, __, __, dl, __, __},
				{tw, __, __, dl, __, __, __, st, __, __, __, dl, __, __, tw},
				{__, __, dl, __, __, __, dl, __, dl, __, __, __, dl, __, __},
				{__, tl, __, __, __, tl, __, __, __, tl, __, __, __, tl, __},
				{__, __, __, __, dw, __, __, __, __, __, dw, __, __, __, __},
				{dl, __, __, dw, __, __, __, dl, __, __, __, dw, __, __, dl},
				{__, __, dw, __, __, __, dl, __, dl, __, __, __, dw, __, __},
				{__, dw, __, __, __, tl, __, __, __, tl, __, __, __, dw, __},
				{tw, __, __, dl, __, __, __, tw, __, __, __, dl, __, __, tw},
			})
		})
	})

	t.Run(".Neighbours()", func(t *testing.T) {

		b := Board{
			Rows:    3,
			Columns: 3,
			Positions: []BoardPosition{
				{__, &Tile{'A', 1}}, {__, &Tile{'B', 1}}, {__, &Tile{'C', 1}},
				{__, &Tile{'D', 1}}, {__, &Tile{'E', 1}}, {__, &Tile{'F', 1}},
				{__, &Tile{'G', 1}}, {__, &Tile{'H', 1}}, {__, &Tile{'I', 1}},
			},
		}

		t.Run("returns four cardinal neighbours of the specified position", func(t *testing.T) {
			neighbours := b.Neighbours(Coord{1, 1})

			if actual, expected := len(neighbours), 4; actual != expected {
				t.Errorf("Expected %d neighbours but got %d instead", expected, actual)
			} else {
				if actual, expected := neighbours[0], b.Position(Coord{0, 1}); actual != expected {
					t.Errorf("Expected north position to be returned but got %v", actual)
				}
				if actual, expected := neighbours[1], b.Position(Coord{2, 1}); actual != expected {
					t.Errorf("Expected south position to be returned but got %v", actual)
				}
				if actual, expected := neighbours[2], b.Position(Coord{1, 2}); actual != expected {
					t.Errorf("Expected east position to be returned but got %v", actual)
				}
				if actual, expected := neighbours[3], b.Position(Coord{1, 0}); actual != expected {
					t.Errorf("Expected west position to be returned but got %v", actual)
				}
			}
		})

		t.Run("omits neighbours that would be out of bounds", func(t *testing.T) {
			neighbours := b.Neighbours(Coord{0, 0})

			if actual, expected := len(neighbours), 4; actual != expected {
				t.Errorf("Expected %d neighbours but got %d instead", expected, actual)
			} else {
				if actual := neighbours[0]; actual != nil {
					t.Errorf("Expected north position to be nil but found tile %v", actual.Tile)
				}
				if actual, expected := neighbours[1], b.Position(Coord{1, 0}); actual != expected {
					t.Errorf("Expected south position to be returned but found tile %v", actual.Tile)
				}
				if actual, expected := neighbours[2], b.Position(Coord{0, 1}); actual != expected {
					t.Errorf("Expected east position to be returned but found tile %v", actual.Tile)
				}
				if actual := neighbours[3]; actual != nil {
					t.Errorf("Expected west position to be nil but found tile %v", actual.Tile)
				}
			}

			neighbours = b.Neighbours(Coord{2, 2})

			if actual, expected := len(neighbours), 4; actual != expected {
				t.Errorf("Expected %d neighbours but got %d instead", expected, actual)
			} else {
				if actual, expected := neighbours[0], b.Position(Coord{1, 2}); actual != expected {
					t.Errorf("Expected north position to be returned but found tile %v", actual.Tile)
				}
				if actual := neighbours[1]; actual != nil {
					t.Errorf("Expected south position to be nil but found tile %v", actual.Tile)
				}
				if actual := neighbours[2]; actual != nil {
					t.Errorf("Expected east position to be nil but found tile %v", actual.Tile)
				}
				if actual, expected := neighbours[3], b.Position(Coord{2, 1}); actual != expected {
					t.Errorf("Expected west position to be returned but found tile %v", actual.Tile)
				}
			}
		})
	})

	t.Run(".Position()", func(t *testing.T) {

		b := Board{
			Rows:    2,
			Columns: 2,
			Positions: []BoardPosition{
				{__, nil}, {st, nil},
				{dl, nil}, {tw, nil},
			},
		}

		t.Run("returns the specified position", func(t *testing.T) {
			if actual, expected := b.Position(Coord{0, 0}), &b.Positions[0]; actual != expected {
				t.Errorf("Expected 0,0 to correspond to position with '%s' type, but found %+v", expected.Type.Name(), actual)
			}

			if actual, expected := b.Position(Coord{0, 1}), &b.Positions[1]; actual != expected {
				t.Errorf("Expected 0,1 to correspond to position with '%s' type, but found %+v", expected.Type.Name(), actual)
			}

			if actual, expected := b.Position(Coord{1, 0}), &b.Positions[2]; actual != expected {
				t.Errorf("Expected 1,0 to correspond to position with '%s' type, but found %+v", expected.Type.Name(), actual)
			}

			if actual, expected := b.Position(Coord{1, 1}), &b.Positions[3]; actual != expected {
				t.Errorf("Expected 1,1 to correspond to position with '%s' type, but found %+v", expected.Type.Name(), actual)
			}
		})

		t.Run("returns nil when out of bounds", func(t *testing.T) {
			if actual, expected := b.Position(Coord{-1, 0}), (*BoardPosition)(nil); actual != expected {
				t.Errorf("Expected -1,0 to be out of bounds but got position %+v", actual)
			}
			if actual, expected := b.Position(Coord{0, -1}), (*BoardPosition)(nil); actual != expected {
				t.Errorf("Expected 0,-1 to be out of bounds but got position %+v", actual)
			}
			if actual, expected := b.Position(Coord{2, 0}), (*BoardPosition)(nil); actual != expected {
				t.Errorf("Expected 2,0 to be out of bounds but got position %+v", actual)
			}
			if actual, expected := b.Position(Coord{0, 2}), (*BoardPosition)(nil); actual != expected {
				t.Errorf("Expected 0,2 to be out of bounds but got position %+v", actual)
			}
		})
	})
}
