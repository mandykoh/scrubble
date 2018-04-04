package scrubble

import "math"

// TilePlacements represents a set of tile placements on a board.
type TilePlacements []TilePlacement

// Bounds returns the minimum and maximum coordinates spanned by the placements.
func (tp TilePlacements) Bounds() (minRow, minCol, maxRow, maxCol int) {
	minRow = math.MaxInt32
	minCol = math.MaxInt32
	maxRow = math.MinInt32
	maxCol = math.MinInt32

	for _, p := range tp {
		if p.Row < minRow {
			minRow = p.Row
		}
		if p.Row > maxRow {
			maxRow = p.Row
		}
		if p.Column < minCol {
			minCol = p.Column
		}
		if p.Column > maxCol {
			maxCol = p.Column
		}
	}

	return
}
