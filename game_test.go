package scrubble

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGame(t *testing.T) {

	t.Run("zero-value", func(t *testing.T) {
		var game Game

		if actual, expected := game.Phase, SetupPhase; actual != expected {
			t.Errorf("Expected zero-value game to be in %s phase, but was %s", expected, actual)
		}
		if actual, expected := len(game.Seats), 0; actual != expected {
			t.Errorf("Expected zero-value game to have no players/seats, but found %d", actual)
		}
	})

	t.Run(".AddPlayer()", func(t *testing.T) {

		t.Run("adds a new seat for each player", func(t *testing.T) {
			var game Game

			if actual, expected := len(game.Seats), 0; actual != expected {
				t.Errorf("Expected zero seats to begin with but found %d", actual)
			}

			p1 := &Player{"Alice"}
			err := game.AddPlayer(p1)

			if actual, expected := err, error(nil); actual != expected {
				t.Fatalf("Expected adding a player to succeed but got error: %v", err)
			}
			if actual, expected := len(game.Seats), 1; actual != expected {
				t.Errorf("Expected one seat after adding a player but found %d", actual)
			}
			if actual, expected := game.Seats[0].OccupiedBy, p1; actual != expected {
				t.Errorf("Expected first seat to be occupied by player %s but was %+v", expected.Name, actual)
			}

			p2 := &Player{"Bob"}
			err = game.AddPlayer(p2)

			if actual, expected := err, error(nil); actual != expected {
				t.Fatalf("Expected adding a player to succeed but got error: %v", err)
			}
			if actual, expected := len(game.Seats), 2; actual != expected {
				t.Errorf("Expected %d seats after adding another player but found %d", expected, actual)
			}
			if actual, expected := game.Seats[1].OccupiedBy, p2; actual != expected {
				t.Errorf("Expected first seat to be occupied by player %s but was %+v", expected.Name, actual)
			}
		})

		t.Run("returns an error when game is not in setup phase", func(t *testing.T) {
			game := Game{
				Phase: MainPhase,
			}

			err := game.AddPlayer(&Player{"Alice"})

			if actual, expected := err, (GameOutOfPhaseError{SetupPhase, MainPhase}); actual != expected {
				t.Fatalf("Expected error %v but was %v", expected, err)
			}
		})
	})

	t.Run(".Challenge()", func(t *testing.T) {

		setupGame := func() Game {
			game := Game{
				Phase: MainPhase,
				Bag:   BagWithStandardEnglishTiles(),
				Board: BoardWithStandardLayout(),
				Seats: []Seat{
					{
						Rack: Rack{
							{'K', 1},
							{'P', 1},
							{'Q', 1},
							{'Z', 1},
							{'T', 1},
							{'R', 1},
							{'W', 1},
						},
						Score: 123,
					},
					{
						Rack: Rack{
							{'D', 1},
							{'A', 1},
							{'B', 1},
							{'E', 1},
							{'O', 1},
							{'M', 1},
						},
						Score: 456,
					},
				},
				CurrentSeatIndex: 1,
				History: History{
					{
						SeatIndex:   0,
						Score:       123,
						TilesSpent:  []Tile{{'A', 1}, {'D', 1}},
						TilesPlayed: TilePlacements{{Tile{'A', 1}, Coord{0, 0}}, {Tile{'D', 1}, Coord{0, 1}}},
						TilesDrawn:  []Tile{{'R', 1}, {'W', 1}},
						WordsFormed: nil,
					},
				},
			}

			for _, placement := range game.History.Last().TilesPlayed {
				tile := placement.Tile
				game.Board.Position(placement.Coord).Tile = &tile
			}

			return game
		}

		t.Run("returns an error when the game is not in the Main phase", func(t *testing.T) {
			game := setupGame()
			game.Phase = SetupPhase

			err := game.Challenge(game.CurrentSeatIndex, nil)

			if actual, expected := err, (GameOutOfPhaseError{MainPhase, SetupPhase}); actual != expected {
				t.Fatalf("Expected error %v but was %v", expected, err)
			}
		})

		t.Run("when successful", func(t *testing.T) {
			seed := time.Now().UnixNano()

			game := setupGame()
			lastTurn := game.History.Last()

			expectedBag := append(Bag{}, game.Bag...)
			expectedBag = append(expectedBag, lastTurn.TilesDrawn...)
			expectedBag.Shuffle(rand.New(rand.NewSource(seed)))

			err := game.Challenge(game.CurrentSeatIndex, rand.New(rand.NewSource(seed)))

			t.Run("doesn't return an error", func(t *testing.T) {
				if err != nil {
					t.Errorf("Expected success but got error %v", err)
				}
			})

			t.Run("reduces the challenged player's score", func(t *testing.T) {
				if actual, expected := game.prevSeat().Score, 0; actual != expected {
					t.Errorf("Expected challenged player's score to be reduced to %d but was %d", expected, actual)
				}
			})

			t.Run("returns drawn tiles to the bag", func(t *testing.T) {
				expectTiles(t, "bagged", game.Bag, expectedBag...)
			})

			t.Run("withdraws placed tiles from the board", func(t *testing.T) {
				for _, placed := range lastTurn.TilesPlayed {
					pos := game.Board.Position(placed.Coord)

					if pos.Tile != nil {
						t.Errorf("Expected tile in position %v to have been withdrawn but found %v there", placed.Coord, pos.Tile)
					}
				}
			})

			t.Run("restores the player's rack to how it was before the play", func(t *testing.T) {
				expectTiles(t, "racked", game.prevSeat().Rack,
					Tile{'K', 1},
					Tile{'P', 1},
					Tile{'Q', 1},
					Tile{'Z', 1},
					Tile{'T', 1},
					Tile{'A', 1},
					Tile{'D', 1})
			})
		})
	})

	t.Run(".ExchangeTiles()", func(t *testing.T) {
		tilesFromRackValidated := 0

		setupGame := func() Game {
			tilesFromRackValidated = 0

			game := Game{
				Phase: MainPhase,
				Bag:   BagWithStandardEnglishTiles(),
				Board: BoardWithStandardLayout(),
				Seats: []Seat{
					{},
					{
						Rack: Rack{
							{'D', 1},
							{'A', 1},
							{'B', 1},
							{'E', 1},
							{'O', 1},
							{'M', 1},
						},
					},
				},
				CurrentSeatIndex: 1,
				Rules: Rules{
					rackValidator: func(rack Rack, toPlay []Tile) ([]Tile, []Tile, error) {
						tilesFromRackValidated++
						return ValidateTilesFromRack(rack, toPlay)
					},
				},
			}

			return game
		}

		t.Run("returns an error when the game is not in the Main phase", func(t *testing.T) {
			game := setupGame()
			game.Phase = SetupPhase

			err := game.ExchangeTiles([]Tile{
				game.CurrentSeat().Rack[0],
			}, nil)

			if actual, expected := err, (GameOutOfPhaseError{MainPhase, SetupPhase}); actual != expected {
				t.Fatalf("Expected error %v but was %v", expected, err)
			}
		})

		t.Run("returns an error when no tiles are being exchanged", func(t *testing.T) {
			game := setupGame()

			err := game.ExchangeTiles(nil, nil)

			if actual, expected := err, (InvalidTileExchangeError{NoTilesExchangedReason}); actual != expected {
				t.Fatalf("Expected error %v but was %v", expected, err)
			}
		})

		t.Run("returns an error when insufficient tiles are in the bag", func(t *testing.T) {
			game := setupGame()
			game.Bag = game.Bag[:MaxRackTiles-1]

			err := game.ExchangeTiles([]Tile{
				game.CurrentSeat().Rack[0],
			}, nil)

			if actual, expected := err, (InvalidTileExchangeError{InsufficientTilesInBagReason}); actual != expected {
				t.Fatalf("Expected error %v but was %v", expected, err)
			}
		})

		t.Run("with insufficient tiles on the rack", func(t *testing.T) {
			game := setupGame()

			game.Rules.rackValidator = func(rack Rack, toPlay []Tile) ([]Tile, []Tile, error) {
				return nil, rack, errors.New("some error")
			}

			tiles := []Tile{
				{'B', 1},
				{'O', 1},
				{'O', 1},
				{'M', 1},
				{'S', 1},
			}

			err := game.ExchangeTiles(tiles, nil)

			t.Run("returns an error", func(t *testing.T) {
				if err == nil {
					t.Errorf("Expected an error but didn't get one")
				} else {
					if actual, expected := err.Error(), "some error"; actual != expected {
						t.Errorf("Expected an error from tile rack validation but got %v", actual)
					}
				}
			})

			t.Run("does not remove tiles from the player's rack", func(t *testing.T) {
				if actual, expected := len(game.CurrentSeat().Rack), 6; actual != expected {
					t.Errorf("Expected player to still have %d tiles but found %d", expected, actual)
				}
			})
		})

		t.Run("with a valid exchange", func(t *testing.T) {
			seed := time.Now().UnixNano()

			game := setupGame()

			originalBagSize := len(game.Bag)
			originalRackSize := len(game.CurrentSeat().Rack)

			expectedBag := append(Bag{}, game.Bag...)
			nextBagTiles := []Tile{
				expectedBag.DrawTile(),
				expectedBag.DrawTile(),
				expectedBag.DrawTile(),
			}

			tilesExchanged := []Tile{
				{'B', 1},
				{'M', 1},
				{'D', 1},
			}

			expectedBag = append(expectedBag, tilesExchanged...)
			expectedBag.Shuffle(rand.New(rand.NewSource(seed)))
			nextBagTiles = append(nextBagTiles, expectedBag.DrawTile())

			err := game.ExchangeTiles(tilesExchanged, rand.New(rand.NewSource(seed)))

			t.Run("doesn't return an error", func(t *testing.T) {
				if err != nil {
					t.Errorf("Expected play to succeed but got error %v", err)
				}
			})

			t.Run("validates tiles used from the rack using the set rules", func(t *testing.T) {
				if actual, expected := tilesFromRackValidated, 1; actual != expected {
					t.Errorf("Expected tile rack validation to have been invoked once but was called %d times", actual)
				}
			})

			t.Run("replenishes the player's rack from the bag", func(t *testing.T) {
				rack := game.prevSeat().Rack

				if actual, expected := len(rack), MaxRackTiles; actual != expected {
					t.Errorf("Expected player's rack to have been replenished to %d tiles but found %d", expected, actual)
				} else {
					expectTiles(t, "racked", rack,
						Tile{'A', 1},
						Tile{'E', 1},
						Tile{'O', 1},
						nextBagTiles[0],
						nextBagTiles[1],
						nextBagTiles[2],
						nextBagTiles[3],
					)
				}
			})

			t.Run("returns the exchanged tiles to the bag", func(t *testing.T) {
				if actual, expected := len(game.Bag), originalBagSize-(MaxRackTiles-originalRackSize); actual != expected {
					t.Errorf("Expected %d tiles to still be in bag but found %d", expected, actual)
				}
			})

			t.Run("moves to next player's turn", func(t *testing.T) {
				if actual, expected := game.CurrentSeatIndex, 0; actual != expected {
					t.Errorf("Expected turn to move to next player but current seat is %d", game.CurrentSeatIndex)
				}
			})

			t.Run("records a history entry", func(t *testing.T) {
				expectHistory(t, game.History,
					HistoryEntry{1, 0, tilesExchanged, nil, nextBagTiles[len(nextBagTiles)-1:], nil},
				)
			})
		})

		t.Run("with a game-ending play (eg final consecutive scoreless turn)", func(t *testing.T) {
			game := setupGame()
			game.Rules = game.Rules.WithGamePhaseController(func(*Game) GamePhase {
				return EndPhase
			})

			err := game.ExchangeTiles([]Tile{
				game.CurrentSeat().Rack[0],
			}, rand.New(rand.NewSource(0)))

			t.Run("succeeds", func(t *testing.T) {
				if err != nil {
					t.Errorf("Expected success but got error %v", err)
				}
			})

			t.Run("moves the game into the End phase", func(t *testing.T) {
				if actual, expected := game.Phase, EndPhase; actual != expected {
					t.Errorf("Expected game to be in %v phase but was %v", expected, actual)
				}
			})
		})
	})

	t.Run(".Pass()", func(t *testing.T) {

		setupGame := func() Game {
			game := Game{
				Phase: MainPhase,
				Bag:   BagWithStandardEnglishTiles(),
				Board: BoardWithStandardLayout(),
				Seats: []Seat{
					{},
					{},
				},
				CurrentSeatIndex: 1,
			}

			return game
		}

		t.Run("returns an error when the game is not in the Main phase", func(t *testing.T) {
			game := Game{
				Phase: SetupPhase,
			}

			err := game.Pass()

			if actual, expected := err, (GameOutOfPhaseError{MainPhase, SetupPhase}); actual != expected {
				t.Fatalf("Expected error %v but was %v", expected, err)
			}
		})

		t.Run("records a history entry", func(t *testing.T) {
			game := setupGame()
			game.CurrentSeat().Rack = Rack{
				{'A', 1},
				{'B', 1},
				{'C', 1},
				{'D', 1},
				{'E', 1},
				{'F', 1},
				{'G', 1},
			}

			err := game.Pass()

			if err != nil {
				t.Errorf("Expected success but got error %v", err)
			}

			expectHistory(t, game.History,
				HistoryEntry{1, 0, nil, nil, nil, nil},
			)
		})

		t.Run("moves to next player's turn", func(t *testing.T) {
			game := setupGame()

			err := game.Pass()

			if err != nil {
				t.Errorf("Expected success but got error %v", err)
			}
			if actual, expected := game.CurrentSeatIndex, 0; actual != expected {
				t.Errorf("Expected turn to move to next player but current seat is %d", game.CurrentSeatIndex)
			}
		})

		t.Run("with a game-ending play (eg final consecutive scoreless turn)", func(t *testing.T) {
			game := setupGame()
			game.Rules = game.Rules.WithGamePhaseController(func(*Game) GamePhase {
				return EndPhase
			})

			err := game.Pass()

			t.Run("succeeds", func(t *testing.T) {
				if err != nil {
					t.Errorf("Expected success but got error %v", err)
				}
			})

			t.Run("moves the game into the End phase", func(t *testing.T) {
				if actual, expected := game.Phase, EndPhase; actual != expected {
					t.Errorf("Expected game to be in %v phase but was %v", expected, actual)
				}
			})
		})
	})

	t.Run(".Play()", func(t *testing.T) {
		placementsValidated := 0
		tilesFromRackValidated := 0
		wordsScored := 0

		setupGame := func() Game {
			placementsValidated = 0
			tilesFromRackValidated = 0
			wordsScored = 0

			p1 := &Player{Name: "Alice"}
			p2 := &Player{Name: "Bob"}

			rackTiles := []Tile{
				{'A', 1},
				{'B', 1},
				{'E', 1},
				{'O', 1},
				{'D', 1},
				{'M', 1},
			}

			game := Game{
				Phase: MainPhase,
				Bag:   BagWithStandardEnglishTiles(),
				Board: BoardWithStandardLayout(),
				Seats: []Seat{
					{OccupiedBy: p1, Rack: append(Rack{}, rackTiles...)},
					{OccupiedBy: p2, Rack: append(Rack{}, rackTiles...)},
				},
				CurrentSeatIndex: 1,
				Rules: Rules{
					placementValidator: func(placements TilePlacements, board *Board) error {
						placementsValidated++
						return nil
					},
					rackValidator: func(rack Rack, toPlay []Tile) ([]Tile, []Tile, error) {
						tilesFromRackValidated++
						return ValidateTilesFromRack(rack, toPlay)
					},
					wordScorer: func(placements TilePlacements, board *Board, dictionary Dictionary) (score int, words []PlayedWord, err error) {
						wordsScored++
						words = append(words, PlayedWord{Word: "SOMEWORD", Score: 123, CoordRange: placements.Bounds()})
						return 123, words, nil
					},
				},
			}

			return game
		}

		t.Run("returns an error when the game is not in the Main phase", func(t *testing.T) {
			game := Game{
				Phase: SetupPhase,
			}

			_, err := game.Play(TilePlacements{
				{Tile{'A', 1}, Coord{7, 7}},
			})

			if actual, expected := err, (GameOutOfPhaseError{MainPhase, SetupPhase}); actual != expected {
				t.Fatalf("Expected error %v but was %v", expected, err)
			}
		})

		t.Run("with insufficient tiles on the rack", func(t *testing.T) {
			game := setupGame()

			game.Rules.rackValidator = func(rack Rack, toPlay []Tile) ([]Tile, []Tile, error) {
				return nil, rack, errors.New("some error")
			}

			playTiles := []Tile{
				{'B', 1},
				{'O', 1},
				{'O', 1},
				{'M', 1},
				{'S', 1},
			}

			var placements TilePlacements
			for i, t := range playTiles {
				placements = append(placements, TilePlacement{t, Coord{7, 7 + i}})
			}

			_, err := game.Play(placements)

			t.Run("returns an error", func(t *testing.T) {
				if err == nil {
					t.Errorf("Expected an error but didn't get one")
				} else {
					if actual, expected := err.Error(), "some error"; actual != expected {
						t.Errorf("Expected an error from tile rack validation but got %v", actual)
					}
				}
			})

			t.Run("does not remove tiles from the player's rack", func(t *testing.T) {
				expectTiles(t, "racked", game.CurrentSeat().Rack,
					Tile{'A', 1},
					Tile{'B', 1},
					Tile{'E', 1},
					Tile{'O', 1},
					Tile{'D', 1},
					Tile{'M', 1},
				)
			})
		})

		t.Run("with invalid tile placement", func(t *testing.T) {
			game := setupGame()

			game.Rules.placementValidator = func(placements TilePlacements, board *Board) error {
				return errors.New("some error")
			}

			placements := TilePlacements{
				{Tile{'B', 1}, Coord{0, 0}},
				{Tile{'D', 1}, Coord{0, 2}},
			}
			_, err := game.Play(placements)

			t.Run("returns an error", func(t *testing.T) {
				if err == nil {
					t.Errorf("Expected an error but didn't get one")
				} else {
					if actual, expected := err.Error(), "some error"; actual != expected {
						t.Errorf("Expected an error from placement validation but got %v", actual)
					}
				}
			})

			t.Run("does not place tiles on the board", func(t *testing.T) {
				for _, p := range placements {
					if actual := game.Board.Position(p.Coord).Tile; actual != nil {
						t.Errorf("Expected no tiles to be placed but found %v in position %d,%d", actual, p.Row, p.Column)
					}
				}
			})
		})

		t.Run("with a valid play", func(t *testing.T) {
			game := setupGame()
			game.Board.Position(Coord{0, 1}).Tile = &Tile{'A', 1}

			nextBagTiles := []Tile{
				game.Bag[len(game.Bag)-1],
				game.Bag[len(game.Bag)-2],
				game.Bag[len(game.Bag)-3],
			}

			placements := TilePlacements{
				{Tile{'B', 1}, Coord{0, 0}},
				{Tile{'D', 1}, Coord{0, 2}},
			}
			playedWords, err := game.Play(placements)

			t.Run("doesn't return an error", func(t *testing.T) {
				if err != nil {
					t.Errorf("Expected play to succeed but got error %v", err)
				}
			})

			t.Run("validates tiles used from the rack using the set rules", func(t *testing.T) {
				if actual, expected := tilesFromRackValidated, 1; actual != expected {
					t.Errorf("Expected tile rack validation to have been invoked once but was called %d times", actual)
				}
			})

			t.Run("validates tile placement using the set rules", func(t *testing.T) {
				if actual, expected := placementsValidated, 1; actual != expected {
					t.Errorf("Expected tile placement validation to have been invoked once but was called %d times", actual)
				}
			})

			t.Run("scores the formed words using the set rules", func(t *testing.T) {
				if actual, expected := wordsScored, 1; actual != expected {
					t.Errorf("Expected word scoring to have been invoked once but was called %d times", actual)
				}
				if actual, expected := game.prevSeat().Score, 123; actual != expected {
					t.Errorf("Expected word score of %d to have been added to player total but was %d", expected, actual)
				}
			})

			t.Run("places tiles on the board", func(t *testing.T) {
				for _, p := range placements {
					if actual := game.Board.Position(p.Coord).Tile; actual == nil || *actual != p.Tile {
						t.Errorf("Expected tile %v to be in position %d,%d but got %v", p.Tile, p.Row, p.Column, actual)
					}
				}
			})

			t.Run("replenishes the player's rack from the bag", func(t *testing.T) {
				rack := game.prevSeat().Rack

				expectTiles(t, "racked", rack,
					Tile{'A', 1},
					Tile{'E', 1},
					Tile{'O', 1},
					Tile{'M', 1},
					nextBagTiles[0],
					nextBagTiles[1],
					nextBagTiles[2],
				)
			})

			t.Run("moves to next player's turn", func(t *testing.T) {
				if actual, expected := game.CurrentSeatIndex, 0; actual != expected {
					t.Errorf("Expected turn to move to next player but current seat is %d", game.CurrentSeatIndex)
				}
			})

			t.Run("records a history entry", func(t *testing.T) {
				expectHistory(t, game.History,
					HistoryEntry{1, 123, placements.Tiles(), placements, nextBagTiles, []PlayedWord{
						{"SOMEWORD", 123, placements.Bounds()},
					}},
				)
			})

			t.Run("returns the words formed from scoring using the set rules", func(t *testing.T) {
				expectPlayedWords(t, playedWords, PlayedWord{"SOMEWORD", 123, placements.Bounds()})
			})
		})

		t.Run("with a game-ending play", func(t *testing.T) {
			game := setupGame()
			game.Rules = game.Rules.WithGamePhaseController(func(*Game) GamePhase {
				return EndPhase
			})

			game.Board.Position(Coord{0, 1}).Tile = &Tile{'A', 1}

			placements := TilePlacements{
				{Tile{'B', 1}, Coord{0, 0}},
				{Tile{'D', 1}, Coord{0, 2}},
			}
			game.Play(placements)

			t.Run("moves the game into the End phase", func(t *testing.T) {
				if actual, expected := game.Phase, EndPhase; actual != expected {
					t.Errorf("Expected game to be in %v phase but was %v", expected, actual)
				}
			})
		})
	})

	t.Run(".RemovePlayer()", func(t *testing.T) {

		t.Run("removes the seat for the specified player", func(t *testing.T) {
			var game Game

			p1 := &Player{"Alice"}
			game.AddPlayer(p1)

			p2 := &Player{"Bob"}
			game.AddPlayer(p2)

			p3 := &Player{"Carol"}
			game.AddPlayer(p3)

			err := game.RemovePlayer(p2)

			if actual, expected := err, error(nil); actual != expected {
				t.Fatalf("Expected removing a player to succeed but got error: %v", err)
			}
			if actual, expected := len(game.Seats), 2; actual != expected {
				t.Errorf("Expected two seats after removing a player but found %d", actual)
			}
			if actual, expected := game.Seats[0].OccupiedBy, p1; actual != expected {
				t.Errorf("Expected remaining seat to be occupied by player %s but was %+v", expected.Name, actual)
			}
			if actual, expected := game.Seats[1].OccupiedBy, p3; actual != expected {
				t.Errorf("Expected remaining seat to be occupied by player %s but was %+v", expected.Name, actual)
			}

			err = game.RemovePlayer(p1)

			if actual, expected := err, error(nil); actual != expected {
				t.Fatalf("Expected removing a player to succeed but got error: %v", err)
			}
			if actual, expected := len(game.Seats), 1; actual != expected {
				t.Errorf("Expected one seat after removing a player but found %d", actual)
			}
			if actual, expected := game.Seats[0].OccupiedBy, p3; actual != expected {
				t.Errorf("Expected remaining seat to be occupied by player %s but was %+v", expected.Name, actual)
			}

			err = game.RemovePlayer(p3)

			if actual, expected := err, error(nil); actual != expected {
				t.Fatalf("Expected removing a player to succeed but got error: %v", err)
			}
			if actual, expected := len(game.Seats), 0; actual != expected {
				t.Errorf("Expected no seats after removing a player but found %d", actual)
			}
		})

		t.Run("has no effect if the specified player doesn't have a seat", func(t *testing.T) {
			var game Game

			p1 := &Player{"Alice"}
			game.AddPlayer(p1)

			p2 := &Player{"Bob"}
			game.AddPlayer(p2)

			err := game.RemovePlayer(&Player{"Carol"})

			if actual, expected := err, error(nil); actual != expected {
				t.Fatalf("Expected removing a player to succeed but got error: %v", err)
			}
			if actual, expected := len(game.Seats), 2; actual != expected {
				t.Errorf("Expected %d seats to remain but found %d", expected, actual)
			}
			if actual, expected := game.Seats[0].OccupiedBy, p1; actual != expected {
				t.Errorf("Expected seat to still be occupied by player %s but was %+v", expected.Name, actual)
			}
			if actual, expected := game.Seats[1].OccupiedBy, p2; actual != expected {
				t.Errorf("Expected seat to still be occupied by player %s but was %+v", expected.Name, actual)
			}
		})

		t.Run("returns an error when game is not in setup phase", func(t *testing.T) {
			game := Game{
				Phase: MainPhase,
			}

			p := &Player{"Alice"}

			game.AddPlayer(p)
			err := game.RemovePlayer(p)

			if actual, expected := err, (GameOutOfPhaseError{SetupPhase, MainPhase}); actual != expected {
				t.Fatalf("Expected error %v but was %v", expected, err)
			}
		})
	})

	t.Run(".Start()", func(t *testing.T) {
		seed := time.Now().UnixNano()

		expectedRand := rand.New(rand.NewSource(seed))
		expectedStartingSeat := expectedRand.Intn(3)

		expectedBag := BagWithStandardEnglishTiles()
		expectedBag.Shuffle(expectedRand)

		var expectedRacks []Rack
		for i := 0; i < 3; i++ {
			var rack Rack
			for j := 0; j < MaxRackTiles; j++ {
				rack = append(rack, expectedBag.DrawTile())
			}
			expectedRacks = append(expectedRacks, rack)
		}

		p1 := &Player{"Alice"}
		p2 := &Player{"Bob"}
		p3 := &Player{"Carol"}

		game := Game{
			Bag:   BagWithStandardEnglishTiles(),
			Board: BoardWithStandardLayout(),
		}
		game.AddPlayer(p1)
		game.AddPlayer(p2)
		game.AddPlayer(p3)
		err := game.Start(rand.New(rand.NewSource(seed)))

		t.Run("succeeds", func(t *testing.T) {
			if actual, expected := err, error(nil); actual != expected {
				t.Fatalf("Expected game to be started but got error: %v", actual)
			}
		})

		t.Run("sets the phase to Main", func(t *testing.T) {
			if actual, expected := game.Phase, MainPhase; actual != expected {
				t.Errorf("Expected game to be in %s phase but was %s", expected, actual)
			}
		})

		t.Run("picks a random starting seat", func(t *testing.T) {
			if actual, expected := game.CurrentSeatIndex, expectedStartingSeat; actual != expected {
				t.Errorf("Expected starting seat to be %d but was %d", expected, actual)
			}
		})

		t.Run("shuffles the bag", func(t *testing.T) {
			expectTiles(t, "bagged", game.Bag, expectedBag...)
		})

		t.Run("fills the players' racks from the bag", func(t *testing.T) {
			for i := 0; i < 3; i++ {
				expectTiles(t, fmt.Sprintf("racked (player %d)", i), game.Seats[i].Rack, expectedRacks[i]...)
			}
			expectTiles(t, "bagged", game.Bag, expectedBag...)
		})

		t.Run("returns an error if not in Setup phase", func(t *testing.T) {
			game := Game{
				Bag:   BagWithStandardEnglishTiles(),
				Phase: MainPhase,
			}

			err := game.Start(rand.New(rand.NewSource(seed)))

			if actual, expected := err, (GameOutOfPhaseError{SetupPhase, MainPhase}); actual != expected {
				t.Errorf("Expected %v but got %v", expected, actual)
			}

			expectTiles(t, "bagged", game.Bag, BagWithStandardEnglishTiles()...)
		})

		t.Run("returns an error if there are no players", func(t *testing.T) {
			game := Game{
				Bag: BagWithStandardEnglishTiles(),
			}

			err := game.Start(rand.New(rand.NewSource(seed)))

			if actual, expected := err, (NotEnoughPlayersError{GameMinPlayers, 0}); actual != expected {
				t.Errorf("Expected %v but got %v", expected, actual)
			}
			if actual, expected := game.Phase, SetupPhase; actual != expected {
				t.Errorf("Expected game to still be in %s phase but was in %s instead", expected, actual)
			}

			game.AddPlayer(p1)
			err = game.Start(rand.New(rand.NewSource(seed)))

			if actual, expected := err, error(nil); actual != expected {
				t.Errorf("Expected no error but got %v", actual)
			}
		})
	})
}
