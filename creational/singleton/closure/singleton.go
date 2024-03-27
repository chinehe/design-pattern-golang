package closure

// instance 唯一实例
var instance = func() *MyLogger {
	return new(MyLogger)
}()

func GetInstance() *MyLogger {
	return instance
}
