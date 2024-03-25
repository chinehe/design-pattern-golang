package lazy

import "sync"

// instance 唯一实例
var instance *LazybonesSingleton

// once 用于保证并发安全
var once = &sync.Once{}

// LazybonesSingleton 懒汉式单例
type LazybonesSingleton struct {
	Name string
}

// GetInstance 获取唯一实例
func GetInstance() *LazybonesSingleton {
	if instance == nil {
		once.Do(func() {
			instance = &LazybonesSingleton{"懒汉式（双重检测）单例"}
		})
	}
	return instance
}
