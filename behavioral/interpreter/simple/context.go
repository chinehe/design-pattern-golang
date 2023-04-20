package simple

// ExpressionContext 表达式上下文
type ExpressionContext struct {
	Content string
}

func (exp *ExpressionContext) Interpret() (bool, error) {
	return (&LogicExpression{Content: exp.Content}).Interpret()
}



