package visitor

// ResourceFileVisitor 资源文件观察者接口（观察者接口）
type ResourceFileVisitor interface {
	Visit(file ResourceFile) error
}
