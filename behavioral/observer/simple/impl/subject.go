package impl

import (
	"design-pattern-golang/behavioral/observer/simple"
	"fmt"
)

// MySubject 我的主题
type MySubject struct {
	observers []simple.Observer
}

func NewMySubject() *MySubject {
	return &MySubject{
		observers: []simple.Observer{},
	}
}

func (subject *MySubject) Register(observer simple.Observer) {
	subject.observers = append(subject.observers,observer)
}

func (subject *MySubject) Remove(observer simple.Observer) {
	for i := range subject.observers {
		if subject.observers[i]==observer {
			subject.observers = append(subject.observers[:i],subject.observers[i+1:]...)
		}
	}
}

func (subject *MySubject) Notify(message string) {
	fmt.Println("my subject notify:"+message)
	for i := range subject.observers {
		go subject.observers[i].Update(message)
	}
}

