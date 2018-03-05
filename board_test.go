package scrubble

import (
	"testing"
)

func TestBoard(t *testing.T) {

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

				pos := b.Position(row, col)

				if tile := pos.Tile; tile != nil {
					t.Errorf("Expected position %d,%d to be empty but found tile %c(%d)", row, col, tile.Letter, tile.Points)
				}

				if actual, expected := pos.Type, posType; actual != expected {
					t.Errorf("Expected position %d,%d to be of '%s' type but was '%s'", row, col, expected.Name(), actual.Name())
				}
			}
		}
	}

	t.Run("BoardWithLayout()", func(t *testing.T) {

		t.Run("creates an empty board with the specified layout", func(t *testing.T) {
			layout := BoardLayout{}.
				BeginRow().Em().Em().Em().Em().Em().Em().Em().
				BeginRow().Em().Em().Em().St().Em().Em().Em().
				BeginRow().Em().Em().Em().Em().Em().Em().Em()

			board := BoardWithLayout(layout)

			expectEmptyBoardWithLayout(t, board, layout)
		})

		t.Run("always creates a rectangular board by filling out with empties to match the longest column", func(t *testing.T) {
			board := BoardWithLayout(BoardLayout{}.
				BeginRow().Em().Em().Em().Em().Em().Em().Em().
				BeginRow().Em().Em().Em().St().
				BeginRow())

			expectEmptyBoardWithLayout(t, board, BoardLayout{}.
				BeginRow().Em().Em().Em().Em().Em().Em().Em().
				BeginRow().Em().Em().Em().St().Em().Em().Em().
				BeginRow().Em().Em().Em().Em().Em().Em().Em())
		})
	})

	t.Run("BoardWithStandardLayout()", func(t *testing.T) {

		t.Run("creates an empty board with a standardised layout", func(t *testing.T) {
			board := BoardWithStandardLayout()

			expectEmptyBoardWithLayout(t, board, BoardLayout{}.
				BeginRow().TW().Em().Em().DL().Em().Em().Em().TW().Em().Em().Em().DL().Em().Em().TW().
				BeginRow().Em().DW().Em().Em().Em().TL().Em().Em().Em().TL().Em().Em().Em().DW().Em().
				BeginRow().Em().Em().DW().Em().Em().Em().DL().Em().DL().Em().Em().Em().DW().Em().Em().
				BeginRow().DL().Em().Em().DW().Em().Em().Em().DL().Em().Em().Em().DW().Em().Em().DL().
				BeginRow().Em().Em().Em().Em().DW().Em().Em().Em().Em().Em().DW().Em().Em().Em().Em().
				BeginRow().Em().TL().Em().Em().Em().TL().Em().Em().Em().TL().Em().Em().Em().TL().Em().
				BeginRow().Em().Em().DL().Em().Em().Em().DL().Em().DL().Em().Em().Em().DL().Em().Em().
				BeginRow().TW().Em().Em().DL().Em().Em().Em().St().Em().Em().Em().DL().Em().Em().TW().
				BeginRow().Em().Em().DL().Em().Em().Em().DL().Em().DL().Em().Em().Em().DL().Em().Em().
				BeginRow().Em().TL().Em().Em().Em().TL().Em().Em().Em().TL().Em().Em().Em().TL().Em().
				BeginRow().Em().Em().Em().Em().DW().Em().Em().Em().Em().Em().DW().Em().Em().Em().Em().
				BeginRow().DL().Em().Em().DW().Em().Em().Em().DL().Em().Em().Em().DW().Em().Em().DL().
				BeginRow().Em().Em().DW().Em().Em().Em().DL().Em().DL().Em().Em().Em().DW().Em().Em().
				BeginRow().Em().DW().Em().Em().Em().TL().Em().Em().Em().TL().Em().Em().Em().DW().Em().
				BeginRow().TW().Em().Em().DL().Em().Em().Em().TW().Em().Em().Em().DL().Em().Em().TW())
		})
	})

	t.Run(".Position()", func(t *testing.T) {

		t.Run("returns the specified position", func(t *testing.T) {

			b := Board{
				Rows:    2,
				Columns: 2,
				Positions: []BoardPosition{
					{NormalPositionType, nil}, {StartPositionType, nil},
					{DoubleLetterScorePositionType, nil}, {TripleWordScorePositionType, nil},
				},
			}

			if actual, expected := b.Position(0, 0), &b.Positions[0]; actual != expected {
				t.Errorf("Expected 0,0 to correspond to position with '%s' type, but found %+v", expected.Type.Name(), actual)
			}

			if actual, expected := b.Position(0, 1), &b.Positions[1]; actual != expected {
				t.Errorf("Expected 0,1 to correspond to position with '%s' type, but found %+v", expected.Type.Name(), actual)
			}

			if actual, expected := b.Position(1, 0), &b.Positions[2]; actual != expected {
				t.Errorf("Expected 1,0 to correspond to position with '%s' type, but found %+v", expected.Type.Name(), actual)
			}

			if actual, expected := b.Position(1, 1), &b.Positions[3]; actual != expected {
				t.Errorf("Expected 1,1 to correspond to position with '%s' type, but found %+v", expected.Type.Name(), actual)
			}
		})
	})
}
