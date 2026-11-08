package movement

import (
	"chessgo/board"
	"chessgo/movement/movement_utility"
)


func GetBishopMoves(cb *board.ChessBoard) []movement_utility.Move {
	rays := GetBishopAttackRays()
	var moves []movement_utility.Move
	bishops := cb.Pieces[cb.SideToMove][board.Bishop]
	friends := cb.Colours[cb.SideToMove]
	occupied := cb.Occupied

	for bishops != 0 {
		source := board.GetNextPieceSquare(&bishops)

		targets := movement_utility.RayMoves(rays.NE[source], rays.NE, occupied, true) |
			movement_utility.RayMoves(rays.NW[source], rays.NW, occupied, true) |
			movement_utility.RayMoves(rays.SE[source], rays.SE, occupied, false) |
			movement_utility.RayMoves(rays.SW[source], rays.SW, occupied, false)

		targets &^= friends

		for targets != 0 {
			target := board.GetNextPieceSquare(&targets)
			moves = append(moves, movement_utility.Move{
				Source:        source,
				Target:        target,
				SourcePiece:   board.PieceToColouredPiece(board.Bishop, cb.SideToMove),
				CapturedPiece: cb.Mailbox[target],
			})
		}
	}

	return moves
}


func GetBishopAttackTable() []board.Bitboard {
	attacks := make([]board.Bitboard, 64)
	rays := GetBishopAttackRays()
	for i := range attacks {
		attacks[i] = rays.NE[i] | rays.NW[i] | rays.SE[i] | rays.SW[i]
	}
	return attacks
}

func GetBishopAttackRays() board.BishopRays {
	// Individual rays instead of combined so can use lsb/msb per dir
	// O(n^2), could optimize
	rays := board.BishopRays{
		NE: make([]board.Bitboard, 64),
		NW: make([]board.Bitboard, 64),
		SE: make([]board.Bitboard, 64),
		SW: make([]board.Bitboard, 64),
	}

	for i := range 64 {
		currentBit := board.Bitboard(1) << i
		ne := currentBit
		nw := currentBit
		se := currentBit
		sw := currentBit
		for ne != 0 || nw != 0 || se != 0 || sw != 0 {
			if ne != 0 {
				ne = (ne << board.NE) & board.NotFileAMask
				rays.NE[i] |= ne
			}
			if nw != 0 {
				nw = (nw << board.NW) & board.NotFileHMask
				rays.NW[i] |= nw
			}
			if se != 0 {
				se = (se >> -board.SE) & board.NotFileAMask
				rays.SE[i] |= se
			}
			if sw != 0 {
				sw = (sw >> -board.SW) & board.NotFileHMask
				rays.SW[i] |= sw
			}
		}
	}
	return rays
}
