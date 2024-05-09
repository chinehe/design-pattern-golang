package observer

// Subject 订阅主题接口
type Subject interface {
	Register(observer Observer)
	Remove(observer Observer)
	Notify(message string)
}
