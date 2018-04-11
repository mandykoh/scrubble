package scrubble

type tripleLetterScorePositionType struct {
}

func (p *tripleLetterScorePositionType) CountsAsConnected() bool {
	return false
}

func (p *tripleLetterScorePositionType) ModifyTileScore(t Tile) int {
	return t.Points * 3
}

func (p *tripleLetterScorePositionType) Name() string {
	return "Triple Letter Score"
}
