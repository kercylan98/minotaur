package configuration

// Field 配置字段
type Field interface {
	GetID() int
	GetName() string
	GetType() FieldType
	IsIndex() bool
}
