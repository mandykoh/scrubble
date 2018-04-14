package scrubble

// GamePhaseController represents a function which determines the next game
// phase after a turn is played. This is called by Game at the end of each turn.
type GamePhaseController func(lastSeat *Seat, lastScore int, game *Game) (next GamePhase)

// NextGamePhase implements a GamePhaseController with the default game
// progression and ending conditions.
func NextGamePhase(lastSeat *Seat, lastScore int, game *Game) GamePhase {
	if len(lastSeat.Rack) == 0 {
		return EndPhase
	}

	return MainPhase
}
