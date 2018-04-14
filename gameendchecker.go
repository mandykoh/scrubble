package scrubble

// GameEndChecker represents a function which checks the game state to determine
// if the end-of-game conditions have been met. This is called by Game at the
// end of each turn, before moving to the next player.
type GameEndChecker func(lastSeat *Seat, lastScore int, game *Game) bool

// GameEnds implements a GameEndChecker with the default game ending conditions.
func GameEnds(lastSeat *Seat, lastScore int, game *Game) bool {
	return len(lastSeat.Rack) == 0
}
