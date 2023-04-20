package linkedlist

import "fmt"

type ClassMaster struct {
	ApprovalChain
}

func (master *ClassMaster)LeaveApproval(name,reason string,days int)  {
	if days>1 {
		fmt.Println("班主任只能审批不大于 1 天的假期")
		if master.Next==nil{
			fmt.Println("审批驳回")
		}else {
			master.Next.LeaveApproval(name,reason,days)
		}
	}else {
		fmt.Println("审批通过：",name,reason,days)
	}
}

type President struct {
	ApprovalChain
}

func (master *President)LeaveApproval(name,reason string,days int)  {
	if days>3 {
		fmt.Println("院长只能审批不大于 3 天的假期")
		if master.Next==nil{
			fmt.Println("审批驳回")
		}else {
			master.Next.LeaveApproval(name,reason,days)
		}
	}else {
		fmt.Println("审批通过：",name,reason,days)
	}
}

type Principal struct {
	ApprovalChain
}

func (master *Principal)LeaveApproval(name,reason string,days int)  {
	if days>7 {
		fmt.Println("校长只通过不大于 7 天的假期")
		if master.Next==nil{
			fmt.Println("审批驳回")
		}else {
			master.Next.LeaveApproval(name,reason,days)
		}
	}else {
		fmt.Println("审批通过：",name,reason,days)
	}
}
