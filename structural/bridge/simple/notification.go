package simple

// Notification 通知接口
type Notification interface {
	Notify(msg string) error
}
