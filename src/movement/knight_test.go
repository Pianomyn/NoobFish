package movement

import (
	"chessgo/board"
	"math/bits"
	"testing"
)

func TestKnightAttackCounts(t *testing.T) {
	table := GetKnightAttackTable()

	tests := []struct {
		name     string
		square   board.Square
		expected int
	}{
		{"Corner (A1)", board.A1, 2},
		{"Edge (A2)", board.A2, 3},
		{"Near Edge (B2)", board.B2, 4},
		{"Outer Ring (C2)", board.C2, 6},
		{"Center (D4)", board.D4, 8},
	}

	for _, tt := range tests {
		got := bits.OnesCount64(uint64(table[tt.square]))
		if got != tt.expected {
			t.Errorf("%s: expected %d attacks, got %d", tt.name, tt.expected, got)
		}
	}
}

func TestKnightWrapArounds(t *testing.T) {
	table := GetKnightAttackTable()

	tests := []struct {
		name           string
		square         board.Square
		forbiddenMasks board.Bitboard
	}{
		{"A1", board.A1, board.FileGMask | board.FileHMask},
		{"B1", board.B1, board.FileHMask},
		{"H8", board.H8, board.FileAMask | board.FileBMask},
		{"G8", board.G8, board.FileAMask},
	}

	for _, tt := range tests {
		illegalMoves := table[tt.square] & tt.forbiddenMasks
		if illegalMoves != 0 {
			t.Errorf("%s (square %d): detected wrap-around.", tt.name, tt.square)
		}
	}
}
