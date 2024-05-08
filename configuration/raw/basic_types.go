package raw

type FieldType = string

const (
	FieldTypeInt     FieldType = "int"
	FieldTypeInt8    FieldType = "int8"
	FieldTypeInt16   FieldType = "int16"
	FieldTypeInt32   FieldType = "int32"
	FieldTypeInt64   FieldType = "int64"
	FieldTypeUint    FieldType = "uint"
	FieldTypeUint8   FieldType = "uint8"
	FieldTypeUint16  FieldType = "uint16"
	FieldTypeUint32  FieldType = "uint32"
	FieldTypeUint64  FieldType = "uint64"
	FieldTypeFloat32 FieldType = "float32"
	FieldTypeFloat64 FieldType = "float64"
	FieldTypeString  FieldType = "string"
	FieldTypeBool    FieldType = "bool"
)

var basicTypes = map[FieldType]bool{
	FieldTypeInt:     true,
	FieldTypeInt8:    true,
	FieldTypeInt16:   true,
	FieldTypeInt32:   true,
	FieldTypeInt64:   true,
	FieldTypeUint:    true,
	FieldTypeUint8:   true,
	FieldTypeUint16:  true,
	FieldTypeUint32:  true,
	FieldTypeUint64:  true,
	FieldTypeFloat32: true,
	FieldTypeFloat64: true,
	FieldTypeString:  true,
	FieldTypeBool:    true,
}

// IsBasicType 判断是否为基础类型
func IsBasicType(t FieldType) bool {
	return basicTypes[t]
}
