package abstractFactory

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	new(InfoFactory).NewLogger().Println("Hello")
	new(InfoFactory).NewPrompter().Prompt("Hello")
	new(ErrorFactory).NewLogger().Println("Hello")
	new(ErrorFactory).NewPrompter().Prompt("Hello")
}
