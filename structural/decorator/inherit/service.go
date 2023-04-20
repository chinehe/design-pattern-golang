package inherit

// Service 服务接口
type Service interface {
	ServiceA()
	ServiceB()
}

// DecoratedService 装饰后的服务接口
type DecoratedService interface {
	Service
	ServiceC()
}
