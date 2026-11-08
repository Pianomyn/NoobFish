package board

type MoveFlags struct {
	DoublePush  bool
	EnPassant   bool
	CastleKing  bool
	CastleQueen bool
	Promotion   bool
}

type Move struct {
	Source        Square
	Target        Square
	SourcePiece   ColouredPiece
	CapturedPiece ColouredPiece
	PromotionTo   Piece
	Flags         MoveFlags
}
