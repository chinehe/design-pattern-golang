package method

// IRuleConfigParser 解析规则配置的解析器接口
type IRuleConfigParser interface {
	Parse(data []byte)
}

// IRuleConfigParserFactory 工厂方法接口
type IRuleConfigParserFactory interface {
	Create() IRuleConfigParser
}
