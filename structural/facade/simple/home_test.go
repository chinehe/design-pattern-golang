package simple

import "testing"

func TestFacade(t *testing.T) {
	home := NewHome(&SecurityCamera{},&Light{},&Sound{})
	home.Leave()
	home.ComeBack()
}