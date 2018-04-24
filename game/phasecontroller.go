package game

import "github.com/mandykoh/scrubble/history"

// MaxScorelessTurns represents the maximum number of consecutive scoreless
// turns for the game to end.
const MaxScorelessTurns = 6

// PhaseController represents a function which determines the next game phase
// after a turn is played. This is called by Game at the end of each turn.
type PhaseController func(game *Game) (next Phase)

// NextPhase implements a PhaseController with the default game progression
// and ending conditions.
func NextPhase(game *Game) Phase {
	lastTurn := game.History.Last()

	// Last player's rack was empty, which means they played out
	if len(game.Seats[lastTurn.SeatIndex].Rack) == 0 {
		return EndPhase
	}

	scoreless := 0
	for i := len(game.History) - 1; i >= 0; i-- {
		entry := &game.History[i]

		if entry.Type == history.ChallengeSuccessEntryType {
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
