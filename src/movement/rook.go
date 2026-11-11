package movement

import (
	"noobfish/attacks"
	"noobfish/board"
)

func GetRookMoves(cb *board.ChessBoard) []board.Move {
	rays := attacks.RookAttackRays
	moves := make([]board.Move, 0, 64)
	rooks := cb.Pieces[cb.SideToMove][board.Rook]
	friends := cb.Colours[cb.SideToMove]
	occupied := cb.Occupied

	for rooks != 0 {
		source := board.GetNextPieceSquare(&rooks)

		targets := attacks.GetRayAttacks(source, rays.North, occupied, true) |
			attacks.GetRayAttacks(source, rays.East, occupied, true) |
			attacks.GetRayAttacks(source, rays.South, occupied, false) |
			attacks.GetRayAttacks(source, rays.West, occupied, false)

		targets &^= friends

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
