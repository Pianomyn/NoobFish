package movement

import (
	"math/bits"
	"noobfish/attacks"
	"noobfish/board"
)

func GetKingMoves(cb *board.ChessBoard) []board.Move {
	attackTable := attacks.KingAttackTable
	moves := make([]board.Move, 0, 64)

	source := board.Square(bits.TrailingZeros64(uint64(cb.Pieces[cb.SideToMove][board.King])))
	targets := attackTable[source] &^ cb.Colours[cb.SideToMove]
	colouredKing := board.PieceToColouredPiece(board.King, cb.SideToMove)

	for targets != 0 {
		target := board.GetNextPieceSquare(&targets)
		targetPiece := cb.Mailbox[target]
		moves = append(moves, board.Move{
			Source:        source,
			Target:        target,
			SourcePiece:   colouredKing,
			CapturedPiece: targetPiece,
			Flags:         board.MoveFlags{},
		})
	}

	// Castling
	if cb.SideToMove == board.White {
		if cb.CastlingRights.WhiteKingside {

		}
		if cb.CastlingRights.WhiteQueenside {

		}
	} else {
		if cb.CastlingRights.BlackKingside {

		}
		if cb.CastlingRights.BlackQueenside {

		}
	}

	return moves
}
