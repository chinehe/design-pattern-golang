package simple

// ConfigParser 配置解析器接口
type ConfigParser interface {
	Parse(data []byte)
}

// ConfigParserFactory 配置解析器工厂接口
type ConfigParserFactory interface {
	Create(name string) ConfigParser
}