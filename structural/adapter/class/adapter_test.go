package class

import "testing"

func TestClassAdapter(t *testing.T) {
	adapterService := AdapterService{}
	adapterService.ServiceA()
	adapterService.ServiceB()
}
