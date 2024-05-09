package interpreter

// LogicExpression 逻辑表达式
type LogicExpression struct {
	Content string
}

func (exp *LogicExpression) Interpret() (bool, error) {
	for i, char := range exp.Content {
		switch char {
		case '|':
			expRes1, err := (&LogicExpression{Content: exp.Content[:i]}).Interpret()
			if err != nil {
				return false, err
			}
			expRes2, err := (&LogicExpression{Content: exp.Content[i+1:]}).Interpret()
			if err != nil {
				return false, err
			}
			return expRes1 || expRes2, nil
		case '&':
			expRes1, err := (&LogicExpression{Content: exp.Content[:i]}).Interpret()
			if err != nil {
				return false, err
			}
			expRes2, err := (&LogicExpression{Content: exp.Content[i+1:]}).Interpret()
			if err != nil {
				return false, err
			}
			return expRes1 && expRes2, nil
		}
	}
	return (&TerminalExpression{Content: exp.Content}).Interpret()
}
