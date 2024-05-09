package template

import "fmt"

// BankServiceTemplate 银行服务模板
type BankServiceTemplate struct {
	BankService
}

// Service 服务模板方法
// 注意这里需要传入服务子类，因为Go语言中没有抽象方法
func (template *BankServiceTemplate) Service(t BankService) {
	t.TakeNumber()
	t.LineUp()
	t.HandleBusiness()
	t.Rate()
}

func (template BankServiceTemplate) TakeNumber() {
	fmt.Println("取号")
}

func (template BankServiceTemplate) LineUp() {
	fmt.Println("开始排队...")
	fmt.Println("结束排队!")
}

func (template BankServiceTemplate) Rate() {
	fmt.Println("客户评分")
}
