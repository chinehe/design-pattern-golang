package class

import "testing"

func TestClassAdapter(t *testing.T) {
	adapterService := AdapterService{
		OriginService: &OriginService{},
	}
	adapterService.ServiceA()
	adapterService.ServiceB()
}
