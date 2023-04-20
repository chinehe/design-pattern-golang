package simple

import (
	"fmt"
	"strconv"
)

// TerminalExpression 终端表达式
type TerminalExpression struct {
	Content string
}

func (exp *TerminalExpression) Interpret()(bool,error) {
	var paramA string
	var paramB string
	var operation int32
	for i, char := range exp.Content {
		if char=='>'||char=='<'||char=='=' {
			paramA = exp.Content[0:i]
			operation = char
			paramB = exp.Content[i+1:]
			break
		}
	}
	if paramA==""||paramB==""||operation==0 {
		return false,fmt.Errorf("invaild expression")
	}
	a, err := strconv.Atoi(paramA)
	if err != nil {
		return false,fmt.Errorf("invaild expression")
	}
	b, err := strconv.Atoi(paramB)
	if err != nil {
		return false,fmt.Errorf("invaild expression")
	}
	switch operation {
	case '>':
		return a > b,nil
	case '<':
		return a < b,nil
	case '=':
		return a == b,nil
	default:
		return false,fmt.Errorf("invaild expression")
	}
}



