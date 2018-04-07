package scrubble

import "testing"

func TestValidatePlacements(t *testing.T) {

	setupBoard := func() *Board {
		board := BoardWithStandardLayout()
		return &board
	}

	t.Run("returns an error when no tiles are being played", func(t *testing.T) {
		board := setupBoard()

		err := ValidatePlacements(TilePlacements{}, board)

		if actual, expected := err, (InvalidTilePlacementError{NoTilesPlacedReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play zero tiles but got %v", expected, actual)
		}
	})

	t.Run("returns an error when any of the board positions is out of bounds", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{0, 0}).Tile = &Tile{'A', 1}

		err := ValidatePlacements(TilePlacements{{Tile{'B', 1}, Coord{0, -1}}}, board)

		if actual, expected := err, (InvalidTilePlacementError{PlacementOutOfBoundsReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play tiles out of bounds but got %v", expected, actual)
		}

		err = ValidatePlacements(TilePlacements{{Tile{'B', 1}, Coord{board.Rows, 0}}}, board)

		if actual, expected := err, (InvalidTilePlacementError{PlacementOutOfBoundsReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play tiles out of bounds but got %v", expected, actual)
		}
	})

	t.Run("returns an error when any of the board positions is already occupied", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{0, 0}).Tile = &Tile{'A', 1}

		err := ValidatePlacements(TilePlacements{
			{Tile{'B', 1}, Coord{0, 0}},
			{Tile{'A', 1}, Coord{0, 1}},
			{Tile{'D', 1}, Coord{0, 2}},
		}, board)

		if actual, expected := err, (InvalidTilePlacementError{PositionOccupiedReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play tiles on an occupied position but got %v", expected, actual)
		}
	})

	t.Run("returns an error when the placements aren't in a straight line", func(t *testing.T) {
		board := setupBoard()

		err := ValidatePlacements(TilePlacements{
			{Tile{'B', 1}, Coord{0, 0}},
			{Tile{'A', 1}, Coord{0, 1}},
			{Tile{'D', 1}, Coord{1, 1}},
		}, board)

		if actual, expected := err, (InvalidTilePlacementError{PlacementNotLinearReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play tiles non-linearly but got %v", expected, actual)
		}
	})

	t.Run("returns an error when the placements overlap", func(t *testing.T) {
		board := setupBoard()

		err := ValidatePlacements(TilePlacements{
			{Tile{'B', 1}, Coord{0, 0}},
			{Tile{'A', 1}, Coord{0, 1}},
			{Tile{'D', 1}, Coord{0, 0}},
		}, board)

		if actual, expected := err, (InvalidTilePlacementError{PlacementOverlapReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play overlapping tiles but got %v", expected, actual)
		}
	})

	t.Run("returns an error when the placements aren't contiguous and would create gaps", func(t *testing.T) {
		board := setupBoard()

		err := ValidatePlacements(TilePlacements{
			{Tile{'B', 1}, Coord{0, 0}},
			{Tile{'A', 1}, Coord{0, 1}},
			{Tile{'D', 1}, Coord{0, 3}},
		}, board)

		if actual, expected := err, (InvalidTilePlacementError{PlacementNotContiguousReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play non-contiguous tiles but got %v", expected, actual)
		}
	})

	t.Run("returns an error when the placements aren't connected to at least one existing tile or on a starting position", func(t *testing.T) {
		board := setupBoard()
		board.Position(Coord{0, 0}).Tile = &Tile{'A', 1}

		err := ValidatePlacements(TilePlacements{
			{Tile{'M', 1}, Coord{2, 0}},
			{Tile{'A', 1}, Coord{2, 1}},
			{Tile{'D', 1}, Coord{2, 2}},
		}, board)

		if actual, expected := err, (InvalidTilePlacementError{PlacementNotConnectedReason}); actual != expected {
			t.Errorf("Expected %v when attempting to play non-connected tiles but got %v", expected, actual)
		}

		if actual, expected := board.Position(Coord{7, 7}).Type, startPositionTypeInstance; actual != expected {
			t.Fatalf("Expected starting position at 7,7 but found %v", actual)
		}

		err = ValidatePlacements(TilePlacements{
			{Tile{'M', 1}, Coord{7, 6}},
			{Tile{'A', 1}, Coord{7, 7}},
			{Tile{'D', 1}, Coord{7, 8}},
		}, board)

		if actual := err; actual != nil {
			t.Errorf("Expected success when playing tiles on a start position but got error %v", actual)
		}
	})
}
