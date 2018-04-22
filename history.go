package scrubble

// History represents a game's history of turns and scoring.
type History []HistoryEntry

// AppendChallengeFail adds an entry to the history representing an unsuccessful challenge.
func (h *History) AppendChallengeFail(challengerSeatIndex int) {
	*h = append(*h, HistoryEntry{
		Type:      ChallengeFailHistoryEntryType,
		SeatIndex: challengerSeatIndex,
	})
}

// AppendChallengeSuccess adds an entry to the history representing a successful challenge.
func (h *History) AppendChallengeSuccess(challengerSeatIndex int) {
	*h = append(*h, HistoryEntry{
		Type:      ChallengeSuccessHistoryEntryType,
		SeatIndex: challengerSeatIndex,
	})
}

// AppendExchange adds an entry to the history representing a turn where tiles
// were successfully exchanged with the bag.
func (h *History) AppendExchange(seatIndex int, tilesSpent []Tile, tilesDrawn []Tile) {
	*h = append(*h, HistoryEntry{
		Type:       ExchangeTilesHistoryEntryType,
		SeatIndex:  seatIndex,
		TilesSpent: tilesSpent,
		TilesDrawn: tilesDrawn,
	})
}

// AppendPass adds an entry to the history representing a turn which was passed.
func (h *History) AppendPass(seatIndex int) {
	*h = append(*h, HistoryEntry{
		Type:      PassHistoryEntryType,
		SeatIndex: seatIndex,
	})
}

// AppendPlay adds an entry to the history representing a turn where tiles were
// successfully played.
func (h *History) AppendPlay(seatIndex int, score int, tilesSpent []Tile, tilesPlayed TilePlacements, tilesDrawn []Tile, wordsFormed []PlayedWord) {
	*h = append(*h, HistoryEntry{
		Type:        PlayHistoryEntryType,
		SeatIndex:   seatIndex,
		Score:       score,
		TilesSpent:  tilesSpent,
		TilesPlayed: tilesPlayed,
		TilesDrawn:  tilesDrawn,
		WordsFormed: wordsFormed,
	})
}

// Last returns last entry in the history.
func (h *History) Last() *HistoryEntry {
	return &(*h)[len(*h)-1]
}
