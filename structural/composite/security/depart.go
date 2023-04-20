package transparent

// Depart 部门（树叶节点）
type Depart struct {
	Name string
}

func (d *Depart) Count() int {
	return 1
}



