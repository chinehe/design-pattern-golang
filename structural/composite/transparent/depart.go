package transparent

import "fmt"

// Depart 部门（树叶节点）
type Depart struct {
	Name string
}

func (d *Depart) Add(organization Organization) {
	fmt.Println("invalid operation")
}

func (d *Depart) Count() int {
	return 1
}



