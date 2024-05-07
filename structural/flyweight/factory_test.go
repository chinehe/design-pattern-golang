package flyweight

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

var count int

func getNext() string {
	if count%2 == 0 {
		count++
		return "black"
	} else {
		count++
		return "white"
	}
}

func TestFlyweight(t *testing.T) {
	factory := ChessPieceFactory{}
	chessBoard := NewChessBoard()
	chessBoard.Print()
	for {
		time.Sleep(time.Second)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Println(input)
		input = strings.TrimRight(string(input), "\n")
		split := strings.Split(input, " ")
		if len(split) != 2 {
			fmt.Println("input error")
			continue
		}
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		chessBoard.Add(&Piece{
			ChessPiece: factory.Create(getNext()),
			X:          x - 1,
			Y:          y - 1,
		})
	}
}
