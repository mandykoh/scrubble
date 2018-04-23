package scrubble

import "github.com/mandykoh/scrubble/positiontype"

// Board represents a game board, which is a grid of positions on which tiles
// can be placed. The zero-value of a Board is a zero-sized board.
type Board struct {
	Rows      int
	Columns   int
	Positions []BoardPosition
}

// BoardWithLayout creates a board with no tiles, with the specified layout.
func BoardWithLayout(layout positiontype.Layout) Board {
	normal, _, _, _, _, _ := positiontype.All()

	rows := len(layout)
	columns := layout.WidestRow()

	b := Board{
		Rows:      rows,
		Columns:   columns,
		Positions: make([]BoardPosition, rows*columns),
	}

	for row, lRow := range layout {

		// Set position types according to the specified layout
		for col, posType := range lRow {
			b.Position(Coord{row, col}).Type = posType
		}

		// Fill in any unspecified remainder of the row with Normal positions
		for col := len(lRow); col < columns; col++ {
			b.Position(Coord{row, col}).Type = normal
		}
	}

	return b
}

// BoardWithStandardLayout returns an empty Board with a standardised layout.
func BoardWithStandardLayout() Board {
	__, st, dl, dw, tl, tw := positiontype.All()

	return BoardWithLayout(positiontype.Layout{
		{tw, __, __, dl, __, __, __, tw, __, __, __, dl, __, __, tw},
		{__, dw, __, __, __, tl, __, __, __, tl, __, __, __, dw, __},
		{__, __, dw, __, __, __, dl, __, dl, __, __, __, dw, __, __},
		{dl, __, __, dw, __, __, __, dl, __, __, __, dw, __, __, dl},
		{__, __, __, __, dw, __, __, __, __, __, dw, __, __, __, __},
		{__, tl, __, __, __, tl, __, __, __, tl, __, __, __, tl, __},
		{__, __, dl, __, __, __, dl, __, dl, __, __, __, dl, __, __},
		{tw, __, __, dl, __, __, __, st, __, __, __, dl, __, __, tw},
		{__, __, dl, __, __, __, dl, __, dl, __, __, __, dl, __, __},
		{__, tl, __, __, __, tl, __, __, __, tl, __, __, __, tl, __},
		{__, __, __, __, dw, __, __, __, __, __, dw, __, __, __, __},
		{dl, __, __, dw, __, __, __, dl, __, __, __, dw, __, __, dl},
		{__, __, dw, __, __, __, dl, __, dl, __, __, __, dw, __, __},
		{__, dw, __, __, __, tl, __, __, __, tl, __, __, __, dw, __},
		{tw, __, __, dl, __, __, __, tw, __, __, __, dl, __, __, tw},
	})
}

// Neighbours returns the cardinal neighbouring positions to the specified
// coordinate. If a neighbour would be out of bounds, nil is returned in its
// place. Neighbours are always returned in North, South, East, West order.
func (b *Board) Neighbours(c Coord) [4]*BoardPosition {
	return [4]*BoardPosition{
		b.Position(c.North()),
		b.Position(c.South()),
		b.Position(c.East()),
		b.Position(c.West()),
	}
}

// Position returns the board position related to the specified coordinate.
// If the requested position is out of bounds, nil is returned.
func (b *Board) Position(c Coord) *BoardPosition {
	if c.Row < 0 || c.Row >= b.Rows || c.Column < 0 || c.Column >= b.Columns {
		return nil
	}
	return &b.Positions[c.Row*b.Columns+c.Column]
}

func (b *Board) neighbourHasTile(c Coord) bool {
	neighbours := b.Neighbours(c)
	for _, n := range neighbours {
		if n != nil && n.Tile != nil {
			return true
		}
	}
	return false
}

func (b *Board) placeTiles(placements TilePlacements) {
	for _, p := range placements {
		tile := p.Tile
		b.Position(p.Coord).Tile = &tile
	}
}
