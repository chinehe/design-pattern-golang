package class

import "fmt"

// AdapterService 适配器服务
type AdapterService struct {
	OriginService *OriginService
}

func (service *AdapterService) ServiceA() {
	service.OriginService.ServiceA()
}

func (service *AdapterService) ServiceB() {
	fmt.Printf("AdapterService ServiceB\n")
}

