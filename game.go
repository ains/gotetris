package gotetris

import "fmt"

type Game struct {
	board        [20]uint16
	columnHeight [10]int
	LinesCleared int
	CurrentPiece int
}

func (game *Game) isSet(x, y int) bool {
	return game.board[y]&(1<<uint(x)) != 0
}

func (game *Game) setBlock(x, y int) {
	game.board[y] |= (1 << uint(x))
}

func (game *Game) clearBlock(x, y int) {
	game.board[y] &= ^(1 << uint(x))
}

func (game *Game) clearRow(y int) {
	game.board[y] = 0
}

func (game *Game) removeCompleteLines() {
	for _, height := range game.columnHeight {
		if height == 0 {
			return
		}
	}

	linesRemoved := 0
	for y := 0; y < 20; y++ {
		if game.board[y] == 0x3FF {
			game.removeRow(y)

			linesRemoved++
			y--
		}
	}

	game.LinesCleared += linesRemoved
}

func (game *Game) removeRow(row int) {
	game.clearRow(row)

	for y := row; y < 19; y++ {
		game.board[y] = game.board[y+1]
	}

	for i := 0; i < 10; i++ {
		game.columnHeight[i]--
	}
}

func DropPiece(newGame Game, piece *Piece, shift int, rot int) Game {
	var yShift, column, columnHeight int

	pieceCoordinates := piece.Coordinates[rot]
	lowestPoints := piece.LowestPoints[rot]
	blockCount := len(pieceCoordinates)

	for i := 0; i < blockCount; i++ {
		x := pieceCoordinates[i][0]
		column = x + shift

		columnHeight = newGame.columnHeight[column] - lowestPoints[x]
		if columnHeight > yShift {
			yShift = columnHeight
		}
	}

	for i := 0; i < blockCount; i++ {
		x := pieceCoordinates[i][0] + shift
		y := pieceCoordinates[i][1] + yShift

		if y+1 > newGame.columnHeight[x] {
			newGame.columnHeight[x] = y + 1
		}

		newGame.setBlock(x, y)

	}

	newGame.CurrentPiece++
	newGame.removeCompleteLines()

	return newGame
}

func (game *Game) OutputBoard() {
	for y := 19; y >= 0; y-- {
		for x := 0; x < 10; x++ {
			if game.isSet(x, y) {
				fmt.Print("# ")
			} else {
				fmt.Print("_ ")
			}
		}
		fmt.Println("")
	}
}
