package scrubble

import "github.com/mandykoh/scrubble/coord"

// PlayedWord represents a word formed on the board by playing tiles.
type PlayedWord struct {
	Word  string
	Score int
	coord.Range
}
