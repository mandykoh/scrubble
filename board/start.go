package board

type start struct {
}

func (p *start) CountsAsConnected() bool {
	return true
}

func (p *start) ModifyTileScore(score int) int {
	return score
}

func (p *start) ModifyWordScore(score int) int {
	return score * 2
}

func (p *start) Name() string {
	return "Start"
}
