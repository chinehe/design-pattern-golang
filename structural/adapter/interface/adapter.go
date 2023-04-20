package _interface

import "fmt"

// AdapterService 适配器服务
type AdapterService struct {
}

func (a *AdapterService) ServiceA() {
	fmt.Println("AdapterService ServiceA")
}

func (a *AdapterService) ServiceB() {
	fmt.Println("AdapterService ServiceB")
}

func (a *AdapterService) ServiceC() {
	fmt.Println("AdapterService ServiceC")
}




