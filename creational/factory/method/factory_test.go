package method

import (
	"testing"
)

func TestNewIRuleConfigParser(t *testing.T) {
	jsonParserFactory := NewIRuleConfigParserFactory("json")
	yamlParserFactory := NewIRuleConfigParserFactory("yaml")
	jsonParserFactory.Create().Parse([]byte("json"))
	yamlParserFactory.Create().Parse([]byte("yaml"))
}
