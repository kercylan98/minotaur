package table

type Type = any

type BasicType struct {
	Type string
}

type ArrayType struct {
	ElementType Type
}

type StructType struct {
	Name   string
	Fields []*StructField
}

type StructField struct {
	Name string
	Type Type
}

type StructFieldWrapper struct {
	StructField
	Desc string
}
