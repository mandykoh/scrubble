package scrubble

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
func ValidatePlacements(placements TilePlacements, board *Board) error {
	placementsLeft := len(placements)
	if placementsLeft == 0 {
		return InvalidTilePlacementError{Reason: NoTilesPlacedReason}
	}

	bounds := placements.Bounds()
	if !bounds.IsLinear() {
		return InvalidTilePlacementError{Reason: PlacementNotLinearReason}
	}

	connected := false

	err := bounds.EachCoord(func(c Coord) error {
		position := board.Position(c)
		if position == nil {
			return InvalidTilePlacementError{Reason: PlacementOutOfBoundsReason}
		}

		if placement := placements.Find(c); placement != nil {
			if position.Tile != nil {
				return InvalidTilePlacementError{Reason: PositionOccupiedReason}
			}

			connected = connected || position.Type.CountsAsConnected() || board.neighbourHasTile(c)
			placementsLeft--

		} else if position.Tile == nil {
			return InvalidTilePlacementError{Reason: PlacementNotContiguousReason}
		}

		return nil
	})
	if err != nil {
		return err
	}

	if placementsLeft != 0 {
		return InvalidTilePlacementError{Reason: PlacementOverlapReason}
	}
	if !connected {
		return InvalidTilePlacementError{Reason: PlacementNotConnectedReason}
	}

	return nil
}
