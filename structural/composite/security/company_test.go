package transparent

import (
	"fmt"
	"testing"
)

func TestSecurityComposite(t *testing.T) {
	company := NewCompany("XX集团")

	wuxiCompany := NewCompany("无锡研发中心")
	shanghaiCompany := NewCompany("上海总部")

	wuxiCompany.Add(&Depart{"测试部"})
	wuxiCompany.Add(&Depart{"开发部"})
	shanghaiCompany.Add(&Depart{Name: "项目管理部"})
	shanghaiCompany.Add(&Depart{Name: "人力资源部"})

	company.Add(wuxiCompany)
	company.Add(shanghaiCompany)
	fmt.Println(company.Count())
}
