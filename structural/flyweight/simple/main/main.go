package main

import (
	"bufio"
	"design-pattern-golang/structural/flyweight/simple"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var count int

func getNext() string {
	if count%2==0 {
		count++
		return "black"
	}else {
		count++
		return "white"
	}
}
func main() {
	factory := simple.ChessPieceFactory{}
	chessBoard := simple.NewChessBoard()
	chessBoard.Print()
	for  {
		time.Sleep(time.Second)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Println(input)
		input = strings.TrimRight(string(input),"\n")
		split := strings.Split(input, " ")
		if len(split)!=2 {
			fmt.Println("input error")
			continue
		}
		x ,_:= strconv.Atoi(split[0])
		y ,_:= strconv.Atoi(split[1])
		chessBoard.Add(&simple.Piece{
			ChessPiece: factory.Create(getNext()),
			X:          x-1,
			Y:          y-1,
		})
	}
}