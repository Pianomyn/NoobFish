package movement

import (
	"noobfish/attacks"
	"noobfish/board"
)

func GetQueenMoves(cb *board.ChessBoard) []board.Move {
	bishopRays := attacks.BishopAttackRays
	rookRays := attacks.RookAttackRays
	var moves []board.Move
	queens := cb.Pieces[cb.SideToMove][board.Queen]
	friends := cb.Colours[cb.SideToMove]
	occupied := cb.Occupied

	for queens != 0 {
		source := board.GetNextPieceSquare(&queens)

		targets := attacks.GetRayAttacks(source, bishopRays.NE, occupied, true) |
			attacks.GetRayAttacks(source, bishopRays.NW, occupied, true) |
			attacks.GetRayAttacks(source, bishopRays.SE, occupied, false) |
			attacks.GetRayAttacks(source, bishopRays.SW, occupied, false) |
			attacks.GetRayAttacks(source, rookRays.North, occupied, true) |
			attacks.GetRayAttacks(source, rookRays.East, occupied, true) |
			attacks.GetRayAttacks(source, rookRays.South, occupied, false) |
			attacks.GetRayAttacks(source, rookRays.West, occupied, false)

		targets &^= friends

		for targets != 0 {
			target := board.GetNextPieceSquare(&targets)
			moves = append(moves, board.Move{
				Source:        source,
				Target:        target,
				SourcePiece:   board.PieceToColouredPiece(board.Queen, cb.SideToMove),
				CapturedPiece: cb.Mailbox[target],
			})
		}
	}

	return moves

}
