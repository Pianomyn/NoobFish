package movement

import (
	"chessgo/board"
	"math/bits"
	"testing"
)

func TestQueenAttackCounts(t *testing.T) {
	table := GetQueenAttackTable()
	tests := []struct {
		name     string
		square   board.Square
		expected int
	}{
		{"A1", board.A1, 21},
		{"A2", board.A2, 21},
		{"A4", board.A4, 21},
		{"B2", board.B2, 23},
		{"B3", board.B3, 23},
		{"B4", board.B4, 23},
		{"C3", board.C3, 25},
		{"C4", board.C5, 25},
		{"D4", board.D4, 27},
	}

	for _, tt := range tests {
		got := bits.OnesCount64(uint64(table[tt.square]))
		if got != tt.expected {
			t.Errorf("%s: expected %d attacks, got %d", tt.name, tt.expected, got)
		}
	}
}
