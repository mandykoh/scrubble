package scrubble

// BoardLayout represents a layout for creating a Board. Layouts are specified
// from the top row down, from the leftmost column to  the rightmost.
type BoardLayout [][]PositionType

// BeginRow specifies that a new row in the layout should be started.
func (b BoardLayout) BeginRow() BoardLayout {
	return append(b, []PositionType{})
}

// DL specifies that a double-letter score layout position should be added.
func (b BoardLayout) DL() BoardLayout {
	return b.appendToLastRow(DoubleLetterScorePositionType)
}

// DW specifies that a double-word score layout position should be added.
func (b BoardLayout) DW() BoardLayout {
	return b.appendToLastRow(DoubleWordScorePositionType)
}

// Em specifies that an empty layout position should be added.
func (b BoardLayout) Em() BoardLayout {
	return b.appendToLastRow(NormalPositionType)
}

// St specifies that a starting position should be added.
func (b BoardLayout) St() BoardLayout {
	return b.appendToLastRow(StartPositionType)
}

// TL specifies that a triple-letter score layout position should be added.
func (b BoardLayout) TL() BoardLayout {
	return b.appendToLastRow(TripleLetterScorePositionType)
}

// TW specifies that a triple-word score layout position should be added.
func (b BoardLayout) TW() BoardLayout {
	return b.appendToLastRow(TripleWordScorePositionType)
}

func (b BoardLayout) appendToLastRow(positionType PositionType) BoardLayout {
	b[len(b)-1] = append(b[len(b)-1], positionType)
	return b
}

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
