package linkedlist

import "testing"

func TestResponsibilityChain(t *testing.T) {
	classMaster := &ClassMaster{}
	president := &President{}
	principal := &Principal{}
	classMaster.Next = president
	president.Next = principal
	classMaster.LeaveApproval("Chine","啦啦啦",8)
}
