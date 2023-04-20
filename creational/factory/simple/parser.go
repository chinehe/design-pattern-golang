package simple

// IRuleConfigParser 解析规则配置的解析器接口
type IRuleConfigParser interface {
	Parse(data []byte)
}
