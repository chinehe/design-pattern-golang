package abstractFactory

import (
	"fmt"
)

// Prompter 提示器
type Prompter interface {
	Prompt(msg string)
}

// ErrorPrompter 错误日志打印器
type ErrorPrompter struct {
}

func (l *ErrorPrompter) Prompt(msg string) {
	fmt.Printf("EEEEEEEEEEEEEEEEEError:%s\n", msg)
}

// InfoPrompter 信息日志打印器
type InfoPrompter struct {
}

func (l *InfoPrompter) Prompt(msg string) {
	fmt.Printf("IIIIIIIIIIIIIIIIIIInfo:%s\n", msg)
}
