package simple

import (
	"testing"
)

func TestNewIRuleConfigParser(t *testing.T) {
	jsonParser := NewIRuleConfigParser("json")
	yamlParser := NewIRuleConfigParser("yaml")
	jsonParser.Parse([]byte("json"))
	yamlParser.Parse([]byte("yaml"))
}
