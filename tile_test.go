package scrubble

import (
	"testing"

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
