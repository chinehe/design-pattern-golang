package hungry

// initializationHungrySingletonInstance 唯一单例
// 如果要外部不能修改该唯一单例，可以将该变量首字母小写，并提供GetInstance方法来获取唯一实例
var initializationHungrySingletonInstance *InitializationHungrySingleton

// InitializationHungrySingleton 带初始化函数的饿汉式单例
type InitializationHungrySingleton struct {
	Name string
}

// 初始化函数
func init() {
	initializationHungrySingletonInstance = &InitializationHungrySingleton{"带初始化函数的饿汉式单例"}
}

// GetInitializationHungrySingletonInstance 获取唯一实例
// 方法一般命名为GetInstance
// 如果不介意有外部代码修改唯一单例的风险，可以将该变量首字母大写，就不需要该方法了
func GetInitializationHungrySingletonInstance() *InitializationHungrySingleton {
	return initializationHungrySingletonInstance
}
