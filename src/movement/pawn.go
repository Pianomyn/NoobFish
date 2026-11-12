package movement

import (
	"noobfish/attacks"
	"noobfish/board"
)

func GetPawnMoves(cb *board.ChessBoard) []board.Move {
	sideToMove := cb.SideToMove
	takes := attacks.PawnAttackTables
	pushes := attacks.PawnPushTables

	pawns := cb.Pieces[sideToMove][board.Pawn]

	for pawns > 0 {
		current := board.GetNextPieceSquare(&pawns)
		currentBit := 1 << current

		targets := pushes[sideToMove][currentBit] | (takes[sideToMove][currentBit] & cb.Colours[board.GetOtherSide(sideToMove)])

		for targets > 0 {
			// TODO
			targets--

		}
	}
	return []board.Move{}

}
