package scrubble

// TripleWordScorePositionType is the board position type indicating a triple
// word score bonus.
var TripleWordScorePositionType PositionType = &tripleWordScorePositionType{}

type tripleWordScorePositionType struct {
}

func (p *tripleWordScorePositionType) Name() string {
	return "Triple Word Score"
}
