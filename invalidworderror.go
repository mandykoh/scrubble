package scrubble

import "fmt"

// InvalidWordError indicates that one or more formed words is not valid for
// play and cannot be scored.
type InvalidWordError struct {
	WordSpans []CoordRange
}

func (e InvalidWordError) Error() string {
	return fmt.Sprintf("%#v", e)
}
