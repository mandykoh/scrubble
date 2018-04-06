package scrubble

type tripleWordScorePositionType struct {
}

func (p *tripleWordScorePositionType) CountsAsConnected() bool {
	return false
}

func (p *tripleWordScorePositionType) Name() string {
	return "Triple Word Score"
}
