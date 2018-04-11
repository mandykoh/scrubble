package scrubble

type startPositionType struct {
}

func (p *startPositionType) CountsAsConnected() bool {
	return true
}

func (p *startPositionType) ModifyTileScore(score int) int {
	return score
}

func (p *startPositionType) ModifyWordScore(score int) int {
	return score * 2
}

func (p *startPositionType) Name() string {
	return "Start"
}
