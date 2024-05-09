package impl

import (
	"design-pattern-golang/behavioral/template"
	"fmt"
)

// BankDepositBusiness 银行存款服务
type BankDepositBusiness struct {
	template.BankServiceTemplate
}

func (business *BankDepositBusiness) HandleBusiness() {
	fmt.Println("存款成功")
}
