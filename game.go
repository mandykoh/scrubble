package scrubble

import "math/rand"

// Game represents the rules and simulation for a single game. The zero-value of
// a Game is a game in the SetupPhase.
type Game struct {
	Phase            GamePhase
	Seats            []Seat
	Bag              Bag
	Board            Board
	CurrentSeatIndex int
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

// Start begins the game by shuffling the bag, picking a random seat for the
// first turn, filling all players' tile racks from the bag, and moving the game
// into the MainPhase.
//
// The supplied random number generator is used to determine the bag shuffling
// and the starting player.
func (g *Game) Start(r *rand.Rand) {
	g.CurrentSeatIndex = r.Intn(len(g.Seats))
	g.Bag.Shuffle(r)

	for i := range g.Seats {
		g.Seats[i].Rack.FillFromBag(&g.Bag)
	}

	g.Phase = MainPhase
}
