package challenge

import (
	"strings"
	"testing"

	"github.com/mandykoh/scrubble/play"
)

func TestIsSuccessful(t *testing.T) {
	dictionary := func(word string) (valid bool) {
		return strings.HasPrefix(word, "VALIDWORD")
	}

	t.Run("returns success when any played words are invalid", func(t *testing.T) {
		success := IsSuccessful([]play.Word{
			{Word: "VALIDWORD1"},
			{Word: "INVALIDWORD"},
		}, dictionary)

		if !success {
			t.Errorf("Expected challenge to succeed but it will fail")
		}
	})

	t.Run("returns failure when all played words are valid", func(t *testing.T) {
		success := IsSuccessful([]play.Word{
			{Word: "VALIDWORD1"},
			{Word: "VALIDWORD2"},
			{Word: "VALIDWORD3"},
		}, dictionary)

		if success {
			t.Errorf("Expected challenge to fail but it will succeed")
		}
	})
}
