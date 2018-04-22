package scrubble

import "testing"

func expectHistoryEntry(t *testing.T, entry HistoryEntry, expected HistoryEntry) {
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
