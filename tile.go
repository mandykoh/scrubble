package scrubble

import "fmt"

// Tile represents a game tile which can be placed on a Board. Each tile has a
// letter and an associated number of points.
type Tile struct {
	Letter rune
	Points int
}

func (t Tile) String() string {
	return fmt.Sprintf("%c(%d)", t.Letter, t.Points)
}
