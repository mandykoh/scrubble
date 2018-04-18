package scrubble

// HistoryEntry represents an entry for one turn in a game's history of turns.
type HistoryEntry struct {
	SeatIndex   int
	Score       int
	TilesSpent  []Tile
	TilesPlayed TilePlacements
	WordsFormed []PlayedWord
}
