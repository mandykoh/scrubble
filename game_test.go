package scrubble

import (
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

	t.Run(".Play()", func(t *testing.T) {

		setupGame := func() (Game, *Player) {
			p := &Player{Name: "Alice"}

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
				Board: BoardWithStandardLayout(),
				Seats: []Seat{
					{
						OccupiedBy: p,
						Rack:       Rack(rackTiles),
					},
				},
			}

			return game, p
		}

		t.Run("returns an error when the game is not in the Main phase", func(t *testing.T) {
			game := Game{
				Phase: SetupPhase,
			}

			err := game.Play([]TilePlacement{})

			if actual, expected := err, (GameOutOfPhaseError{MainPhase, SetupPhase}); actual != expected {
				t.Fatalf("Expected error %v but was %v", expected, err)
			}
		})

		t.Run("returns an error when no tiles are being played", func(t *testing.T) {
			game, _ := setupGame()

			err := game.Play([]TilePlacement{})

			if actual, expected := err, (InvalidTilePlacementError{}); actual != expected {
				t.Errorf("Expected %v when attempting to play zero tiles but got %v", expected, actual)
			}
		})

		t.Run("returns an error when the current player doesn't have the required tiles", func(t *testing.T) {
			game, _ := setupGame()

			playTiles := []Tile{
				{'B', 1},
				{'O', 1},
				{'O', 1},
				{'M', 1},
				{'S', 1},
			}

			var placements []TilePlacement
			for i, t := range playTiles {
				placements = append(placements, TilePlacement{t, 0, i})
			}

			err := game.Play(placements)

			if insufficientError, ok := err.(InsufficientTilesError); !ok {
				t.Errorf("Expected InsufficientTilesError but got %v", err)
			} else {
				if actual, expected := len(insufficientError.Missing), 2; actual != expected {
					t.Errorf("Expected error to indicate one missing tile but was %d", actual)
				} else {
					if actual, expected := insufficientError.Missing[0], (Tile{'O', 1}); actual != expected {
						t.Errorf("Expected missing tile %c(%d) but was %c(%d)", expected.Letter, expected.Points, actual.Letter, actual.Points)
					}
					if actual, expected := insufficientError.Missing[1], (Tile{'S', 1}); actual != expected {
						t.Errorf("Expected missing tile %c(%d) but was %c(%d)", expected.Letter, expected.Points, actual.Letter, actual.Points)
					}
				}
			}

			if actual, expected := len(game.Seats[0].Rack), 6; actual != expected {
				t.Errorf("Expected player to still have %d tiles but found %d", expected, actual)
			}
		})

		t.Run("places tiles on the board for a valid play", func(t *testing.T) {
			game, _ := setupGame()

			placements := []TilePlacement{
				{Tile{'B', 1}, 0, 0},
				{Tile{'A', 1}, 0, 1},
				{Tile{'D', 1}, 0, 2},
			}
			err := game.Play(placements)

			if err != nil {
				t.Errorf("Expected play to succeed but got error %v", err)
			}

			rack := game.currentSeat().Rack

			if actual, expected := len(rack), 3; actual != expected {
				t.Errorf("Expected %d tiles to remain in the player's rack but found %d", expected, actual)
			} else {
				if actual, expected := rack[0], (Tile{'E', 1}); actual != expected {
					t.Errorf("Expected first remaining tile in rack to be %c(%d) but found %c(%d)", expected.Letter, expected.Points, actual.Letter, actual.Points)
				}
				if actual, expected := rack[1], (Tile{'O', 1}); actual != expected {
					t.Errorf("Expected second remaining tile in rack to be %c(%d) but found %c(%d)", expected.Letter, expected.Points, actual.Letter, actual.Points)
				}
				if actual, expected := rack[2], (Tile{'M', 1}); actual != expected {
					t.Errorf("Expected third remaining tile in rack to be %c(%d) but found %c(%d)", expected.Letter, expected.Points, actual.Letter, actual.Points)
				}
			}

			for _, p := range placements {
				if actual := game.Board.Position(p.Row, p.Column).Tile; actual == nil || *actual != p.Tile {
					if actual == nil {
						t.Errorf("Expected tile %c(%d) to be in position %d,%d but got <nil>", p.Tile.Letter, p.Tile.Points, p.Row, p.Column)
					} else {
						t.Errorf("Expected tile %c(%d) to be in position %d,%d but got %c(%d)", p.Tile.Letter, p.Tile.Points, p.Row, p.Column, actual.Letter, actual.Points)
					}
				}
			}
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
			for i, expected := range expectedBag {
				if actual := game.Bag[i]; actual != expected {
					t.Errorf("Expected tile %d to be %c(%d) but was %c(%d)", i, expected.Letter, expected.Points, actual.Letter, actual.Points)
				}
			}
		})

		t.Run("fills the players' racks from the bag", func(t *testing.T) {
			for i := 0; i < 3; i++ {
				if actual, expected := len(game.Seats[i].Rack), len(expectedRacks[i]); actual != expected {
					t.Fatalf("Expected player '%s' to have a full rack of %d tiles but found %d", game.Seats[i].OccupiedBy.Name, expected, actual)
				}
				for j, expected := range expectedRacks[i] {
					if actual := game.Seats[i].Rack[j]; actual != expected {
						t.Errorf("Expected tile %d of %s's rack to be %c(%d) but was %c(%d)", j, game.Seats[i].OccupiedBy.Name, expected.Letter, expected.Points, actual.Letter, actual.Points)
					}
				}

				if actual, expected := len(game.Bag), len(expectedBag); actual != expected {
					t.Errorf("Expected bag to have %d tiles after filling racks but found %d", expected, actual)
				}
			}
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

			if actual, expected := len(game.Bag), len(BagWithStandardEnglishTiles()); actual != expected {
				t.Errorf("Expected bag to still have %d tiles after error but found %d", expected, actual)
			}
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
