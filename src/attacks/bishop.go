package attacks

import (
	"noobfish/board"
)

func GetBishopAttackRays() BishopRays {
	// Individual rays instead of combined so can use lsb/msb per dir
	// O(n^2), could optimize
	rays := BishopRays{
		NE: make([]board.Bitboard, 64),
		NW: make([]board.Bitboard, 64),
		SE: make([]board.Bitboard, 64),
		SW: make([]board.Bitboard, 64),
	}

	for i := range 64 {
		row := i / 8
		col := i % 8

		for offset := 1; offset < 8; offset++ {
			if r, c := row+offset, col+offset; r < 8 && c < 8 {
				rays.NE[i] |= board.Bitboard(1<<(r*8 + c))
			}
			if r, c := row+offset, col-offset; r < 8 && c > -1 {
				rays.NW[i] |= board.Bitboard(1<<(r*8 + c))
			}
			if r, c := row-offset, col+offset; r > -1 && c < 8 {
				rays.SE[i] |= board.Bitboard(1<<(r*8 + c))
			}
			if r, c := row-offset, col-offset; r > -1 && c > -1 {
				rays.SW[i] |= board.Bitboard(1<<(r*8 + c))
			}

		}
	}
	return rays
}
