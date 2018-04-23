package scrubble

import (
	"strings"
	"testing"

	"github.com/mandykoh/scrubble/coord"
	"github.com/mandykoh/scrubble/play"
	"github.com/mandykoh/scrubble/positiontype"
	"github.com/mandykoh/scrubble/tile"
)

func TestIsChallengeSuccessful(t *testing.T) {
	dictionary := func(word string) (valid bool) {
		return strings.HasPrefix(word, "VALIDWORD")
	}

	t.Run("returns success when any played words are invalid", func(t *testing.T) {
		success := IsChallengeSuccessful([]play.Word{
			{Word: "VALIDWORD1"},
			{Word: "INVALIDWORD"},
		}, dictionary)

		if !success {
			t.Errorf("Expected challenge to succeed but it will fail")
		}
	})

	t.Run("returns failure when all played words are valid", func(t *testing.T) {
		success := IsChallengeSuccessful([]play.Word{
			{Word: "VALIDWORD1"},
			{Word: "VALIDWORD2"},
			{Word: "VALIDWORD3"},
		}, dictionary)

		if success {
			t.Errorf("Expected challenge to fail but it will succeed")
		}
	})
}

func TestValidatePlacements(t *testing.T) {

	setupBoard := func() *Board {
		board := BoardWithStandardLayout()
		return &board
	}

	t.Run("returns an error when no tiles are being played", func(t *testing.T) {
		board := setupBoard()

		err := ValidatePlacements(play.Tiles{}, board)

		if actual, expected := err, (play.InvalidTilePlacementError{Reason: play.NoTilesPlacedReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play zero tiles but got %v", expected, actual)
		}
	})

	t.Run("returns an error when any of the board positions is out of bounds", func(t *testing.T) {
		board := setupBoard()
		board.Position(coord.Make(0, 0)).Tile = &tile.Tile{Letter: 'A', Points: 1}

		err := ValidatePlacements(play.Tiles{{tile.Tile{Letter: 'B', Points: 1}, coord.Make(0, -1)}}, board)

		if actual, expected := err, (play.InvalidTilePlacementError{Reason: play.PlacementOutOfBoundsReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play tiles out of bounds but got %v", expected, actual)
		}

		err = ValidatePlacements(play.Tiles{{tile.Make('B', 1), coord.Make(board.Rows, 0)}}, board)

		if actual, expected := err, (play.InvalidTilePlacementError{Reason: play.PlacementOutOfBoundsReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play tiles out of bounds but got %v", expected, actual)
		}
	})

	t.Run("returns an error when any of the board positions is already occupied", func(t *testing.T) {
		board := setupBoard()
		board.Position(coord.Make(0, 0)).Tile = &tile.Tile{Letter: 'A', Points: 1}

		err := ValidatePlacements(play.Tiles{
			{tile.Make('B', 1), coord.Make(0, 0)},
			{tile.Make('A', 1), coord.Make(0, 1)},
			{tile.Make('D', 1), coord.Make(0, 2)},
		}, board)

		if actual, expected := err, (play.InvalidTilePlacementError{Reason: play.PositionOccupiedReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play tiles on an occupied position but got %v", expected, actual)
		}
	})

	t.Run("returns an error when the placements aren't in a straight line", func(t *testing.T) {
		board := setupBoard()

		err := ValidatePlacements(play.Tiles{
			{tile.Make('B', 1), coord.Make(0, 0)},
			{tile.Make('A', 1), coord.Make(0, 1)},
			{tile.Make('D', 1), coord.Make(1, 1)},
		}, board)

		if actual, expected := err, (play.InvalidTilePlacementError{Reason: play.PlacementNotLinearReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play tiles non-linearly but got %v", expected, actual)
		}
	})

	t.Run("returns an error when the placements overlap", func(t *testing.T) {
		board := setupBoard()

		err := ValidatePlacements(play.Tiles{
			{tile.Make('B', 1), coord.Make(0, 0)},
			{tile.Make('A', 1), coord.Make(0, 1)},
			{tile.Make('D', 1), coord.Make(0, 0)},
		}, board)

		if actual, expected := err, (play.InvalidTilePlacementError{Reason: play.PlacementOverlapReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play overlapping tiles but got %v", expected, actual)
		}
	})

	t.Run("returns an error when the placements aren't contiguous and would create gaps", func(t *testing.T) {
		board := setupBoard()

		err := ValidatePlacements(play.Tiles{
			{tile.Make('B', 1), coord.Make(0, 0)},
			{tile.Make('A', 1), coord.Make(0, 1)},
			{tile.Make('D', 1), coord.Make(0, 3)},
		}, board)

		if actual, expected := err, (play.InvalidTilePlacementError{Reason: play.PlacementNotContiguousReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play non-contiguous tiles but got %v", expected, actual)
		}
	})

	t.Run("returns an error when the placements aren't connected to at least one existing tile or on a starting position", func(t *testing.T) {
		board := setupBoard()
		board.Position(coord.Make(0, 0)).Tile = &tile.Tile{Letter: 'A', Points: 1}

		err := ValidatePlacements(play.Tiles{
			{tile.Make('M', 1), coord.Make(2, 0)},
			{tile.Make('A', 1), coord.Make(2, 1)},
			{tile.Make('D', 1), coord.Make(2, 2)},
		}, board)

		if actual, expected := err, (play.InvalidTilePlacementError{Reason: play.PlacementNotConnectedReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play non-connected tiles but got %v", expected, actual)
		}

		_, startPositionType, _, _, _, _ := positiontype.All()

		if actual, expected := board.Position(coord.Make(7, 7)).Type, startPositionType; actual != expected {
			t.Fatalf("Expected starting position at 7,7 but found %v", actual)
		}

		err = ValidatePlacements(play.Tiles{
			{tile.Make('M', 1), coord.Make(7, 6)},
			{tile.Make('A', 1), coord.Make(7, 7)},
			{tile.Make('D', 1), coord.Make(7, 8)},
		}, board)

		if actual := err; actual != nil {
			t.Errorf("Expected success when playing tiles on a start position but got error %v", actual)
		}
	})
}

func TestValidateTilesFromRack(t *testing.T) {

	expectRackContains := func(t *testing.T, r tile.Rack, letters ...rune) {
		if actual, expected := len(r), len(letters); actual != expected {
			t.Fatalf("Expected rack to contain %d tiles but found %d", expected, actual)
		}

		for i, expected := range letters {
			if actual := r[i].Letter; actual != expected {
				t.Errorf("Expected letter '%c' on the rack but found '%c' instead", expected, actual)
			}
		}
	}

	t.Run("returns missing tiles when the rack has insufficient tiles for the play", func(t *testing.T) {
		r := tile.Rack{
			{'A', 1},
			{'B', 1},
			{'O', 1},
			{'M', 1},
		}

		used, remaining, err := ValidateTilesFromRack(r, []tile.Tile{
			{'B', 1},
			{'O', 1},
			{'O', 1},
			{'M', 1},
			{'S', 1},
		})

		switch e := err.(type) {

		case InsufficientTilesError:
			expectTiles(t, "missing", e.Missing,
				tile.Tile{'O', 1},
				tile.Tile{'S', 1},
			)

		default:
			t.Errorf("Expected an InsufficientTilesError but got %v", err)
		}

		expectRackContains(t, used, 'B', 'O', 'M')
		expectRackContains(t, remaining, 'A')
		expectRackContains(t, r, 'A', 'B', 'O', 'M')
	})

	t.Run("returns no missing tiles and the remainder if successful", func(t *testing.T) {
		r := tile.Rack{
			{'A', 1},
			{'O', 1},
			{'M', 1},
			{'B', 1},
			{'O', 1},
		}

		used, remaining, err := ValidateTilesFromRack(r, []tile.Tile{
			{'B', 1},
			{'O', 1},
			{'O', 1},
			{'M', 1},
		})

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			expectRackContains(t, used, 'B', 'O', 'O', 'M')
			expectRackContains(t, remaining, 'A')
			expectRackContains(t, r, 'A', 'O', 'M', 'B', 'O')
		}
	})

	t.Run("treats zero-point tiles as wildcard tiles", func(t *testing.T) {
		r := tile.Rack{
			{'A', 1},
			{' ', 0},
			{'M', 1},
			{'B', 1},
			{'O', 1},
		}

		used, remaining, err := ValidateTilesFromRack(r, []tile.Tile{
			{'B', 1},
			{'O', 1},
			{'O', 0},
			{'M', 1},
		})

		if err != nil {
			t.Errorf("Expected success but got error %v", err)
		} else {
			expectRackContains(t, used, 'B', 'O', ' ', 'M')
			expectRackContains(t, remaining, 'A')
		}
	})
}
