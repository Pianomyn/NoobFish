package movement

import "chessgo/board"


func GetRookMoves(cb *board.ChessBoard) []board.Bitboard {
	return make([]board.Bitboard, 64)
}

func GetRookAttackTable() []board.Bitboard {
	orthogonal := make([]board.Bitboard, 64)

	for i := 0; i < 64; i++ {
		row := i / 8
		col := i % 8

		rankMask := board.Rank1Mask << (8 * row)
		colMask := board.FileAMask << col

		orthogonal[i] |= (rankMask | colMask) &^ (1 << i)
	}

	return orthogonal
}
