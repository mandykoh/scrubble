package scrubble

type tripleLetterScorePositionType struct {
}

func (p *tripleLetterScorePositionType) CountsAsConnected() bool {
	return false
}

func (p *tripleLetterScorePositionType) ModifyTileScore(score int) int {
	return score * 3
}

func (p *tripleLetterScorePositionType) ModifyWordScore(score int) int {
	return score
}

func (p *tripleLetterScorePositionType) Name() string {
	return "Triple Letter Score"
}
