package slice

import "fmt"

// Context 上下文
type Context struct {
	handlers   []LeaveApproval // 处理函数链
	index      uint8         // 当前处理函数

	Name string
	Reason string
	days int
}

func NewContext(name string, reason string, days int) *Context {
	return &Context{
		handlers: make([]LeaveApproval,1),
		Name:     name,
		Reason:   reason,
		days:     days,
	}
}



//
// Next
//  @Description: 执行处理函数链中的下一个处理函数
//  @receiver c
//
func (c *Context) Next() {
	c.index++
	for c.index < uint8(len(c.handlers)) {
		if c.handlers[c.index].LeaveApproval(c){
			return
		}
		c.index++
	}
	fmt.Println("审批被驳回")
}

//
// Use
//  @Description: 使用处理函数（将处理函数添加到处理函数链中）
//  @receiver c
//  @param handler 处理函数
//
func (c *Context) Use(handler LeaveApproval) {
	c.handlers = append(c.handlers, handler)
}
