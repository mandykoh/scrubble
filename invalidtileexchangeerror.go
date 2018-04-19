package scrubble

import "fmt"

// InvalidTileExchangeError indicates that an attempt to exchange tiles with the
// bag was invalid.
type InvalidTileExchangeError struct {
	Reason InvalidTileExchangeReason
}

func (e InvalidTileExchangeError) Error() string {
	return fmt.Sprintf("%#v", e)
}
