package movement

import (
	"chessgo/board"
	"math/bits"
	"testing"
)

type moveTest struct {
	name            string
	friendlies      board.Bitboard
	enemies         board.Bitboard
	expectedTargets board.Bitboard
}

func buildChessBoard(start board.Square, context moveTest, sideToMove board.Side) *board.ChessBoard {
	enemySide := sideToMove ^ 1
	cb := &board.ChessBoard{SideToMove: sideToMove}
	cb.Pieces[sideToMove][board.Bishop] = board.Bitboard(0).Set(start)
	cb.Pieces[sideToMove][board.Pawn] = context.friendlies
	cb.Pieces[enemySide][board.Pawn] = context.enemies
	cb.Colours[sideToMove] = cb.Pieces[sideToMove][board.Bishop] | cb.Pieces[sideToMove][board.Pawn]
	cb.Colours[enemySide] = cb.Pieces[enemySide][board.Pawn]
	cb.Occupied = cb.Colours[sideToMove] | cb.Colours[enemySide]
	return cb
}

func TestBishopMovesBlockers(t *testing.T) {
	rays := GetBishopAttackRays()
	start := board.Square(board.C4)

	fullRay := rays.NE[start] | rays.NW[start] | rays.SE[start] | rays.SW[start]

	tests := []moveTest{
		{
			name:            "NE - Friend at D5 blocks entire ray",
			friendlies:      board.Bitboard(0).Set(board.D5),
			enemies:         board.Bitboard(0),
			expectedTargets: fullRay &^ rays.NE[start],
		},
		{
			name:            "NE - Enemy at D5 allows capture only",
			friendlies:      board.Bitboard(0),
			enemies:         board.Bitboard(0).Set(board.D5),
			expectedTargets: fullRay &^ rays.NE[start] | board.Bitboard(0).Set(board.D5),
		},
		{
			name:            "NE - Friend at E6 allows D5 only",
			friendlies:      board.Bitboard(0).Set(board.E6),
			enemies:         board.Bitboard(0),
			expectedTargets: fullRay &^ rays.NE[start] | board.Bitboard(0).Set(board.D5),
		},
		{
			name:            "NE - Enemy at E6 allows D5 and capture E6",
			friendlies:      board.Bitboard(0),
			enemies:         board.Bitboard(0).Set(board.E6),
			expectedTargets: fullRay &^ rays.NE[start] | board.Bitboard(0).Set(board.D5).Set(board.E6),
		},
		{
			name:            "SW - Friend at B3 blocks entire ray",
			friendlies:      board.Bitboard(0).Set(board.B3),
			enemies:         board.Bitboard(0),
			expectedTargets: fullRay &^ rays.SW[start],
		},
		{
			name:            "SW - Enemy at A2 allows B3 and capture A2",
			friendlies:      board.Bitboard(0),
			enemies:         board.Bitboard(0).Set(board.A2),
			expectedTargets: fullRay,
		},
		{
			name:            "NW - Friend at B5 blocks entire ray",
			friendlies:      board.Bitboard(0).Set(board.B5),
			enemies:         board.Bitboard(0),
			expectedTargets: fullRay &^ rays.NW[start],
		},
		{
			name: "NE - Multiple blockers",
			friendlies: board.Bitboard(0).Set(board.F7),
			enemies: board.Bitboard(0).Set(board.E6),
			expectedTargets: fullRay &^ rays.NE[start] | board.Bitboard(0).Set(board.D5).Set(board.E6),
		},
		{
			name: "SW - Multiple blockers",
			friendlies: board.Bitboard(0).Set(board.B3),
			enemies: board.Bitboard(0).Set(board.A2),
			expectedTargets: fullRay &^ rays.SW[start],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cb := buildChessBoard(start, tt, board.White)
			moves := GetBishopMoves(cb)

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

func TestBishopMovesNoBlockers(t *testing.T) {
	rays := GetBishopAttackRays()
	start := board.Square(board.C4)
	fullRay := rays.NE[start] | rays.NW[start] | rays.SE[start] | rays.SW[start]

	tests := []moveTest{
		{
			name:            "Empty Board",
			friendlies:      board.Bitboard(0),
			enemies:         board.Bitboard(0),
			expectedTargets: fullRay,
		},
		{
			name:            "Unreachable - All Friends",
			friendlies:      ^fullRay &^ board.Bitboard(0).Set(start),
			enemies:         board.Bitboard(0),
			expectedTargets: fullRay,
		},
		{
			name:            "Unreachable - All Enemies",
			friendlies:      board.Bitboard(0),
			enemies:         ^fullRay &^ board.Bitboard(0).Set(start),
			expectedTargets: fullRay,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cb := buildChessBoard(start, tt, board.White)
			moves := GetBishopMoves(cb)

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

// Ray tests
func TestBishopAttackCounts(t *testing.T) {
	rays := GetBishopAttackRays()
	tests := []struct {
		name     string
		square   board.Square
		expected int
	}{
		{"A1", board.A1, 7},
		{"A2", board.A2, 7},
		{"A4", board.A4, 7},
		{"B2", board.B2, 9},
		{"B3", board.B3, 9},
		{"B4", board.B4, 9},
		{"C3", board.C3, 11},
		{"C4", board.C4, 11},
		{"D4", board.D4, 13},
	}

	for _, tt := range tests {
		combined := rays.NE[tt.square] |
			rays.NW[tt.square] |
			rays.SE[tt.square] |
			rays.SW[tt.square]

		got := bits.OnesCount64(uint64(combined))
		if got != tt.expected {
			t.Errorf("%s: expected %d attacks, got %d", tt.name, tt.expected, got)
		}
	}
}

func TestBishopWraparounds(t *testing.T) {
	rays := GetBishopAttackRays()

	tests := []struct {
		name      string
		square    board.Square
		forbidden []board.Square
	}{
		{
			name:   "H-File NE/SE Wrap",
			square: board.H4,
			forbidden: []board.Square{
				board.A5,
				board.A3,
			},
		},
		{
			name:   "A-File NW/SW Wrap",
			square: board.A4,
			forbidden: []board.Square{
				board.H5,
				board.H3,
			},
		},
		{
			name:   "H1 Corner Wrap",
			square: board.H1,
			forbidden: []board.Square{
				board.A2,
			},
		},
		{
			name:   "A8 Corner Wrap",
			square: board.A8,
			forbidden: []board.Square{
				board.H7,
			},
		},
	}

	for _, tt := range tests {
		combined := rays.NE[tt.square] |
			rays.NW[tt.square] |
			rays.SE[tt.square] |
			rays.SW[tt.square]

		for _, f := range tt.forbidden {
			if (combined & (board.Bitboard(1) << f)) != 0 {
				t.Errorf("%s: square %d erroneously included (wraparound)", tt.name, f)
			}
		}
	}
}
