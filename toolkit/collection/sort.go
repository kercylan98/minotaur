package collection

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"github.com/kercylan98/minotaur/toolkit/random"
	"sort"
)

// DescBy 返回降序比较结果
func DescBy[Sort constraints.Ordered](a, b Sort) bool {
	return a > b
}

// AscBy 返回升序比较结果
func AscBy[Sort constraints.Ordered](a, b Sort) bool {
	return a < b
}

// Desc 对切片进行降序排序
func Desc[S ~[]V, V any, Sort constraints.Ordered](slice *S, getter func(index int) Sort) {
	sort.Slice(*slice, func(i, j int) bool {
		return getter(i) > getter(j)
	})
}

// DescByClone 对切片进行降序排序，返回排序后的切片
func DescByClone[S ~[]V, V any, Sort constraints.Ordered](slice S, getter func(index int) Sort) S {
	result := CloneSlice(slice)
	Desc(&result, getter)
	return result

}

// Asc 对切片进行升序排序
func Asc[S ~[]V, V any, Sort constraints.Ordered](slice *S, getter func(index int) Sort) {
	sort.Slice(*slice, func(i, j int) bool {
		return getter(i) < getter(j)
	})
}

// AscByClone 对切片进行升序排序，返回排序后的切片
func AscByClone[S ~[]V, V any, Sort constraints.Ordered](slice S, getter func(index int) Sort) S {
	result := CloneSlice(slice)
	Asc(&result, getter)
	return result
}

// Shuffle 对切片进行随机排序
func Shuffle[S ~[]V, V any](slice *S) {
	sort.Slice(*slice, func(i, j int) bool {
		return random.Bool()
	})
}

// ShuffleByClone 对切片进行随机排序，返回排序后的切片
func ShuffleByClone[S ~[]V, V any](slice S) S {
	result := CloneSlice(slice)
	Shuffle(&result)
	return result
}
