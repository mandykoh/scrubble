package scrubble

// NormalPositionType is the standard board position type with no special
// properties.
var NormalPositionType PositionType = &normalPositionType{}

type normalPositionType struct {
}

func (p *normalPositionType) Name() string {
	return "Normal"
}
