package attacks

import (
	"noobfish/board"
)

func GetKnightAttackTable() []board.Bitboard {
	knightAttacks := make([]board.Bitboard, 64)
	for i := range knightAttacks {
		bit := board.Bitboard(1) << i

		knightAttacks[i] |= (bit & board.NotFileABMask) << board.NoWeWe
		knightAttacks[i] |= (bit & board.NotFileAMask) << board.NoNoWe
		knightAttacks[i] |= (bit & board.NotFileGHMask) << board.NoEaEa
		knightAttacks[i] |= (bit & board.NotFileHMask) << board.NoNoEa

		// Cant negative bitshift
		knightAttacks[i] |= (bit & board.NotFileABMask) >> (-board.SoWeWe)
		knightAttacks[i] |= (bit & board.NotFileAMask) >> (-board.SoSoWe)
		knightAttacks[i] |= (bit & board.NotFileGHMask) >> (-board.SoEaEa)
		knightAttacks[i] |= (bit & board.NotFileHMask) >> (-board.SoSoEa)
	}

	return knightAttacks
}
