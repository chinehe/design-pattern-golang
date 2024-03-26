package simple

import (
	"fmt"
)

// ErrorLogger 错误日志打印器
type ErrorLogger struct {
}

func (l ErrorLogger) Println(msg string) {
	fmt.Printf("Error:%s\n", msg)
}
