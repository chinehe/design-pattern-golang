package simple

import (
	"testing"
)

func TestResourcePoolConfigBuilder(t *testing.T) {
	builder := NewResourcePoolConfigBuilder()
	builder.SetName("resource pool config")
	builder.SetMaxIdle(-1)
	builder.SetMaxIdle(10)
	builder.SetMinIdle(10)
	config := builder.Build()
	t.Logf("config:%+v",config)
}
