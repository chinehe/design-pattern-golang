package impl


import (
	"design-pattern-golang/behavioral/template/simple"
	"fmt"
)

// BankWithdrawBusiness 银行取款服务
type BankWithdrawBusiness struct {
	simple.BankServiceTemplate
}

func (business *BankWithdrawBusiness) HandleBusiness() {
	fmt.Println("取款成功")
}

