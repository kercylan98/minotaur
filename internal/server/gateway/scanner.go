package gateway

import "time"

// Scanner 端点扫描器
type Scanner interface {
	// GetEndpoints 获取端点列表
	GetEndpoints() ([]*Endpoint, error)
	// GetInterval 获取扫描间隔
	GetInterval() time.Duration
}
