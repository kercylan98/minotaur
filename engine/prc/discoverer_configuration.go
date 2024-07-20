package prc

import "github.com/kercylan98/minotaur/toolkit/random"

// newDiscovererConfiguration 创建一个新的 DiscovererConfiguration
func newDiscovererConfiguration() *DiscovererConfiguration {
	return &DiscovererConfiguration{
		name: random.HostName(),
	}
}

// DiscovererConfiguration 是用于配置 Discoverer 的数据结构
type DiscovererConfiguration struct {
	advertiseAddr string            // 广播地址
	advertisePort int               // 广播端口
	name          string            // 名称
	userMetadata  map[string]string // 用户元数据
	joinNodes     []PhysicalAddress // 默认加入的节点
}

// WithJoinNodes 设置默认加入的节点。
func (c *DiscovererConfiguration) WithJoinNodes(nodes ...PhysicalAddress) *DiscovererConfiguration {
	c.joinNodes = append(c.joinNodes, nodes...)
	return c
}

// WithUserMetadataKey 设置用户元数据。
func (c *DiscovererConfiguration) WithUserMetadataKey(key, value string) *DiscovererConfiguration {
	if c.userMetadata == nil {
		c.userMetadata = make(map[string]string)
	}
	c.userMetadata[key] = value
	return c
}

// WithUserMetadata 设置用户元数据，重复的设置将会导致已有 key 被覆盖。
func (c *DiscovererConfiguration) WithUserMetadata(metadata map[string]string) *DiscovererConfiguration {
	if c.userMetadata == nil {
		c.userMetadata = make(map[string]string)
	}
	for k, v := range metadata {
		c.userMetadata[k] = v
	}
	return c
}

// WithName 指定资源控制器在网络中唯一的名称。
func (c *DiscovererConfiguration) WithName(name string) *DiscovererConfiguration {
	c.name = name
	return c
}

// WithAdvertise 设置广播地址和端口。
func (c *DiscovererConfiguration) WithAdvertise(addr string, port int) *DiscovererConfiguration {
	c.advertiseAddr = addr
	c.advertisePort = port
	return c
}
