package parameter

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
			DoorOpenCommand(door)()
		case "close door":
			DoorCloseCommand(door)()
		case "on light":
			LightOnCommand(light)()
		case "off light":
			LightOffCommand(light)()
		}
	}
}
