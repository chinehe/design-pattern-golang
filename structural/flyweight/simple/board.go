package simple

import "fmt"

// Piece 棋盘上的棋子
type Piece struct {
	ChessPiece *ChessPiece
	X int
	Y int
}

// ChessBoard 棋盘
type ChessBoard struct {
	ChessPieces [15][15]*Piece
}

func NewChessBoard() *ChessBoard {
	return &ChessBoard{ChessPieces: [15][15]*Piece{}}
}

func (board *ChessBoard)Print()  {
	fmt.Println("-----------------------------------------")
	header := "0"
	for i := 1; i < 16; i++ {
		header += fmt.Sprintf(" %v ",i%10)
	}
	fmt.Println(header)
	for line := range board.ChessPieces {
		lineRes := fmt.Sprintf("%v",(line+1)%10)
		for column := range board.ChessPieces[line] {
			if board.ChessPieces[line][column]==nil {
				lineRes+=" + "
				continue
			}
			if board.ChessPieces[line][column].ChessPiece.Color=="white" {
				lineRes+=" ⚪ "
			}else {
				lineRes+=" ● "
			}
		}
		fmt.Println(lineRes)
	}
}

func (board *ChessBoard)Add(piece *Piece)  {
	board.ChessPieces[piece.X][piece.Y] = piece
	board.Print()
}
