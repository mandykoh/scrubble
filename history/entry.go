package history

import (
	"github.com/mandykoh/scrubble/play"
	"github.com/mandykoh/scrubble/tile"
)

// Entry represents an entry for one turn in a game's history of turns.
type Entry struct {
	Type        EntryType
	SeatIndex   int
	Score       int
	TilesSpent  []tile.Tile
	TilesPlayed play.Tiles
	TilesDrawn  []tile.Tile
	WordsFormed []play.Word
}
