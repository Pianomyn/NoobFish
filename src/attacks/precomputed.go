package attacks

import (
	"noobfish/board"
)

var (
	BishopAttackRays BishopRays
	BishopAttackTable []board.Bitboard
	KingAttackTable []board.Bitboard
	KnightAttackTable []board.Bitboard
	PawnBlackAttackTable []board.Bitboard
	PawnWhiteAttackTable []board.Bitboard
	QueenAttackTable []board.Bitboard
	RookAttackTable []board.Bitboard
)

func init() {
	BishopAttackRays = GetBishopAttackRays()
	BishopAttackTable = GetBishopAttackTable()
	KingAttackTable = GetKingAttackTable()
	KnightAttackTable = GetKnightAttackTable()
	PawnBlackAttackTable = GetPawnAttackTable(board.Black)
	PawnWhiteAttackTable = GetPawnAttackTable(board.White)
	QueenAttackTable = GetQueenAttackTable()
	RookAttackTable = GetRookAttackTable()
}
