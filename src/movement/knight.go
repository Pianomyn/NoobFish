package movement

import (
	"noobfish/attacks"
	"noobfish/board"
)

func GetKnightMoves(cb *board.ChessBoard) []board.Move {
	moves := make([]board.Move, 0, 64)
	attackTable := attacks.KnightAttackTable

	knights := cb.Pieces[cb.SideToMove][board.Knight]

	for knights != 0 {
		source := board.GetNextPieceSquare(&knights)

		moveBitboard := attackTable[source] &^ cb.Colours[cb.SideToMove]

		for moveBitboard != 0 {
			target := board.GetNextPieceSquare(&moveBitboard)
			moves = append(
				moves,
				board.Move{
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
