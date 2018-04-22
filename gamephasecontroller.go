package scrubble

// MaxScorelessTurns represents the maximum number of consecutive scoreless
// turns for the game to end.
const MaxScorelessTurns = 6

// GamePhaseController represents a function which determines the next game
// phase after a turn is played. This is called by Game at the end of each turn.
type GamePhaseController func(game *Game) (next GamePhase)

// NextGamePhase implements a GamePhaseController with the default game
// progression and ending conditions.
func NextGamePhase(game *Game) GamePhase {
	lastTurn := game.History.Last()

	// Last player's rack was empty, which means they played out
	if len(game.Seats[lastTurn.SeatIndex].Rack) == 0 {
		return EndPhase
	}

	scoreless := 0
	for i := len(game.History) - 1; i >= 0; i-- {
		entry := &game.History[i]

		if entry.Type == ChallengeSuccessHistoryEntryType {
			i--
		} else if entry.Score > 0 {
			break
		}

		scoreless++
		if scoreless >= MaxScorelessTurns {
			return EndPhase
		}
	}

	return MainPhase
}
