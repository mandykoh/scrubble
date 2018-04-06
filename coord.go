package scrubble

// Coord represents the coordinate of a position on the board.
//
// Coordinate rows and columns are zero-indexed, with the first row and column
// being the upper left.
type Coord struct {
	Row    int
	Column int
}

// East returns the coordinate which is one step to the east.
func (c Coord) East() Coord {
	return Coord{c.Row, c.Column + 1}
}

// Max returns the component-wise maximum of this coordinate and the other.
func (c Coord) Max(other Coord) (max Coord) {
	if c.Row > other.Row {
		max.Row = c.Row
	} else {
		max.Row = other.Row
	}

	if c.Column > other.Column {
		max.Column = c.Column
	} else {
		max.Column = other.Column
	}

	return
}

// Min returns the component-wise minimum of this coordinate and the other.
func (c Coord) Min(other Coord) (min Coord) {
	if c.Row < other.Row {
		min.Row = c.Row
	} else {
		min.Row = other.Row
	}

	if c.Column < other.Column {
		min.Column = c.Column
	} else {
		min.Column = other.Column
	}

	return
}

// North returns the coordinate which is one step to the north.
func (c Coord) North() Coord {
	return Coord{c.Row - 1, c.Column}
}

// South returns the coordinate which is one step to the south.
func (c Coord) South() Coord {
	return Coord{c.Row + 1, c.Column}
}

// West returns the coordinate which is one step to the west.
func (c Coord) West() Coord {
	return Coord{c.Row, c.Column - 1}
}
