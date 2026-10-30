package board

import (
	"math/bits"
)

type Side uint8
type Piece uint8
type ColouredPiece uint8

const (
	White Side = iota
	Black
)

func (s Side) String() string {
	if s == White {
		return "White"
	}
	return "Black"
}

func (cb *ChessBoard) changeSideToMove() {
	cb.SideToMove ^= 1
}

const (
	Pawn Piece = iota
	Knight
	Bishop
	Rook
	Queen
	King
	NoPiece = 255
)

const (
	WhitePawn ColouredPiece = iota
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing
	BlackPawn
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing
	NoColouredPiece = 255
)

func PieceToColouredPiece(p Piece, s Side) ColouredPiece {
	return ColouredPiece(6*uint8(s) + uint8(p))
}

func (c Piece) String() string {
	return []string{
		"Pawn",
		"Knight",
		"Bishop",
		"Rook",
		"Queen",
		"King",
	}[c]
}

type CastlingRights struct {
	WhiteKingside  bool
	WhiteQueenside bool
	BlackKingside  bool
	BlackQueenside bool
}

type ChessBoard struct {
	Pieces  [2][6]Bitboard    // Piece to square
	Mailbox [64]ColouredPiece // Square to piece

	Colours         [2]Bitboard
	Occupied        Bitboard
	SideToMove      Side
	CastlingRights  CastlingRights
	EnPassantSquare Square
}

func NewChessBoard() *ChessBoard {
	cb := &ChessBoard{
		SideToMove: White,
		CastlingRights: CastlingRights{
			WhiteKingside:  true,
			WhiteQueenside: true,
			BlackKingside:  true,
			BlackQueenside: true,
		},
		EnPassantSquare: NoSquare,
	}

	// Pawns
	cb.Pieces[White][Pawn] |= Rank2Mask
	cb.Pieces[Black][Pawn] |= Rank7Mask

	// Knights
	cb.Pieces[White][Knight] = cb.Pieces[White][Knight].Set(B1).Set(G1)
	cb.Pieces[Black][Knight] = cb.Pieces[Black][Knight].Set(B8).Set(G8)

	// Bishops
	cb.Pieces[White][Bishop] = cb.Pieces[White][Bishop].Set(C1).Set(F1)
	cb.Pieces[Black][Bishop] = cb.Pieces[Black][Bishop].Set(C8).Set(F8)

	// Rooks
	cb.Pieces[White][Rook] = cb.Pieces[White][Rook].Set(A1).Set(H1)
	cb.Pieces[Black][Rook] = cb.Pieces[Black][Rook].Set(A8).Set(H8)

	// Queens
	cb.Pieces[White][Queen] = cb.Pieces[White][Queen].Set(D1)
	cb.Pieces[Black][Queen] = cb.Pieces[Black][Queen].Set(D8)

	// Kings
	cb.Pieces[White][King] = cb.Pieces[White][King].Set(E1)
	cb.Pieces[Black][King] = cb.Pieces[Black][King].Set(E8)

	cb.Sync()
	return cb
}

func (cb *ChessBoard) Sync() {
	cb.Colours[White] = 0
	cb.Colours[Black] = 0
	for i := range cb.Mailbox {
		cb.Mailbox[i] = NoColouredPiece
	}

	for side := White; side <= Black; side++ {
		for p := Pawn; p <= King; p++ {
			cb.Colours[side] |= cb.Pieces[side][p]

			instances := cb.Pieces[side][p]
			for instances > 0 {
				sq := GetNextPieceSquare(&instances)
				cb.Mailbox[sq] = PieceToColouredPiece(p, side)
			}
		}
	}

	cb.Occupied = cb.Colours[White] | cb.Colours[Black]
}

func GetNextPieceSquare(pieces *Bitboard) Square {
	/*
		For a bitboard containing locations of all pieces of a particular type,
		yield the location of the next piece
	*/

	lsbIndex := Square(bits.TrailingZeros64(uint64(*pieces)))

	*pieces = pieces.Clear(lsbIndex)
	return lsbIndex
}
