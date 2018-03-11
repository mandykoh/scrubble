package scrubble

// Seat represents an active player’s seat and their status within a game.
type Seat struct {
	OccupiedBy *Player
	Score      int
	Rack       Rack
}
