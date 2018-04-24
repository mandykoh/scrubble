package textscrubble

import (
	"strings"

	"strconv"

	"math/rand"

	gt "github.com/buger/goterm"
	"github.com/mandykoh/scrubble/board"
	"github.com/mandykoh/scrubble/coord"
	"github.com/mandykoh/scrubble/game"
	"github.com/mandykoh/scrubble/play"
	"github.com/mandykoh/scrubble/tile"
)

type Player struct {
	Name string
}

func Challenge(g *game.Game, rng *rand.Rand) {
	challengerIndex := g.CurrentSeatIndex

	success, err := g.Challenge(challengerIndex, rng)
	if err != nil {
		gt.Println(gt.Color(err.Error(), gt.RED))
	} else if success {
		gt.Printf(gt.Color("\n\nPlay successfully challenged!", gt.GREEN))
	} else {
		gt.Printf(gt.Color("\n\nChallenge failed! All words are valid", gt.RED))
	}
}

func ExchangeTiles(letters string, g *game.Game, rng *rand.Rand) {
	seat := g.CurrentSeat()
	tiles := LettersToRackTiles(letters, seat.Rack)

	err := g.ExchangeTiles(tiles, rng)
	if err != nil {
		gt.Println(gt.Color(err.Error(), gt.RED))
	} else {
		DrawRack(seat.Rack)
		gt.Printf("\n\nTiles exchanged")
	}
}

func LettersToPlacements(rowDir, colDir, row, col int, letters string, rack tile.Rack, b *board.Board) play.Tiles {
	var placements play.Tiles
	tiles := LettersToRackTiles(letters, rack)

	for _, t := range tiles {
		pos := b.Position(coord.Make(row, col))

		for pos != nil && pos.Tile != nil {
			row += rowDir
			col += colDir
			pos = b.Position(coord.Make(row, col))
		}

		placements = append(placements, play.TilePlacement{
			Tile:  t,
			Coord: coord.Make(row, col),
		})
		row += rowDir
		col += colDir
	}

	return placements
}

func LettersToRackTiles(letters string, rack tile.Rack) (tiles []tile.Tile) {
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

		tiles = append(tiles, tile.Tile{Letter: letter, Points: 0})
	}
	return
}

func Pass(g *game.Game) {
	err := g.Pass()
	if err != nil {
		gt.Println(gt.Color(err.Error(), gt.RED))
	}
}

func PlayTiles(dir, row, col, letters string, g *game.Game) {
	rowDir, colDir := 1, 0
	if dir == "across" {
		rowDir, colDir = 0, 1
	}

	rowNum, _ := strconv.Atoi(row)
	colNum, _ := strconv.Atoi(col)

	seat := g.CurrentSeat()
	placements := LettersToPlacements(rowDir, colDir, rowNum, colNum, letters, seat.Rack, &g.Board)

	_, err := g.Play(placements)
	if err != nil {
		gt.Println(gt.Color(err.Error(), gt.RED))
	} else {
		DrawRack(seat.Rack)
		if len(g.History.Last().TilesDrawn) > 0 {
			gt.Printf("\n\nTiles replenished from bag")
		}
	}
}

func ShuffleRack(g *game.Game, rng *rand.Rand) {
	seat := g.CurrentSeat()

	rng.Shuffle(len(seat.Rack), func(i, j int) {
		seat.Rack[i], seat.Rack[j] = seat.Rack[j], seat.Rack[i]
	})
	DrawRack(seat.Rack)
}
