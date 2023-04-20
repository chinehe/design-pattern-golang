package method

import "fmt"

// JsonParser json解析器
type JsonParser struct {
}

func (parser *JsonParser) Parse(data []byte) {
	fmt.Printf("json parser:%v\n", string(data))
}

type JsonParserFactory struct {

}

func (factory *JsonParserFactory) Create() IRuleConfigParser {
	return &JsonParser{}
}
