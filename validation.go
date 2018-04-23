package scrubble

import (
	"github.com/mandykoh/scrubble/coord"
	"github.com/mandykoh/scrubble/play"
	"github.com/mandykoh/scrubble/tile"
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

// ValidateTilesFromRack checks the intended usage of tiles from a rack for
// legality. This includes that the rack actually has the tiles to be placed.
//
// If not, InsufficientTilesError is returned with the missing tiles in the
// Missing field.
//
// Otherwise, the tiles used and the remainder (after the placed tiles have been
// removed from the rack) is returned with no error, indicating that it would be
// safe to update the rack for placement.
//
// Zero-point tiles are treated as wildcards: a zero-point tile being placed
// matches any zero-point tile in the rack regardless of letter. This implies
// that wildcards need to have their letters replaced with the desired letter
// when being placed (so that any resulting words are valid).
func ValidateTilesFromRack(rack tile.Rack, toPlay []tile.Tile) (used, remaining []tile.Tile, err error) {
	var missing []tile.Tile
	used = make([]tile.Tile, 0, len(toPlay))
	remaining = append(remaining, rack...)

Placements:
	for _, p := range toPlay {
		for i, t := range remaining {
			if (t.Points == 0 && p.Points == 0) || t == p {
				used = append(used, t)
				remaining = append(remaining[:i], remaining[i+1:]...)
				continue Placements
			}
		}

		missing = append(missing, p)
	}

	if len(missing) != 0 {
		err = InsufficientTilesError{missing}
	}

	return
}
