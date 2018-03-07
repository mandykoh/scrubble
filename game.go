package scrubble

// Game represents the rules and simulation for a single game.
type Game struct {
	Seats []Seat
	Bag   Bag
	Board Board
}
