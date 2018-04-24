package board

// PositionType represents a type of board position, which may be a starting
// position, or a position with a score bonus, etc.
type PositionType interface {
	CountsAsConnected() bool
	ModifyTileScore(score int) int
	ModifyWordScore(score int) int
	Name() string
}
