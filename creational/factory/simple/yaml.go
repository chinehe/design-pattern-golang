package simple

import "fmt"

// YamlParser yaml解析器
type YamlParser struct {
}

func (parser *YamlParser) Parse(data []byte) {
	fmt.Printf("yaml parser:%v\n", string(data))
}
