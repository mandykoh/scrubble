package scrubble

// Board represents a game board, which is a grid of positions on which tiles
// can be placed.
type Board struct {
	Rows      int
	Columns   int
	Positions []BoardPosition
}

// BoardWithLayout creates a board with no tiles, with the specified layout.
func BoardWithLayout(layout BoardLayout) Board {
	rows := len(layout)
	columns := layout.widestRow()

	b := Board{
		Rows:      rows,
		Columns:   columns,
		Positions: make([]BoardPosition, rows*columns),
	}

	for row, lRow := range layout {

		// Set position types according to the specified layout
		for col, posType := range lRow {
			b.Position(row, col).Type = posType
		}

		// Fill in any unspecified remainder of the row with Normal positions
		for col := len(lRow); col < columns; col++ {
			b.Position(row, col).Type = NormalPositionType
		}
	}

	return b
}

// BoardWithStandardLayout returns an empty Board with a standardised layout.
func BoardWithStandardLayout() Board {
	return BoardWithLayout(BoardLayout{}.
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
}

// Position returns the board position related to the specified row and column.
// Rows and columns are zero-indexed.
func (b *Board) Position(row, col int) *BoardPosition {
	return &b.Positions[row*b.Columns+col]
}
