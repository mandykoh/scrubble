package scrubble

import "fmt"

// InvalidWordError indicates that a formed word is not valid for play and
// cannot be scored.
type InvalidWordError struct {
	Reason InvalidWordReason
}

func (e InvalidWordError) Error() string {
	return fmt.Sprintf("%#v", e)
}
