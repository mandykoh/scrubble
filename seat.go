package scrubble

// Seat represents an active player’s seat and their status within a game. The
// zero-value of a Seat is an unoccupied seat with no score and empty rack.
type Seat struct {
	OccupiedBy *Player
	Score      int
	Rack       Rack
}
