package movement

import (
	"noobfish/attacks"
	"noobfish/board"
	"testing"
)

type rookMoveTest struct {
	name            string
	friendlies      board.Bitboard
	enemies         board.Bitboard
	expectedTargets board.Bitboard
}

func buildChessBoardForRook(start board.Square, context rookMoveTest, sideToMove board.Side) *board.ChessBoard {
	enemySide := sideToMove ^ 1
	cb := &board.ChessBoard{SideToMove: sideToMove}
	cb.Pieces[sideToMove][board.Rook] = board.Bitboard(0).Set(start)
	cb.Pieces[sideToMove][board.Pawn] = context.friendlies
	cb.Pieces[enemySide][board.Pawn] = context.enemies
	cb.Colours[sideToMove] = cb.Pieces[sideToMove][board.Rook] | cb.Pieces[sideToMove][board.Pawn]
	cb.Colours[enemySide] = cb.Pieces[enemySide][board.Pawn]
	cb.Occupied = cb.Colours[sideToMove] | cb.Colours[enemySide]
	return cb
}

func TestRookMovesBlockers(t *testing.T) {
	rays := attacks.RookAttackRays
	start := board.Square(board.C4)

	fullRay := rays.ToAttackTable()[start]
	println(fullRay)

	tests := []rookMoveTest{
		{
			name:            "N - Friend at C5 blocks entire ray",
			friendlies:      board.Bitboard(0).Set(board.C5),
			enemies:         board.Bitboard(0),
			expectedTargets: fullRay ^ rays.North[start],
		},
		{
			name:            "N - Enemy at C5 allows capture only",
			friendlies:      board.Bitboard(0),
			enemies:         board.Bitboard(0).Set(board.C5),
			expectedTargets: fullRay^rays.North[start] | board.Bitboard(0).Set(board.C5),
		},
		{
			name:            "N - Friend at C6 allows C5 only",
			friendlies:      board.Bitboard(0).Set(board.C6),
			enemies:         board.Bitboard(0),
			expectedTargets: fullRay&^rays.North[start] | board.Bitboard(0).Set(board.C5),
		},
		{
			name:            "N - Enemy at C6 allows C5 and capture C6",
			friendlies:      board.Bitboard(0),
			enemies:         board.Bitboard(0).Set(board.C6),
			expectedTargets: fullRay&^rays.North[start] | board.Bitboard(0).Set(board.C5).Set(board.C6),
		},
		{
			name:            "E - Friend at D4 blocks entire ray",
			friendlies:      board.Bitboard(0).Set(board.D4),
			enemies:         board.Bitboard(0),
			expectedTargets: fullRay &^ rays.East[start],
		},
		{
			name:            "W - Friend at B4 blocks entire ray",
			friendlies:      board.Bitboard(0).Set(board.B4),
			enemies:         board.Bitboard(0),
			expectedTargets: fullRay &^ rays.West[start],
		},
		{
			name:            "S - Friend at C3 blocks entire ray",
			friendlies:      board.Bitboard(0).Set(board.C3),
			enemies:         board.Bitboard(0),
			expectedTargets: fullRay &^ rays.South[start],
		},
		{
			name:            "N - Multiple blockers",
			friendlies:      board.Bitboard(0).Set(board.C8),
			enemies:         board.Bitboard(0).Set(board.C7),
			expectedTargets: fullRay&^rays.North[start] | board.Bitboard(0).Set(board.C5).Set(board.C6).Set(board.C7),
		},
		{
			name:            "S - Multiple blockers",
			friendlies:      board.Bitboard(0).Set(board.C3),
			enemies:         board.Bitboard(0).Set(board.C2),
			expectedTargets: fullRay &^ rays.South[start],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cb := buildChessBoardForRook(start, tt, board.White)
			moves := GetRookMoves(cb)

			var gotTargets board.Bitboard
			for _, m := range moves {
				gotTargets = gotTargets.Set(m.Target)
			}

			if gotTargets != tt.expectedTargets {
				t.Errorf("got targets %064b, expected %064b", gotTargets, tt.expectedTargets)
			}
		})
	}
}

func TestRookMovesNoBlockers(t *testing.T) {
	rays := attacks.GetRookAttackRays()
	start := board.Square(board.C4)
	fullRay := rays.ToAttackTable()[start]

	tests := []rookMoveTest{
		{
			name:            "Empty Board",
			friendlies:      board.Bitboard(0),
			enemies:         board.Bitboard(0),
			expectedTargets: fullRay,
		},
		{
			name:            "Unreachable - All Friends",
			friendlies:      ^fullRay,
			enemies:         board.Bitboard(0),
			expectedTargets: fullRay,
		},
		{
			name:            "Unreachable - All Enemies",
			friendlies:      board.Bitboard(0),
			enemies:         ^fullRay,
			expectedTargets: fullRay,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cb := buildChessBoardForRook(start, tt, board.White)
			moves := GetRookMoves(cb)

			var gotTargets board.Bitboard
			for _, m := range moves {
				gotTargets = gotTargets.Set(m.Target)
			}

			if gotTargets != tt.expectedTargets {
				t.Errorf("got targets %064b, expected %064b", gotTargets, tt.expectedTargets)
			}
		})
	}
}
