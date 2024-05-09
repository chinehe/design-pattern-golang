package impl

import (
	"design-pattern-golang/behavioral/mediator"
	"fmt"
)

// Buyer 买家
type Buyer struct {
	Name     string
	Mediator mediator.Mediator
}

func NewBuyer(name string, mediator mediator.Mediator) *Buyer {
	return &Buyer{Name: name, Mediator: mediator}
}

func (b *Buyer) SendMessage(msg string) {
	fmt.Println(b.Name, "send msg:", msg)
	b.Mediator.Relay(msg)
}

func (b *Buyer) ReceiveMessage(msg string) {
	fmt.Println(b.Name, "receive msg:", msg)
}
