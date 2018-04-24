package textscrubble

import (
	gt "github.com/buger/goterm"
	"github.com/mandykoh/scrubble"
	"github.com/mandykoh/scrubble/board"
	"github.com/mandykoh/scrubble/coord"
	"github.com/mandykoh/scrubble/tile"
)

func DrawBoard(b *board.Board) {
	_, st, dl, dw, tl, tw := board.AllPositionTypes()

	for r := 0; r < b.Rows; r++ {
		offsetY := r*2 + 1

		for c := 0; c < b.Columns; c++ {
			offsetX := c*4 + 1

			gt.MoveCursor(offsetX, offsetY)
			gt.Print(gt.Color(".", gt.GREEN))
			gt.MoveCursorDown(1)
			gt.MoveCursorBackward(1)
			gt.Print(gt.Color("|", gt.GREEN))

			pos := b.Position(coord.Make(r, c))

			bg := gt.BLACK
			if pos.Tile != nil {
				bg = gt.WHITE
			}

			gt.MoveCursor(offsetX+1, offsetY+1)
			switch pos.Type {
			case st:
				gt.Print(gt.Background(gt.Color("★", gt.WHITE), bg))
			case dl:
				gt.Print(gt.Background(gt.Color("dl", gt.BLUE), bg))
			case dw:
				gt.Print(gt.Background(gt.Color("dw", gt.RED), bg))
			case tl:
				gt.Print(gt.Background(gt.Color("tl", gt.GREEN), bg))
			case tw:
				gt.Print(gt.Background(gt.Color("tw", gt.YELLOW), bg))
			default:
				gt.Print(gt.Background(" ", bg))
			}

			if pos.Tile != nil {
				gt.MoveCursor(offsetX+1, offsetY)
				gt.Printf(gt.Bold(gt.Background(gt.Color("%c  ", gt.BLACK), gt.WHITE)), pos.Tile.Letter)
				gt.MoveCursor(offsetX+2, offsetY+1)
				gt.Printf(gt.Background(gt.Color("%2d", gt.BLACK), gt.WHITE), pos.Tile.Points)
			}
		}

		gt.MoveCursor(b.Columns*4+1, offsetY)
		gt.Printf(gt.Color("| %d", gt.GREEN), r)
		gt.MoveCursor(b.Columns*4+1, offsetY+1)
		gt.Print(gt.Color("|", gt.GREEN))
	}

	for i := 0; i < b.Rows; i++ {
		gt.MoveCursor(i*4+2, b.Rows*2+1)
		gt.Printf(gt.Color("%d", gt.GREEN), i)
	}
}

func DrawGame(g *scrubble.Game, players []Player) {
	gt.Clear()
	DrawBoard(&g.Board)
	DrawStats(g, players)

	gt.MoveCursor(0, g.Board.Rows*2+3)

	if g.Phase == scrubble.EndPhase {
		gt.Println("Game over")
	} else {
		gt.Printf("%s’s turn (? for help, Enter to clear output):", players[g.CurrentSeatIndex].Name)
	}

	gt.Flush()
}

func DrawRack(r tile.Rack) {
	gt.Println()

	for _, t := range r {
		letter := t.Letter
		if letter == ' ' {
			letter = '_'
		}

		gt.MoveCursorUp(1)
		gt.MoveCursorForward(2)
		gt.Printf(gt.Bold(gt.Background(gt.Color("%c  ", gt.BLACK), gt.WHITE)), letter)
		gt.MoveCursorDown(1)
		gt.MoveCursorBackward(3)
		gt.Printf(gt.Background(gt.Color("%3d", gt.BLACK), gt.WHITE), t.Points)
	}
}

func DrawStats(g *scrubble.Game, players []Player) {
	gt.MoveCursor(g.Board.Columns*4+7, 1)
	gt.Printf("%d tiles in bag", len(g.Bag))

	for i, s := range g.Seats {
		gt.MoveCursor(g.Board.Columns*4+7, i+3)
		gt.Printf("%s %d", players[i].Name, s.Score)
	}
}
