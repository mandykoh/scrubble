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

// Position returns the board position related to the specified row and column.
// Rows and columns are zero-indexed.
func (b *Board) Position(row, col int) *BoardPosition {
	return &b.Positions[row * b.Columns+ col]
}
