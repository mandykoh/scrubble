package scrubble

// GamePhaseController represents a function which determines the next game
// phase after a turn is played. This is called by Game at the end of each turn.
type GamePhaseController func(game *Game) (next GamePhase)

// NextGamePhase implements a GamePhaseController with the default game
// progression and ending conditions.
func NextGamePhase(game *Game) GamePhase {
	lastTurn := game.History.Last()

	if len(game.Seats[lastTurn.SeatIndex].Rack) == 0 {
		return EndPhase
	}

	return MainPhase
}
