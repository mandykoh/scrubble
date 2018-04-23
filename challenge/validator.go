package challenge

import (
	"github.com/mandykoh/scrubble/dict"
	"github.com/mandykoh/scrubble/play"
)

// Validator represents a function which determines whether a challenge to a
// play is successful.
type Validator func(formedWords []play.Word, dictionary dict.Dictionary) bool
