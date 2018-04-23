package play

import "fmt"

// InvalidTilePlacementError indicates that a play called for placing tiles in
// an invalid manner.
type InvalidTilePlacementError struct {
	Reason InvalidTilePlacementReason
}

func (e InvalidTilePlacementError) Error() string {
	return fmt.Sprintf("%#v", e)
}
