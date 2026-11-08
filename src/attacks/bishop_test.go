package attacks

import (
	"noobfish/board"
	"math/bits"
	"testing"
)

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

