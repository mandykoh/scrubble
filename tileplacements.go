package scrubble

import (
	"math"

	"github.com/mandykoh/scrubble/tile"
)

// TilePlacements represents a set of tile placements on a board.
type TilePlacements []TilePlacement

// Bounds returns the minimum and maximum coordinates spanned by the placements.
func (tp TilePlacements) Bounds() CoordRange {
	bounds := CoordRange{
		Coord{math.MaxInt32, math.MaxInt32},
		Coord{math.MinInt32, math.MinInt32},
	}

	for _, p := range tp {
		bounds = bounds.Include(p.Coord)
	}

	return bounds
}

// Find returns the first placement corresponding to the given coordinate, or
// nil if no matching placement exists.
func (tp TilePlacements) Find(c Coord) *TilePlacement {
	for i := range tp {
		if tp[i].Coord == c {
			return &tp[i]
		}
	}
	return nil
}

// Tiles returns the collection of tiles being placed.
func (tp TilePlacements) Tiles() []tile.Tile {
	tiles := make([]tile.Tile, len(tp))
	for i, p := range tp {
		tiles[i] = p.Tile
	}
	return tiles
}

func (tp *TilePlacements) take(c Coord) *TilePlacement {
	for i := range *tp {
		p := (*tp)[i]
		if p.Coord == c {
			*tp = append((*tp)[:i], (*tp)[i+1:]...)
			return &p
		}
	}
	return nil
}

func (tp *TilePlacements) takeLast() *TilePlacement {
	length := len(*tp)
	if length == 0 {
		return nil
	}

	p := &(*tp)[length-1]
	*tp = (*tp)[:length-1]
	return p
}
