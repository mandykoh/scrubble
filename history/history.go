package history

import (
	"github.com/mandykoh/scrubble/play"
	"github.com/mandykoh/scrubble/tile"
)

// History represents a game's history of turns and scoring.
type History []Entry

// AppendChallengeFail adds an entry to the history representing an unsuccessful challenge.
func (h *History) AppendChallengeFail(challengerSeatIndex int) {
	*h = append(*h, Entry{
		Type:      ChallengeFailEntryType,
		SeatIndex: challengerSeatIndex,
	})
}

// AppendChallengeSuccess adds an entry to the history representing a successful challenge.
func (h *History) AppendChallengeSuccess(challengerSeatIndex int) {
	*h = append(*h, Entry{
		Type:      ChallengeSuccessEntryType,
		SeatIndex: challengerSeatIndex,
	})
}

// AppendExchange adds an entry to the history representing a turn where tiles
// were successfully exchanged with the bag.
func (h *History) AppendExchange(seatIndex int, tilesSpent, tilesDrawn []tile.Tile) {
	*h = append(*h, Entry{
		Type:       ExchangeTilesEntryType,
		SeatIndex:  seatIndex,
		TilesSpent: tilesSpent,
		TilesDrawn: tilesDrawn,
	})
}

// AppendPass adds an entry to the history representing a turn which was passed.
func (h *History) AppendPass(seatIndex int) {
	*h = append(*h, Entry{
		Type:      PassEntryType,
		SeatIndex: seatIndex,
	})
}

// AppendPlay adds an entry to the history representing a turn where tiles were
// successfully played.
func (h *History) AppendPlay(seatIndex int, score int, tilesSpent []tile.Tile, tilesPlayed play.Tiles, tilesDrawn []tile.Tile, wordsFormed []play.Word) {
	*h = append(*h, Entry{
		Type:        PlayEntryType,
		SeatIndex:   seatIndex,
		Score:       score,
		TilesSpent:  tilesSpent,
		TilesPlayed: tilesPlayed,
		TilesDrawn:  tilesDrawn,
		WordsFormed: wordsFormed,
	})
}

// Last returns last entry in the history.
func (h *History) Last() *Entry {
	return &(*h)[len(*h)-1]
}
