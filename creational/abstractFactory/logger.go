package abstractFactory

import (
	"fmt"
)

// Logger 日志器
type Logger interface {
	Println(msg string)
}

// ErrorLogger 错误日志打印器
type ErrorLogger struct {
}

func (l ErrorLogger) Println(msg string) {
	fmt.Printf("Error:%s\n", msg)
}

// InfoLogger 信息日志打印器
type InfoLogger struct {
}

func (l InfoLogger) Println(msg string) {
	fmt.Printf("Info:%s\n", msg)
}
