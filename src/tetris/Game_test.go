package tetris

import "testing"

func Test_emptyBoard(t *testing.T) {
	game := Game{}
	if game.board != [20]uint16{} {
		t.Error("Game constructor produces non empty board.")
	}
}

func Test_dropSinglePiece(t *testing.T) {
	game := Game{}
	expectedBoard := [20]uint16{56, 32}

	game = DropPiece(game, PieceMap['L'], 0, 0)
	if game.board != expectedBoard {
		t.Error("Single piece drop failed with board")
		t.Log(game.board)
	}
}

func Test_dropSuccess(t *testing.T) {
	game := Game{}
	expectedBoard := [20]uint16{56, 56, 8, 8}

	game = DropPiece(game, PieceMap['L'], 0, 0)
	game = DropPiece(game, PieceMap['L'], -1, 1)

	if game.board != expectedBoard {
		t.Error("Two piece drop produced an incorrect board.")
		t.Log(game.board)
	}
}

func Test_clearLine(t *testing.T) {
	game := Game{}
	expectedBoard := [20]uint16{959, 523, 8}

	game = DropPiece(game, PieceMap['L'], -3, 0)
	game = DropPiece(game, PieceMap['I'], -2, 1)
	game = DropPiece(game, PieceMap['T'], 3, 0)
	game = DropPiece(game, PieceMap['O'], 0, 0)
	game = DropPiece(game, PieceMap['O'], -4, 0)
	game = DropPiece(game, PieceMap['T'], 5, 3)

	if game.board != expectedBoard || game.LinesCleared != 1 {
		t.Error("Single line clearing produced an incorrect board.")
		t.Log(game.board)
	}

}

func Test_clearMultiLine(t *testing.T) {
	game := Game{}
	game = DropPiece(game, PieceMap['L'], -3, 0)
	game = DropPiece(game, PieceMap['I'], -2, 1)
	game = DropPiece(game, PieceMap['T'], 3, 0)
	game = DropPiece(game, PieceMap['O'], 0, 0)
	game = DropPiece(game, PieceMap['O'], -4, 0)
	game = DropPiece(game, PieceMap['Z'], 2, 1)
	game = DropPiece(game, PieceMap['I'], 3, 1)
	game = DropPiece(game, PieceMap['O'], 0, 0)
	game = DropPiece(game, PieceMap['J'], -3, 2)
	game = DropPiece(game, PieceMap['Z'], 2, 1)

	expectedBoard := [20]uint16{511, 511, 511, 511, 448, 128}
	if game.board != expectedBoard {
		t.Error("Setting up multi-line clear produced an incorrect board.")
	}

	game = DropPiece(game, PieceMap['I'], 4, 1)

	expectedBoard = [20]uint16{448, 128}
	if game.board != expectedBoard || game.LinesCleared != 4 {
		t.Error("Multi-line clear produced an incorrect board.")
	}

	t.Log(game.board)
}
