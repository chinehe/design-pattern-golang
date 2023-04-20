package method

// NewIRuleConfigParserFactory 工厂方法的工厂方法
func NewIRuleConfigParserFactory(parserName string) IRuleConfigParserFactory {
	switch parserName {
	case "json":
		return &JsonParserFactory{}
	case "yaml":
		return &YamlParserFactory{}
	default:
		return nil
	}
}
