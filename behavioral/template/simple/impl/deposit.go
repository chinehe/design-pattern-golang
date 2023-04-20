package impl

import (
	"design-pattern-golang/behavioral/template/simple"
	"fmt"
)

// BankDepositBusiness 银行存款服务
type BankDepositBusiness struct {
	simple.BankServiceTemplate
}

func (business *BankDepositBusiness) HandleBusiness() {
	fmt.Println("存款成功")
}
