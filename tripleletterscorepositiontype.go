package scrubble

type tripleLetterScorePositionType struct {
}

func (p *tripleLetterScorePositionType) CountsAsConnected() bool {
	return false
}

func (p *tripleLetterScorePositionType) Name() string {
	return "Triple Letter Score"
}
