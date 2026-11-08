package movement

import (
	"noobfish/board"
	"math/bits"
)

func GetKingMoves(cb *board.ChessBoard) []Move {
	attackTable := GetKingAttackTable()
	var moves []Move

	source := board.Square(bits.TrailingZeros64(uint64(cb.Pieces[cb.SideToMove][board.King])))
	targets := attackTable[source] &^ cb.Colours[cb.SideToMove]
	colouredKing := board.PieceToColouredPiece(board.King, cb.SideToMove)

	for targets != 0 {
		target := board.GetNextPieceSquare(&targets)
		targetPiece := cb.Mailbox[target]
		moves = append(moves, Move{
			Source:        source,
			Target:        target,
			SourcePiece:   colouredKing,
			CapturedPiece: targetPiece,
			Flags:         MoveFlags{},
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

func GetKingAttackTable() []board.Bitboard {
	kingAttacks := make([]board.Bitboard, 64)
	for i := range kingAttacks {
		bit := board.Bitboard(1) << i

		kingAttacks[i] |= (bit & board.NotRank8Mask) << (board.N)
		kingAttacks[i] |= (bit & board.NotFileHMask & board.NotRank8Mask) << (board.NE)
		kingAttacks[i] |= (bit & board.NotFileHMask) << (board.E)
		kingAttacks[i] |= (bit & board.NotFileHMask & board.NotRank1Mask) >> (-board.SE)
		kingAttacks[i] |= (bit & board.NotRank1Mask) >> (-board.S)
		kingAttacks[i] |= (bit & board.NotFileAMask & board.NotRank1Mask) >> (-board.SW)
		kingAttacks[i] |= (bit & board.NotFileAMask) >> (-board.W)
		kingAttacks[i] |= (bit & board.NotFileAMask & board.NotRank8Mask) << (board.NW)
	}

	return kingAttacks
}
