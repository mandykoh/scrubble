package scrubble

type normalPositionType struct {
}

func (p *normalPositionType) CountsAsConnected() bool {
	return false
}

func (p *normalPositionType) Name() string {
	return "Normal"
}
