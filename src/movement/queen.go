package movement

import "chessgo/board"

func GetQueenAttackTable() []board.Bitboard {
	queenAttacks := make([]board.Bitboard, 64)
	rookAttacks := GetRookAttackTable()
	bishopAttacks := GetBishopAttackRays()
	for i := range queenAttacks {
		bishopAttacks := bishopAttacks.NE[i] |
			bishopAttacks.NW[i] |
			bishopAttacks.SW[i] |
			bishopAttacks.SE[i]

		queenAttacks[i] = bishopAttacks | rookAttacks[i]
	}
	return queenAttacks
}
