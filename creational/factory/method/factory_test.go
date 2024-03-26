package simple

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	new(InfoLoggerFactory).NewLogger().Println("Hello")
	new(ErrorLoggerFactory).NewLogger().Println("Hello")
}
