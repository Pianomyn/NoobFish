package movement

import (
	"noobfish/attacks"
	"noobfish/board"
)

func GetBishopMoves(cb *board.ChessBoard) []board.Move {
	rays := attacks.BishopAttackRays
	var moves []board.Move
	bishops := cb.Pieces[cb.SideToMove][board.Bishop]
	friends := cb.Colours[cb.SideToMove]
	occupied := cb.Occupied

	for bishops != 0 {
		source := board.GetNextPieceSquare(&bishops)

		targets := attacks.GetRayAttacks(rays.NE[source], rays.NE, friends, occupied, true) |
			attacks.GetRayAttacks(rays.NW[source], rays.NW, friends, occupied, true) |
			attacks.GetRayAttacks(rays.SE[source], rays.SE, friends, occupied, false) |
			attacks.GetRayAttacks(rays.SW[source], rays.SW, friends, occupied, false)

		for targets != 0 {
			target := board.GetNextPieceSquare(&targets)
			moves = append(moves, board.Move{
				Source:        source,
				Target:        target,
				SourcePiece:   board.PieceToColouredPiece(board.Bishop, cb.SideToMove),
				CapturedPiece: cb.Mailbox[target],
			})
		}
	}

	return moves
}
