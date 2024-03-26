package simple

// LoggerFactory 抽象工厂
type LoggerFactory interface {
	NewLogger() Logger
}

// InfoLoggerFactory 信息日志工厂
type InfoLoggerFactory struct {
}

func (i *InfoLoggerFactory) NewLogger() Logger {
	return &InfoLogger{}
}

// ErrorLoggerFactory 错误日志工厂
type ErrorLoggerFactory struct {
}

func (i *ErrorLoggerFactory) NewLogger() Logger {
	return &ErrorLogger{}
}
