package simple

import "fmt"

// YamlRuleConfigParser yaml规则配置解析器
type YamlRuleConfigParser struct {

}

func (parser *YamlRuleConfigParser) Parse(data []byte) {
	fmt.Printf("yaml rule parser:%v\n", string(data))
}

// YamlSystemConfigParser yaml系统配置解析器
type YamlSystemConfigParser struct {

}

func (parser *YamlSystemConfigParser) Parse(data []byte) {
	fmt.Printf("yaml system parser:%v\n", string(data))
}

// YamlParserFactory yaml解析器工厂
type YamlParserFactory struct {

}

func (factory *YamlParserFactory) Create(name string) ConfigParser{
	switch name {
	case "rule":
		return &YamlRuleConfigParser{}
	case "system":
		return &YamlSystemConfigParser{}
	default:
		return nil
	}
}

