package configuration

type Field interface {
	GetID() int
	GetName() string
	GetType() FieldType
}
