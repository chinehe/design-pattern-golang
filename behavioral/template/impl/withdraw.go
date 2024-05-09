package impl

import (
	"design-pattern-golang/behavioral/template"
	"fmt"
)

// BankWithdrawBusiness 银行取款服务
type BankWithdrawBusiness struct {
	template.BankServiceTemplate
}

func (business *BankWithdrawBusiness) HandleBusiness() {
	fmt.Println("取款成功")
}
