package challenge

import (
	"github.com/mandykoh/scrubble/dict"
	"github.com/mandykoh/scrubble/history"
)

// Validate determines whether the challenge to a play is legal, and whether it
// would then be successful. A challenge succeeds if any of the words formed by
// the play are invalid according to the dictionary.
func Validate(lastPlay *history.Entry, isWordValid dict.Dictionary) (success bool, err error) {
	if lastPlay == nil {
		return false, InvalidChallengeError{NoPlayToChallengeReason}
	}

	switch lastPlay.Type {
	case history.ChallengeFailEntryType, history.ChallengeSuccessEntryType:
		return false, InvalidChallengeError{PlayAlreadyChallengedReason}

	case history.PlayEntryType:
		break

	default:
		return false, InvalidChallengeError{NoPlayToChallengeReason}
	}

	for _, w := range lastPlay.WordsFormed {
		if !isWordValid(w.Word) {
			return true, nil
		}
	}
	return false, nil
}
