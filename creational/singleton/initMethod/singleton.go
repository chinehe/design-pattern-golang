package initMethod

// instance 唯一实例
var instance = new(MyLogger)

func GetInstance() *MyLogger {
	return instance
}
