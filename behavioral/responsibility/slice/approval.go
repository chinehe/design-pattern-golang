package slice

// LeaveApproval 请假审批接口
type LeaveApproval interface {
	LeaveApproval(context *Context) bool
}
