package storage

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/str"
	"reflect"
	"strings"
)

const (
	DefaultBodyFieldName = "Data"
)

type SetOption[PrimaryKey generic.Ordered, Body any] func(set *Set[PrimaryKey, Body])

// WithIndex 添加一个索引
//   - 索引将会在数据结构体中创建一个字段，这个字段必须可以在 Body 内部找到，用于对查找功能的拓展
func WithIndex[PrimaryKey generic.Ordered, Index generic.Ordered, Body any](name string, getValue func(data Body) Index) SetOption[PrimaryKey, Body] {
	return func(set *Set[PrimaryKey, Body]) {
		WithTagIndex[PrimaryKey, Index, Body](name, nil, getValue)(set)
	}
}

// WithTagIndex 添加一个带 tag 的索引
//   - 同 WithIndex，但是可以自定义索引的 tag
func WithTagIndex[PrimaryKey generic.Ordered, Index generic.Ordered, Body any](name string, tags []string, getValue func(data Body) Index) SetOption[PrimaryKey, Body] {
	return func(set *Set[PrimaryKey, Body]) {
		value := getValue(set.zero)
		upperName := str.FirstUpper(name)
		if set.getIndexValue == nil {
			set.getIndexValue = map[string]func(data Body) any{}
		}
		set.getIndexValue[upperName] = func(data Body) any {
			return getValue(data)
		}
		var tag string
		if len(tags) > 0 {
			tag = strings.Join(tags, " ")
		}
		set.indexes = append(set.indexes, reflect.StructField{
			Name: upperName,
			Type: reflect.TypeOf(value),
			Tag:  reflect.StructTag(tag),
		})
	}
}

// WithBodyName 设置 Body 字段名称
//   - 默认字段名称为 DefaultBodyFieldName
func WithBodyName[PrimaryKey generic.Ordered, Body any](name string) SetOption[PrimaryKey, Body] {
	return func(set *Set[PrimaryKey, Body]) {
		if len(name) == 0 {
			return
		}
		set.bodyField.Name = str.FirstUpper(name)
	}
}

// WithBodyTag 设置 Body 字段标签
//   - 如果有多个标签，将会以空格分隔，例如：`json:"data" yaml:"data"`
func WithBodyTag[PrimaryKey generic.Ordered, Body any](tags ...string) SetOption[PrimaryKey, Body] {
	return func(set *Set[PrimaryKey, Body]) {
		set.bodyField.Tag = reflect.StructTag(strings.Join(tags, " "))
	}
}
