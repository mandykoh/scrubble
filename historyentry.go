package scrubble

// HistoryEntry represents an entry for one turn in a game's history of turns.
type HistoryEntry struct {
	Type        HistoryEntryType
	SeatIndex   int
	Score       int
	TilesSpent  []Tile
	TilesPlayed TilePlacements
	TilesDrawn  []Tile
	WordsFormed []PlayedWord
}
