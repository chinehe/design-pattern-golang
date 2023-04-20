package simple

import "fmt"

// Home 家(外观类)
type Home struct {
	SecurityCamera *SecurityCamera
	Light *Light
	Sound *Sound
}

func NewHome(securityCamera *SecurityCamera, light *Light, sound *Sound) *Home {
	return &Home{SecurityCamera: securityCamera, Light: light, Sound: sound}
}

func (home *Home) ComeBack()  {
	fmt.Println("come back...")
	home.Light.On()
	home.SecurityCamera.Off()
	home.Sound.On()
	home.Sound.Start()
}

func (home *Home) Leave()  {
	fmt.Println("leave...")
	home.Light.Off()
	home.SecurityCamera.On()
	home.Sound.Stop()
	home.Sound.Off()
}
