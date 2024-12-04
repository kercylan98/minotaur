package datasheet

type TypeMarshaler interface {
	Marshal(t Type) (string, error)
}
