package simple

import "fmt"

// SecurityCamera 安全摄像头
type SecurityCamera struct {

}

func (camera *SecurityCamera) On()  {
	fmt.Println("security camera power on")
}

func (camera *SecurityCamera) Off()  {
	fmt.Println("security camera power off")
}
