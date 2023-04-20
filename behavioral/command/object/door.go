package object

import "fmt"

// GarageDoor 大门
type GarageDoor struct {
	IsOpen bool
}

func (door *GarageDoor) Open()  {
	if door.IsOpen {
		fmt.Println("门本来就开着")
	}else {
		door.IsOpen = true
		fmt.Println("打开门了...")
	}
}


func (door *GarageDoor) Close()  {
	if door.IsOpen {
		door.IsOpen = false
		fmt.Println("关闭门了...")
	}else {
		fmt.Println("门本来就关着")
	}
}
