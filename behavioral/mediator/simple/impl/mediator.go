package impl

import (
	"design-pattern-golang/behavioral/mediator/simple"
	"fmt"
)

// HouseMediator 中介
type HouseMediator struct {
	Customers []simple.Customer
}

func NewHouseMediator() *HouseMediator {
	return &HouseMediator{
		Customers: []simple.Customer{},
	}
}

func (h *HouseMediator) RegisterCustomer(customer simple.Customer) {
	h.Customers = append(h.Customers,customer)
}

func (h *HouseMediator) Relay( msg string) {
	fmt.Println("mediator relay msg:",msg)
	for i := range h.Customers {
		h.Customers[i].ReceiveMessage(msg)
	}
}

