package main

import (
	"os"

	"time"

	"fmt"
	"math/rand"

	"bufio"

	"regexp"

	"strings"

	gt "github.com/buger/goterm"
	"github.com/mandykoh/scrubble"
	"github.com/mandykoh/scrubble/cmd/textscrubble"
)

func lettersToRackTiles(letters []string, rack scrubble.Rack) (tiles []scrubble.Tile) {

LetterSearch:
	for _, l := range letters {
		letter := []rune(l)[0]

		for _, t := range rack {
			if (letter == '_' && t.Letter == ' ') || t.Letter == letter {
				tiles = append(tiles, t)
				continue LetterSearch
			}
		}

		tiles = append(tiles, scrubble.Tile{Letter: letter, Points: 1})
	}
	return
}

func main() {
	cmdExchangePattern := regexp.MustCompile(`^exchange ([a-zA-Z_]+)$`)

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

		} else if matches := cmdExchangePattern.FindStringSubmatch(line); matches != nil {
			lettersToExchange := strings.Split(matches[1], "")
			tiles := lettersToRackTiles(lettersToExchange, seat.Rack)
			err := game.ExchangeTiles(tiles, rng)
			if err != nil {
				gt.Println(gt.Color(err.Error(), gt.RED))
			} else {
				gt.Printf("Tiles exchanged")
				textscrubble.DrawRack(seat.Rack)
			}

		} else if line == "?" {
			gt.Println("      rack - show rack")
			gt.Println("    across - play tiles across, eg: across 1 3 DOG")
			gt.Println("      down - play tiles down, eg: down 4 2 DOG")
			gt.Println("      pass - forfeit turn")
			gt.Println("  exchange - exchange tiles, eg: exchange DG")
			gt.Println(" challenge - challenge the last play")
		}
	}
}
