package play

import "github.com/mandykoh/scrubble/coord"

// Word represents a word formed on the board by playing tiles.
type Word struct {
	Word  string
	Score int
	coord.Range
}
