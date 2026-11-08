package movement

import (
	"noobfish/attacks"
	"noobfish/board"
)

func GetBishopMoves(cb *board.ChessBoard) []Move {
	rays := attacks.GetBishopAttackRays()
	var moves []Move
	bishops := cb.Pieces[cb.SideToMove][board.Bishop]
	friends := cb.Colours[cb.SideToMove]
	occupied := cb.Occupied

	for bishops != 0 {
		source := board.GetNextPieceSquare(&bishops)

		targets := attacks.RayMoves(rays.NE[source], rays.NE, occupied, true) |
			attacks.RayMoves(rays.NW[source], rays.NW, occupied, true) |
			attacks.RayMoves(rays.SE[source], rays.SE, occupied, false) |
			attacks.RayMoves(rays.SW[source], rays.SW, occupied, false)

		targets &^= friends

		for targets != 0 {
			target := board.GetNextPieceSquare(&targets)
			moves = append(moves, Move{
				Source:        source,
				Target:        target,
				SourcePiece:   board.PieceToColouredPiece(board.Bishop, cb.SideToMove),
				CapturedPiece: cb.Mailbox[target],
			})
		}
	}

	return moves
}
