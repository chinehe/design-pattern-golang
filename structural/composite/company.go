package composite

import (
	"fmt"
	"strings"
)

// Company 公司
type Company struct {
	Name        string
	Departments []*Department
}

func (c *Company) Add(e *Department) {
	c.Departments = append(c.Departments, e)
}

func (c *Company) Display() string {
	builder := strings.Builder{}
	builder.WriteString(c.Name)
	builder.WriteString(":\n")
	for i, department := range c.Departments {
		builder.WriteString(fmt.Sprintf("%d-%s\n", i+1, department.Display()))
	}
	return builder.String()
}
