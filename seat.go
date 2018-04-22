package scrubble

// Seat represents an active player’s seat and their status within a game. The
// zero-value of a Seat is a seat with no score and an empty rack.
type Seat struct {
	Score int
	Rack  Rack
}
