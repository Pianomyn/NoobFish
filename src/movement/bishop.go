package movement

import (
	"noobfish/attacks"
	"noobfish/board"
)

func GetBishopMoves(cb *board.ChessBoard) []board.Move {
	rays := attacks.BishopAttackRays
	moves := make([]board.Move, 0, 64)
	bishops := cb.Pieces[cb.SideToMove][board.Bishop]
	friends := cb.Colours[cb.SideToMove]
	occupied := cb.Occupied

	for bishops != 0 {
		source := board.GetNextPieceSquare(&bishops)

		targets := attacks.GetRayAttacks(source, rays.NE, occupied, true) |
			attacks.GetRayAttacks(source, rays.NW, occupied, true) |
			attacks.GetRayAttacks(source, rays.SE, occupied, false) |
			attacks.GetRayAttacks(source, rays.SW, occupied, false)

		targets &^= friends

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
