package movement

import (
	"noobfish/board"
	"noobfish/attacks"
)

func GetRookMoves(cb *board.ChessBoard) []board.Move {
	rays := attacks.RookAttackRays
	var moves []board.Move
	rooks := cb.Pieces[cb.SideToMove][board.Rook]
	friends := cb.Colours[cb.SideToMove]
	occupied := cb.Occupied

	for rooks != 0 {
		source := board.GetNextPieceSquare(&rooks)

		targets := attacks.GetRayAttacks(source, rays.North, friends, occupied, true) |
			attacks.GetRayAttacks(source, rays.East, friends, occupied, true) |
			attacks.GetRayAttacks(source, rays.South, friends, occupied, false) |
			attacks.GetRayAttacks(source, rays.West, friends, occupied, false)

		for targets != 0 {
			target := board.GetNextPieceSquare(&targets)
			moves = append(moves, board.Move{
				Source:        source,
				Target:        target,
				SourcePiece:   board.PieceToColouredPiece(board.Rook, cb.SideToMove),
				CapturedPiece: cb.Mailbox[target],
			})
		}
	}

	return moves
}
