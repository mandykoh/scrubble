package scrubble

type normalPositionType struct {
}

func (p *normalPositionType) CountsAsConnected() bool {
	return false
}

func (p *normalPositionType) ModifyTileScore(t Tile) int {
	return t.Points
}

func (p *normalPositionType) Name() string {
	return "Normal"
}
