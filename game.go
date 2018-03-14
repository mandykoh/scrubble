package scrubble

// Game represents the rules and simulation for a single game. The zero-value of
// a Game is a game in the SetupPhase.
type Game struct {
	Phase GamePhase
	Seats []Seat
	Bag   Bag
	Board Board
}

// AddPlayer adds a seat for a new player to the game.
func (g *Game) AddPlayer(p *Player) {
	g.Seats = append(g.Seats, Seat{OccupiedBy: p})
}

// RemovePlayer removes the seat occupied by the specified player. If no such
// seat exists, this has no effect.
func (g *Game) RemovePlayer(p *Player) {
	for i, s := range g.Seats {
		if s.OccupiedBy == p {
			g.Seats = append(g.Seats[:i], g.Seats[i+1:]...)
			break
		}
	}
}
