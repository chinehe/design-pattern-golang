package simple

// ErrorNotification 错误通知
type ErrorNotification struct {
	Sender MessageSender
}

func (notification *ErrorNotification) Notify(msg string) error {
	return notification.Sender.Send("error:\n"+msg)
}



