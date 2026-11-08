package movement

import (
	"chessgo/board"
	"math/bits"
	"testing"
)

func TestPawnPushCounts(t *testing.T) {
	whitePushes := GetPawnPushTable(board.White)
	blackPushes := GetPawnPushTable(board.Black)

	tests := []struct {
		name     string
		table    []board.Bitboard
		square   board.Square
		expected int
	}{
		// White Pawns
		{"A2", whitePushes, board.A2, 2},
		{"D4", whitePushes, board.D4, 1},
		{"H7", whitePushes, board.H7, 1},

		// Black Pawns
		{"E7", blackPushes, board.E7, 2},
		{"E6", blackPushes, board.E6, 1},
		{"A2", blackPushes, board.A2, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bits.OnesCount64(uint64(tt.table[tt.square]))
			if got != tt.expected {
				t.Errorf("%s: expected %d pushes, got %d", tt.name, tt.expected, got)
			}
		})
	}
}

func TestPawnAttackCounts(t *testing.T) {
	whiteTable := GetPawnAttackTable(board.White)
	blackTable := GetPawnAttackTable(board.Black)

	tests := []struct {
		name     string
		table    []board.Bitboard
		square   board.Square
		expected int
	}{
		{"A2", whiteTable, board.A2, 1},
		{"H2", whiteTable, board.H2, 1},
		{"D2", whiteTable, board.D2, 2},
		{"D4", whiteTable, board.D4, 2},

		{"A7", blackTable, board.A7, 1},
		{"H7", blackTable, board.H7, 1},
		{"D7", blackTable, board.D7, 2},
		{"D5", blackTable, board.D5, 2},

		{"White_A8_Invalid", whiteTable, board.A8, 0},
		{"Black_A1_Invalid", blackTable, board.A1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bits.OnesCount64(uint64(tt.table[tt.square]))
			if got != tt.expected {
				t.Errorf("Square %v: expected %d attacks, got %d", tt.square, tt.expected, got)
			}
		})
	}
}
