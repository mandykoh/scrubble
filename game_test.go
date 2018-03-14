package scrubble

import (
	"math/rand"
	"testing"
	"time"
)

func TestGame(t *testing.T) {

	t.Run("zero-value is in SetupPhase", func(t *testing.T) {
		var game Game

		if actual, expected := game.Phase, SetupPhase; actual != expected {
			t.Errorf("Expected zero-value game to be in %s phase, but was %s", expected, actual)
		}
	})

	t.Run(".AddPlayer()", func(t *testing.T) {

		t.Run("adds a new seat for each player", func(t *testing.T) {
			var game Game

			if actual, expected := len(game.Seats), 0; actual != expected {
				t.Errorf("Expected zero seats to begin with but found %d", actual)
			}

			p1 := &Player{"Alice"}
			game.AddPlayer(p1)

			if actual, expected := len(game.Seats), 1; actual != expected {
				t.Errorf("Expected one seat after adding a player but found %d", actual)
			}
			if actual, expected := game.Seats[0].OccupiedBy, p1; actual != expected {
				t.Errorf("Expected first seat to be occupied by player %s but was %+v", expected.Name, actual)
			}

			p2 := &Player{"Bob"}
			game.AddPlayer(p2)

			if actual, expected := len(game.Seats), 2; actual != expected {
				t.Errorf("Expected %d seats after adding another player but found %d", expected, actual)
			}
			if actual, expected := game.Seats[1].OccupiedBy, p2; actual != expected {
				t.Errorf("Expected first seat to be occupied by player %s but was %+v", expected.Name, actual)
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

			game.RemovePlayer(p2)

			if actual, expected := len(game.Seats), 2; actual != expected {
				t.Errorf("Expected two seats after removing a player but found %d", actual)
			}
			if actual, expected := game.Seats[0].OccupiedBy, p1; actual != expected {
				t.Errorf("Expected remaining seat to be occupied by player %s but was %+v", expected.Name, actual)
			}
			if actual, expected := game.Seats[1].OccupiedBy, p3; actual != expected {
				t.Errorf("Expected remaining seat to be occupied by player %s but was %+v", expected.Name, actual)
			}

			game.RemovePlayer(p1)

			if actual, expected := len(game.Seats), 1; actual != expected {
				t.Errorf("Expected one seat after removing a player but found %d", actual)
			}
			if actual, expected := game.Seats[0].OccupiedBy, p3; actual != expected {
				t.Errorf("Expected remaining seat to be occupied by player %s but was %+v", expected.Name, actual)
			}

			game.RemovePlayer(p3)

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

			game.RemovePlayer(&Player{"Carol"})

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
		game.Start(rand.New(rand.NewSource(seed)))

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

		t.Run("fills the players' racks", func(t *testing.T) {
			for i := 0; i < 3; i++ {
				if actual, expected := len(game.Seats[i].Rack), len(expectedRacks[i]); actual != expected {
					t.Fatalf("Expected player '%s' to have a full rack of %d tiles but found %d", game.Seats[i].OccupiedBy.Name, expected, actual)
				}
				for j, expected := range expectedRacks[i] {
					if actual := game.Seats[i].Rack[j]; actual != expected {
						t.Errorf("Expected tile %d of %s's rack to be %c(%d) but was %c(%d)", j, game.Seats[i].OccupiedBy.Name, expected.Letter, expected.Points, actual.Letter, actual.Points)
					}
				}
			}
		})
	})
}
