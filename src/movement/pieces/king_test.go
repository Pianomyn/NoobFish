package movement

import (
	"chessgo/board"
	"math/bits"
	"testing"
)

func TestKingAttackCounts(t *testing.T) {
	table := GetKingAttackTable()
	tests := []struct {
		name     string
		square   board.Square
		expected int
	}{

		{"Corner (A1)", board.A1, 3},
		{"Edge (A2)", board.A2, 5},
		{"Center (D4)", board.D4, 8},
	}

	for _, tt := range tests {
		got := bits.OnesCount64(uint64(table[tt.square]))
		if got != tt.expected {
			t.Errorf("%s: expected %d attacks, got %d", tt.name, tt.expected, got)
		}
	}
}

func TestKingWrapArounds(t *testing.T) {
	table := GetKingAttackTable()

	tests := []struct {
		name           string
		square         board.Square
		forbiddenMasks board.Bitboard
	}{
		{"A1", board.A1, board.FileHMask},
		{"A4", board.A4, board.FileHMask},
		{"H1", board.H1, board.FileAMask},
		{"H8", board.H8, board.FileAMask},
	}

	for _, tt := range tests {
		illegalMoves := table[tt.square] & tt.forbiddenMasks
		if illegalMoves != 0 {
			t.Errorf("%s (square %d): detected wrap-around.", tt.name, tt.square)
		}
	}
}
