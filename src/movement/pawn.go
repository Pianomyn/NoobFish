package movement

import (
	"chessgo/board"
)

func GetPawnPushTable(side board.Side) []board.Bitboard {
	pushes := make([]board.Bitboard, 64)
	for i := range pushes {
		currentBit := board.Bitboard(1) << i
		if side == board.White {
			if currentBit&(board.Rank1Mask|board.Rank8Mask) != 0 {
				continue
			} else if currentBit&board.Rank2Mask != 0 {
				pushes[i] |= currentBit<<board.N | currentBit<<(board.N*2)
			} else {
				pushes[i] |= currentBit << board.N
			}
		} else {
			if currentBit&(board.Rank1Mask|board.Rank8Mask) != 0 {
				continue
			} else if currentBit&board.Rank7Mask != 0 {
				pushes[i] |= currentBit>>-board.S | currentBit>>-(board.S*2)
			} else {
				pushes[i] |= currentBit >> -board.S
			}
		}
	}
	return pushes
}

func GetPawnAttackTable(side board.Side) []board.Bitboard {
	attacks := make([]board.Bitboard, 64)
	for i := range attacks {
		currentBit := board.Bitboard(1) << i
		if side == board.White {
			if currentBit&(board.Rank1Mask|board.Rank8Mask) != 0 {
				continue
			} else if currentBit&board.FileAMask != 0 {
				attacks[i] |= currentBit << board.NE
			} else if currentBit&board.FileHMask != 0 {
				attacks[i] |= currentBit << board.NW
			} else {
				attacks[i] |= currentBit<<board.NE | currentBit<<board.NW
			}
		} else {
			if currentBit&(board.Rank1Mask|board.Rank8Mask) != 0 {
				continue
			} else if currentBit&board.FileAMask != 0 {
				attacks[i] |= currentBit >> -board.SE
			} else if currentBit&board.FileHMask != 0 {
				attacks[i] |= currentBit >> -board.SW
			} else {
				attacks[i] |= currentBit>>-board.SE | currentBit>>-board.SW
			}
		}
	}
	return attacks
}
