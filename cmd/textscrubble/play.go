package textscrubble

import (
	"strings"

	"github.com/mandykoh/scrubble"
)

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
	lettersToFind := strings.Split(letters, "")

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
