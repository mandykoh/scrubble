package scrubble

// BoardLayout represents a layout for creating a Board. Layouts are specified
// from the top row down, from the leftmost column to  the rightmost.
type BoardLayout [][]PositionType

// BeginRow specifies that a new row in the layout should be started.
func (b BoardLayout) BeginRow() BoardLayout {
	return append(b, []PositionType{})
}

// DLetter specifies that a double-letter score layout position should be added.
func (b BoardLayout) DLetter() BoardLayout {
	return b.appendToLastRow(DoubleLetterScorePositionType)
}

// DWord specifies that a double-word score layout position should be added.
func (b BoardLayout) DWord() BoardLayout {
	return b.appendToLastRow(DoubleWordScorePositionType)
}

// Empty specifies that an empty layout position should be added.
func (b BoardLayout) Empty() BoardLayout {
	return b.appendToLastRow(NormalPositionType)
}

// Start specifies that a starting position should be added.
func (b BoardLayout) Start() BoardLayout {
	return b.appendToLastRow(StartPositionType)
}

// TLetter specifies that a triple-letter score layout position should be added.
func (b BoardLayout) TLetter() BoardLayout {
	return b.appendToLastRow(TripleLetterScorePositionType)
}

// TWord specifies that a triple-word score layout position should be added.
func (b BoardLayout) TWord() BoardLayout {
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
