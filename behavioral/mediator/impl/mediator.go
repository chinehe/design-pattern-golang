package impl

import (
	"design-pattern-golang/behavioral/mediator"
	"fmt"
)

// HouseMediator 中介
type HouseMediator struct {
	Customers []mediator.Customer
}

func NewHouseMediator() *HouseMediator {
	return &HouseMediator{
		Customers: []mediator.Customer{},
	}
}

func (h *HouseMediator) RegisterCustomer(customer mediator.Customer) {
	h.Customers = append(h.Customers, customer)
}

func (h *HouseMediator) Relay(msg string) {
	fmt.Println("mediator relay msg:", msg)
	for i := range h.Customers {
		h.Customers[i].ReceiveMessage(msg)
	}
}
