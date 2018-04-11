package scrubble

type doubleLetterScorePositionType struct {
}

func (p *doubleLetterScorePositionType) CountsAsConnected() bool {
	return false
}

func (p *doubleLetterScorePositionType) ModifyTileScore(score int) int {
	return score * 2
}

func (p *doubleLetterScorePositionType) ModifyWordScore(score int) int {
	return score
}

func (p *doubleLetterScorePositionType) Name() string {
	return "Double Letter Score"
}
