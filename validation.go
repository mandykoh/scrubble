package scrubble

import (
	"github.com/mandykoh/scrubble/coord"
	"github.com/mandykoh/scrubble/play"
)

// ValidatePlacements checks the intended placement of tiles on a board for
// legality. This includes: that at least one tile is placed, that tiles are
// placed contiguously, that tiles are placed only in a straight line, that
// there are no gaps created in the result, that tiles do not overlap with each
// other or with tiles already on the board, and that no tiles are placed out of
// bounds.
//
// If any violations are detected, InvalidTilePlacementError is returned with
// the reason indicating the violation.
//
// Otherwise, nil is returned, indicating that it would be safe to place the
// given tiles on the board (word validity not withstanding).
func ValidatePlacements(placements play.Tiles, board *Board) error {
	placementsLeft := len(placements)
	if placementsLeft == 0 {
		return play.InvalidTilePlacementError{Reason: play.NoTilesPlacedReason}
	}

	bounds := placements.Bounds()
	if !bounds.IsLinear() {
		return play.InvalidTilePlacementError{Reason: play.PlacementNotLinearReason}
	}

	connected := false

	err := bounds.Each(func(c coord.Coord) error {
		position := board.Position(c)
		if position == nil {
			return play.InvalidTilePlacementError{Reason: play.PlacementOutOfBoundsReason}
		}

		if placement := placements.Find(c); placement != nil {
			if position.Tile != nil {
				return play.InvalidTilePlacementError{Reason: play.PositionOccupiedReason}
			}

			connected = connected || position.Type.CountsAsConnected() || board.neighbourHasTile(c)
			placementsLeft--

		} else if position.Tile == nil {
			return play.InvalidTilePlacementError{Reason: play.PlacementNotContiguousReason}
		}

		return nil
	})
	if err != nil {
		return err
	}

	if placementsLeft != 0 {
		return play.InvalidTilePlacementError{Reason: play.PlacementOverlapReason}
	}
	if !connected {
		return play.InvalidTilePlacementError{Reason: play.PlacementNotConnectedReason}
	}

	return nil
}
