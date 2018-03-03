package scrubble

// TripleLetterScorePositionType is the board position type indicating a triple
// letter score bonus.
var TripleLetterScorePositionType PositionType = &tripleLetterScorePositionType{}

type tripleLetterScorePositionType struct {
}

func (p *tripleLetterScorePositionType) Name() string {
	return "Triple Letter Score"
}
