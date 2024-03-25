package simple

import (
	"fmt"
)

// InfoLogger 信息日志打印器
type InfoLogger struct {
}

func (l InfoLogger) Println(msg string) {
	fmt.Printf("Info:%s\n", msg)
}
