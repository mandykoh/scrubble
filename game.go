package scrubble

// Game represents the rules and simulation for a single game.
type Game struct {
	Seats []Seat
	Bag   Bag
	Board Board
}

// GameWithDefaults returns a game with a default board layout and bag setup,
// with the specified players.
func GameWithDefaults(p1 *Player, p2 *Player, others ...*Player) Game {
	return GameWithSettings(
		BagWithStandardEnglishTiles(),
		BoardWithStandardLayout(),
		p1, p2, others...)
}

// GameWithSettings returns a game set up with seats for the given players,
// using the specified bag and board.
func GameWithSettings(bag Bag, board Board, p1 *Player, p2 *Player, others ...*Player) Game {
	seats := []Seat{
		{OccupiedBy: p1},
		{OccupiedBy: p2},
	}

	for _, p := range others {
		seats = append(seats, Seat{OccupiedBy: p})
	}

	return Game{
		Seats: seats,
		Bag:   bag,
		Board: board,
	}
}
