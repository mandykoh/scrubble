package scrubble

// Board represents a game board, which is a grid of positions on which tiles
// can be placed. The zero-value of a Board is a zero-sized board.
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
			b.Position(row, col).Type = normalPositionTypeInstance
		}
	}

	return b
}

// BoardWithStandardLayout returns an empty Board with a standardised layout.
func BoardWithStandardLayout() Board {
	__, st, dl, dw, tl, tw := BoardPositionTypes()

	return BoardWithLayout(BoardLayout{
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
}

// Position returns the board position related to the specified row and column.
// Rows and columns are zero-indexed.
func (b *Board) Position(row, col int) *BoardPosition {
	return &b.Positions[row*b.Columns+col]
}
