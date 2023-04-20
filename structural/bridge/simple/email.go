package simple

import "fmt"

// EmailSender 邮件消息发送器
type EmailSender struct {
	UserName string
	Email string
}

func (sender *EmailSender)Send(msg string) error  {
	fmt.Printf("user(userName=%s,email=%s) send message:%s",sender.UserName,sender.Email,msg)
	return nil
}