package configuration

import "github.com/kercylan98/minotaur/configuration/raw"

// Scanner 用于扫描配置的扫描器接口
type Scanner interface {
	// StructScan 配置表结构扫描
	StructScan() (raw.Config, error)

	// DataScan 配置表数据扫描
	DataScan(fields []raw.Field) (any, error)
}
