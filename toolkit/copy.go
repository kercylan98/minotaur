package toolkit

import (
	"reflect"
)

func DeepCopy(v reflect.Value) reflect.Value {
	if !v.IsValid() {
		return v
	}

	switch v.Kind() {
	case reflect.Ptr:
		if v.IsZero() {
			return v
		}
		ptr := reflect.New(v.Elem().Type())
		ptr.Elem().Set(DeepCopy(v.Elem()))
		return ptr
	case reflect.Interface:
		return DeepCopy(v.Elem())
	case reflect.Struct:
		structCopy := reflect.New(v.Type()).Elem()
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			//if !field.CanSet() {
			//	continue // 跳过不可导出的字段
			//}
			copyValue := DeepCopy(field)
			structCopy.Field(i).Set(copyValue)
		}
		return structCopy
	case reflect.Map:
		mapping := reflect.MakeMapWithSize(v.Type(), v.Len())
		for _, key := range v.MapKeys() {
			copyValue := DeepCopy(v.MapIndex(key))
			mapping.SetMapIndex(key, copyValue)
		}
		return mapping
	case reflect.Slice:
		slice := reflect.MakeSlice(v.Type(), v.Len(), v.Cap())
		for i := 0; i < v.Len(); i++ {
			slice.Index(i).Set(DeepCopy(v.Index(i)))
		}
		return slice
	case reflect.Array:
		array := reflect.New(v.Type()).Elem()
		for i := 0; i < v.Len(); i++ {
			array.Index(i).Set(DeepCopy(v.Index(i)))
		}
		return array
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128,
		reflect.String, reflect.Bool:
		return v
	default:
		return v
	}
}
