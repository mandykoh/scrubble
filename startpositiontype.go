package scrubble

// StartPositionType is the board position type indicating a starting position.
var StartPositionType PositionType = &startPositionType{}

type startPositionType struct {
}

func (p *startPositionType) Name() string {
	return "Start"
}
