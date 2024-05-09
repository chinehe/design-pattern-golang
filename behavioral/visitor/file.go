package visitor

// ResourceFile 资源文件接口（被观察者接口）
type ResourceFile interface {
	Accept(printer ResourceFileVisitor) error
}
