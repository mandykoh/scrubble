package scrubble

// LayoutPositionTypes returns a set of position types which can be used to
// conveniently specify BoardLayouts. The position types returned are:
// normal/empty, start, double letter score, double word score, triple letter
// score, triple word score.
//
// See BoardWithLayout for example usage.
func LayoutPositionTypes() (__, st, dl, dw, tl, tw PositionType) {
	return NormalPositionType,
		StartPositionType,
		DoubleLetterScorePositionType,
		DoubleWordScorePositionType,
		TripleLetterScorePositionType,
		TripleWordScorePositionType
}

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
