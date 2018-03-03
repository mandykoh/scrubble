package scrubble

// PositionType represents a type of board position, which may be a starting
// position, or a position with a score bonus, etc.
type PositionType interface {
	Name() string
}
