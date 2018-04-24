package board

type tripleWordScore struct {
}

func (p *tripleWordScore) CountsAsConnected() bool {
	return false
}

func (p *tripleWordScore) ModifyTileScore(score int) int {
	return score
}

func (p *tripleWordScore) ModifyWordScore(score int) int {
	return score * 3
}

func (p *tripleWordScore) Name() string {
	return "Triple Word Score"
}
