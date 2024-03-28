package builder

// resourcePoolConfigBuilder 资源池配置建造器
type resourcePoolConfigBuilder struct {
	resourcePoolConfig *ResourcePoolConfig
}

func (builder *resourcePoolConfigBuilder) SetName(name string) {
	builder.resourcePoolConfig.Name = name
}

func (builder *resourcePoolConfigBuilder) SetMaxTotal(maxTotal int) {
	if maxTotal < 0 {
		maxTotal = defaultMaxTotal
	}
	builder.resourcePoolConfig.MaxTotal = maxTotal
}

func (builder *resourcePoolConfigBuilder) SetMaxIdle(maxIdle int) {
	if maxIdle < 0 {
		maxIdle = defaultMaxIdle
	}
	if maxIdle < builder.resourcePoolConfig.MinIdle {
		builder.resourcePoolConfig.MinIdle = maxIdle
	}
	builder.resourcePoolConfig.MaxIdle = maxIdle
}

func (builder *resourcePoolConfigBuilder) SetMinIdle(minIdle int) {
	if minIdle < 0 {
		minIdle = defaultMinIdle
	}
	if minIdle > builder.resourcePoolConfig.MaxIdle {
		builder.resourcePoolConfig.MaxIdle = minIdle
	}
	builder.resourcePoolConfig.MaxIdle = minIdle
}

// Build 建造方法
func (builder *resourcePoolConfigBuilder) Build() *ResourcePoolConfig {
	return builder.resourcePoolConfig
}

func NewResourcePoolConfigBuilder() *resourcePoolConfigBuilder {
	builder := &resourcePoolConfigBuilder{
		resourcePoolConfig: &ResourcePoolConfig{
			MaxTotal: defaultMaxTotal,
			MaxIdle:  defaultMaxIdle,
			MinIdle:  defaultMinIdle,
		},
	}
	return builder
}
