package collection

type ComparisonHandler[V any] func(source, target V) bool

// InSlice 检查 v 是否被包含在 slice 中，当 handler 返回 true 时，表示 v 和 slice 中的某个元素相匹配
func InSlice[S ~[]V, V any](slice S, v V, handler ComparisonHandler[V]) bool {
	if len(slice) == 0 {
		return false
	}
	for _, value := range slice {
		if handler(v, value) {
			return true
		}
	}
	return false
}

// InComparableSlice 检查 v 是否被包含在 slice 中
func InComparableSlice[S ~[]V, V comparable](slice S, v V) bool {
	if slice == nil {
		return false
	}
	for _, value := range slice {
		if value == v {
			return true
		}
	}
	return false
}

// AllInSlice 检查 values 中的所有元素是否均被包含在 slice 中，当 handler 返回 true 时，表示 values 中的某个元素和 slice 中的某个元素相匹配
//   - 在所有 values 中的元素都被包含在 slice 中时，返回 true
//   - 当 values 长度为 0 或为 nil 时，将返回 true
func AllInSlice[S ~[]V, V any](slice S, values []V, handler ComparisonHandler[V]) bool {
	if len(slice) == 0 {
		return false
	}
	for _, value := range values {
		if !InSlice(slice, value, handler) {
			return false
		}
	}
	return true
}

// AllInComparableSlice 检查 values 中的所有元素是否均被包含在 slice 中
//   - 在所有 values 中的元素都被包含在 slice 中时，返回 true
//   - 当 values 长度为 0 或为 nil 时，将返回 true
func AllInComparableSlice[S ~[]V, V comparable](slice S, values []V) bool {
	if len(slice) == 0 {
		return false
	}
	for _, value := range values {
		if !InComparableSlice(slice, value) {
			return false
		}
	}
	return true
}

// AnyInSlice 检查 values 中的任意一个元素是否被包含在 slice 中，当 handler 返回 true 时，表示 value 中的某个元素和 slice 中的某个元素相匹配
//   - 当 values 中的任意一个元素被包含在 slice 中时，返回 true
func AnyInSlice[S ~[]V, V any](slice S, values []V, handler ComparisonHandler[V]) bool {
	if len(slice) == 0 {
		return false
	}
	for _, value := range values {
		if InSlice(slice, value, handler) {
			return true
		}
	}
	return false
}

// AnyInComparableSlice 检查 values 中的任意一个元素是否被包含在 slice 中
//   - 当 values 中的任意一个元素被包含在 slice 中时，返回 true
func AnyInComparableSlice[S ~[]V, V comparable](slice S, values []V) bool {
	if len(slice) == 0 {
		return false
	}
	for _, value := range values {
		if InComparableSlice(slice, value) {
			return true
		}
	}
	return false
}

// InSlices 通过将多个切片合并后检查 v 是否被包含在 slices 中，当 handler 返回 true 时，表示 v 和 slices 中的某个元素相匹配
//   - 当传入的 v 被包含在 slices 的任一成员中时，返回 true
func InSlices[S ~[]V, V any](slices []S, v V, handler ComparisonHandler[V]) bool {
	return InSlice(MergeSlices(slices...), v, handler)
}

// InComparableSlices 通过将多个切片合并后检查 v 是否被包含在 slices 中
//   - 当传入的 v 被包含在 slices 的任一成员中时，返回 true
func InComparableSlices[S ~[]V, V comparable](slices []S, v V) bool {
	return InComparableSlice(MergeSlices(slices...), v)
}

// AllInSlices 通过将多个切片合并后检查 values 中的所有元素是否被包含在 slices 中，当 handler 返回 true 时，表示 values 中的某个元素和 slices 中的某个元素相匹配
//   - 当 values 中的所有元素都被包含在 slices 的任一成员中时，返回 true
func AllInSlices[S ~[]V, V any](slices []S, values []V, handler ComparisonHandler[V]) bool {
	return AllInSlice(MergeSlices(slices...), values, handler)
}

// AllInComparableSlices 通过将多个切片合并后检查 values 中的所有元素是否被包含在 slices 中
//   - 当 values 中的所有元素都被包含在 slices 的任一成员中时，返回 true
func AllInComparableSlices[S ~[]V, V comparable](slices []S, values []V) bool {
	return AllInComparableSlice(MergeSlices(slices...), values)
}

// AnyInSlices 通过将多个切片合并后检查 values 中的任意一个元素是否被包含在 slices 中，当 handler 返回 true 时，表示 values 中的某个元素和 slices 中的某个元素相匹配
//   - 当 values 中的任意一个元素被包含在 slices 的任一成员中时，返回 true
func AnyInSlices[S ~[]V, V any](slices []S, values []V, handler ComparisonHandler[V]) bool {
	return AnyInSlice(MergeSlices(slices...), values, handler)
}

// AnyInComparableSlices 通过将多个切片合并后检查 values 中的任意一个元素是否被包含在 slices 中
//   - 当 values 中的任意一个元素被包含在 slices 的任一成员中时，返回 true
func AnyInComparableSlices[S ~[]V, V comparable](slices []S, values []V) bool {
	return AnyInComparableSlice(MergeSlices(slices...), values)
}

// InAllSlices 检查 v 是否被包含在 slices 的每一项元素中，当 handler 返回 true 时，表示 v 和 slices 中的某个元素相匹配
//   - 当 v 被包含在 slices 的每一项元素中时，返回 true
func InAllSlices[S ~[]V, V any](slices []S, v V, handler ComparisonHandler[V]) bool {
	if len(slices) == 0 {
		return false
	}
	for _, slice := range slices {
		if !InSlice(slice, v, handler) {
			return false
		}
	}
	return true
}

// InAllComparableSlices 检查 v 是否被包含在 slices 的每一项元素中
//   - 当 v 被包含在 slices 的每一项元素中时，返回 true
func InAllComparableSlices[S ~[]V, V comparable](slices []S, v V) bool {
	if len(slices) == 0 {
		return false
	}
	for _, slice := range slices {
		if !InComparableSlice(slice, v) {
			return false
		}
	}
	return true
}

// AnyInAllSlices 检查 slices 中的每一个元素是否均包含至少任意一个 values 中的元素，当 handler 返回 true 时，表示 value 中的某个元素和 slices 中的某个元素相匹配
//   - 当 slices 中的每一个元素均包含至少任意一个 values 中的元素时，返回 true
func AnyInAllSlices[S ~[]V, V any](slices []S, values []V, handler ComparisonHandler[V]) bool {
	if len(slices) == 0 {
		return false
	}
	for _, slice := range slices {
		if !AnyInSlice(slice, values, handler) {
			return false
		}
	}
	return true
}

// AnyInAllComparableSlices 检查 slices 中的每一个元素是否均包含至少任意一个 values 中的元素
//   - 当 slices 中的每一个元素均包含至少任意一个 values 中的元素时，返回 true
func AnyInAllComparableSlices[S ~[]V, V comparable](slices []S, values []V) bool {
	if len(slices) == 0 {
		return false
	}
	for _, slice := range slices {
		if !AnyInComparableSlice(slice, values) {
			return false
		}
	}
	return true
}

// KeyInMap 检查 m 中是否包含特定 key
func KeyInMap[M ~map[K]V, K comparable, V any](m M, key K) bool {
	_, ok := m[key]
	return ok
}

// ValueInMap 检查 m 中是否包含特定 value，当 handler 返回 true 时，表示 value 和 m 中的某个元素相匹配
func ValueInMap[M ~map[K]V, K comparable, V any](m M, value V, handler ComparisonHandler[V]) bool {
	if len(m) == 0 {
		return false
	}
	for _, v := range m {
		if handler(value, v) {
			return true
		}
	}
	return false
}

// AllKeyInMap 检查 m 中是否包含 keys 中所有的元素
func AllKeyInMap[M ~map[K]V, K comparable, V any](m M, keys ...K) bool {
	if len(m) < len(keys) {
		return false
	}
	for _, key := range keys {
		if !KeyInMap(m, key) {
			return false
		}
	}
	return true
}

// AllValueInMap 检查 m 中是否包含 values 中所有的元素，当 handler 返回 true 时，表示 values 中的某个元素和 m 中的某个元素相匹配
func AllValueInMap[M ~map[K]V, K comparable, V any](m M, values []V, handler ComparisonHandler[V]) bool {
	if len(m) == 0 {
		return false
	}
	for _, value := range values {
		if !ValueInMap(m, value, handler) {
			return false
		}
	}
	return true
}

// AnyKeyInMap 检查 m 中是否包含 keys 中任意一个元素
func AnyKeyInMap[M ~map[K]V, K comparable, V any](m M, keys ...K) bool {
	if len(m) == 0 {
		return false
	}
	for _, key := range keys {
		if KeyInMap(m, key) {
			return true
		}
	}
	return false
}

// AnyValueInMap 检查 m 中是否包含 values 中任意一个元素，当 handler 返回 true 时，表示 values 中的某个元素和 m 中的某个元素相匹配
func AnyValueInMap[M ~map[K]V, K comparable, V any](m M, values []V, handler ComparisonHandler[V]) bool {
	if len(m) == 0 {
		return false
	}
	for _, value := range values {
		if ValueInMap(m, value, handler) {
			return true
		}
	}
	return false
}

// AllKeyInMaps 检查 maps 中的每一个元素是否均包含 keys 中所有的元素
func AllKeyInMaps[M ~map[K]V, K comparable, V any](maps []M, keys ...K) bool {
	if len(maps) == 0 {
		return false
	}
	for _, m := range maps {
		if !AllKeyInMap(m, keys...) {
			return false
		}
	}
	return true
}

// AllValueInMaps 检查 maps 中的每一个元素是否均包含 value 中所有的元素，当 handler 返回 true 时，表示 value 中的某个元素和 maps 中的某个元素相匹配
func AllValueInMaps[M ~map[K]V, K comparable, V any](maps []M, values []V, handler ComparisonHandler[V]) bool {
	if len(maps) == 0 {
		return false
	}
	for _, m := range maps {
		if !AllValueInMap(m, values, handler) {
			return false
		}
	}
	return true
}

// AnyKeyInMaps 检查 keys 中的任意一个元素是否被包含在 maps 中的任意一个元素中
//   - 当 keys 中的任意一个元素被包含在 maps 中的任意一个元素中时，返回 true
func AnyKeyInMaps[M ~map[K]V, K comparable, V any](maps []M, keys ...K) bool {
	if len(maps) == 0 {
		return false
	}
	for _, m := range maps {
		if AnyKeyInMap(m, keys...) {
			return true
		}
	}
	return false
}

// AnyValueInMaps 检查 maps 中的任意一个元素是否包含 value 中的任意一个元素，当 handler 返回 true 时，表示 value 中的某个元素和 maps 中的某个元素相匹配
//   - 当 maps 中的任意一个元素包含 value 中的任意一个元素时，返回 true
func AnyValueInMaps[M ~map[K]V, K comparable, V any](maps []M, values []V, handler ComparisonHandler[V]) bool {
	if len(maps) == 0 {
		return false
	}
	for _, m := range maps {
		if !AnyValueInMap(m, values, handler) {
			return false
		}
	}
	return true
}

// KeyInAllMaps 检查 key 是否被包含在 maps 的每一个元素中
func KeyInAllMaps[M ~map[K]V, K comparable, V any](maps []M, key K) bool {
	if len(maps) == 0 {
		return false
	}
	for _, m := range maps {
		if !KeyInMap(m, key) {
			return false
		}
	}
	return true
}

// AnyKeyInAllMaps 检查 maps 中的每一个元素是否均包含 keys 中任意一个元素
//   - 当 maps 中的每一个元素均包含 keys 中任意一个元素时，返回 true
func AnyKeyInAllMaps[M ~map[K]V, K comparable, V any](maps []M, keys []K) bool {
	if len(maps) == 0 {
		return false
	}
	for _, m := range maps {
		if !AnyKeyInMap(m, keys...) {
			return false
		}
	}
	return true
}
