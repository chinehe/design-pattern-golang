package impl

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	depositBusiness := &BankDepositBusiness{}
	withdrawBusiness := &BankWithdrawBusiness{}
	depositBusiness.Service(depositBusiness)
	withdrawBusiness.Service(withdrawBusiness)
}
