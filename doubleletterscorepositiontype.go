package scrubble

type doubleLetterScorePositionType struct {
}

func (p *doubleLetterScorePositionType) CountsAsConnected() bool {
	return false
}

func (p *doubleLetterScorePositionType) Name() string {
	return "Double Letter Score"
}
