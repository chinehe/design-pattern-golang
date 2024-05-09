package impl

import "testing"

func TestMediator(t *testing.T) {
	mediator := NewHouseMediator()
	buyer1 := NewBuyer("AAA", mediator)
	buyer2 := NewBuyer("BBB", mediator)
	buyer3 := NewBuyer("CCC", mediator)
	buyer4 := NewBuyer("DDD", mediator)
	mediator.RegisterCustomer(buyer1)
	mediator.RegisterCustomer(buyer2)
	mediator.RegisterCustomer(buyer3)
	mediator.RegisterCustomer(buyer4)
	buyer1.SendMessage("AAAAAAA")
}
