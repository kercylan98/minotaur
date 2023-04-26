package configuration

type Configuration interface {
	GetName() string
	GetFields() []Field
	AddField(field Field)
}
