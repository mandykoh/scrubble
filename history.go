package scrubble

// History represents a game's history of turns and scoring.
type History []HistoryEntry

// AppendPlay adds an entry to the history representing a turn where tiles were
// successfully played.
func (h *History) AppendPlay(seatIndex int, score int, tilesSpent []Tile, tilesPlayed TilePlacements, wordsFormed []PlayedWord) {
	*h = append(*h, HistoryEntry{
		SeatIndex:   seatIndex,
		Score:       score,
		TilesSpent:  tilesSpent,
		TilesPlayed: tilesPlayed,
		WordsFormed: wordsFormed,
	})
}

// Last returns last entry in the history.
func (h *History) Last() *HistoryEntry {
	return &(*h)[len(*h)-1]
}
