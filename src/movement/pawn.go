package movement

import (
	"noobfish/attacks"
	"noobfish/board"
)

func GetPawnMoves(cb *board.ChessBoard) []board.Move {
	moves := make([]board.Move, 64)
	sideToMove := cb.SideToMove
	takes := attacks.PawnAttackTables
	pushes := attacks.PawnPushTables

	pawns := cb.Pieces[sideToMove][board.Pawn]

	for pawns > 0 {
		source := board.GetNextPieceSquare(&pawns)

		targets := pushes[sideToMove][source] | (takes[sideToMove][source] & cb.Colours[board.GetOtherSide(sideToMove)])

		for targets > 0 {
			target := board.GetNextPieceSquare(&targets)
			moves = append(moves, board.Move{
				Source:        source,
				Target:        target,
				SourcePiece:   cb.Mailbox[source],
				CapturedPiece: cb.Mailbox[target],
			})

		}
		if cb.EnPassantSquare != board.NoSquare {
			if attacks.PawnAttackTables[cb.SideToMove][source]&
				board.Bitboard(uint64(cb.EnPassantSquare)) != 0 {
				var capturedPieceSquare board.Square

				// Safe, won't be out of bounds in standard chess
				if cb.SideToMove == board.White {
					capturedPieceSquare = board.Square(int8(cb.EnPassantSquare) + board.S)
				} else {
					capturedPieceSquare = board.Square(int8(cb.EnPassantSquare) + board.N)
				}

				moves = append(moves, board.Move{
					Source:        source,
					Target:        cb.EnPassantSquare,
					SourcePiece:   cb.Mailbox[source],
					CapturedPiece: cb.Mailbox[capturedPieceSquare],
				})

			}
		}
	}
	return moves
}
