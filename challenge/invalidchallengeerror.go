package challenge

import "fmt"

// InvalidChallengeError indicates that an attempt to challenge a play was
// invalid.
type InvalidChallengeError struct {
	Reason InvalidChallengeReason
}

func (e InvalidChallengeError) Error() string {
	return fmt.Sprintf("%#v", e)
}
