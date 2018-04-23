package positiontype

// Interface represents a type of board position, which may be a starting
// position, or a position with a score bonus, etc.
type Interface interface {
	CountsAsConnected() bool
	ModifyTileScore(score int) int
	ModifyWordScore(score int) int
	Name() string
}
