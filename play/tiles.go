package play

import (
	"math"

	"github.com/mandykoh/scrubble/board"
	"github.com/mandykoh/scrubble/coord"
	"github.com/mandykoh/scrubble/tile"
)

// Tiles represents a set of tile placements on a board.
type Tiles []TilePlacement

// Bounds returns the minimum and maximum coordinates spanned by the placements.
func (tp Tiles) Bounds() coord.Range {
	bounds := coord.Range{
		Min: coord.Make(math.MaxInt32, math.MaxInt32),
		Max: coord.Make(math.MinInt32, math.MinInt32),
	}

	for _, p := range tp {
		bounds = bounds.Include(p.Coord)
	}

	return bounds
}

// Find returns the first placement corresponding to the given coordinate, or
// nil if no matching placement exists.
func (tp Tiles) Find(c coord.Coord) *TilePlacement {
	for i := range tp {
		if tp[i].Coord == c {
			return &tp[i]
		}
	}
	return nil
}

// Place sets the specified board positions to the specified tiles.
func (tp *Tiles) Place(b *board.Board) {
	for _, p := range *tp {
		t := p.Tile
		b.Position(p.Coord).Tile = &t
	}
}

// Take removes and returns a placement for the specified coordinate, returning
// nil if no tile is being placed at that coordinate.
func (tp *Tiles) Take(c coord.Coord) *TilePlacement {
	for i := range *tp {
		p := (*tp)[i]
		if p.Coord == c {
			*tp = append((*tp)[:i], (*tp)[i+1:]...)
			return &p
		}
	}
	return nil
}

// TakeLast removes the last placement and returns it, or nil if there are no
// placements left.
func (tp *Tiles) TakeLast() *TilePlacement {
	length := len(*tp)
	if length == 0 {
		return nil
	}

	p := &(*tp)[length-1]
	*tp = (*tp)[:length-1]
	return p
}

// Tiles returns the collection of tiles being placed.
func (tp Tiles) Tiles() []tile.Tile {
	tiles := make([]tile.Tile, len(tp))
	for i, p := range tp {
		tiles[i] = p.Tile
	}
	return tiles
}
