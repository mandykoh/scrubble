package scrubble

// DoubleLetterScorePositionType is the board position type indicating a double
// letter score bonus.
var DoubleLetterScorePositionType PositionType = &doubleLetterScorePositionType{}

type doubleLetterScorePositionType struct {
}

func (p *doubleLetterScorePositionType) Name() string {
	return "Double Letter Score"
}
