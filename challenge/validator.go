package challenge

import (
	"github.com/mandykoh/scrubble/dict"
	"github.com/mandykoh/scrubble/history"
)

// Validator represents a function which determines whether a challenge to a
// play is legal and whether it would succeed.
type Validator func(lastPlay *history.Entry, dictionary dict.Dictionary) (success bool, err error)
