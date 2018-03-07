package scrubble

// Player represents a single player, who may hold seats at multiple games.
type Player struct {
	Name  string
	Seats []*Seat
}
