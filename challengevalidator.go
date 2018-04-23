package scrubble

import "github.com/mandykoh/scrubble/dict"

// ChallengeValidator represents a function which determines whether a challenge
// to a play is successful.
type ChallengeValidator func(formedWords []PlayedWord, dictionary dict.Dictionary) bool
