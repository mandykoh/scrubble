package scrubble

import "github.com/mandykoh/scrubble/dict"

// WordScorer represents a function which can determine the words formed from
// playing tiles on a board, and their scores.
type WordScorer func(placements TilePlacements, board *Board, isWordValid dict.Dictionary) (score int, words []PlayedWord, err error)
