package syncOnce

import (
	"sync"
)

// instance 唯一实例
var instance *MyLogger

var once sync.Once

func GetInstance() *MyLogger {
	once.Do(func() {
		instance = new(MyLogger)
	})
	return instance
}
