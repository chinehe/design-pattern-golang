package composite

import (
	"fmt"
	"strings"
)

// Department 部门
type Department struct {
	Name      string
	Employees []*Employee
}

func (d *Department) Add(e *Employee) {
	d.Employees = append(d.Employees, e)
}

func (d *Department) Display() string {
	builder := strings.Builder{}
	builder.WriteString(d.Name)
	builder.WriteString(":\n")
	for i, employee := range d.Employees {
		builder.WriteString(fmt.Sprintf("%d-%s\n", i+1, employee.Display()))
	}
	return builder.String()
}
