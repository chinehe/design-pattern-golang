package simple

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	NewLogger("Error").Println("Hello")
	NewLogger("Info").Println("Hello")
}
