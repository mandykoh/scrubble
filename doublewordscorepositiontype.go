package scrubble

type doubleWordScorePositionType struct {
}

func (p *doubleWordScorePositionType) CountsAsConnected() bool {
	return false
}

func (p *doubleWordScorePositionType) Name() string {
	return "Double Word Score"
}
