package scrubble

type tripleWordScorePositionType struct {
}

func (p *tripleWordScorePositionType) CountsAsConnected() bool {
	return false
}

func (p *tripleWordScorePositionType) ModifyTileScore(t Tile) int {
	return t.Points
}

func (p *tripleWordScorePositionType) Name() string {
	return "Triple Word Score"
}
