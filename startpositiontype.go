package scrubble

type startPositionType struct {
}

func (p *startPositionType) CountsAsConnected() bool {
	return true
}

func (p *startPositionType) Name() string {
	return "Start"
}
