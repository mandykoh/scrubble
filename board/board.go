package board

import (
	"github.com/mandykoh/scrubble/coord"
)

// Board represents a game board, which is a grid of positions on which tiles
// can be placed. The zero-value of a Board is a zero-sized board.
type Board struct {
	Rows      int
	Columns   int
	Positions []Position
}

// WithLayout creates a board with no tiles, with the specified layout.
func WithLayout(layout Layout) Board {
	normal, _, _, _, _, _ := AllPositionTypes()

	rows := len(layout)
	columns := layout.WidestRow()

	b := Board{
		Rows:      rows,
		Columns:   columns,
		Positions: make([]Position, rows*columns),
	}

	for row, lRow := range layout {

		// Set position types according to the specified layout
		for col, posType := range lRow {
			b.Position(coord.Make(row, col)).Type = posType
		}

		// Fill in any unspecified remainder of the row with Normal positions
		for col := len(lRow); col < columns; col++ {
			b.Position(coord.Make(row, col)).Type = normal
		}
	}

	return b
}

// WithStandardLayout returns an empty Board with a standardised layout.
func WithStandardLayout() Board {
	__, st, dl, dw, tl, tw := AllPositionTypes()

	return WithLayout(Layout{
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

// NeighbourHasTile returns true if any neighbouring position of the position at
// the specified coordinate has a tile on it.
func (b *Board) NeighbourHasTile(c coord.Coord) bool {
	neighbours := b.Neighbours(c)
	for _, n := range neighbours {
		if n != nil && n.Tile != nil {
			return true
		}
	}
	return false
}

// Neighbours returns the cardinal neighbouring positions to the specified
// coordinate. If a neighbour would be out of bounds, nil is returned in its
// place. Neighbours are always returned in North, South, East, West order.
func (b *Board) Neighbours(c coord.Coord) [4]*Position {
	return [4]*Position{
		b.Position(c.North()),
		b.Position(c.South()),
		b.Position(c.East()),
		b.Position(c.West()),
	}
}

// Position returns the board position related to the specified coordinate.
// If the requested position is out of bounds, nil is returned.
func (b *Board) Position(c coord.Coord) *Position {
	if c.Row < 0 || c.Row >= b.Rows || c.Column < 0 || c.Column >= b.Columns {
		return nil
	}
	return &b.Positions[c.Row*b.Columns+c.Column]
}
