package positiontype

// Layout represents a layout for creating a Board. Layouts are specified
// from the top row down, from the leftmost column to  the rightmost.
type Layout [][]Interface

// WidestRow returns the number of columns in the widest row of the layout.
func (l Layout) WidestRow() int {
	columns := 0
	for _, row := range l {
		cols := len(row)
		if cols > columns {
			columns = cols
		}
	}

	return columns
}
