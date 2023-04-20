package simple

const (
	defaultMaxTotal = 10
	defaultMaxIdle = 9
	defaultMinIdle = 1
)

// ResourcePoolConfig 资源池配置（需要建造的对象）
type ResourcePoolConfig struct {
	Name string
	MaxTotal int
	MaxIdle int
	MinIdle int
}

