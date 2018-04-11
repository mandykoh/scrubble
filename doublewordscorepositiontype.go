package scrubble

type doubleWordScorePositionType struct {
}

func (p *doubleWordScorePositionType) CountsAsConnected() bool {
	return false
}

func (p *doubleWordScorePositionType) ModifyTileScore(score int) int {
	return score
}

func (p *doubleWordScorePositionType) ModifyWordScore(score int) int {
	return score * 2
}

func (p *doubleWordScorePositionType) Name() string {
	return "Double Word Score"
}
