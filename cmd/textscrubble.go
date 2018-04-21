package main

import (
	"os"

	"time"

	"fmt"
	"math/rand"

	"bufio"

	"regexp"

	"strings"

	"strconv"

	gt "github.com/buger/goterm"
	"github.com/mandykoh/scrubble"
	"github.com/mandykoh/scrubble/cmd/textscrubble"
)

func lettersToPlacements(rowDir, colDir, row, col int, letters string, rack scrubble.Rack, board *scrubble.Board) scrubble.TilePlacements {
	var placements scrubble.TilePlacements
	tiles := lettersToRackTiles(letters, rack)

	row -= rowDir
	col -= colDir

	for _, tile := range tiles {
		var pos *scrubble.BoardPosition
		for pos == nil || pos.Tile != nil {
			row += rowDir
			col += colDir
			pos = board.Position(scrubble.Coord{Row: row, Column: col})
		}

		placements = append(placements, scrubble.TilePlacement{
			Tile:  tile,
			Coord: scrubble.Coord{Row: row, Column: col},
		})
	}

	return placements
}

func lettersToRackTiles(letters string, rack scrubble.Rack) (tiles []scrubble.Tile) {
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

func main() {
	cmdExchangePattern := regexp.MustCompile(`^exchange ([a-zA-Z_]+)$`)
	cmdPlayPattern := regexp.MustCompile(`^(across|down) (\d+) (\d+) ([a-zA-Z_]+)$`)

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	game := &scrubble.Game{
		Board: scrubble.BoardWithStandardLayout(),
		Bag:   scrubble.BagWithStandardEnglishTiles(),
	}

	for i := 1; i < len(os.Args); i++ {
		player := &scrubble.Player{Name: os.Args[i]}
		game.AddPlayer(player)
	}

	err := game.Start(rng)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting game: %v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		seat := game.CurrentSeat()
		textscrubble.DrawGame(game)

		gt.Println()

		if game.Phase == scrubble.EndPhase {
			break
		}

		scanner.Scan()
		line := scanner.Text()

		if line == "rack" {
			textscrubble.DrawRack(seat.Rack)

		} else if line == "pass" {
			err := game.Pass()
			if err != nil {
				gt.Println(gt.Color(err.Error(), gt.RED))
			}

		} else if line == "shuffle" {
			rng.Shuffle(len(seat.Rack), func(i, j int) {
				seat.Rack[i], seat.Rack[j] = seat.Rack[j], seat.Rack[i]
			})
			textscrubble.DrawRack(seat.Rack)

		} else if matches := cmdExchangePattern.FindStringSubmatch(line); matches != nil {
			tiles := lettersToRackTiles(matches[1], seat.Rack)
			err := game.ExchangeTiles(tiles, rng)
			if err != nil {
				gt.Println(gt.Color(err.Error(), gt.RED))
			} else {
				textscrubble.DrawRack(seat.Rack)
				gt.Printf("\n\nTiles exchanged")
			}

		} else if matches := cmdPlayPattern.FindStringSubmatch(line); matches != nil {
			rowDir, colDir := 1, 0
			if matches[1] == "across" {
				rowDir, colDir = 0, 1
			}

			row, _ := strconv.Atoi(matches[2])
			col, _ := strconv.Atoi(matches[3])

			placements := lettersToPlacements(rowDir, colDir, row, col, matches[4], seat.Rack, &game.Board)
			_, err := game.Play(placements)
			if err != nil {
				gt.Println(gt.Color(err.Error(), gt.RED))
			} else {
				textscrubble.DrawRack(seat.Rack)
			}

		} else if line == "?" {
			gt.Println("      rack - show rack")
			gt.Println("    across - play tiles across, eg: across 1 3 DOG")
			gt.Println("      down - play tiles down, eg: down 4 2 DOG")
			gt.Println("      pass - forfeit turn")
			gt.Println("   shuffle - shuffle rack")
			gt.Println("  exchange - exchange tiles, eg: exchange DG")
			gt.Println(" challenge - challenge the last play")
		}
	}
}
