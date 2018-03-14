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
