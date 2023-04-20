package _interface

import "fmt"

// MyService 实际编写的服务
type MyService struct {
	AdapterService
}

func (service MyService) ServiceA() {
	fmt.Println("MyService ServiceA")
}


