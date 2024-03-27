package declareAndInit

// instance 唯一实例
var instance *MyLogger

func init() {
	instance = new(MyLogger)
}

func GetInstance() *MyLogger {
	return instance
}
