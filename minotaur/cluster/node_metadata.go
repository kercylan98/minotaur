package cluster

import (
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"net"
	"strings"
)

type Metadata struct {
	ClusterName string `json:"cluster_name,omitempty"` // 集群名称
	Region      string `json:"region,omitempty"`       // 地域
	Zone        string `json:"zone,omitempty"`         // 可用区
	ShardId     uint16 `json:"shard_id,omitempty"`     // 分片ID
	Weight      uint64 `json:"weight,omitempty"`       // 权重
}

// Name 返回节点名称
func (m Metadata) Name(host string, port int) string {
	var parts []string
	if m.ClusterName != "" {
		parts = append(parts, m.ClusterName)
	}
	if m.Region != "" {
		parts = append(parts, m.Region)
	}
	if m.Zone != "" {
		parts = append(parts, m.Zone)
	}
	if m.ShardId != 0 {
		parts = append(parts, convert.Uint16ToString(m.ShardId))
	}
	parts = append(parts, net.JoinHostPort(host, convert.IntToString(port)))
	parts = append(parts, toolkit.Hostname())

	return strings.Join(parts, "-")
}

// NodeMeta 返回节点元数据
func (m Metadata) NodeMeta(limit int) []byte {
	return toolkit.MarshalJSON(m)
}
