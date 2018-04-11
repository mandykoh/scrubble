package scrubble

type normalPositionType struct {
}

func (p *normalPositionType) CountsAsConnected() bool {
	return false
}

func (p *normalPositionType) ModifyTileScore(score int) int {
	return score
}

func (p *normalPositionType) ModifyWordScore(score int) int {
	return score
}

func (p *normalPositionType) Name() string {
	return "Normal"
}
