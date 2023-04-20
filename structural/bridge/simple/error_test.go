package simple

import "testing"

func TestErrorNotification_Notify(t *testing.T) {
	notification := &ErrorNotification{Sender: &EmailSender{
		UserName: "chinehe",
		Email:    "chinehe@chinehe.com",
	}}
	notification.Notify("你好")
}
