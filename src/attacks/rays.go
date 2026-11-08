package attacks

import (
	"chessgo/board"
	"math/bits"
)

func RayMoves(ray board.Bitboard, allRays []board.Bitboard, occupied board.Bitboard, positive bool) board.Bitboard {
	/*
		For diagonal rays:
			positive = top left and top right
		For orthogonal rays:
			positive = top and right
	*/
	blockers := ray & occupied
	if blockers == 0 {
		return ray
	}

	var first int
	if positive {
		first = bits.TrailingZeros64(uint64(blockers))
	} else {
		// Counting from MSB downwards, flip/convert to get Little-Endian bit
		first = 63 - bits.LeadingZeros64(uint64(blockers))
	}

	return ray ^ allRays[first]
}
