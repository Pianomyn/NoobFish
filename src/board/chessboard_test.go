package board

import (
	"fmt"
	"testing"
)

func TestNewChessBoard(t *testing.T) {
	cb := NewChessBoard()

	if cb.SideToMove != White {
		t.Error("ChessBoard initial side to move should be White")
	}

	pieceTests := []struct {
		pos      Square
		expected ColouredPiece
	}{
		{A1, WhiteRook}, {B1, WhiteKnight}, {C1, WhiteBishop}, {D1, WhiteQueen},
		{E1, WhiteKing}, {F1, WhiteBishop}, {G1, WhiteKnight}, {H1, WhiteRook},
		{A8, BlackRook}, {B8, BlackKnight}, {C8, BlackBishop}, {D8, BlackQueen},
		{E8, BlackKing}, {F8, BlackBishop}, {G8, BlackKnight}, {H8, BlackRook},
	}
	for i := 0; i < 8; i++ {
		pieceTests = append(pieceTests,
			struct {
				pos      Square
				expected ColouredPiece
			}{A2 + Square(i), WhitePawn},
			struct {
				pos      Square
				expected ColouredPiece
			}{A7 + Square(i), BlackPawn},
		)
	}

	for _, tt := range pieceTests {
		t.Run(fmt.Sprintf("Square_%v", tt.pos), func(t *testing.T) {
			if cb.Mailbox[tt.pos] != tt.expected {
				t.Errorf("Mailbox at %v: expected %v, got %v", tt.pos, tt.expected, cb.Mailbox[tt.pos])
			}

			if tt.expected != NoColouredPiece {
				side := Side(uint8(tt.expected) / 6)
				piece := Piece(uint8(tt.expected) % 6)

				if !cb.Pieces[side][piece].Get(tt.pos) {
					t.Errorf("Bitboard for %v missing at %v", piece, tt.pos)
				}
				if !cb.Colours[side].Get(tt.pos) {
					t.Errorf("Colour bitboard for %v missing at %v", side, tt.pos)
				}
			}
		})
	}
}

func TestSyncBasic(t *testing.T) {
	cb := &ChessBoard{}
	cb.Pieces[White][Rook] = cb.Pieces[White][Rook].Clear(A1)
	cb.Pieces[White][Rook] = cb.Pieces[White][Rook].Set(A3)
	cb.Pieces[Black][Queen] = cb.Pieces[Black][Queen].Clear(H8)
	cb.Pieces[Black][Queen] = cb.Pieces[Black][Queen].Set(H8)

	cb.Sync()

	if cb.Colours[White].Get(A1) {
		t.Errorf("White colour bitboard didn't clear rook")
	}
	if !cb.Colours[White].Get(A3) {
		t.Errorf("White colour bitboard didn't place rook")
	}
	if !cb.Colours[Black].Get(H8) {
		t.Errorf("Black colour bitboard missing queen square")
	}
	if cb.Occupied.Get(A1) || !cb.Occupied.Get(A3) || !cb.Occupied.Get(H8) {
		t.Errorf("Occupied bitboard incorrect")
	}
}

func TestPieceToColouredPiece(t *testing.T) {
	cases := []struct {
		piece    Piece
		side     Side
		expected ColouredPiece
	}{
		{Pawn, White, WhitePawn},
		{Pawn, Black, BlackPawn},
		{Rook, White, WhiteRook},
		{Rook, Black, BlackRook},
		{Knight, White, WhiteKnight},
		{Knight, Black, BlackKnight},
		{Bishop, White, WhiteBishop},
		{Bishop, Black, BlackBishop},
		{King, White, WhiteKing},
		{King, Black, BlackKing},
		{Queen, White, WhiteQueen},
		{Queen, Black, BlackQueen},
		{NoPiece, White, NoColouredPiece},
		{NoPiece, Black, NoColouredPiece},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Piece to Coloured Piece %v %v", c.side, c.piece), func(t *testing.T) {
			colouredPiece := PieceToColouredPiece(c.piece, c.side)
			if colouredPiece != c.expected {
				t.Errorf("Expected %v, got %v", c.expected, colouredPiece)
			}
		})
	}
}

func TestColouredPieceToPiece(t *testing.T) {
	cases := []struct {
		expected      Piece
		colouredPiece ColouredPiece
	}{
		{Pawn, WhitePawn},
		{Pawn, BlackPawn},
		{Rook, WhiteRook},
		{Rook, BlackRook},
		{Knight, WhiteKnight},
		{Knight, BlackKnight},
		{Bishop, WhiteBishop},
		{Bishop, BlackBishop},
		{King, WhiteKing},
		{King, BlackKing},
		{Queen, WhiteQueen},
		{Queen, BlackQueen},
		{NoPiece, NoColouredPiece},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("Coloured Piece to Piece %v %v", c.colouredPiece, c.expected), func(t *testing.T) {
			piece := ColouredPieceToPiece(c.colouredPiece)
			if piece != c.expected {
				t.Errorf("Expected %v, got %v", c.expected, piece)
			}
		})
	}

}
