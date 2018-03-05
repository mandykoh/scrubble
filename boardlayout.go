package scrubble

// BoardLayout represents a layout for creating a Board. Layouts are specified
// from the top row down, from the leftmost column to  the rightmost.
type BoardLayout [][]PositionType

func (b BoardLayout) widestRow() int {
	columns := 0
	for _, row := range b {
		cols := len(row)
		if cols > columns {
			columns = cols
		}
	}

	return columns
}
