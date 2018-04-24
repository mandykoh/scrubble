package scoring

import (
	"github.com/mandykoh/scrubble/board"
	"github.com/mandykoh/scrubble/dict"
	"github.com/mandykoh/scrubble/play"
)

// WordScorer represents a function which can determine the words formed from
// playing tiles on a board, and their scores.
type WordScorer func(placements play.Tiles, board *board.Board, isWordValid dict.Dictionary) (score int, words []play.Word, err error)
