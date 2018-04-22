package textscrubble

import (
	"strings"

	"strconv"

	"math/rand"

	gt "github.com/buger/goterm"
	"github.com/mandykoh/scrubble"
)

type Player struct {
	Name string
}

func Challenge(game *scrubble.Game, rng *rand.Rand) {
	challengerIndex := game.CurrentSeatIndex

	success, err := game.Challenge(challengerIndex, rng)
	if err != nil {
		gt.Println(gt.Color(err.Error(), gt.RED))
	} else if success {
		gt.Printf(gt.Color("\n\nPlay successfully challenged!", gt.GREEN))
	} else {
		gt.Printf(gt.Color("\n\nChallenge failed! All words are valid", gt.RED))
	}
}

func ExchangeTiles(letters string, game *scrubble.Game, rng *rand.Rand) {
	seat := game.CurrentSeat()
	tiles := LettersToRackTiles(letters, seat.Rack)

	err := game.ExchangeTiles(tiles, rng)
	if err != nil {
		gt.Println(gt.Color(err.Error(), gt.RED))
	} else {
		DrawRack(seat.Rack)
		gt.Printf("\n\nTiles exchanged")
	}
}

func LettersToPlacements(rowDir, colDir, row, col int, letters string, rack scrubble.Rack, board *scrubble.Board) scrubble.TilePlacements {
	var placements scrubble.TilePlacements
	tiles := LettersToRackTiles(letters, rack)

	for _, tile := range tiles {
		pos := board.Position(scrubble.Coord{Row: row, Column: col})

		for pos != nil && pos.Tile != nil {
			row += rowDir
			col += colDir
			pos = board.Position(scrubble.Coord{Row: row, Column: col})
		}

		placements = append(placements, scrubble.TilePlacement{
			Tile:  tile,
			Coord: scrubble.Coord{Row: row, Column: col},
		})
		row += rowDir
		col += colDir
	}

	return placements
}

func LettersToRackTiles(letters string, rack scrubble.Rack) (tiles []scrubble.Tile) {
	lettersToFind := strings.Split(strings.ToUpper(letters), "")

LetterSearch:
	for _, l := range lettersToFind {
		letter := []rune(l)[0]

		for _, t := range rack {
			if (letter == '_' && t.Letter == ' ') || t.Letter == letter {
				tiles = append(tiles, t)
				continue LetterSearch
			}
		}

		tiles = append(tiles, scrubble.Tile{Letter: letter, Points: 0})
	}
	return
}

func Pass(game *scrubble.Game) {
	err := game.Pass()
	if err != nil {
		gt.Println(gt.Color(err.Error(), gt.RED))
	}
}

func PlayTiles(dir, row, col, letters string, game *scrubble.Game) {
	rowDir, colDir := 1, 0
	if dir == "across" {
		rowDir, colDir = 0, 1
	}

	rowNum, _ := strconv.Atoi(row)
	colNum, _ := strconv.Atoi(col)

	seat := game.CurrentSeat()
	placements := LettersToPlacements(rowDir, colDir, rowNum, colNum, letters, seat.Rack, &game.Board)

	_, err := game.Play(placements)
	if err != nil {
		gt.Println(gt.Color(err.Error(), gt.RED))
	} else {
		DrawRack(seat.Rack)
		if len(game.History.Last().TilesDrawn) > 0 {
			gt.Printf("\n\nTiles replenished from bag")
		}
	}
}

func ShuffleRack(game *scrubble.Game, rng *rand.Rand) {
	seat := game.CurrentSeat()

	rng.Shuffle(len(seat.Rack), func(i, j int) {
		seat.Rack[i], seat.Rack[j] = seat.Rack[j], seat.Rack[i]
	})
	DrawRack(seat.Rack)
}
