package slice

import "fmt"

type ClassMaster struct {
}

func (master *ClassMaster) LeaveApproval(context *Context) bool {
	if context.days>1 {
		fmt.Println("班主任只能审批不大于 1 天的假期")
		return false
	}
	fmt.Println("班主任审批通过!")
	return true
}


type President struct {

}

func (president *President) LeaveApproval(context *Context) bool {
	if context.days>3 {
		fmt.Println("院长只能审批不大于 3 天的假期")
		return false
	}
	fmt.Println("院长审批通过!")
	return true
}

type Principal struct {

}

func (principal *Principal) LeaveApproval(context *Context) bool {
	if context.days>7 {
		fmt.Println("校长只通过不大于 7 天的假期")
		return false
	}
	fmt.Println("校长长审批通过!")
	return true
}