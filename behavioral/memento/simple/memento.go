package simple

// FileMementoStack 文件备忘录栈（记录者）
type FileMementoStack struct {
	contents []string
}

func NewFileMementoStack() *FileMementoStack {
	return &FileMementoStack{contents: []string{}}
}

func (stack *FileMementoStack) Push(content string)  {
	stack.contents = append(stack.contents,content)
}

func (stack *FileMementoStack) Pop() (content string) {
	if len(stack.contents)<=0 {
		return ""
	}
	content = stack.contents[len(stack.contents)-1]
	stack.contents = stack.contents[:len(stack.contents)-1]
	return
}