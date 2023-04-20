package simple

import "testing"

func TestInterpreter(t *testing.T) {
	exp := "1=1&2=2|3=2"
	context := &ExpressionContext{Content: exp}
	t.Log(context.Interpret())
}
