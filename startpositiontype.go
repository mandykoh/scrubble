package scrubble

type startPositionType struct {
}

func (p *startPositionType) CountsAsConnected() bool {
	return true
}

func (p *startPositionType) ModifyTileScore(t Tile) int {
	return t.Points
}

func (p *startPositionType) Name() string {
	return "Start"
}
