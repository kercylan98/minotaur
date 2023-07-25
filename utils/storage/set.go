package storage

import (
	jsonIter "github.com/json-iterator/go"
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/str"
	"reflect"
)

var json = jsonIter.ConfigCompatibleWithStandardLibrary

func NewSet[PrimaryKey generic.Ordered, Body any](zero Body, getIndex func(data Body) PrimaryKey, options ...SetOption[PrimaryKey, Body]) *Set[PrimaryKey, Body] {
	set := &Set[PrimaryKey, Body]{
		zero:     zero,
		tf:       reflect.Indirect(reflect.ValueOf(zero)).Type(),
		getIndex: getIndex,
		bodyField: reflect.StructField{
			Name: DefaultBodyFieldName,
			Type: reflect.TypeOf(str.None),
		},
		items: make(map[PrimaryKey]*Data[Body]),
	}
	for _, option := range options {
		option(set)
	}
	return set
}

// Set 数据集合
type Set[PrimaryKey generic.Ordered, Body any] struct {
	storage       Storage[PrimaryKey, Body]
	zero          Body
	tf            reflect.Type
	getIndex      func(data Body) PrimaryKey
	bodyField     reflect.StructField
	indexes       []reflect.StructField
	getIndexValue map[string]func(data Body) any
	items         map[PrimaryKey]*Data[Body]
}

// New 创建一份新数据
//   - 这份数据不会被存储
func (slf *Set[PrimaryKey, Body]) New() *Data[Body] {
	var data = reflect.New(slf.tf).Interface().(Body)
	return newData(data)
}

// Get 通过主键获取数据
//   - 优先从内存中加载，如果数据不存在，则尝试从存储中加载
func (slf *Set[PrimaryKey, Body]) Get(index PrimaryKey) (*Data[Body], error) {
	if data, exist := slf.items[index]; exist {
		return data, nil
	}
	body, err := slf.storage.Load(index)
	if err != nil {
		return nil, err
	}
	data := newData(body)
	slf.items[index] = data
	return data, nil
}

// Set 设置数据
//   - 该方法会将数据存储到内存中
func (slf *Set[PrimaryKey, Body]) Set(data *Data[Body]) {
	slf.items[slf.getIndex(data.body)] = data
}

// Save 保存数据
//   - 该方法会将数据存储到存储器中
func (slf *Set[PrimaryKey, Body]) Save(data *Data[Body]) error {
	return slf.storage.Save(slf, slf.getIndex(data.body), data.body)
}

// Struct 将数据存储转换为结构体
func (slf *Set[PrimaryKey, Body]) Struct(data *Data[Body]) (any, error) {
	var fields = make([]reflect.StructField, 0, len(slf.indexes)+1)
	for _, field := range append(slf.indexes, slf.bodyField) {
		fields = append(fields, field)
	}
	instance := reflect.New(reflect.StructOf(fields))
	value := instance.Elem()
	for _, field := range slf.indexes {
		get, exist := slf.getIndexValue[field.Name]
		if !exist {
			continue
		}
		value.FieldByName(field.Name).Set(reflect.ValueOf(get(data.body)))
	}
	bytes, err := json.Marshal(data.body)
	if err != nil {
		return nil, err
	}

	value.FieldByName(slf.bodyField.Name).Set(reflect.ValueOf(string(bytes)))
	return value.Interface(), nil
}
