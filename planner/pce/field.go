package pce

import (
	"reflect"
	"strings"
)

// Field 基本字段类型接口
type Field interface {
	// TypeName 字段类型名称
	TypeName() string
	// Zero 获取零值
	Zero() any
	// Parse 解析
	Parse(value string) any
}

// GetFieldGolangType 获取字段的 Golang 类型
func GetFieldGolangType(field Field) string {
	typeOf := reflect.TypeOf(field).Elem()
	kind := strings.ToLower(typeOf.Kind().String())
	return kind
}
