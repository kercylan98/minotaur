package table

// FieldScanner 字段扫描器
type FieldScanner interface {
	Next() Field
}
