package impl

import (
	"design-pattern-golang/behavioral/mediator/simple"
	"fmt"
)

// Buyer 买家
type Buyer struct {
	Name string
	Mediator simple.Mediator
}

func NewBuyer(name string, mediator simple.Mediator) *Buyer {
	return &Buyer{Name: name, Mediator: mediator}
}

func (b *Buyer) SendMessage(msg string) {
	fmt.Println(b.Name,"send msg:",msg)
	b.Mediator.Relay(msg)
}

func (b *Buyer) ReceiveMessage(msg string) {
	fmt.Println(b.Name,"receive msg:",msg)
}

