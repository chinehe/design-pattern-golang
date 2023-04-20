package slice

import "testing"

func TestResponsibilityChain(t *testing.T) {
	context := NewContext("ChineHe","123",8)
	context.Use(new(ClassMaster))
	context.Use(new(President))
	context.Use(new(Principal))
	context.Next()
}
