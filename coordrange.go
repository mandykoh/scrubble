package scrubble

// CoordRange represents a bounded, rectangular set of coordinates.
type CoordRange struct {
	Min Coord
	Max Coord
}

// EachCoord runs the specified function for each coordinate in the range,
// stopping if an error is returned and returning the error.
func (r *CoordRange) EachCoord(op func(Coord) error) error {
	for row := r.Min.Row; row <= r.Max.Row; row++ {
		for col := r.Min.Column; col <= r.Max.Column; col++ {
			err := op(Coord{row, col})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Include returns a new CoordRange which is the result of including the
// specified coordinate within the current range.
func (r *CoordRange) Include(c Coord) CoordRange {
	return CoordRange{r.Min.Min(c), r.Max.Max(c)}
}

// Includes returns true if the specified coordinate falls within this range.
func (r *CoordRange) Includes(c Coord) bool {
	return r.Min.Row <= c.Row && c.Row <= r.Max.Row &&
		r.Min.Column <= c.Column && c.Column <= r.Max.Column
}

// IsLinear returns true if the range represents a straight line of coordinates
// (as opposed to a rectangular area).
func (r *CoordRange) IsLinear() bool {
	return r.Min.Row == r.Max.Row || r.Min.Column == r.Max.Column
}
