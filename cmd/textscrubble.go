package main

import (
	"os"

	"time"

	"fmt"
	"math/rand"

	"bufio"

	"regexp"

	"strconv"

	gt "github.com/buger/goterm"
	"github.com/mandykoh/scrubble"
	"github.com/mandykoh/scrubble/cmd/textscrubble"
)

func main() {
	if len(os.Args) < 3 || (os.Args[1] != "simple" && os.Args[1] != "challenge") {
		fmt.Fprintf(os.Stderr, "Usage: textscrubble <mode> <player1_name> [player2_name] ... [playerN_name]\n")
		fmt.Fprintf(os.Stderr, "\n  <mode> can be:\n\n")
		fmt.Fprintf(os.Stderr, "     simple - words are automatically validated against the dictionary (only valid words can be played)\n")
		fmt.Fprintf(os.Stderr, "  challenge - players can manually challenge a play (which is then validated with a dictionary)\n")
		os.Exit(1)
	}

	challengeEnabled := os.Args[1] == "challenge"

	cmdExchangePattern := regexp.MustCompile(`^exchange ([a-zA-Z_]+)$`)
	cmdPlayPattern := regexp.MustCompile(`^(across|down) (\d+) (\d+) ([a-zA-Z_]+)$`)

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	game := scrubble.NewGameWithDefaults()
	game.Rules = game.Rules.WithDictionaryForScoring(!challengeEnabled)

	for i := 2; i < len(os.Args); i++ {
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
			tiles := textscrubble.LettersToRackTiles(matches[1], seat.Rack)
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

			placements := textscrubble.LettersToPlacements(rowDir, colDir, row, col, matches[4], seat.Rack, &game.Board)
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

			if challengeEnabled {
				gt.Println(" challenge - challenge the last play")
			}
		}
	}
}
