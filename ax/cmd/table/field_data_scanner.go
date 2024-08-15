package table

// FieldDataScanner 字段数据扫描器
type FieldDataScanner interface {
	Next() string
}
