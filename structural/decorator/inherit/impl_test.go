package inherit

import "testing"

func TestDecorate(t *testing.T) {
	new(MyService).ServiceA()
	new(MyService).ServiceB()
	new(MyDecoratedService).ServiceA()
	new(MyDecoratedService).ServiceB()
	new(MyDecoratedService).ServiceC()
}
