package scrubble

import "github.com/mandykoh/scrubble/tile"

// HistoryEntry represents an entry for one turn in a game's history of turns.
type HistoryEntry struct {
	Type        HistoryEntryType
	SeatIndex   int
	Score       int
	TilesSpent  []tile.Tile
	TilesPlayed TilePlacements
	TilesDrawn  []tile.Tile
	WordsFormed []PlayedWord
}
