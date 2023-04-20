package class

import "fmt"

// AdapterService 适配器服务
type AdapterService struct {
	*OriginService
}

func (service *AdapterService) ServiceB() {
	fmt.Printf("AdapterService ServiceB\n")
}

