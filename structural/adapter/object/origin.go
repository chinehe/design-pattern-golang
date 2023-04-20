package class

import "fmt"

// OriginService 要被适配的类
type OriginService struct {
}

// ServiceA OriginService原本有的服务方法
func (service *OriginService) ServiceA() {
	fmt.Printf("OriginService ServiceA\n")
}


