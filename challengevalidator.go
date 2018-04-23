package scrubble

import (
	"github.com/mandykoh/scrubble/dict"
	"github.com/mandykoh/scrubble/play"
)

// ChallengeValidator represents a function which determines whether a challenge
// to a play is successful.
type ChallengeValidator func(formedWords []play.Word, dictionary dict.Dictionary) bool
