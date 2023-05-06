package configuration

const (
	FieldTypeString = FieldType(iota)
	FieldTypeInt
	FieldTypeInt8
	FieldTypeInt16
	FieldTypeInt32
	FieldTypeInt64
	FieldTypeUint
	FieldTypeUint8
	FieldTypeUint16
	FieldTypeUint32
	FieldTypeUint64
	FieldTypeFloat32
	FieldTypeFloat64
	FieldTypeByte
	FieldTypeBool
	FieldTypeRune
	FieldTypeSliceString
	FieldTypeSliceInt
	FieldTypeSliceInt8
	FieldTypeSliceInt16
	FieldTypeSliceInt32
	FieldTypeSliceInt64
	FieldTypeSliceUint
	FieldTypeSliceUint8
	FieldTypeSliceUint16
	FieldTypeSliceUint32
	FieldTypeSliceUint64
	FieldTypeSliceFloat32
	FieldTypeSliceFloat64
	FieldTypeSliceByte
	FieldTypeSliceBool
	FieldTypeSliceRune
)

type FieldType byte
