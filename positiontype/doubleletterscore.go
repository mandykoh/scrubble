package positiontype

type doubleLetterScore struct {
}

func (p *doubleLetterScore) CountsAsConnected() bool {
	return false
}

func (p *doubleLetterScore) ModifyTileScore(score int) int {
	return score * 2
}

func (p *doubleLetterScore) ModifyWordScore(score int) int {
	return score
}

func (p *doubleLetterScore) Name() string {
	return "Double Letter Score"
}
