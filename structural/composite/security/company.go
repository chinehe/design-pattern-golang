package transparent

// Company 公司（树枝节点）
type Company struct {
	Name string
	SubOrganizations []Organization
}

func NewCompany(name string) *Company {
	company := new(Company)
	company.SubOrganizations = make([]Organization,0)
	return company
}

func (c *Company) Add(organization Organization) {
	c.SubOrganizations = append(c.SubOrganizations,organization)
}

func (c *Company) Count() int {
	count := 1
	for _,s := range c.SubOrganizations {
		count += s.Count()
	}
	return count
}




