package configuration

// Configuration 配置
type Configuration interface {
	GetName() string
	GetFields() []Field
	AddField(field Field)
}
