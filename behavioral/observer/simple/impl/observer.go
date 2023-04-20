package impl

import (
	"fmt"
)

// MyObserverA 我的观察者A
type MyObserverA struct {

}

func (observer *MyObserverA) Update(msg string) {
	fmt.Println("my observer A update:"+msg)
}

// MyObserverB 我的观察者B
type MyObserverB struct {

}

func (observer *MyObserverB) Update(msg string) {
	fmt.Println("my observer B update:"+msg)
}

// MyObserverC 我的观察者C
type MyObserverC struct {

}

func (observer *MyObserverC) Update(msg string) {
	fmt.Println("my observer C update:"+msg)
}

