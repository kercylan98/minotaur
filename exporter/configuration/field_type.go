package configuration

import "fmt"

const (
	FieldTypeString = FieldType(iota)
	FieldTypeInt
	FieldTypeInt64
	FieldTypeFloat32
	FieldTypeFloat64
	FieldTypeByte
	FieldTypeBool
)

type FieldType byte

func (slf FieldType) String() string {
	switch slf {
	case FieldTypeString:
		return "string"
	case FieldTypeInt:
		return "int"
	case FieldTypeInt64:
		return "int64"
	case FieldTypeFloat32:
		return "float32"
	case FieldTypeFloat64:
		return "float64"
	case FieldTypeByte:
		return "byte"
	case FieldTypeBool:
		return "bool"
	}

	panic(fmt.Errorf("not support field type %v", byte(slf)))
}
