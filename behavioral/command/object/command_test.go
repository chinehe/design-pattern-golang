package object

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCommand(t *testing.T) {
	door := new(GarageDoor)
	light := new(Light)

	events := []string{"open door", "close door", "on light", "off light"}
	rand.Seed(time.Now().Unix())

	for i := 0; i < 100; i++ {
		event := events[rand.Intn(4)]
		fmt.Println("--------")
		fmt.Println(event)
		switch event {
		case "open door":
			NewDoorOpenCommand(door).Execute()
		case "close door":
			NewDoorCloseCommand(door).Execute()
		case "on light":
			NewLightOnCommand(light).Execute()
		case "off light":
			NewLightOffCommand(light).Execute()
		}
	}

}
