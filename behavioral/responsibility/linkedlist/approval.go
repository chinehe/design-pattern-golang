package linkedlist

// LeaveApproval 请假审批接口
type LeaveApproval interface {
	LeaveApproval(name,reason string,days int)
}
