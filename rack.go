package scrubble

const MaxRackTiles = 7

// Rack represents a player’s rack of tiles which are available to play. The
// zero-value of a Rack is an empty rack.
type Rack []Tile

// FillFromBag fills the rack up to MaxRackTiles tiles by drawing from the
// specified bag. If the bag holds less than the required number of tiles, all
// are added to the rack.
func (r *Rack) FillFromBag(b *Bag) {
	for len(*r) < MaxRackTiles && len(*b) > 0 {
		*r = append(*r, b.DrawTile())
	}
}

func (r *Rack) tryPlayTiles(placements TilePlacements) (remaining Rack, missing []Tile) {
	remaining = append(remaining, *r...)

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

	return
}
