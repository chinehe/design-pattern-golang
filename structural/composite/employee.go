package composite

import (
	"fmt"
)

// Employee 员工
type Employee struct {
	Name string
}

func (e *Employee) Display() string {
	return fmt.Sprintf("%s", e.Name)
}
