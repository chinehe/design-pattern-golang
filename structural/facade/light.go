package facade

import "fmt"

// Light 电灯
type Light struct {
}

func (light *Light) On() {
	fmt.Println("light power on")
}

func (light *Light) Off() {
	fmt.Println("light power off")
}
