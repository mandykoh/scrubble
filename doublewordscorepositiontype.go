package scrubble

// DoubleWordScorePositionType is the board position type indicating a double
// word score bonus.
var DoubleWordScorePositionType PositionType = &doubleWordScorePositionType{}

type doubleWordScorePositionType struct {
}

func (p *doubleWordScorePositionType) Name() string {
	return "Double Word Score"
}
