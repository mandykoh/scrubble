package challenge

import (
	"strings"
	"testing"

	"github.com/mandykoh/scrubble/history"
	"github.com/mandykoh/scrubble/play"
)

func TestValidate(t *testing.T) {
	dictionary := func(word string) (valid bool) {
		return strings.HasPrefix(word, "VALIDWORD")
	}

	t.Run("returns an error when there is no play to challenge", func(t *testing.T) {
		_, err := Validate(nil, dictionary)
		if actual, expected := err, (InvalidChallengeError{NoPlayToChallengeReason}); actual != expected {
			t.Fatalf("Expected error %v but was %v", expected, err)
		}

		_, err = Validate(&history.Entry{Type: history.PassEntryType}, dictionary)
		if actual, expected := err, (InvalidChallengeError{NoPlayToChallengeReason}); actual != expected {
			t.Fatalf("Expected error %v but was %v", expected, err)
		}
	})

	t.Run("returns an error when the last play was already challenged", func(t *testing.T) {
		_, err := Validate(&history.Entry{Type: history.ChallengeFailEntryType}, dictionary)
		if actual, expected := err, (InvalidChallengeError{PlayAlreadyChallengedReason}); actual != expected {
			t.Fatalf("Expected error %v but was %v", expected, err)
		}

		_, err = Validate(&history.Entry{Type: history.ChallengeSuccessEntryType}, dictionary)
		if actual, expected := err, (InvalidChallengeError{PlayAlreadyChallengedReason}); actual != expected {
			t.Fatalf("Expected error %v but was %v", expected, err)
		}
	})

	t.Run("returns success when any played words are invalid", func(t *testing.T) {
		success, err := Validate(&history.Entry{
			Type: history.PlayEntryType,
			WordsFormed: []play.Word{
				{Word: "VALIDWORD1"},
				{Word: "INVALIDWORD"},
			},
		}, dictionary)

		if err != nil {
			t.Errorf("Expected no errors but got %v", err)
		}
		if !success {
			t.Errorf("Expected challenge to succeed but it will fail")
		}
	})

	t.Run("returns failure when all played words are valid", func(t *testing.T) {
		success, err := Validate(&history.Entry{
			Type: history.PlayEntryType,
			WordsFormed: []play.Word{
				{Word: "VALIDWORD1"},
				{Word: "VALIDWORD2"},
				{Word: "VALIDWORD3"},
			},
		}, dictionary)

		if err != nil {
			t.Errorf("Expected no errors but got %v", err)
		}
		if success {
			t.Errorf("Expected challenge to fail but it will succeed")
		}
	})
}
