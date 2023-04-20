package simple

// MessageSender 消息发送器
type MessageSender interface {
	Send(msg string) error
}