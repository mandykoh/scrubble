package main

import (
	"os"

	"time"

	"fmt"
	"math/rand"

	"bufio"

	"regexp"

	gt "github.com/buger/goterm"
	"github.com/mandykoh/scrubble/cmd/textscrubble/textscrubble"
	"github.com/mandykoh/scrubble/game"
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

	g := game.NewWithDefaults()
	g.Rules = g.Rules.WithDictionaryForScoring(!challengeEnabled)

	var players []textscrubble.Player

	for i := 2; i < len(os.Args); i++ {
		players = append(players, textscrubble.Player{Name: os.Args[i]})
		g.AddPlayer()
	}

	err := g.Start(rng)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting game: %v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		s := g.CurrentSeat()
		textscrubble.DrawGame(g, players)

		gt.Println()

		scanner.Scan()
		line := scanner.Text()

		if line == "rack" {
			textscrubble.DrawRack(s.Rack)

		} else if line == "pass" {
			textscrubble.Pass(g)

		} else if line == "shuffle" {
			textscrubble.ShuffleRack(g, rng)

		} else if challengeEnabled && line == "challenge" {
			textscrubble.Challenge(g, rng)

		} else if matches := cmdExchangePattern.FindStringSubmatch(line); matches != nil {
			textscrubble.ExchangeTiles(matches[1], g, rng)

		} else if matches := cmdPlayPattern.FindStringSubmatch(line); matches != nil {
			textscrubble.PlayTiles(matches[1], matches[2], matches[3], matches[4], g)

		} else if line == "?" {
			gt.Println("      rack - show rack")
			gt.Println("    across - play tiles across from a starting row/col, eg: across 1 3 dg")
			gt.Println("      down - play tiles down from a starting row/col, eg: down 4 2 dg")
			gt.Println("      pass - forfeit turn")
			gt.Println("   shuffle - shuffle rack")
			gt.Println("  exchange - exchange tiles, eg: exchange dg")

			if challengeEnabled {
				gt.Println(" challenge - challenge the last play")
			}

			gt.Println("\n  When specifying tiles, blank tiles will be matched if no other tiles match")
		}
	}
}
