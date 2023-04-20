package hungry

// SimpleHungrySingletonInstance 简单饿汉式单例的唯一实例
// 在项目启动时初始化
var simpleHungrySingletonInstance = &SimpleHungrySingleton{"简单饿汉式单例"}

// SimpleHungrySingleton 简单饿汉式单例
type SimpleHungrySingleton struct {
	Name string
}

// GetSimpleHungrySingletonInstance 获取唯一实例
// 方法一般命名为GetInstance
// 如果不介意有外部代码修改唯一单例的风险，可以将该变量首字母大写，就不需要该方法了
func GetSimpleHungrySingletonInstance() *SimpleHungrySingleton {
	return simpleHungrySingletonInstance
}
