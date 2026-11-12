package attacks

import (
	"noobfish/board"
)

var (
	BishopAttackRays  BishopRays
	KingAttackTable   []board.Bitboard
	KnightAttackTable []board.Bitboard
	PawnAttackTables  [2][]board.Bitboard
	PawnPushTables    [2][]board.Bitboard
	QueenAttackTable  []board.Bitboard
	RookAttackRays    RookRays
)

func init() {
	BishopAttackRays = GetBishopAttackRays()
	KingAttackTable = GetKingAttackTable()
	KnightAttackTable = GetKnightAttackTable()
	PawnAttackTables = GetPawnAttackTables()
	PawnPushTables = GetPawnPushTables()
	QueenAttackTable = GetQueenAttackTable()
	RookAttackRays = GetRookAttackRays()
}
