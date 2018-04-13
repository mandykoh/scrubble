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

// ValidateTilesFromRack checks the intended usage of tiles from a rack for
// legality. This includes that the rack actually has the tiles to be placed.
//
// If not, InsufficientTilesError is returned with the missing tiles in the
// Missing field.
//
// Otherwise, the remainder (after the placed tiles have been removed from the
// rack) is returned with no error, indicating that it would be safe to update
// the rack for placement.
func ValidateTilesFromRack(rack Rack, placements TilePlacements) (remaining Rack, err error) {
	var missing []Tile
	remaining = append(remaining, rack...)

Placements:
	for _, p := range placements {
		for i, t := range remaining {
			if t == p.Tile {
				remaining = append(remaining[:i], remaining[i+1:]...)
				continue Placements
			}
		}

		missing = append(missing, p.Tile)
	}

	if len(missing) != 0 {
		err = InsufficientTilesError{missing}
	}

	return
}
