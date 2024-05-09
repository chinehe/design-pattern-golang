package observer

// Observer 观察者接口
type Observer interface {
	Update(msg string)
}
