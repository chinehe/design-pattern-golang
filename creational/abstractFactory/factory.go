package abstractFactory

// Factory 抽象工厂
type Factory interface {
	NewLogger() Logger
	NewPrompter() Prompter
}

// InfoFactory 信息工厂
type InfoFactory struct {
}

func (i *InfoFactory) NewLogger() Logger {
	return &InfoLogger{}
}

func (i *InfoFactory) NewPrompter() Prompter {
	return &InfoPrompter{}
}

// ErrorFactory 错误工厂
type ErrorFactory struct {
}

func (i *ErrorFactory) NewLogger() Logger {
	return &ErrorLogger{}
}

func (i *ErrorFactory) NewPrompter() Prompter {
	return &ErrorPrompter{}
}
