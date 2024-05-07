package composite

import (
	"fmt"
	"testing"
)

func TestComponent(t *testing.T) {
	e1 := &Employee{"E1"}
	e2 := &Employee{"E2"}
	e3 := &Employee{"E3"}
	e4 := &Employee{"E4"}
	e5 := &Employee{"E5"}
	e6 := &Employee{"E6"}
	e7 := &Employee{"E7"}
	e8 := &Employee{"E8"}
	e9 := &Employee{"E9"}

	d1 := &Department{Name: "研发部"}
	d1.Add(e1)
	d1.Add(e2)
	d1.Add(e3)

	d2 := &Department{Name: "产品部"}
	d2.Add(e4)
	d2.Add(e5)
	d2.Add(e6)

	d3 := &Department{Name: "项目部"}
	d3.Add(e7)
	d3.Add(e8)
	d3.Add(e9)

	c := Company{
		Name: "大苹果有限公司",
	}
	c.Add(d1)
	c.Add(d2)
	c.Add(d3)

	fmt.Println(c.Display())
}
