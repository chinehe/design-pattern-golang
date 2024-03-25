package simple

// NewLogger 创建日志器的简单工厂
func NewLogger(level string) Logger {
	switch level {
	case "Error":
		return &ErrorLogger{}
	case "Info":
		return &InfoLogger{}
	default:
		return nil
	}
}
