package scrubble

type tripleWordScorePositionType struct {
}

func (p *tripleWordScorePositionType) CountsAsConnected() bool {
	return false
}

func (p *tripleWordScorePositionType) ModifyTileScore(score int) int {
	return score
}

func (p *tripleWordScorePositionType) ModifyWordScore(score int) int {
	return score * 3
}

func (p *tripleWordScorePositionType) Name() string {
	return "Triple Word Score"
}
