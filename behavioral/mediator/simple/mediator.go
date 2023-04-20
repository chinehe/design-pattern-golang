package simple

// Mediator 中介者接口
type Mediator interface {
	RegisterCustomer(customer Customer)
	Relay(msg string)
}
