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

		targets := attacks.GetRayAttacks(source, bishopRays.NE, friends, occupied, true) |
			attacks.GetRayAttacks(source, bishopRays.NW, friends, occupied, true) |
			attacks.GetRayAttacks(source, bishopRays.SE, friends, occupied, false) |
			attacks.GetRayAttacks(source, bishopRays.SW, friends, occupied, false) |
			attacks.GetRayAttacks(source, rookRays.North, friends, occupied, true) |
			attacks.GetRayAttacks(source, rookRays.East, friends, occupied, true) |
			attacks.GetRayAttacks(source, rookRays.South, friends, occupied, false) |
			attacks.GetRayAttacks(source, rookRays.West, friends, occupied, false)

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
