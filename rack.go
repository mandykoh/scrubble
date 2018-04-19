package scrubble

// MaxRackTiles is the maximum number of tiles a Rack should hold.
const MaxRackTiles = 7

// Rack represents a player’s rack of tiles which are available to play. The
// zero-value of a Rack is an empty rack.
type Rack []Tile

// FillFromBag fills the rack up to MaxRackTiles tiles by drawing from the
// specified bag. If the bag holds less than the required number of tiles, all
// are added to the rack.
func (r *Rack) FillFromBag(b *Bag) (drawn []Tile) {
	for needed := MaxRackTiles - len(*r); needed > 0 && len(*b) > 0; needed-- {
		drawn = append(drawn, b.DrawTile())
	}
	*r = append(*r, drawn...)
	return
}

// Remove removes the specified tiles from the rack. Any tiles that are not in
// the rack are ignored.
func (r *Rack) Remove(tiles ...Tile) {
	for _, t := range tiles {
		for i, rt := range *r {
			if rt == t {
				*r = append((*r)[0:i], (*r)[i+1:]...)
				break
			}
		}
	}
}
