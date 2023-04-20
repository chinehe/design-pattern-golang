package method

import "fmt"

// YamlParser yaml解析器
type YamlParser struct {
}

func (parser *YamlParser) Parse(data []byte) {
	fmt.Printf("yaml parser:%v\n", string(data))
}

type YamlParserFactory struct {

}

func (factory *YamlParserFactory) Create() IRuleConfigParser {
	return &YamlParser{}
}

