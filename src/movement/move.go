package movement

import (
	"chessgo/board"
)

type MoveFlags struct {
	DoublePush  bool
	EnPassant   bool
	CastleKing  bool
	CastleQueen bool
	Promotion   bool
}

type Move struct {
	Source        board.Square
	Target        board.Square
	SourcePiece   board.ColouredPiece
	CapturedPiece board.ColouredPiece
	PromotionTo   board.Piece
	Flags         MoveFlags
}
