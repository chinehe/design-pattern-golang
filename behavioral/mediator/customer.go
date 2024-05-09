package mediator

// Customer 客户接口
type Customer interface {
	SendMessage(msg string)
	ReceiveMessage(msg string)
}
