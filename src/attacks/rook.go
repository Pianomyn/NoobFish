package attacks

import (
	"noobfish/board"
)

func GetRookAttackRays() RookRays {
	rays := RookRays{
		North: make([]board.Bitboard, 64),
		East:  make([]board.Bitboard, 64),
		South: make([]board.Bitboard, 64),
		West:  make([]board.Bitboard, 64),
	}

	for i := 0; i < 64; i++ {
		row := i / 8
		col := i % 8

		for offset := 1; offset < 8; offset++ {
			if r := row + offset; r < 8 {
				rays.North[i] |= board.Bitboard(1<<(r*8 + col))
			}
			if r := row - offset; r > -1 {
				rays.South[i] |= board.Bitboard(1<<(r*8 + col))
			}
			if c := col + offset; c < 8 {
				rays.East[i] |= board.Bitboard(1<<(row*8 + c))
			}
			if c := col - offset; c > -1 {
				rays.West[i] |= board.Bitboard(1<<(row*8 + c))
			}
		}
	}

	return rays
}
