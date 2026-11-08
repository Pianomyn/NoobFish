package movement

import "noobfish/board"

func GetRookMoves(cb *board.ChessBoard) []board.Bitboard {
	return make([]board.Bitboard, 64)
}
