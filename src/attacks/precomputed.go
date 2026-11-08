package attacks

import (
	"noobfish/board"
)

var (
	BishopAttackRays BishopRays
	KingAttackTable []board.Bitboard
	KnightAttackTable []board.Bitboard
	PawnBlackAttackTable []board.Bitboard
	PawnWhiteAttackTable []board.Bitboard
	QueenAttackTable []board.Bitboard
	RookAttackTable []board.Bitboard
)

func init() {
	BishopAttackRays = GetBishopAttackRays()
	KingAttackTable = GetKingAttackTable()
	KnightAttackTable = GetKnightAttackTable()
	PawnBlackAttackTable = GetPawnAttackTable(board.Black)
	PawnWhiteAttackTable = GetPawnAttackTable(board.White)
	QueenAttackTable = GetQueenAttackTable()
	RookAttackTable = GetRookAttackTable()
}
