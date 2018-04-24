package board

type tripleLetterScore struct {
}

func (p *tripleLetterScore) CountsAsConnected() bool {
	return false
}

func (p *tripleLetterScore) ModifyTileScore(score int) int {
	return score * 3
}

func (p *tripleLetterScore) ModifyWordScore(score int) int {
	return score
}

func (p *tripleLetterScore) Name() string {
	return "Triple Letter Score"
}
