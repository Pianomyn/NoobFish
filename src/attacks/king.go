package attacks

import (
	"noobfish/board"
)

func GetKingAttackTable() []board.Bitboard {
	kingAttacks := make([]board.Bitboard, 64)
	for i := range kingAttacks {
		bit := board.Bitboard(1) << i

		kingAttacks[i] |= (bit & board.NotRank8Mask) << (board.N)
		kingAttacks[i] |= (bit & board.NotFileHMask & board.NotRank8Mask) << (board.NE)
		kingAttacks[i] |= (bit & board.NotFileHMask) << (board.E)
		kingAttacks[i] |= (bit & board.NotFileHMask & board.NotRank1Mask) >> (-board.SE)
		kingAttacks[i] |= (bit & board.NotRank1Mask) >> (-board.S)
		kingAttacks[i] |= (bit & board.NotFileAMask & board.NotRank1Mask) >> (-board.SW)
		kingAttacks[i] |= (bit & board.NotFileAMask) >> (-board.W)
		kingAttacks[i] |= (bit & board.NotFileAMask & board.NotRank8Mask) << (board.NW)
	}

	return kingAttacks
}
