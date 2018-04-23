package positiontype

type normal struct {
}

func (p *normal) CountsAsConnected() bool {
	return false
}

func (p *normal) ModifyTileScore(score int) int {
	return score
}

func (p *normal) ModifyWordScore(score int) int {
	return score
}

func (p *normal) Name() string {
	return "Normal"
}
