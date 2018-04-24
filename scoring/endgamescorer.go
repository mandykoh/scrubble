package scoring

import (
	"github.com/mandykoh/scrubble/history"
	"github.com/mandykoh/scrubble/seat"
)

type EndGameScorer func(lastPlay *history.Entry, seats []seat.Seat) (finalScores []int)
