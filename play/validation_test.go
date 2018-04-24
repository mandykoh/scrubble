package play

import (
	"testing"

	"github.com/mandykoh/scrubble/board"
	"github.com/mandykoh/scrubble/coord"
	"github.com/mandykoh/scrubble/tile"
)

func TestValidatePlacements(t *testing.T) {

	setupBoard := func() *board.Board {
		b := board.WithStandardLayout()
		return &b
	}

	t.Run("returns an error when no tiles are being played", func(t *testing.T) {
		b := setupBoard()

		err := ValidatePlacements(Tiles{}, b)

		if actual, expected := err, (InvalidTilePlacementError{Reason: NoTilesPlacedReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play zero tiles but got %v", expected, actual)
		}
	})

	t.Run("returns an error when any of the board positions is out of bounds", func(t *testing.T) {
		b := setupBoard()
		b.Position(coord.Make(0, 0)).Tile = &tile.Tile{Letter: 'A', Points: 1}

		err := ValidatePlacements(Tiles{{tile.Tile{Letter: 'B', Points: 1}, coord.Make(0, -1)}}, b)

		if actual, expected := err, (InvalidTilePlacementError{Reason: PlacementOutOfBoundsReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play tiles out of bounds but got %v", expected, actual)
		}

		err = ValidatePlacements(Tiles{{tile.Make('B', 1), coord.Make(b.Rows, 0)}}, b)

		if actual, expected := err, (InvalidTilePlacementError{Reason: PlacementOutOfBoundsReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play tiles out of bounds but got %v", expected, actual)
		}
	})

	t.Run("returns an error when any of the board positions is already occupied", func(t *testing.T) {
		b := setupBoard()
		b.Position(coord.Make(0, 0)).Tile = &tile.Tile{Letter: 'A', Points: 1}

		err := ValidatePlacements(Tiles{
			{tile.Make('B', 1), coord.Make(0, 0)},
			{tile.Make('A', 1), coord.Make(0, 1)},
			{tile.Make('D', 1), coord.Make(0, 2)},
		}, b)

		if actual, expected := err, (InvalidTilePlacementError{Reason: PositionOccupiedReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play tiles on an occupied position but got %v", expected, actual)
		}
	})

	t.Run("returns an error when the placements aren't in a straight line", func(t *testing.T) {
		b := setupBoard()

		err := ValidatePlacements(Tiles{
			{tile.Make('B', 1), coord.Make(0, 0)},
			{tile.Make('A', 1), coord.Make(0, 1)},
			{tile.Make('D', 1), coord.Make(1, 1)},
		}, b)

		if actual, expected := err, (InvalidTilePlacementError{Reason: PlacementNotLinearReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play tiles non-linearly but got %v", expected, actual)
		}
	})

	t.Run("returns an error when the placements overlap", func(t *testing.T) {
		b := setupBoard()

		err := ValidatePlacements(Tiles{
			{tile.Make('B', 1), coord.Make(0, 0)},
			{tile.Make('A', 1), coord.Make(0, 1)},
			{tile.Make('D', 1), coord.Make(0, 0)},
		}, b)

		if actual, expected := err, (InvalidTilePlacementError{Reason: PlacementOverlapReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play overlapping tiles but got %v", expected, actual)
		}
	})

	t.Run("returns an error when the placements aren't contiguous and would create gaps", func(t *testing.T) {
		b := setupBoard()

		err := ValidatePlacements(Tiles{
			{tile.Make('B', 1), coord.Make(0, 0)},
			{tile.Make('A', 1), coord.Make(0, 1)},
			{tile.Make('D', 1), coord.Make(0, 3)},
		}, b)

		if actual, expected := err, (InvalidTilePlacementError{Reason: PlacementNotContiguousReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play non-contiguous tiles but got %v", expected, actual)
		}
	})

	t.Run("returns an error when the placements aren't connected to at least one existing tile or on a starting position", func(t *testing.T) {
		b := setupBoard()
		b.Position(coord.Make(0, 0)).Tile = &tile.Tile{Letter: 'A', Points: 1}

		err := ValidatePlacements(Tiles{
			{tile.Make('M', 1), coord.Make(2, 0)},
			{tile.Make('A', 1), coord.Make(2, 1)},
			{tile.Make('D', 1), coord.Make(2, 2)},
		}, b)

		if actual, expected := err, (InvalidTilePlacementError{Reason: PlacementNotConnectedReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play non-connected tiles but got %v", expected, actual)
		}

		_, startPositionType, _, _, _, _ := board.AllPositionTypes()

		if actual, expected := b.Position(coord.Make(7, 7)).Type, startPositionType; actual != expected {
			t.Fatalf("Expected starting position at 7,7 but found %v", actual)
		}

		err = ValidatePlacements(Tiles{
			{tile.Make('M', 1), coord.Make(7, 6)},
			{tile.Make('A', 1), coord.Make(7, 7)},
			{tile.Make('D', 1), coord.Make(7, 8)},
		}, b)

		if actual := err; actual != nil {
			t.Errorf("Expected success when playing tiles on a start position but got error %v", actual)
		}
	})
}
