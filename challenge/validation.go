package challenge

import (
	"github.com/mandykoh/scrubble/dict"
	"github.com/mandykoh/scrubble/play"
)

// IsSuccessful determines whether the challenge to a play is successful. A
// challenge succeeds if any of the words formed by the play are invalid
// according to the dictionary.
func IsSuccessful(formedWords []play.Word, isWordValid dict.Dictionary) bool {
	for _, w := range formedWords {
		if !isWordValid(w.Word) {
			return true
		}
	}
	return false
}
