package attacks

import (
	"noobfish/board"
)

func GetQueenAttackTable() []board.Bitboard {
	queenAttacks := make([]board.Bitboard, 64)
	rookAttackRays := GetRookAttackRays()
	bishopAttackRays := GetBishopAttackRays()
	for i := range queenAttacks {
		bishopAttacks := bishopAttackRays.NE[i] |
			bishopAttackRays.NW[i] |
			bishopAttackRays.SW[i] |
			bishopAttackRays.SE[i]
		rookAttacks := rookAttackRays.North[i] |
			rookAttackRays.East[i] |
			rookAttackRays.South[i] |
			rookAttackRays.West[i]

		queenAttacks[i] = bishopAttacks | rookAttacks
	}
	return queenAttacks
}
