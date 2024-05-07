package decorator

import "fmt"

// MyService 服务接口实现类
type MyService struct {
}

func (m *MyService) ServiceA() {
	fmt.Printf("myservice serviceA\n")
}

func (m *MyService) ServiceB() {
	fmt.Printf("myservice serviceB\n")
}

// MyDecoratedService 装饰后的服务接口的实现类
type MyDecoratedService struct {
}

func (m *MyDecoratedService) ServiceA() {
	fmt.Printf("MyDecoratedService serviceA\n")
}

func (m *MyDecoratedService) ServiceB() {
	fmt.Printf("MyDecoratedService serviceB\n")
}

func (m *MyDecoratedService) ServiceC() {
	fmt.Printf("MyDecoratedService serviceC\n")
}
