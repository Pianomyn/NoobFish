package attacks

import (
	"noobfish/board"
	"math/bits"
)

func GetRayAttacks(ray board.Bitboard, allRays []board.Bitboard, friends board.Bitboard, occupied board.Bitboard, positive bool) board.Bitboard {
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

	return ray ^ allRays[first] &^ friends
}

type BishopRays struct {
	NE []board.Bitboard
	NW []board.Bitboard
	SE []board.Bitboard
	SW []board.Bitboard
}

type RookRays struct {
	North []board.Bitboard
	East []board.Bitboard
	South []board.Bitboard
	West []board.Bitboard
}

func (br BishopRays) ToAttackTable() []board.Bitboard {
	attackTable := make([]board.Bitboard, 64)
	for i := range 64 {
		attackTable[i] = br.NE[i] | br.NW[i] | br.SE[i] | br.SW[i]
	}
	return attackTable
}

func (rr RookRays) ToAttackTable() []board.Bitboard {
	attackTable := make([]board.Bitboard, 64)
	for i := range 64 {
		attackTable[i] = rr.North[i] | rr.East[i] | rr.South[i] | rr.West[i]
	}
	return attackTable
}
