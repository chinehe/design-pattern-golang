package impl

import "testing"

func TestObserver(t *testing.T) {
	subject := NewMySubject()
	subject.Register(&MyObserverA{})
	subject.Register(&MyObserverB{})
	subject.Register(&MyObserverC{})
	subject.Notify("Hello Observer")
}
