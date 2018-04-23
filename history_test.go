package scrubble

import (
	"testing"

	"github.com/mandykoh/scrubble/history"
)

func expectHistory(t *testing.T, history history.History, expected ...history.Entry) {
	t.Helper()

	if actual, expectedLen := len(history), len(expected); actual != expectedLen {
		t.Errorf("Expected there to be %d history entries but found %d", expectedLen, actual)

	} else {
		for i, e := range expected {
			expectHistoryEntry(t, history[i], e)
		}
	}
}

func expectHistoryEntry(t *testing.T, entry history.Entry, expected history.Entry) {
	t.Helper()

	if actual, expected := entry.Type, expected.Type; actual != expected {
		t.Errorf("Expected history entry type of %v but was %v", expected, actual)
	}

	if actual, expected := entry.SeatIndex, expected.SeatIndex; actual != expected {
		t.Errorf("Expected history entry to record seat index %d but was %d", expected, actual)
	}

	if actual, expected := entry.Score, expected.Score; actual != expected {
		t.Errorf("Expected history entry to record score of %d but was %d", expected, actual)
	}

	expectTiles(t, "spent", entry.TilesSpent, expected.TilesSpent...)
	expectTilePlacements(t, entry.TilesPlayed, expected.TilesPlayed...)
	expectTiles(t, "drawn", entry.TilesDrawn, expected.TilesDrawn...)
	expectPlayedWords(t, entry.WordsFormed, expected.WordsFormed...)
}
