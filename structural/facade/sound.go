package facade

import "fmt"

// Sound 音响
type Sound struct {
}

func (sound *Sound) On() {
	fmt.Println("sound power on")
}

func (sound *Sound) Off() {
	fmt.Println("sound power off")
}

func (sound *Sound) Start() {
	fmt.Println("sound start sing")
}

func (sound *Sound) Stop() {
	fmt.Println("sound stop sing")
}
