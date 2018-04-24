package game

import (
	"testing"

	"github.com/mandykoh/scrubble/history"
	"github.com/mandykoh/scrubble/play"
	"github.com/mandykoh/scrubble/tile"
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

func expectPlayedWords(t *testing.T, words []play.Word, expected ...play.Word) {
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

func expectTiles(t *testing.T, descriptor string, tiles []tile.Tile, expected ...tile.Tile) {
	t.Helper()

	if actual, expectedLen := len(tiles), len(expected); actual != expectedLen {
		t.Errorf("Expected there to be %d tiles %s but found %d", expectedLen, descriptor, actual)

	} else {
		for i, e := range expected {
			if tiles[i] != e {
				t.Errorf("Expected %s tile %v in position %d but found %v instead", descriptor, e, i, tiles[i])
			}
		}
	}
}

func expectTilePlacements(t *testing.T, placements play.Tiles, expected ...play.TilePlacement) {
	t.Helper()

	if actual, expectedLen := len(placements), len(expected); actual != expectedLen {
		t.Errorf("Expected there to be %d tiles placed but found %d", expectedLen, actual)

	} else {
		for i, e := range expected {
			if placements[i] != e {
				t.Errorf("Expected tile placement %v in position %d but found %v instead", e, i, placements[i])
			}
		}
	}
}
