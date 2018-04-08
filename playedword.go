package scrubble

// PlayedWord represents a word formed on the board by playing tiles.
type PlayedWord struct {
	Word  string
	Score int
	CoordRange
}
