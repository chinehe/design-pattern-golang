package simple

// NewConfigParser 抽象工厂模式的工厂方法
func NewConfigParser(format,environment string) ConfigParser  {
	switch format {
	case "json":
		return (&JsonParserFactory{}).Create(environment)
	case "yaml":
		return (&YamlParserFactory{}).Create(environment)
	default:
		return nil
	}
}
