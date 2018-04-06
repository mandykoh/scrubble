package scrubble

import "math"

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
