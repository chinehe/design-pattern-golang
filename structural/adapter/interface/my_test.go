package _interface

import "testing"

func TestInterfaceAdapter(t *testing.T) {
	service := MyService{}
	service.ServiceA()
	service.ServiceB()
	service.ServiceC()
}
