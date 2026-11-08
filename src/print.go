// https://www.chessprogramming.org/Knight_Pattern
package main

import (
	"noobfish/board"
	"noobfish/attacks"
	"fmt"
)

func PrintMoves(sourceSquare board.Square, bb board.Bitboard) {
	fmt.Println("  +-----------------+")
	for r := 7; r >= 0; r-- {
		fmt.Printf("%d | ", r+1) // Rank number
		for f := 0; f < 8; f++ {
			sq := board.Square(r*8 + f)

			if sq == sourceSquare {
				fmt.Print("O ")
				continue
			}

			if (uint64(bb)>>uint8(sq))&1 == 1 {
				fmt.Print("x ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("  +-----------------+")
	fmt.Println("    a b c d e f g h")
}

func PrintAllMoves(piece board.Piece) {
	var allMoves []board.Bitboard
	switch piece {
	case board.King:
		allMoves = attacks.KingAttackTable
	case board.Knight:
		allMoves = attacks.KnightAttackTable
	case board.Rook:
		allMoves = attacks.RookAttackTable
	case board.Bishop:
		allMoves = attacks.BishopAttackTable
	case board.Queen:
		allMoves = attacks.QueenAttackTable
	default:
		panic("Invalid Piece")
	}
	for i, bb := range allMoves {
		PrintMoves(board.Square(i), bb)
	}
}
