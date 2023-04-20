package simple

// NewIRuleConfigParser 工厂方法
func NewIRuleConfigParser(parserName string) IRuleConfigParser {
	switch parserName {
	case "json":
		return &JsonParser{}
	case "yaml":
		return &YamlParser{}
	default:
		return nil
	}
}
