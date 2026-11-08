package movement

import (
	"chessgo/board"
	"chessgo/movement/movement_utility"
)

func GetKnightMoves(cb *board.ChessBoard) []movement_utility.Move {
	var moves []movement_utility.Move
	attackTable := GetKnightAttackTable()

	knights := cb.Pieces[cb.SideToMove][board.Knight]

	for knights != 0 {
		source := board.GetNextPieceSquare(&knights)

		moveBitboard := attackTable[source] &^ cb.Colours[cb.SideToMove]
		
		for moveBitboard != 0 {
			target := board.GetNextPieceSquare(&moveBitboard)
			moves = append(
				moves,
				movement_utility.Move{
					Source:        source,
					Target:        target,
					SourcePiece:   board.PieceToColouredPiece(board.Knight, cb.SideToMove),
					CapturedPiece: cb.Mailbox[target],
				},
			)
		}
	}
	return moves
}

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
