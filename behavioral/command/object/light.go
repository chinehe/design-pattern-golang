package object

import "fmt"

// Light 电灯
type Light struct {
	IsOn bool
}

func (light *Light) On()  {
	if light.IsOn {
		fmt.Println("电灯本来就亮着")
	}else {
		light.IsOn = true
		fmt.Println("打开电灯了...")
	}
}


func (light *Light) Off()  {
	if light.IsOn {
		light.IsOn = false
		fmt.Println("关闭电灯了...")
	}else {
		fmt.Println("电灯本来就关着")
	}
}
