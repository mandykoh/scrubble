package scrubble

import "testing"

func expectPlayedWords(t *testing.T, words []PlayedWord, expected ...PlayedWord) {
	t.Helper()

	if actual, expectedLen := len(words), len(expected); actual != expectedLen {
		t.Errorf("Expected there to be %d words formed but found %d", expectedLen, actual)

	} else {
		for i, e := range expected {
			if words[i] != e {
				t.Errorf("Expected word %#v in position %d but found %#v instead", e, i, words[i])
			}
		}
	}
}
