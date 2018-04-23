package scrubble

import (
	"testing"

	"github.com/mandykoh/scrubble/play"
	"github.com/mandykoh/scrubble/tile"
)

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
