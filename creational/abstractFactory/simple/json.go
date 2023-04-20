package simple

import "fmt"

// JsonRuleConfigParser json的规则配置解析器
type JsonRuleConfigParser struct {

}

func (parser *JsonRuleConfigParser) Parse(data []byte) {
	fmt.Printf("json rule parser:%v\n", string(data))
}

// JsonSystemConfigParser json的系统配置解析器
type JsonSystemConfigParser struct {

}

func (parser *JsonSystemConfigParser) Parse(data []byte) {
	fmt.Printf("json system parser:%v\n", string(data))
}

// JsonParserFactory json解析器工厂
type JsonParserFactory struct {

}

func (factory *JsonParserFactory) Create(name string) ConfigParser{
	switch name {
	case "rule":
		return &JsonRuleConfigParser{}
	case "system":
		return &JsonSystemConfigParser{}
	default:
		return nil
	}
}

