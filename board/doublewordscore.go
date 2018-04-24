package board

type doubleWordScore struct {
}

func (p *doubleWordScore) CountsAsConnected() bool {
	return false
}

func (p *doubleWordScore) ModifyTileScore(score int) int {
	return score
}

func (p *doubleWordScore) ModifyWordScore(score int) int {
	return score * 2
}

func (p *doubleWordScore) Name() string {
	return "Double Word Score"
}
