package tile

// ValidateFromRack checks the intended usage of tiles from a rack for
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
func ValidateFromRack(rack Rack, toPlay []Tile) (used, remaining []Tile, err error) {
	var missing []Tile
	used = make([]Tile, 0, len(toPlay))
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
