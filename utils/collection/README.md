# Collection

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

collection 定义了各种对于集合操作有用的各种函数


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[CloneSlice](#CloneSlice)|通过创建一个新切片并将 slice 的元素复制到新切片的方式来克隆切片
|[CloneMap](#CloneMap)|通过创建一个新 map 并将 m 的元素复制到新 map 的方式来克隆 map
|[CloneSliceN](#CloneSliceN)|通过创建一个新切片并将 slice 的元素复制到新切片的方式来克隆切片为 n 个切片
|[CloneMapN](#CloneMapN)|通过创建一个新 map 并将 m 的元素复制到新 map 的方式来克隆 map 为 n 个 map
|[CloneSlices](#CloneSlices)|对 slices 中的每一项元素进行克隆，最终返回一个新的二维切片
|[CloneMaps](#CloneMaps)|对 maps 中的每一项元素进行克隆，最终返回一个新的 map 切片
|[InSlice](#InSlice)|检查 v 是否被包含在 slice 中，当 handler 返回 true 时，表示 v 和 slice 中的某个元素相匹配
|[InComparableSlice](#InComparableSlice)|检查 v 是否被包含在 slice 中
|[AllInSlice](#AllInSlice)|检查 values 中的所有元素是否均被包含在 slice 中，当 handler 返回 true 时，表示 values 中的某个元素和 slice 中的某个元素相匹配
|[AllInComparableSlice](#AllInComparableSlice)|检查 values 中的所有元素是否均被包含在 slice 中
|[AnyInSlice](#AnyInSlice)|检查 values 中的任意一个元素是否被包含在 slice 中，当 handler 返回 true 时，表示 value 中的某个元素和 slice 中的某个元素相匹配
|[AnyInComparableSlice](#AnyInComparableSlice)|检查 values 中的任意一个元素是否被包含在 slice 中
|[InSlices](#InSlices)|通过将多个切片合并后检查 v 是否被包含在 slices 中，当 handler 返回 true 时，表示 v 和 slices 中的某个元素相匹配
|[InComparableSlices](#InComparableSlices)|通过将多个切片合并后检查 v 是否被包含在 slices 中
|[AllInSlices](#AllInSlices)|通过将多个切片合并后检查 values 中的所有元素是否被包含在 slices 中，当 handler 返回 true 时，表示 values 中的某个元素和 slices 中的某个元素相匹配
|[AllInComparableSlices](#AllInComparableSlices)|通过将多个切片合并后检查 values 中的所有元素是否被包含在 slices 中
|[AnyInSlices](#AnyInSlices)|通过将多个切片合并后检查 values 中的任意一个元素是否被包含在 slices 中，当 handler 返回 true 时，表示 values 中的某个元素和 slices 中的某个元素相匹配
|[AnyInComparableSlices](#AnyInComparableSlices)|通过将多个切片合并后检查 values 中的任意一个元素是否被包含在 slices 中
|[InAllSlices](#InAllSlices)|检查 v 是否被包含在 slices 的每一项元素中，当 handler 返回 true 时，表示 v 和 slices 中的某个元素相匹配
|[InAllComparableSlices](#InAllComparableSlices)|检查 v 是否被包含在 slices 的每一项元素中
|[AnyInAllSlices](#AnyInAllSlices)|检查 slices 中的每一个元素是否均包含至少任意一个 values 中的元素，当 handler 返回 true 时，表示 value 中的某个元素和 slices 中的某个元素相匹配
|[AnyInAllComparableSlices](#AnyInAllComparableSlices)|检查 slices 中的每一个元素是否均包含至少任意一个 values 中的元素
|[KeyInMap](#KeyInMap)|检查 m 中是否包含特定 key
|[ValueInMap](#ValueInMap)|检查 m 中是否包含特定 value，当 handler 返回 true 时，表示 value 和 m 中的某个元素相匹配
|[AllKeyInMap](#AllKeyInMap)|检查 m 中是否包含 keys 中所有的元素
|[AllValueInMap](#AllValueInMap)|检查 m 中是否包含 values 中所有的元素，当 handler 返回 true 时，表示 values 中的某个元素和 m 中的某个元素相匹配
|[AnyKeyInMap](#AnyKeyInMap)|检查 m 中是否包含 keys 中任意一个元素
|[AnyValueInMap](#AnyValueInMap)|检查 m 中是否包含 values 中任意一个元素，当 handler 返回 true 时，表示 values 中的某个元素和 m 中的某个元素相匹配
|[AllKeyInMaps](#AllKeyInMaps)|检查 maps 中的每一个元素是否均包含 keys 中所有的元素
|[AllValueInMaps](#AllValueInMaps)|检查 maps 中的每一个元素是否均包含 value 中所有的元素，当 handler 返回 true 时，表示 value 中的某个元素和 maps 中的某个元素相匹配
|[AnyKeyInMaps](#AnyKeyInMaps)|检查 keys 中的任意一个元素是否被包含在 maps 中的任意一个元素中
|[AnyValueInMaps](#AnyValueInMaps)|检查 maps 中的任意一个元素是否包含 value 中的任意一个元素，当 handler 返回 true 时，表示 value 中的某个元素和 maps 中的某个元素相匹配
|[KeyInAllMaps](#KeyInAllMaps)|检查 key 是否被包含在 maps 的每一个元素中
|[AnyKeyInAllMaps](#AnyKeyInAllMaps)|检查 maps 中的每一个元素是否均包含 keys 中任意一个元素
|[ConvertSliceToAny](#ConvertSliceToAny)|将切片转换为任意类型的切片
|[ConvertSliceToIndexMap](#ConvertSliceToIndexMap)|将切片转换为索引为键的映射
|[ConvertSliceToIndexOnlyMap](#ConvertSliceToIndexOnlyMap)|将切片转换为索引为键的映射
|[ConvertSliceToMap](#ConvertSliceToMap)|将切片转换为值为键的映射
|[ConvertSliceToBoolMap](#ConvertSliceToBoolMap)|将切片转换为值为键的映射
|[ConvertMapKeysToSlice](#ConvertMapKeysToSlice)|将映射的键转换为切片
|[ConvertMapValuesToSlice](#ConvertMapValuesToSlice)|将映射的值转换为切片
|[InvertMap](#InvertMap)|将映射的键和值互换
|[ConvertMapValuesToBool](#ConvertMapValuesToBool)|将映射的值转换为布尔值
|[ReverseSlice](#ReverseSlice)|将切片反转
|[ClearSlice](#ClearSlice)|清空切片
|[ClearMap](#ClearMap)|清空 map
|[DropSliceByIndices](#DropSliceByIndices)|删除切片中特定索引的元素
|[DropSliceByCondition](#DropSliceByCondition)|删除切片中符合条件的元素
|[DropSliceOverlappingElements](#DropSliceOverlappingElements)|删除切片中与另一个切片重叠的元素
|[DeduplicateSliceInPlace](#DeduplicateSliceInPlace)|去除切片中的重复元素
|[DeduplicateSlice](#DeduplicateSlice)|去除切片中的重复元素，返回新切片
|[DeduplicateSliceInPlaceWithCompare](#DeduplicateSliceInPlaceWithCompare)|去除切片中的重复元素，使用自定义的比较函数
|[DeduplicateSliceWithCompare](#DeduplicateSliceWithCompare)|去除切片中的重复元素，使用自定义的比较函数，返回新的切片
|[FilterOutByIndices](#FilterOutByIndices)|过滤切片中特定索引的元素，返回过滤后的切片
|[FilterOutByCondition](#FilterOutByCondition)|过滤切片中符合条件的元素，返回过滤后的切片
|[FilterOutByKey](#FilterOutByKey)|过滤 map 中特定的 key，返回过滤后的 map
|[FilterOutByValue](#FilterOutByValue)|过滤 map 中特定的 value，返回过滤后的 map
|[FilterOutByKeys](#FilterOutByKeys)|过滤 map 中多个 key，返回过滤后的 map
|[FilterOutByValues](#FilterOutByValues)|过滤 map 中多个 values，返回过滤后的 map
|[FilterOutByMap](#FilterOutByMap)|过滤 map 中符合条件的元素，返回过滤后的 map
|[FindLoopedNextInSlice](#FindLoopedNextInSlice)|返回 i 的下一个数组成员，当 i 达到数组长度时从 0 开始
|[FindLoopedPrevInSlice](#FindLoopedPrevInSlice)|返回 i 的上一个数组成员，当 i 为 0 时从数组末尾开始
|[FindCombinationsInSliceByRange](#FindCombinationsInSliceByRange)|获取给定数组的所有组合，且每个组合的成员数量限制在指定范围内
|[FindFirstOrDefaultInSlice](#FindFirstOrDefaultInSlice)|判断切片中是否存在元素，返回第一个元素，不存在则返回默认值
|[FindOrDefaultInSlice](#FindOrDefaultInSlice)|判断切片中是否存在某个元素，返回第一个匹配的索引和元素，不存在则返回默认值
|[FindOrDefaultInComparableSlice](#FindOrDefaultInComparableSlice)|判断切片中是否存在某个元素，返回第一个匹配的索引和元素，不存在则返回默认值
|[FindInSlice](#FindInSlice)|判断切片中是否存在某个元素，返回第一个匹配的索引和元素，不存在则索引返回 -1
|[FindIndexInSlice](#FindIndexInSlice)|判断切片中是否存在某个元素，返回第一个匹配的索引，不存在则索引返回 -1
|[FindInComparableSlice](#FindInComparableSlice)|判断切片中是否存在某个元素，返回第一个匹配的索引和元素，不存在则索引返回 -1
|[FindIndexInComparableSlice](#FindIndexInComparableSlice)|判断切片中是否存在某个元素，返回第一个匹配的索引，不存在则索引返回 -1
|[FindMinimumInComparableSlice](#FindMinimumInComparableSlice)|获取切片中的最小值
|[FindMinimumInSlice](#FindMinimumInSlice)|获取切片中的最小值
|[FindMaximumInComparableSlice](#FindMaximumInComparableSlice)|获取切片中的最大值
|[FindMaximumInSlice](#FindMaximumInSlice)|获取切片中的最大值
|[FindMin2MaxInComparableSlice](#FindMin2MaxInComparableSlice)|获取切片中的最小值和最大值
|[FindMin2MaxInSlice](#FindMin2MaxInSlice)|获取切片中的最小值和最大值
|[FindMinFromComparableMap](#FindMinFromComparableMap)|获取 map 中的最小值
|[FindMinFromMap](#FindMinFromMap)|获取 map 中的最小值
|[FindMaxFromComparableMap](#FindMaxFromComparableMap)|获取 map 中的最大值
|[FindMaxFromMap](#FindMaxFromMap)|获取 map 中的最大值
|[FindMin2MaxFromComparableMap](#FindMin2MaxFromComparableMap)|获取 map 中的最小值和最大值
|[FindMin2MaxFromMap](#FindMin2MaxFromMap)|获取 map 中的最小值和最大值
|[SwapSlice](#SwapSlice)|将切片中的两个元素进行交换
|[MappingFromSlice](#MappingFromSlice)|将切片中的元素进行转换
|[MappingFromMap](#MappingFromMap)|将 map 中的元素进行转换
|[MergeSlices](#MergeSlices)|合并切片
|[MergeMaps](#MergeMaps)|合并 map，当多个 map 中存在相同的 key 时，后面的 map 中的 key 将会覆盖前面的 map 中的 key
|[MergeMapsWithSkip](#MergeMapsWithSkip)|合并 map，当多个 map 中存在相同的 key 时，后面的 map 中的 key 将会被跳过
|[ChooseRandomSliceElementRepeatN](#ChooseRandomSliceElementRepeatN)|返回 slice 中的 n 个可重复随机元素
|[ChooseRandomIndexRepeatN](#ChooseRandomIndexRepeatN)|返回 slice 中的 n 个可重复随机元素的索引
|[ChooseRandomSliceElement](#ChooseRandomSliceElement)|返回 slice 中随机一个元素，当 slice 长度为 0 时将会得到 V 的零值
|[ChooseRandomIndex](#ChooseRandomIndex)|返回 slice 中随机一个元素的索引，当 slice 长度为 0 时将会得到 -1
|[ChooseRandomSliceElementN](#ChooseRandomSliceElementN)|返回 slice 中的 n 个不可重复的随机元素
|[ChooseRandomIndexN](#ChooseRandomIndexN)|获取切片中的 n 个随机元素的索引
|[ChooseRandomMapKeyRepeatN](#ChooseRandomMapKeyRepeatN)|获取 map 中的 n 个随机 key，允许重复
|[ChooseRandomMapValueRepeatN](#ChooseRandomMapValueRepeatN)|获取 map 中的 n 个随机 n，允许重复
|[ChooseRandomMapKeyAndValueRepeatN](#ChooseRandomMapKeyAndValueRepeatN)|获取 map 中的 n 个随机 key 和 v，允许重复
|[ChooseRandomMapKey](#ChooseRandomMapKey)|获取 map 中的随机 key
|[ChooseRandomMapValue](#ChooseRandomMapValue)|获取 map 中的随机 value
|[ChooseRandomMapKeyN](#ChooseRandomMapKeyN)|获取 map 中的 inputN 个随机 key
|[ChooseRandomMapValueN](#ChooseRandomMapValueN)|获取 map 中的 n 个随机 value
|[ChooseRandomMapKeyAndValue](#ChooseRandomMapKeyAndValue)|获取 map 中的随机 key 和 v
|[ChooseRandomMapKeyAndValueN](#ChooseRandomMapKeyAndValueN)|获取 map 中的 inputN 个随机 key 和 v
|[DescBy](#DescBy)|返回降序比较结果
|[AscBy](#AscBy)|返回升序比较结果
|[Desc](#Desc)|对切片进行降序排序
|[DescByClone](#DescByClone)|对切片进行降序排序，返回排序后的切片
|[Asc](#Asc)|对切片进行升序排序
|[AscByClone](#AscByClone)|对切片进行升序排序，返回排序后的切片
|[Shuffle](#Shuffle)|对切片进行随机排序
|[ShuffleByClone](#ShuffleByClone)|对切片进行随机排序，返回排序后的切片


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[ComparisonHandler](#struct_ComparisonHandler)|用于比较 `source` 和 `target` 两个值是否相同的比较函数
|`STRUCT`|[OrderedValueGetter](#struct_OrderedValueGetter)|用于获取 v 的可排序字段值的函数

</details>


***
## 详情信息
#### func CloneSlice\[S ~[]V, V any\](slice S) S
<span id="CloneSlice"></span>
> 通过创建一个新切片并将 slice 的元素复制到新切片的方式来克隆切片

**示例代码：**

slice 克隆后将会得到一个新的 slice result，而 result 和 slice 将不会有任何关联，但是如果 slice 中的元素是引用类型，那么 result 中的元素将会和 slice 中的元素指向同一个地址
  - 示例中的结果将会输出 [1 2 3]


```go

func ExampleCloneSlice() {
	var slice = []int{1, 2, 3}
	var result = collection.CloneSlice(slice)
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestCloneSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{{"TestCloneSlice_NonEmptySlice", []int{1, 2, 3}, []int{1, 2, 3}}, {"TestCloneSlice_EmptySlice", []int{}, []int{}}, {"TestCloneSlice_NilSlice", nil, nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.CloneSlice(c.input)
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
			for i := 0; i < len(actual); i++ {
				if actual[i] != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the inputV of input is not equal")
				}
			}
		})
	}
}

```


</details>


***
#### func CloneMap\[M ~map[K]V, K comparable, V any\](m M) M
<span id="CloneMap"></span>
> 通过创建一个新 map 并将 m 的元素复制到新 map 的方式来克隆 map
>   - 当 m 为空时，将会返回 nil

**示例代码：**

map 克隆后将会得到一个新的 map result，而 result 和 map 将不会有任何关联，但是如果 map 中的元素是引用类型，那么 result 中的元素将会和 map 中的元素指向同一个地址
  - 示例中的结果将会输出 3


```go

func ExampleCloneMap() {
	var m = map[int]int{1: 1, 2: 2, 3: 3}
	var result = collection.CloneMap(m)
	fmt.Println(len(result))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestCloneMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]int
	}{{"TestCloneMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]int{1: 1, 2: 2, 3: 3}}, {"TestCloneMap_EmptyMap", map[int]int{}, map[int]int{}}, {"TestCloneMap_NilMap", nil, nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.CloneMap(c.input)
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of map is not equal")
			}
			for k, v := range actual {
				if v != c.expected[k] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the inputV of map is not equal")
				}
			}
		})
	}
}

```


</details>


***
#### func CloneSliceN\[S ~[]V, V any\](slice S, n int) []S
<span id="CloneSliceN"></span>
> 通过创建一个新切片并将 slice 的元素复制到新切片的方式来克隆切片为 n 个切片
>   - 当 slice 为空时，将会返回 nil，当 n <= 0 时，将会返回零值切片

**示例代码：**

slice 克隆为 2 个新的 slice，将会得到一个新的 slice result，而 result 和 slice 将不会有任何关联，但是如果 slice 中的元素是引用类型，那么 result 中的元素将会和 slice 中的元素指向同一个地址
  - result 的结果为 [[1 2 3] [1 2 3]]
  - 示例中的结果将会输出 2


```go

func ExampleCloneSliceN() {
	var slice = []int{1, 2, 3}
	var result = collection.CloneSliceN(slice, 2)
	fmt.Println(len(result))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestCloneSliceN(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		inputN   int
		expected [][]int
	}{{"TestCloneSliceN_NonEmptySlice", []int{1, 2, 3}, 2, [][]int{{1, 2, 3}, {1, 2, 3}}}, {"TestCloneSliceN_EmptySlice", []int{}, 2, [][]int{{}, {}}}, {"TestCloneSliceN_NilSlice", nil, 2, nil}, {"TestCloneSliceN_NonEmptySlice_ZeroN", []int{1, 2, 3}, 0, [][]int{}}, {"TestCloneSliceN_EmptySlice_ZeroN", []int{}, 0, [][]int{}}, {"TestCloneSliceN_NilSlice_ZeroN", nil, 0, nil}, {"TestCloneSliceN_NonEmptySlice_NegativeN", []int{1, 2, 3}, -1, [][]int{}}, {"TestCloneSliceN_EmptySlice_NegativeN", []int{}, -1, [][]int{}}, {"TestCloneSliceN_NilSlice_NegativeN", nil, -1, nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.CloneSliceN(c.input, c.inputN)
			if actual == nil {
				if c.expected != nil {
					t.Fatalf("%s failed, expected: %v, actual: %v, inputN: %d, error: %s", c.name, c.expected, c.inputN, actual, "after clone, the expected is nil")
				}
				return
			}
			for a, i := range actual {
				for b, v := range i {
					if v != c.expected[a][b] {
						t.Fatalf("%s failed, expected: %v, actual: %v, inputN: %d, error: %s", c.name, c.expected, c.inputN, actual, "after clone, the inputV of input is not equal")
					}
				}
			}
		})
	}
}

```


</details>


***
#### func CloneMapN\[M ~map[K]V, K comparable, V any\](m M, n int) []M
<span id="CloneMapN"></span>
> 通过创建一个新 map 并将 m 的元素复制到新 map 的方式来克隆 map 为 n 个 map
>   - 当 m 为空时，将会返回 nil，当 n <= 0 时，将会返回零值切片

**示例代码：**

map 克隆为 2 个新的 map，将会得到一个新的 map result，而 result 和 map 将不会有任何关联，但是如果 map 中的元素是引用类型，那么 result 中的元素将会和 map 中的元素指向同一个地址
  - result 的结果为 [map[1:1 2:2 3:3] map[1:1 2:2 3:3]] `无序的 Key-Value 对`
  - 示例中的结果将会输出 2


```go

func ExampleCloneMapN() {
	var m = map[int]int{1: 1, 2: 2, 3: 3}
	var result = collection.CloneMapN(m, 2)
	fmt.Println(len(result))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestCloneMapN(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		inputN   int
		expected []map[int]int
	}{{"TestCloneMapN_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, 2, []map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 3: 3}}}, {"TestCloneMapN_EmptyMap", map[int]int{}, 2, []map[int]int{{}, {}}}, {"TestCloneMapN_NilMap", nil, 2, nil}, {"TestCloneMapN_NonEmptyMap_ZeroN", map[int]int{1: 1, 2: 2, 3: 3}, 0, []map[int]int{}}, {"TestCloneMapN_EmptyMap_ZeroN", map[int]int{}, 0, []map[int]int{}}, {"TestCloneMapN_NilMap_ZeroN", nil, 0, nil}, {"TestCloneMapN_NonEmptyMap_NegativeN", map[int]int{1: 1, 2: 2, 3: 3}, -1, []map[int]int{}}, {"TestCloneMapN_EmptyMap_NegativeN", map[int]int{}, -1, []map[int]int{}}, {"TestCloneMapN_NilMap_NegativeN", nil, -1, nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.CloneMapN(c.input, c.inputN)
			if actual == nil {
				if c.expected != nil {
					t.Fatalf("%s failed, expected: %v, actual: %v, inputN: %d, error: %s", c.name, c.expected, actual, c.inputN, "after clone, the expected is nil")
				}
				return
			}
			for a, i := range actual {
				for b, v := range i {
					if v != c.expected[a][b] {
						t.Fatalf("%s failed, expected: %v, actual: %v, inputN: %d, error: %s", c.name, c.expected, actual, c.inputN, "after clone, the inputV of map is not equal")
					}
				}
			}
		})
	}
}

```


</details>


***
#### func CloneSlices\[S ~[]V, V any\](slices ...S) []S
<span id="CloneSlices"></span>
> 对 slices 中的每一项元素进行克隆，最终返回一个新的二维切片
>   - 当 slices 为空时，将会返回 nil
>   - 该函数相当于使用 CloneSlice 函数一次性对多个切片进行克隆

**示例代码：**

slice 克隆为 2 个新的 slice，将会得到一个新的 slice result，而 result 和 slice 将不会有任何关联，但是如果 slice 中的元素是引用类型，那么 result 中的元素将会和 slice 中的元素指向同一个地址
  - result 的结果为 [[1 2 3] [1 2 3]]


```go

func ExampleCloneSlices() {
	var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 2, 3}
	var result = collection.CloneSlices(slice1, slice2)
	fmt.Println(len(result))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestCloneSlices(t *testing.T) {
	var cases = []struct {
		name     string
		input    [][]int
		expected [][]int
	}{{"TestCloneSlices_NonEmptySlices", [][]int{{1, 2, 3}, {1, 2, 3}}, [][]int{{1, 2, 3}, {1, 2, 3}}}, {"TestCloneSlices_EmptySlices", [][]int{{}, {}}, [][]int{{}, {}}}, {"TestCloneSlices_NilSlices", nil, nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.CloneSlices(c.input...)
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
			for a, i := range actual {
				for b, v := range i {
					if v != c.expected[a][b] {
						t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the inputV of input is not equal")
					}
				}
			}
		})
	}
}

```


</details>


***
#### func CloneMaps\[M ~map[K]V, K comparable, V any\](maps ...M) []M
<span id="CloneMaps"></span>
> 对 maps 中的每一项元素进行克隆，最终返回一个新的 map 切片
>   - 当 maps 为空时，将会返回 nil
>   - 该函数相当于使用 CloneMap 函数一次性对多个 map 进行克隆

**示例代码：**

map 克隆为 2 个新的 map，将会得到一个新的 map result，而 result 和 map 将不会有任何关联，但是如果 map 中的元素是引用类型，那么 result 中的元素将会和 map 中的元素指向同一个地址
  - result 的结果为 [map[1:1 2:2 3:3] map[1:1 2:2 3:3]] `无序的 Key-Value 对`


```go

func ExampleCloneMaps() {
	var m1 = map[int]int{1: 1, 2: 2, 3: 3}
	var m2 = map[int]int{1: 1, 2: 2, 3: 3}
	var result = collection.CloneMaps(m1, m2)
	fmt.Println(len(result))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestCloneMaps(t *testing.T) {
	var cases = []struct {
		name     string
		input    []map[int]int
		expected []map[int]int
	}{{"TestCloneMaps_NonEmptyMaps", []map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 3: 3}}, []map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 3: 3}}}, {"TestCloneMaps_EmptyMaps", []map[int]int{{}, {}}, []map[int]int{{}, {}}}, {"TestCloneMaps_NilMaps", nil, nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.CloneMaps(c.input...)
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of maps is not equal")
			}
			for a, i := range actual {
				for b, v := range i {
					if v != c.expected[a][b] {
						t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the inputV of maps is not equal")
					}
				}
			}
		})
	}
}

```


</details>


***
#### func InSlice\[S ~[]V, V any\](slice S, v V, handler ComparisonHandler[V]) bool
<span id="InSlice"></span>
> 检查 v 是否被包含在 slice 中，当 handler 返回 true 时，表示 v 和 slice 中的某个元素相匹配

**示例代码：**

```go

func ExampleInSlice() {
	result := collection.InSlice([]int{1, 2, 3}, 2, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		inputV   int
		expected bool
	}{{"TestInSlice_NonEmptySliceIn", []int{1, 2, 3}, 1, true}, {"TestInSlice_NonEmptySliceNotIn", []int{1, 2, 3}, 4, false}, {"TestInSlice_EmptySlice", []int{}, 1, false}, {"TestInSlice_NilSlice", nil, 1, false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.InSlice(c.input, c.inputV, func(source, target int) bool {
				return source == target
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func InComparableSlice\[S ~[]V, V comparable\](slice S, v V) bool
<span id="InComparableSlice"></span>
> 检查 v 是否被包含在 slice 中

**示例代码：**

```go

func ExampleInComparableSlice() {
	result := collection.InComparableSlice([]int{1, 2, 3}, 2)
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInComparableSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		inputV   int
		expected bool
	}{{"TestInComparableSlice_NonEmptySliceIn", []int{1, 2, 3}, 1, true}, {"TestInComparableSlice_NonEmptySliceNotIn", []int{1, 2, 3}, 4, false}, {"TestInComparableSlice_EmptySlice", []int{}, 1, false}, {"TestInComparableSlice_NilSlice", nil, 1, false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.InComparableSlice(c.input, c.inputV)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func AllInSlice\[S ~[]V, V any\](slice S, values []V, handler ComparisonHandler[V]) bool
<span id="AllInSlice"></span>
> 检查 values 中的所有元素是否均被包含在 slice 中，当 handler 返回 true 时，表示 values 中的某个元素和 slice 中的某个元素相匹配
>   - 在所有 values 中的元素都被包含在 slice 中时，返回 true
>   - 当 values 长度为 0 或为 nil 时，将返回 true

**示例代码：**

```go

func ExampleAllInSlice() {
	result := collection.AllInSlice([]int{1, 2, 3}, []int{1, 2}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAllInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		inputV   []int
		expected bool
	}{{"TestAllInSlice_NonEmptySliceIn", []int{1, 2, 3}, []int{1, 2}, true}, {"TestAllInSlice_NonEmptySliceNotIn", []int{1, 2, 3}, []int{1, 4}, false}, {"TestAllInSlice_EmptySlice", []int{}, []int{1, 2}, false}, {"TestAllInSlice_NilSlice", nil, []int{1, 2}, false}, {"TestAllInSlice_EmptyValueSlice", []int{1, 2, 3}, []int{}, true}, {"TestAllInSlice_NilValueSlice", []int{1, 2, 3}, nil, true}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AllInSlice(c.input, c.inputV, func(source, target int) bool {
				return source == target
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func AllInComparableSlice\[S ~[]V, V comparable\](slice S, values []V) bool
<span id="AllInComparableSlice"></span>
> 检查 values 中的所有元素是否均被包含在 slice 中
>   - 在所有 values 中的元素都被包含在 slice 中时，返回 true
>   - 当 values 长度为 0 或为 nil 时，将返回 true

**示例代码：**

```go

func ExampleAllInComparableSlice() {
	result := collection.AllInComparableSlice([]int{1, 2, 3}, []int{1, 2})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAllInComparableSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		inputV   []int
		expected bool
	}{{"TestAllInComparableSlice_NonEmptySliceIn", []int{1, 2, 3}, []int{1, 2}, true}, {"TestAllInComparableSlice_NonEmptySliceNotIn", []int{1, 2, 3}, []int{1, 4}, false}, {"TestAllInComparableSlice_EmptySlice", []int{}, []int{1, 2}, false}, {"TestAllInComparableSlice_NilSlice", nil, []int{1, 2}, false}, {"TestAllInComparableSlice_EmptyValueSlice", []int{1, 2, 3}, []int{}, true}, {"TestAllInComparableSlice_NilValueSlice", []int{1, 2, 3}, nil, true}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AllInComparableSlice(c.input, c.inputV)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "not as expected")
			}
		})
	}
}

```


</details>


***
#### func AnyInSlice\[S ~[]V, V any\](slice S, values []V, handler ComparisonHandler[V]) bool
<span id="AnyInSlice"></span>
> 检查 values 中的任意一个元素是否被包含在 slice 中，当 handler 返回 true 时，表示 value 中的某个元素和 slice 中的某个元素相匹配
>   - 当 values 中的任意一个元素被包含在 slice 中时，返回 true

**示例代码：**

```go

func ExampleAnyInSlice() {
	result := collection.AnyInSlice([]int{1, 2, 3}, []int{1, 2}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAnyInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		inputV   []int
		expected bool
	}{{"TestAnyInSlice_NonEmptySliceIn", []int{1, 2, 3}, []int{1, 2}, true}, {"TestAnyInSlice_NonEmptySliceNotIn", []int{1, 2, 3}, []int{1, 4}, true}, {"TestAnyInSlice_EmptySlice", []int{}, []int{1, 2}, false}, {"TestAnyInSlice_NilSlice", nil, []int{1, 2}, false}, {"TestAnyInSlice_EmptyValueSlice", []int{1, 2, 3}, []int{}, false}, {"TestAnyInSlice_NilValueSlice", []int{1, 2, 3}, nil, false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyInSlice(c.input, c.inputV, func(source, target int) bool {
				return source == target
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func AnyInComparableSlice\[S ~[]V, V comparable\](slice S, values []V) bool
<span id="AnyInComparableSlice"></span>
> 检查 values 中的任意一个元素是否被包含在 slice 中
>   - 当 values 中的任意一个元素被包含在 slice 中时，返回 true

**示例代码：**

```go

func ExampleAnyInComparableSlice() {
	result := collection.AnyInComparableSlice([]int{1, 2, 3}, []int{1, 2})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAnyInComparableSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		inputV   []int
		expected bool
	}{{"TestAnyInComparableSlice_NonEmptySliceIn", []int{1, 2, 3}, []int{1, 2}, true}, {"TestAnyInComparableSlice_NonEmptySliceNotIn", []int{1, 2, 3}, []int{1, 4}, true}, {"TestAnyInComparableSlice_EmptySlice", []int{}, []int{1, 2}, false}, {"TestAnyInComparableSlice_NilSlice", nil, []int{1, 2}, false}, {"TestAnyInComparableSlice_EmptyValueSlice", []int{1, 2, 3}, []int{}, false}, {"TestAnyInComparableSlice_NilValueSlice", []int{1, 2, 3}, nil, false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyInComparableSlice(c.input, c.inputV)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "not as expected")
			}
		})
	}
}

```


</details>


***
#### func InSlices\[S ~[]V, V any\](slices []S, v V, handler ComparisonHandler[V]) bool
<span id="InSlices"></span>
> 通过将多个切片合并后检查 v 是否被包含在 slices 中，当 handler 返回 true 时，表示 v 和 slices 中的某个元素相匹配
>   - 当传入的 v 被包含在 slices 的任一成员中时，返回 true

**示例代码：**

```go

func ExampleInSlices() {
	result := collection.InSlices([][]int{{1, 2, 3}, {4, 5, 6}}, 2, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInSlices(t *testing.T) {
	var cases = []struct {
		name     string
		input    [][]int
		inputV   int
		expected bool
	}{{"TestInSlices_NonEmptySliceIn", [][]int{{1, 2}, {3, 4}}, 1, true}, {"TestInSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, 5, false}, {"TestInSlices_EmptySlice", [][]int{{}, {}}, 1, false}, {"TestInSlices_NilSlice", nil, 1, false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.InSlices(c.input, c.inputV, func(source, target int) bool {
				return source == target
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func InComparableSlices\[S ~[]V, V comparable\](slices []S, v V) bool
<span id="InComparableSlices"></span>
> 通过将多个切片合并后检查 v 是否被包含在 slices 中
>   - 当传入的 v 被包含在 slices 的任一成员中时，返回 true

**示例代码：**

```go

func ExampleInComparableSlices() {
	result := collection.InComparableSlices([][]int{{1, 2, 3}, {4, 5, 6}}, 2)
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInComparableSlices(t *testing.T) {
	var cases = []struct {
		name     string
		input    [][]int
		inputV   int
		expected bool
	}{{"TestInComparableSlices_NonEmptySliceIn", [][]int{{1, 2}, {3, 4}}, 1, true}, {"TestInComparableSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, 5, false}, {"TestInComparableSlices_EmptySlice", [][]int{{}, {}}, 1, false}, {"TestInComparableSlices_NilSlice", nil, 1, false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.InComparableSlices(c.input, c.inputV)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "")
			}
		})
	}
}

```


</details>


***
#### func AllInSlices\[S ~[]V, V any\](slices []S, values []V, handler ComparisonHandler[V]) bool
<span id="AllInSlices"></span>
> 通过将多个切片合并后检查 values 中的所有元素是否被包含在 slices 中，当 handler 返回 true 时，表示 values 中的某个元素和 slices 中的某个元素相匹配
>   - 当 values 中的所有元素都被包含在 slices 的任一成员中时，返回 true

**示例代码：**

```go

func ExampleAllInSlices() {
	result := collection.AllInSlices([][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAllInSlices(t *testing.T) {
	var cases = []struct {
		name        string
		input       [][]int
		inputValues []int
		expected    bool
	}{{"TestAllInSlices_NonEmptySliceIn", [][]int{{1, 2}, {3, 4}}, []int{1, 2}, true}, {"TestAllInSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, []int{1, 5}, false}, {"TestAllInSlices_EmptySlice", [][]int{{}, {}}, []int{1, 2}, false}, {"TestAllInSlices_NilSlice", nil, []int{1, 2}, false}, {"TestAllInSlices_EmptyValueSlice", [][]int{{1, 2}, {3, 4}}, []int{}, true}, {"TestAllInSlices_NilValueSlice", [][]int{{1, 2}, {3, 4}}, nil, true}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AllInSlices(c.input, c.inputValues, func(source, target int) bool {
				return source == target
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func AllInComparableSlices\[S ~[]V, V comparable\](slices []S, values []V) bool
<span id="AllInComparableSlices"></span>
> 通过将多个切片合并后检查 values 中的所有元素是否被包含在 slices 中
>   - 当 values 中的所有元素都被包含在 slices 的任一成员中时，返回 true

**示例代码：**

```go

func ExampleAllInComparableSlices() {
	result := collection.AllInComparableSlices([][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAllInComparableSlices(t *testing.T) {
	var cases = []struct {
		name        string
		input       [][]int
		inputValues []int
		expected    bool
	}{{"TestAllInComparableSlices_NonEmptySliceIn", [][]int{{1, 2}, {3, 4}}, []int{1, 2}, true}, {"TestAllInComparableSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, []int{1, 5}, false}, {"TestAllInComparableSlices_EmptySlice", [][]int{{}, {}}, []int{1, 2}, false}, {"TestAllInComparableSlices_NilSlice", nil, []int{1, 2}, false}, {"TestAllInComparableSlices_EmptyValueSlice", [][]int{{1, 2}, {3, 4}}, []int{}, true}, {"TestAllInComparableSlices_NilValueSlice", [][]int{{1, 2}, {3, 4}}, nil, true}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AllInComparableSlices(c.input, c.inputValues)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "")
			}
		})
	}
}

```


</details>


***
#### func AnyInSlices\[S ~[]V, V any\](slices []S, values []V, handler ComparisonHandler[V]) bool
<span id="AnyInSlices"></span>
> 通过将多个切片合并后检查 values 中的任意一个元素是否被包含在 slices 中，当 handler 返回 true 时，表示 values 中的某个元素和 slices 中的某个元素相匹配
>   - 当 values 中的任意一个元素被包含在 slices 的任一成员中时，返回 true

**示例代码：**

```go

func ExampleAnyInSlices() {
	result := collection.AnyInSlices([][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAnyInSlices(t *testing.T) {
	var cases = []struct {
		name     string
		slices   [][]int
		values   []int
		expected bool
	}{{"TestAnyInSlices_NonEmptySliceIn", [][]int{{1, 2}, {3, 4}}, []int{1, 2}, true}, {"TestAnyInSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, []int{1, 5}, true}, {"TestAnyInSlices_EmptySlice", [][]int{{}, {}}, []int{1, 2}, false}, {"TestAnyInSlices_NilSlice", nil, []int{1, 2}, false}, {"TestAnyInSlices_EmptyValueSlice", [][]int{{1, 2}, {3, 4}}, []int{}, false}, {"TestAnyInSlices_NilValueSlice", [][]int{{1, 2}, {3, 4}}, nil, false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyInSlices(c.slices, c.values, func(source, target int) bool {
				return source == target
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func AnyInComparableSlices\[S ~[]V, V comparable\](slices []S, values []V) bool
<span id="AnyInComparableSlices"></span>
> 通过将多个切片合并后检查 values 中的任意一个元素是否被包含在 slices 中
>   - 当 values 中的任意一个元素被包含在 slices 的任一成员中时，返回 true

**示例代码：**

```go

func ExampleAnyInComparableSlices() {
	result := collection.AnyInComparableSlices([][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAnyInComparableSlices(t *testing.T) {
	var cases = []struct {
		name     string
		slices   [][]int
		values   []int
		expected bool
	}{{"TestAnyInComparableSlices_NonEmptySliceIn", [][]int{{1, 2}, {3, 4}}, []int{1, 2}, true}, {"TestAnyInComparableSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, []int{1, 5}, true}, {"TestAnyInComparableSlices_EmptySlice", [][]int{{}, {}}, []int{1, 2}, false}, {"TestAnyInComparableSlices_NilSlice", nil, []int{1, 2}, false}, {"TestAnyInComparableSlices_EmptyValueSlice", [][]int{{1, 2}, {3, 4}}, []int{}, false}, {"TestAnyInComparableSlices_NilValueSlice", [][]int{{1, 2}, {3, 4}}, nil, false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyInComparableSlices(c.slices, c.values)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "")
			}
		})
	}
}

```


</details>


***
#### func InAllSlices\[S ~[]V, V any\](slices []S, v V, handler ComparisonHandler[V]) bool
<span id="InAllSlices"></span>
> 检查 v 是否被包含在 slices 的每一项元素中，当 handler 返回 true 时，表示 v 和 slices 中的某个元素相匹配
>   - 当 v 被包含在 slices 的每一项元素中时，返回 true

**示例代码：**

```go

func ExampleInAllSlices() {
	result := collection.InAllSlices([][]int{{1, 2, 3}, {4, 5, 6}}, 2, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInAllSlices(t *testing.T) {
	var cases = []struct {
		name     string
		slices   [][]int
		value    int
		expected bool
	}{{"TestInAllSlices_NonEmptySliceIn", [][]int{{1, 2}, {1, 3}}, 1, true}, {"TestInAllSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, 5, false}, {"TestInAllSlices_EmptySlice", [][]int{{}, {}}, 1, false}, {"TestInAllSlices_NilSlice", nil, 1, false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.InAllSlices(c.slices, c.value, func(source, target int) bool {
				return source == target
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func InAllComparableSlices\[S ~[]V, V comparable\](slices []S, v V) bool
<span id="InAllComparableSlices"></span>
> 检查 v 是否被包含在 slices 的每一项元素中
>   - 当 v 被包含在 slices 的每一项元素中时，返回 true

**示例代码：**

```go

func ExampleInAllComparableSlices() {
	result := collection.InAllComparableSlices([][]int{{1, 2, 3}, {4, 5, 6}}, 2)
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInAllComparableSlices(t *testing.T) {
	var cases = []struct {
		name     string
		slices   [][]int
		value    int
		expected bool
	}{{"TestInAllComparableSlices_NonEmptySliceIn", [][]int{{1, 2}, {1, 3}}, 1, true}, {"TestInAllComparableSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, 5, false}, {"TestInAllComparableSlices_EmptySlice", [][]int{{}, {}}, 1, false}, {"TestInAllComparableSlices_NilSlice", nil, 1, false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.InAllComparableSlices(c.slices, c.value)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "")
			}
		})
	}
}

```


</details>


***
#### func AnyInAllSlices\[S ~[]V, V any\](slices []S, values []V, handler ComparisonHandler[V]) bool
<span id="AnyInAllSlices"></span>
> 检查 slices 中的每一个元素是否均包含至少任意一个 values 中的元素，当 handler 返回 true 时，表示 value 中的某个元素和 slices 中的某个元素相匹配
>   - 当 slices 中的每一个元素均包含至少任意一个 values 中的元素时，返回 true

**示例代码：**

```go

func ExampleAnyInAllSlices() {
	result := collection.AnyInAllSlices([][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAnyInAllSlices(t *testing.T) {
	var cases = []struct {
		name     string
		slices   [][]int
		values   []int
		expected bool
	}{{"TestAnyInAllSlices_NonEmptySliceIn", [][]int{{1, 2}, {1, 3}}, []int{1, 2}, true}, {"TestAnyInAllSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, []int{1, 5}, false}, {"TestAnyInAllSlices_EmptySlice", [][]int{{}, {}}, []int{1, 2}, false}, {"TestAnyInAllSlices_NilSlice", nil, []int{1, 2}, false}, {"TestAnyInAllSlices_EmptyValueSlice", [][]int{{1, 2}, {3, 4}}, []int{}, false}, {"TestAnyInAllSlices_NilValueSlice", [][]int{{1, 2}, {3, 4}}, nil, false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyInAllSlices(c.slices, c.values, func(source, target int) bool {
				return source == target
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func AnyInAllComparableSlices\[S ~[]V, V comparable\](slices []S, values []V) bool
<span id="AnyInAllComparableSlices"></span>
> 检查 slices 中的每一个元素是否均包含至少任意一个 values 中的元素
>   - 当 slices 中的每一个元素均包含至少任意一个 values 中的元素时，返回 true

**示例代码：**

```go

func ExampleAnyInAllComparableSlices() {
	result := collection.AnyInAllComparableSlices([][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAnyInAllComparableSlices(t *testing.T) {
	var cases = []struct {
		name     string
		slices   [][]int
		values   []int
		expected bool
	}{{"TestAnyInAllComparableSlices_NonEmptySliceIn", [][]int{{1, 2}, {1, 3}}, []int{1, 2}, true}, {"TestAnyInAllComparableSlices_NonEmptySliceNotIn", [][]int{{1, 2}, {3, 4}}, []int{1, 5}, false}, {"TestAnyInAllComparableSlices_EmptySlice", [][]int{{}, {}}, []int{1, 2}, false}, {"TestAnyInAllComparableSlices_NilSlice", nil, []int{1, 2}, false}, {"TestAnyInAllComparableSlices_EmptyValueSlice", [][]int{{1, 2}, {3, 4}}, []int{}, false}, {"TestAnyInAllComparableSlices_NilValueSlice", [][]int{{1, 2}, {3, 4}}, nil, false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyInAllComparableSlices(c.slices, c.values)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "")
			}
		})
	}
}

```


</details>


***
#### func KeyInMap\[M ~map[K]V, K comparable, V any\](m M, key K) bool
<span id="KeyInMap"></span>
> 检查 m 中是否包含特定 key

**示例代码：**

```go

func ExampleKeyInMap() {
	result := collection.KeyInMap(map[string]int{"a": 1, "b": 2}, "a")
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestKeyInMap(t *testing.T) {
	var cases = []struct {
		name     string
		m        map[int]int
		key      int
		expected bool
	}{{"TestKeyInMap_NonEmptySliceIn", map[int]int{1: 1, 2: 2}, 1, true}, {"TestKeyInMap_NonEmptySliceNotIn", map[int]int{1: 1, 2: 2}, 3, false}, {"TestKeyInMap_EmptySlice", map[int]int{}, 1, false}, {"TestKeyInMap_NilSlice", nil, 1, false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.KeyInMap(c.m, c.key)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func ValueInMap\[M ~map[K]V, K comparable, V any\](m M, value V, handler ComparisonHandler[V]) bool
<span id="ValueInMap"></span>
> 检查 m 中是否包含特定 value，当 handler 返回 true 时，表示 value 和 m 中的某个元素相匹配

**示例代码：**

```go

func ExampleValueInMap() {
	result := collection.ValueInMap(map[string]int{"a": 1, "b": 2}, 2, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestValueInMap(t *testing.T) {
	var cases = []struct {
		name     string
		m        map[int]int
		value    int
		handler  collection.ComparisonHandler[int]
		expected bool
	}{{"TestValueInMap_NonEmptySliceIn", map[int]int{1: 1, 2: 2}, 1, intComparisonHandler, true}, {"TestValueInMap_NonEmptySliceNotIn", map[int]int{1: 1, 2: 2}, 3, intComparisonHandler, false}, {"TestValueInMap_EmptySlice", map[int]int{}, 1, intComparisonHandler, false}, {"TestValueInMap_NilSlice", nil, 1, intComparisonHandler, false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.ValueInMap(c.m, c.value, c.handler)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func AllKeyInMap\[M ~map[K]V, K comparable, V any\](m M, keys ...K) bool
<span id="AllKeyInMap"></span>
> 检查 m 中是否包含 keys 中所有的元素

**示例代码：**

```go

func ExampleAllKeyInMap() {
	result := collection.AllKeyInMap(map[string]int{"a": 1, "b": 2}, "a", "b")
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAllKeyInMap(t *testing.T) {
	var cases = []struct {
		name     string
		m        map[int]int
		keys     []int
		expected bool
	}{{"TestAllKeyInMap_NonEmptySliceIn", map[int]int{1: 1, 2: 2}, []int{1, 2}, true}, {"TestAllKeyInMap_NonEmptySliceNotIn", map[int]int{1: 1, 2: 2}, []int{1, 3}, false}, {"TestAllKeyInMap_EmptySlice", map[int]int{}, []int{1, 2}, false}, {"TestAllKeyInMap_NilSlice", nil, []int{1, 2}, false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AllKeyInMap(c.m, c.keys...)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func AllValueInMap\[M ~map[K]V, K comparable, V any\](m M, values []V, handler ComparisonHandler[V]) bool
<span id="AllValueInMap"></span>
> 检查 m 中是否包含 values 中所有的元素，当 handler 返回 true 时，表示 values 中的某个元素和 m 中的某个元素相匹配

**示例代码：**

```go

func ExampleAllValueInMap() {
	result := collection.AllValueInMap(map[string]int{"a": 1, "b": 2}, []int{1}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAllValueInMap(t *testing.T) {
	var cases = []struct {
		name     string
		m        map[int]int
		values   []int
		expected bool
	}{{"TestAllValueInMap_NonEmptySliceIn", map[int]int{1: 1, 2: 2}, []int{1, 2}, true}, {"TestAllValueInMap_NonEmptySliceNotIn", map[int]int{1: 1, 2: 2}, []int{1, 3}, false}, {"TestAllValueInMap_EmptySlice", map[int]int{}, []int{1, 2}, false}, {"TestAllValueInMap_NilSlice", nil, []int{1, 2}, false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AllValueInMap(c.m, c.values, intComparisonHandler)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after clone, the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func AnyKeyInMap\[M ~map[K]V, K comparable, V any\](m M, keys ...K) bool
<span id="AnyKeyInMap"></span>
> 检查 m 中是否包含 keys 中任意一个元素

**示例代码：**

```go

func ExampleAnyKeyInMap() {
	result := collection.AnyKeyInMap(map[string]int{"a": 1, "b": 2}, "a", "b")
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAnyKeyInMap(t *testing.T) {
	var cases = []struct {
		name     string
		m        map[int]int
		keys     []int
		expected bool
	}{{"TestAnyKeyInMap_NonEmptySliceIn", map[int]int{1: 1, 2: 2}, []int{1, 2}, true}, {"TestAnyKeyInMap_NonEmptySliceNotIn", map[int]int{1: 1, 2: 2}, []int{1, 3}, true}, {"TestAnyKeyInMap_EmptySlice", map[int]int{}, []int{1, 2}, false}, {"TestAnyKeyInMap_NilSlice", nil, []int{1, 2}, false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyKeyInMap(c.m, c.keys...)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "not as expected")
			}
		})
	}
}

```


</details>


***
#### func AnyValueInMap\[M ~map[K]V, K comparable, V any\](m M, values []V, handler ComparisonHandler[V]) bool
<span id="AnyValueInMap"></span>
> 检查 m 中是否包含 values 中任意一个元素，当 handler 返回 true 时，表示 values 中的某个元素和 m 中的某个元素相匹配

**示例代码：**

```go

func ExampleAnyValueInMap() {
	result := collection.AnyValueInMap(map[string]int{"a": 1, "b": 2}, []int{1}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAnyValueInMap(t *testing.T) {
	var cases = []struct {
		name     string
		m        map[int]int
		values   []int
		expected bool
	}{{"TestAnyValueInMap_NonEmptySliceIn", map[int]int{1: 1, 2: 2}, []int{1, 2}, true}, {"TestAnyValueInMap_NonEmptySliceNotIn", map[int]int{1: 1, 2: 2}, []int{1, 3}, true}, {"TestAnyValueInMap_EmptySlice", map[int]int{}, []int{1, 2}, false}, {"TestAnyValueInMap_NilSlice", nil, []int{1, 2}, false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyValueInMap(c.m, c.values, intComparisonHandler)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "not as expected")
			}
		})
	}
}

```


</details>


***
#### func AllKeyInMaps\[M ~map[K]V, K comparable, V any\](maps []M, keys ...K) bool
<span id="AllKeyInMaps"></span>
> 检查 maps 中的每一个元素是否均包含 keys 中所有的元素

**示例代码：**

```go

func ExampleAllKeyInMaps() {
	result := collection.AllKeyInMaps([]map[string]int{{"a": 1, "b": 2}, {"a": 1, "b": 2}}, "a", "b")
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAllKeyInMaps(t *testing.T) {
	var cases = []struct {
		name     string
		maps     []map[int]int
		keys     []int
		expected bool
	}{{"TestAllKeyInMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 2}, false}, {"TestAllKeyInMaps_NonEmptySliceNotIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 3}, false}, {"TestAllKeyInMaps_EmptySlice", []map[int]int{{1: 1, 2: 2}, {}}, []int{1, 2}, false}, {"TestAllKeyInMaps_NilSlice", []map[int]int{{}, {}}, []int{1, 2}, false}, {"TestAllKeyInMaps_EmptySlice", []map[int]int{}, []int{1, 2}, false}, {"TestAllKeyInMaps_NilSlice", nil, []int{1, 2}, false}, {"TestAllKeyInMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 4: 4}}, []int{1, 2}, true}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AllKeyInMaps(c.maps, c.keys...)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "not as expected")
			}
		})
	}
}

```


</details>


***
#### func AllValueInMaps\[M ~map[K]V, K comparable, V any\](maps []M, values []V, handler ComparisonHandler[V]) bool
<span id="AllValueInMaps"></span>
> 检查 maps 中的每一个元素是否均包含 value 中所有的元素，当 handler 返回 true 时，表示 value 中的某个元素和 maps 中的某个元素相匹配

**示例代码：**

```go

func ExampleAllValueInMaps() {
	result := collection.AllValueInMaps([]map[string]int{{"a": 1, "b": 2}, {"a": 1, "b": 2}}, []int{1}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAllValueInMaps(t *testing.T) {
	var cases = []struct {
		name     string
		maps     []map[int]int
		values   []int
		expected bool
	}{{"TestAllValueInMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 2}, false}, {"TestAllValueInMaps_NonEmptySliceNotIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 3}, false}, {"TestAllValueInMaps_EmptySlice", []map[int]int{{1: 1, 2: 2}, {}}, []int{1, 2}, false}, {"TestAllValueInMaps_NilSlice", []map[int]int{{}, {}}, []int{1, 2}, false}, {"TestAllValueInMaps_EmptySlice", []map[int]int{}, []int{1, 2}, false}, {"TestAllValueInMaps_NilSlice", nil, []int{1, 2}, false}, {"TestAllValueInMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 4: 4}}, []int{1, 2}, true}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AllValueInMaps(c.maps, c.values, intComparisonHandler)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "not as expected")
			}
		})
	}
}

```


</details>


***
#### func AnyKeyInMaps\[M ~map[K]V, K comparable, V any\](maps []M, keys ...K) bool
<span id="AnyKeyInMaps"></span>
> 检查 keys 中的任意一个元素是否被包含在 maps 中的任意一个元素中
>   - 当 keys 中的任意一个元素被包含在 maps 中的任意一个元素中时，返回 true

**示例代码：**

```go

func ExampleAnyKeyInMaps() {
	result := collection.AnyKeyInMaps([]map[string]int{{"a": 1, "b": 2}, {"a": 1, "b": 2}}, "a", "b")
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAnyKeyInMaps(t *testing.T) {
	var cases = []struct {
		name     string
		maps     []map[int]int
		keys     []int
		expected bool
	}{{"TestAnyKeyInMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 2}, true}, {"TestAnyKeyInMaps_NonEmptySliceNotIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 3}, true}, {"TestAnyKeyInMaps_EmptySlice", []map[int]int{{1: 1, 2: 2}, {}}, []int{1, 2}, true}, {"TestAnyKeyInMaps_NilSlice", []map[int]int{{}, {}}, []int{1, 2}, false}, {"TestAnyKeyInMaps_EmptySlice", []map[int]int{}, []int{1, 2}, false}, {"TestAnyKeyInMaps_NilSlice", nil, []int{1, 2}, false}, {"TestAnyKeyInMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 4: 4}}, []int{1, 2}, true}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyKeyInMaps(c.maps, c.keys...)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "not as expected")
			}
		})
	}
}

```


</details>


***
#### func AnyValueInMaps\[M ~map[K]V, K comparable, V any\](maps []M, values []V, handler ComparisonHandler[V]) bool
<span id="AnyValueInMaps"></span>
> 检查 maps 中的任意一个元素是否包含 value 中的任意一个元素，当 handler 返回 true 时，表示 value 中的某个元素和 maps 中的某个元素相匹配
>   - 当 maps 中的任意一个元素包含 value 中的任意一个元素时，返回 true

**示例代码：**

```go

func ExampleAnyValueInMaps() {
	result := collection.AnyValueInMaps([]map[string]int{{"a": 1, "b": 2}, {"a": 1, "b": 2}}, []int{1}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAnyValueInMaps(t *testing.T) {
	var cases = []struct {
		name     string
		maps     []map[int]int
		values   []int
		expected bool
	}{{"TestAnyValueInMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 2}, false}, {"TestAnyValueInMaps_NonEmptySliceNotIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 3}, true}, {"TestAnyValueInMaps_EmptySlice", []map[int]int{{1: 1, 2: 2}, {}}, []int{1, 2}, false}, {"TestAnyValueInMaps_NilSlice", []map[int]int{{}, {}}, []int{1, 2}, false}, {"TestAnyValueInMaps_EmptySlice", []map[int]int{}, []int{1, 2}, false}, {"TestAnyValueInMaps_NilSlice", nil, []int{1, 2}, false}, {"TestAnyValueInMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 4: 4}}, []int{1, 2}, true}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyValueInMaps(c.maps, c.values, intComparisonHandler)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "not as expected")
			}
		})
	}
}

```


</details>


***
#### func KeyInAllMaps\[M ~map[K]V, K comparable, V any\](maps []M, key K) bool
<span id="KeyInAllMaps"></span>
> 检查 key 是否被包含在 maps 的每一个元素中

**示例代码：**

```go

func ExampleKeyInAllMaps() {
	result := collection.KeyInAllMaps([]map[string]int{{"a": 1, "b": 2}, {"a": 1, "b": 2}}, "a")
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestKeyInAllMaps(t *testing.T) {
	var cases = []struct {
		name     string
		maps     []map[int]int
		key      int
		expected bool
	}{{"TestKeyInAllMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, 1, false}, {"TestKeyInAllMaps_NonEmptySliceNotIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, 3, false}, {"TestKeyInAllMaps_EmptySlice", []map[int]int{{1: 1, 2: 2}, {}}, 1, false}, {"TestKeyInAllMaps_NilSlice", []map[int]int{{}, {}}, 1, false}, {"TestKeyInAllMaps_EmptySlice", []map[int]int{}, 1, false}, {"TestKeyInAllMaps_NilSlice", nil, 1, false}, {"TestKeyInAllMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 4: 4}}, 1, true}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.KeyInAllMaps(c.maps, c.key)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "not as expected")
			}
		})
	}
}

```


</details>


***
#### func AnyKeyInAllMaps\[M ~map[K]V, K comparable, V any\](maps []M, keys []K) bool
<span id="AnyKeyInAllMaps"></span>
> 检查 maps 中的每一个元素是否均包含 keys 中任意一个元素
>   - 当 maps 中的每一个元素均包含 keys 中任意一个元素时，返回 true

**示例代码：**

```go

func ExampleAnyKeyInAllMaps() {
	result := collection.AnyKeyInAllMaps([]map[string]int{{"a": 1, "b": 2}, {"a": 1, "b": 2}}, []string{"a"})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAnyKeyInAllMaps(t *testing.T) {
	var cases = []struct {
		name     string
		maps     []map[int]int
		keys     []int
		expected bool
	}{{"TestAnyKeyInAllMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 2}, false}, {"TestAnyKeyInAllMaps_NonEmptySliceNotIn", []map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []int{1, 3}, true}, {"TestAnyKeyInAllMaps_EmptySlice", []map[int]int{{1: 1, 2: 2}, {}}, []int{1, 2}, false}, {"TestAnyKeyInAllMaps_NilSlice", []map[int]int{{}, {}}, []int{1, 2}, false}, {"TestAnyKeyInAllMaps_EmptySlice", []map[int]int{}, []int{1, 2}, false}, {"TestAnyKeyInAllMaps_NilSlice", nil, []int{1, 2}, false}, {"TestAnyKeyInAllMaps_NonEmptySliceIn", []map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 4: 4}}, []int{1, 2}, true}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual = collection.AnyKeyInAllMaps(c.maps, c.keys)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v", c.name, c.expected, actual)
			}
		})
	}
}

```


</details>


***
#### func ConvertSliceToAny\[S ~[]V, V any\](s S) []any
<span id="ConvertSliceToAny"></span>
> 将切片转换为任意类型的切片

**示例代码：**

```go

func ExampleConvertSliceToAny() {
	result := collection.ConvertSliceToAny([]int{1, 2, 3})
	fmt.Println(reflect.TypeOf(result).String(), len(result))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestConvertSliceToAny(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []interface{}
	}{{name: "TestConvertSliceToAny_NonEmpty", input: []int{1, 2, 3}, expected: []any{1, 2, 3}}, {name: "TestConvertSliceToAny_Empty", input: []int{}, expected: []any{}}, {name: "TestConvertSliceToAny_Nil", input: nil, expected: nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ConvertSliceToAny(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			for i := 0; i < len(actual); i++ {
				av, ev := actual[i], c.expected[i]
				if reflect.TypeOf(av).Kind() != reflect.TypeOf(ev).Kind() {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
				if av != ev {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
			}
		})
	}
}

```


</details>


***
#### func ConvertSliceToIndexMap\[S ~[]V, V any\](s S) map[int]V
<span id="ConvertSliceToIndexMap"></span>
> 将切片转换为索引为键的映射

**示例代码：**

```go

func ExampleConvertSliceToIndexMap() {
	slice := []int{1, 2, 3}
	result := collection.ConvertSliceToIndexMap(slice)
	for i, v := range slice {
		fmt.Println(result[i], v)
	}
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestConvertSliceToIndexMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected map[int]int
	}{{name: "TestConvertSliceToIndexMap_NonEmpty", input: []int{1, 2, 3}, expected: map[int]int{0: 1, 1: 2, 2: 3}}, {name: "TestConvertSliceToIndexMap_Empty", input: []int{}, expected: map[int]int{}}, {name: "TestConvertSliceToIndexMap_Nil", input: nil, expected: nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ConvertSliceToIndexMap(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			for k, v := range actual {
				if c.expected[k] != v {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
			}
		})
	}
}

```


</details>


***
#### func ConvertSliceToIndexOnlyMap\[S ~[]V, V any\](s S) map[int]struct {}
<span id="ConvertSliceToIndexOnlyMap"></span>
> 将切片转换为索引为键的映射

**示例代码：**

```go

func ExampleConvertSliceToIndexOnlyMap() {
	slice := []int{1, 2, 3}
	result := collection.ConvertSliceToIndexOnlyMap(slice)
	expected := map[int]bool{0: true, 1: true, 2: true}
	for k := range result {
		fmt.Println(expected[k])
	}
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestConvertSliceToIndexOnlyMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected map[int]struct{}
	}{{name: "TestConvertSliceToIndexOnlyMap_NonEmpty", input: []int{1, 2, 3}, expected: map[int]struct{}{0: {}, 1: {}, 2: {}}}, {name: "TestConvertSliceToIndexOnlyMap_Empty", input: []int{}, expected: map[int]struct{}{}}, {name: "TestConvertSliceToIndexOnlyMap_Nil", input: nil, expected: nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ConvertSliceToIndexOnlyMap(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			for k, v := range actual {
				if _, ok := c.expected[k]; !ok {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
				if v != struct{}{} {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
			}
		})
	}
}

```


</details>


***
#### func ConvertSliceToMap\[S ~[]V, V comparable\](s S) map[V]struct {}
<span id="ConvertSliceToMap"></span>
> 将切片转换为值为键的映射

**示例代码：**

```go

func ExampleConvertSliceToMap() {
	slice := []int{1, 2, 3}
	result := collection.ConvertSliceToMap(slice)
	fmt.Println(collection.AllKeyInMap(result, slice...))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestConvertSliceToMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected map[int]struct{}
	}{{name: "TestConvertSliceToMap_NonEmpty", input: []int{1, 2, 3}, expected: map[int]struct{}{1: {}, 2: {}, 3: {}}}, {name: "TestConvertSliceToMap_Empty", input: []int{}, expected: map[int]struct{}{}}, {name: "TestConvertSliceToMap_Nil", input: nil, expected: nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ConvertSliceToMap(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			for k, v := range actual {
				if _, ok := c.expected[k]; !ok {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
				if v != struct{}{} {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
			}
		})
	}
}

```


</details>


***
#### func ConvertSliceToBoolMap\[S ~[]V, V comparable\](s S) map[V]bool
<span id="ConvertSliceToBoolMap"></span>
> 将切片转换为值为键的映射

**示例代码：**

```go

func ExampleConvertSliceToBoolMap() {
	slice := []int{1, 2, 3}
	result := collection.ConvertSliceToBoolMap(slice)
	for _, v := range slice {
		fmt.Println(v, result[v])
	}
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestConvertSliceToBoolMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected map[int]bool
	}{{name: "TestConvertSliceToBoolMap_NonEmpty", input: []int{1, 2, 3}, expected: map[int]bool{1: true, 2: true, 3: true}}, {name: "TestConvertSliceToBoolMap_Empty", input: []int{}, expected: map[int]bool{}}, {name: "TestConvertSliceToBoolMap_Nil", input: nil, expected: nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ConvertSliceToBoolMap(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			for k, v := range actual {
				if c.expected[k] != v {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
			}
		})
	}
}

```


</details>


***
#### func ConvertMapKeysToSlice\[M ~map[K]V, K comparable, V any\](m M) []K
<span id="ConvertMapKeysToSlice"></span>
> 将映射的键转换为切片

**示例代码：**

```go

func ExampleConvertMapKeysToSlice() {
	result := collection.ConvertMapKeysToSlice(map[int]int{1: 1, 2: 2, 3: 3})
	for i, v := range result {
		fmt.Println(i, v)
	}
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestConvertMapKeysToSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected []int
	}{{name: "TestConvertMapKeysToSlice_NonEmpty", input: map[int]int{1: 1, 2: 2, 3: 3}, expected: []int{1, 2, 3}}, {name: "TestConvertMapKeysToSlice_Empty", input: map[int]int{}, expected: []int{}}, {name: "TestConvertMapKeysToSlice_Nil", input: nil, expected: nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ConvertMapKeysToSlice(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			var matchCount = 0
			for _, av := range actual {
				for _, ev := range c.expected {
					if av == ev {
						matchCount++
					}
				}
			}
			if matchCount != len(actual) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
		})
	}
}

```


</details>


***
#### func ConvertMapValuesToSlice\[M ~map[K]V, K comparable, V any\](m M) []V
<span id="ConvertMapValuesToSlice"></span>
> 将映射的值转换为切片

**示例代码：**

```go

func ExampleConvertMapValuesToSlice() {
	result := collection.ConvertMapValuesToSlice(map[int]int{1: 1, 2: 2, 3: 3})
	expected := map[int]bool{1: true, 2: true, 3: true}
	for _, v := range result {
		fmt.Println(expected[v])
	}
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestConvertMapValuesToSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected []int
	}{{name: "TestConvertMapValuesToSlice_NonEmpty", input: map[int]int{1: 1, 2: 2, 3: 3}, expected: []int{1, 2, 3}}, {name: "TestConvertMapValuesToSlice_Empty", input: map[int]int{}, expected: []int{}}, {name: "TestConvertMapValuesToSlice_Nil", input: nil, expected: nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ConvertMapValuesToSlice(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			var matchCount = 0
			for _, av := range actual {
				for _, ev := range c.expected {
					if av == ev {
						matchCount++
					}
				}
			}
			if matchCount != len(actual) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
		})
	}
}

```


</details>


***
#### func InvertMap\[M ~map[K]V, N map[V]K, K comparable, V comparable\](m M) N
<span id="InvertMap"></span>
> 将映射的键和值互换

**示例代码：**

```go

func ExampleInvertMap() {
	result := collection.InvertMap(map[int]string{1: "a", 2: "b", 3: "c"})
	fmt.Println(collection.AllKeyInMap(result, "a", "b", "c"))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestInvertMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]string
		expected map[string]int
	}{{name: "TestInvertMap_NonEmpty", input: map[int]string{1: "1", 2: "2", 3: "3"}, expected: map[string]int{"1": 1, "2": 2, "3": 3}}, {name: "TestInvertMap_Empty", input: map[int]string{}, expected: map[string]int{}}, {name: "TestInvertMap_Nil", input: nil, expected: nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.InvertMap[map[int]string, map[string]int](c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			for k, v := range actual {
				if c.expected[k] != v {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
			}
		})
	}
}

```


</details>


***
#### func ConvertMapValuesToBool\[M ~map[K]V, N map[K]bool, K comparable, V any\](m M) N
<span id="ConvertMapValuesToBool"></span>
> 将映射的值转换为布尔值

**示例代码：**

```go

func ExampleConvertMapValuesToBool() {
	result := collection.ConvertMapValuesToBool(map[int]int{1: 1})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestConvertMapValuesToBool(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]bool
	}{{name: "TestConvertMapValuesToBool_NonEmpty", input: map[int]int{1: 1, 2: 2, 3: 3}, expected: map[int]bool{1: true, 2: true, 3: true}}, {name: "TestConvertMapValuesToBool_Empty", input: map[int]int{}, expected: map[int]bool{}}, {name: "TestConvertMapValuesToBool_Nil", input: nil, expected: nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ConvertMapValuesToBool[map[int]int, map[int]bool](c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			for k, v := range actual {
				if c.expected[k] != v {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
			}
		})
	}
}

```


</details>


***
#### func ReverseSlice\[S ~[]V, V any\](s *S)
<span id="ReverseSlice"></span>
> 将切片反转

**示例代码：**

```go

func ExampleReverseSlice() {
	var s = []int{1, 2, 3}
	collection.ReverseSlice(&s)
	fmt.Println(s)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestReverseSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{{name: "TestReverseSlice_NonEmpty", input: []int{1, 2, 3}, expected: []int{3, 2, 1}}, {name: "TestReverseSlice_Empty", input: []int{}, expected: []int{}}, {name: "TestReverseSlice_Nil", input: nil, expected: nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.ReverseSlice(&c.input)
			if len(c.input) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, c.input)
			}
			for i := 0; i < len(c.input); i++ {
				av, ev := c.input[i], c.expected[i]
				if reflect.TypeOf(av).Kind() != reflect.TypeOf(ev).Kind() {
					t.Errorf("expected: %v, actual: %v", c.expected, c.input)
				}
				if av != ev {
					t.Errorf("expected: %v, actual: %v", c.expected, c.input)
				}
			}
		})
	}
}

```


</details>


***
#### func ClearSlice\[S ~[]V, V any\](slice *S)
<span id="ClearSlice"></span>
> 清空切片

**示例代码：**

```go

func ExampleClearSlice() {
	slice := []int{1, 2, 3, 4, 5}
	ClearSlice(&slice)
	fmt.Println(slice)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestClearSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{{"TestClearSlice_NonEmptySlice", []int{1, 2, 3}, []int{}}, {"TestClearSlice_EmptySlice", []int{}, []int{}}, {"TestClearSlice_NilSlice", nil, nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.ClearSlice(&c.input)
			if len(c.input) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after clear, the length of input is not equal")
			}
			for i := 0; i < len(c.input); i++ {
				if c.input[i] != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after clear, the inputV of input is not equal")
				}
			}
		})
	}
}

```


</details>


***
#### func ClearMap\[M ~map[K]V, K comparable, V any\](m M)
<span id="ClearMap"></span>
> 清空 map

**示例代码：**

```go

func ExampleClearMap() {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	ClearMap(m)
	fmt.Println(m)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestClearMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]int
	}{{"TestClearMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]int{}}, {"TestClearMap_EmptyMap", map[int]int{}, map[int]int{}}, {"TestClearMap_NilMap", nil, nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.ClearMap(c.input)
			if len(c.input) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after clear, the length of map is not equal")
			}
			for k, v := range c.input {
				if v != c.expected[k] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after clear, the inputV of map is not equal")
				}
			}
		})
	}
}

```


</details>


***
#### func DropSliceByIndices\[S ~[]V, V any\](slice *S, indices ...int)
<span id="DropSliceByIndices"></span>
> 删除切片中特定索引的元素

**示例代码：**

```go

func ExampleDropSliceByIndices() {
	slice := []int{1, 2, 3, 4, 5}
	DropSliceByIndices(&slice, 1, 3)
	fmt.Println(slice)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestDropSliceByIndices(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		indices  []int
		expected []int
	}{{"TestDropSliceByIndices_NonEmptySlice", []int{1, 2, 3, 4, 5}, []int{1, 3}, []int{1, 3, 5}}, {"TestDropSliceByIndices_EmptySlice", []int{}, []int{1, 3}, []int{}}, {"TestDropSliceByIndices_NilSlice", nil, []int{1, 3}, nil}, {"TestDropSliceByIndices_NonEmptySlice", []int{1, 2, 3, 4, 5}, []int{}, []int{1, 2, 3, 4, 5}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.DropSliceByIndices(&c.input, c.indices...)
			if len(c.input) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after drop, the length of input is not equal")
			}
			for i := 0; i < len(c.input); i++ {
				if c.input[i] != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after drop, the inputV of input is not equal")
				}
			}
		})
	}
}

```


</details>


***
#### func DropSliceByCondition\[S ~[]V, V any\](slice *S, condition func (v V)  bool)
<span id="DropSliceByCondition"></span>
> 删除切片中符合条件的元素
>   - condition 的返回值为 true 时，将会删除该元素

**示例代码：**

```go

func ExampleDropSliceByCondition() {
	slice := []int{1, 2, 3, 4, 5}
	DropSliceByCondition(&slice, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(slice)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestDropSliceByCondition(t *testing.T) {
	var cases = []struct {
		name      string
		input     []int
		condition func(v int) bool
		expected  []int
	}{{"TestDropSliceByCondition_NonEmptySlice", []int{1, 2, 3, 4, 5}, func(v int) bool {
		return v%2 == 0
	}, []int{1, 3, 5}}, {"TestDropSliceByCondition_EmptySlice", []int{}, func(v int) bool {
		return v%2 == 0
	}, []int{}}, {"TestDropSliceByCondition_NilSlice", nil, func(v int) bool {
		return v%2 == 0
	}, nil}, {"TestDropSliceByCondition_NonEmptySlice", []int{1, 2, 3, 4, 5}, func(v int) bool {
		return v%2 == 1
	}, []int{2, 4}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.DropSliceByCondition(&c.input, c.condition)
			if len(c.input) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after drop, the length of input is not equal")
			}
			for i := 0; i < len(c.input); i++ {
				if c.input[i] != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after drop, the inputV of input is not equal")
				}
			}
		})
	}
}

```


</details>


***
#### func DropSliceOverlappingElements\[S ~[]V, V any\](slice *S, anotherSlice []V, comparisonHandler ComparisonHandler[V])
<span id="DropSliceOverlappingElements"></span>
> 删除切片中与另一个切片重叠的元素

**示例代码：**

```go

func ExampleDropSliceOverlappingElements() {
	slice := []int{1, 2, 3, 4, 5}
	DropSliceOverlappingElements(&slice, []int{1, 3, 5}, func(source, target int) bool {
		return source == target
	})
	fmt.Println(slice)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestDropSliceOverlappingElements(t *testing.T) {
	var cases = []struct {
		name               string
		input              []int
		anotherSlice       []int
		comparisonHandler  collection.ComparisonHandler[int]
		expected           []int
		expectedAnother    []int
		expectedComparison []int
	}{{"TestDropSliceOverlappingElements_NonEmptySlice", []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 6, 7}, func(v1, v2 int) bool {
		return v1 == v2
	}, []int{1, 2}, []int{6, 7}, []int{3, 4, 5}}, {"TestDropSliceOverlappingElements_EmptySlice", []int{}, []int{3, 4, 5, 6, 7}, func(v1, v2 int) bool {
		return v1 == v2
	}, []int{}, []int{3, 4, 5, 6, 7}, []int{}}, {"TestDropSliceOverlappingElements_NilSlice", nil, []int{3, 4, 5, 6, 7}, func(v1, v2 int) bool {
		return v1 == v2
	}, nil, []int{3, 4, 5, 6, 7}, nil}, {"TestDropSliceOverlappingElements_NonEmptySlice", []int{1, 2, 3, 4, 5}, []int{}, func(v1, v2 int) bool {
		return v1 == v2
	}, []int{1, 2, 3, 4, 5}, []int{}, []int{}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.DropSliceOverlappingElements(&c.input, c.anotherSlice, c.comparisonHandler)
			if len(c.input) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after drop, the length of input is not equal")
			}
			for i := 0; i < len(c.input); i++ {
				if c.input[i] != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, c.input, "after drop, the inputV of input is not equal")
				}
			}
		})
	}
}

```


</details>


***
#### func DeduplicateSliceInPlace\[S ~[]V, V comparable\](s *S)
<span id="DeduplicateSliceInPlace"></span>
> 去除切片中的重复元素

**示例代码：**

```go

func ExampleDeduplicateSliceInPlace() {
	slice := []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}
	DeduplicateSliceInPlace(&slice)
	fmt.Println(slice)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestDeduplicateSliceInPlace(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{{name: "TestDeduplicateSliceInPlace_NonEmpty", input: []int{1, 2, 3, 1, 2, 3}, expected: []int{1, 2, 3}}, {name: "TestDeduplicateSliceInPlace_Empty", input: []int{}, expected: []int{}}, {name: "TestDeduplicateSliceInPlace_Nil", input: nil, expected: nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.DeduplicateSliceInPlace(&c.input)
			if len(c.input) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, c.input)
			}
			for i := 0; i < len(c.input); i++ {
				av, ev := c.input[i], c.expected[i]
				if av != ev {
					t.Errorf("expected: %v, actual: %v", c.expected, c.input)
				}
			}
		})
	}
}

```


</details>


***
#### func DeduplicateSlice\[S ~[]V, V comparable\](s S) S
<span id="DeduplicateSlice"></span>
> 去除切片中的重复元素，返回新切片

**示例代码：**

```go

func ExampleDeduplicateSlice() {
	slice := []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}
	fmt.Println(DeduplicateSlice(slice))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestDeduplicateSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{{name: "TestDeduplicateSlice_NonEmpty", input: []int{1, 2, 3, 1, 2, 3}, expected: []int{1, 2, 3}}, {name: "TestDeduplicateSlice_Empty", input: []int{}, expected: []int{}}, {name: "TestDeduplicateSlice_Nil", input: nil, expected: nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.DeduplicateSlice(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			for i := 0; i < len(actual); i++ {
				av, ev := actual[i], c.expected[i]
				if av != ev {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
			}
		})
	}
}

```


</details>


***
#### func DeduplicateSliceInPlaceWithCompare\[S ~[]V, V any\](s *S, compare func (a V)  bool)
<span id="DeduplicateSliceInPlaceWithCompare"></span>
> 去除切片中的重复元素，使用自定义的比较函数

**示例代码：**

```go

func ExampleDeduplicateSliceInPlaceWithCompare() {
	slice := []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}
	DeduplicateSliceInPlaceWithCompare(&slice, func(a, b int) bool {
		return a == b
	})
	fmt.Println(slice)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestDeduplicateSliceInPlaceWithCompare(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{{name: "TestDeduplicateSliceInPlaceWithCompare_NonEmpty", input: []int{1, 2, 3, 1, 2, 3}, expected: []int{1, 2, 3}}, {name: "TestDeduplicateSliceInPlaceWithCompare_Empty", input: []int{}, expected: []int{}}, {name: "TestDeduplicateSliceInPlaceWithCompare_Nil", input: nil, expected: nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.DeduplicateSliceInPlaceWithCompare(&c.input, func(a, b int) bool {
				return a == b
			})
			if len(c.input) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, c.input)
			}
			for i := 0; i < len(c.input); i++ {
				av, ev := c.input[i], c.expected[i]
				if av != ev {
					t.Errorf("expected: %v, actual: %v", c.expected, c.input)
				}
			}
		})
	}
}

```


</details>


***
#### func DeduplicateSliceWithCompare\[S ~[]V, V any\](s S, compare func (a V)  bool) S
<span id="DeduplicateSliceWithCompare"></span>
> 去除切片中的重复元素，使用自定义的比较函数，返回新的切片

**示例代码：**

```go

func ExampleDeduplicateSliceWithCompare() {
	slice := []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}
	fmt.Println(DeduplicateSliceWithCompare(slice, func(a, b int) bool {
		return a == b
	}))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestDeduplicateSliceWithCompare(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{{name: "TestDeduplicateSliceWithCompare_NonEmpty", input: []int{1, 2, 3, 1, 2, 3}, expected: []int{1, 2, 3}}, {name: "TestDeduplicateSliceWithCompare_Empty", input: []int{}, expected: []int{}}, {name: "TestDeduplicateSliceWithCompare_Nil", input: nil, expected: nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.DeduplicateSliceWithCompare(c.input, func(a, b int) bool {
				return a == b
			})
			if len(actual) != len(c.expected) {
				t.Errorf("expected: %v, actual: %v", c.expected, actual)
			}
			for i := 0; i < len(actual); i++ {
				av, ev := actual[i], c.expected[i]
				if av != ev {
					t.Errorf("expected: %v, actual: %v", c.expected, actual)
				}
			}
		})
	}
}

```


</details>


***
#### func FilterOutByIndices\[S []V, V any\](slice S, indices ...int) S
<span id="FilterOutByIndices"></span>
> 过滤切片中特定索引的元素，返回过滤后的切片

**示例代码：**

```go

func ExampleFilterOutByIndices() {
	var slice = []int{1, 2, 3, 4, 5}
	var result = collection.FilterOutByIndices(slice, 1, 3)
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFilterOutByIndices(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		indices  []int
		expected []int
	}{{"TestFilterOutByIndices_NonEmptySlice", []int{1, 2, 3, 4, 5}, []int{1, 3}, []int{1, 3, 5}}, {"TestFilterOutByIndices_EmptySlice", []int{}, []int{1, 3}, []int{}}, {"TestFilterOutByIndices_NilSlice", nil, []int{1, 3}, nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FilterOutByIndices(c.input, c.indices...)
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the length of input is not equal")
			}
			for i := 0; i < len(actual); i++ {
				if actual[i] != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the inputV of input is not equal")
				}
			}
		})
	}
}

```


</details>


***
#### func FilterOutByCondition\[S ~[]V, V any\](slice S, condition func (v V)  bool) S
<span id="FilterOutByCondition"></span>
> 过滤切片中符合条件的元素，返回过滤后的切片
>   - condition 的返回值为 true 时，将会过滤掉该元素

**示例代码：**

```go

func ExampleFilterOutByCondition() {
	var slice = []int{1, 2, 3, 4, 5}
	var result = collection.FilterOutByCondition(slice, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFilterOutByCondition(t *testing.T) {
	var cases = []struct {
		name      string
		input     []int
		condition func(int) bool
		expected  []int
	}{{"TestFilterOutByCondition_NonEmptySlice", []int{1, 2, 3, 4, 5}, func(v int) bool {
		return v%2 == 0
	}, []int{1, 3, 5}}, {"TestFilterOutByCondition_EmptySlice", []int{}, func(v int) bool {
		return v%2 == 0
	}, []int{}}, {"TestFilterOutByCondition_NilSlice", nil, func(v int) bool {
		return v%2 == 0
	}, nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FilterOutByCondition(c.input, c.condition)
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the length of input is not equal")
			}
			for i := 0; i < len(actual); i++ {
				if actual[i] != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the inputV of input is not equal")
				}
			}
		})
	}
}

```


</details>


***
#### func FilterOutByKey\[M ~map[K]V, K comparable, V any\](m M, key K) M
<span id="FilterOutByKey"></span>
> 过滤 map 中特定的 key，返回过滤后的 map

**示例代码：**

```go

func ExampleFilterOutByKey() {
	var m = map[string]int{"a": 1, "b": 2, "c": 3}
	var result = collection.FilterOutByKey(m, "b")
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFilterOutByKey(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		key      int
		expected map[int]int
	}{{"TestFilterOutByKey_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, 1, map[int]int{2: 2, 3: 3}}, {"TestFilterOutByKey_EmptyMap", map[int]int{}, 1, map[int]int{}}, {"TestFilterOutByKey_NilMap", nil, 1, nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FilterOutByKey(c.input, c.key)
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the length of map is not equal")
			}
			for k, v := range actual {
				if v != c.expected[k] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the inputV of map is not equal")
				}
			}
		})
	}
}

```


</details>


***
#### func FilterOutByValue\[M ~map[K]V, K comparable, V any\](m M, value V, handler ComparisonHandler[V]) M
<span id="FilterOutByValue"></span>
> 过滤 map 中特定的 value，返回过滤后的 map

**示例代码：**

```go

func ExampleFilterOutByValue() {
	var m = map[string]int{"a": 1, "b": 2, "c": 3}
	var result = collection.FilterOutByValue(m, 2, func(source, target int) bool {
		return source == target
	})
	fmt.Println(len(result))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFilterOutByValue(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		value    int
		expected map[int]int
	}{{"TestFilterOutByValue_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, 1, map[int]int{2: 2, 3: 3}}, {"TestFilterOutByValue_EmptyMap", map[int]int{}, 1, map[int]int{}}, {"TestFilterOutByValue_NilMap", nil, 1, nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FilterOutByValue(c.input, c.value, func(source, target int) bool {
				return source == target
			})
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the length of map is not equal")
			}
			for k, v := range actual {
				if v != c.expected[k] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the inputV of map is not equal")
				}
			}
		})
	}
}

```


</details>


***
#### func FilterOutByKeys\[M ~map[K]V, K comparable, V any\](m M, keys ...K) M
<span id="FilterOutByKeys"></span>
> 过滤 map 中多个 key，返回过滤后的 map

**示例代码：**

```go

func ExampleFilterOutByKeys() {
	var m = map[string]int{"a": 1, "b": 2, "c": 3}
	var result = collection.FilterOutByKeys(m, "a", "c")
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFilterOutByKeys(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		keys     []int
		expected map[int]int
	}{{"TestFilterOutByKeys_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, []int{1, 3}, map[int]int{2: 2}}, {"TestFilterOutByKeys_EmptyMap", map[int]int{}, []int{1, 3}, map[int]int{}}, {"TestFilterOutByKeys_NilMap", nil, []int{1, 3}, nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FilterOutByKeys(c.input, c.keys...)
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the length of map is not equal")
			}
			for k, v := range actual {
				if v != c.expected[k] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the inputV of map is not equal")
				}
			}
		})
	}
}

```


</details>


***
#### func FilterOutByValues\[M ~map[K]V, K comparable, V any\](m M, values []V, handler ComparisonHandler[V]) M
<span id="FilterOutByValues"></span>
> 过滤 map 中多个 values，返回过滤后的 map

**示例代码：**

```go

func ExampleFilterOutByValues() {
	var m = map[string]int{"a": 1, "b": 2, "c": 3}
	var result = collection.FilterOutByValues(m, []int{1}, func(source, target int) bool {
		return source == target
	})
	for i, s := range []string{"a", "b", "c"} {
		fmt.Println(i, result[s])
	}
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFilterOutByValues(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		values   []int
		expected map[int]int
	}{{"TestFilterOutByValues_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, []int{1, 3}, map[int]int{2: 2}}, {"TestFilterOutByValues_EmptyMap", map[int]int{}, []int{1, 3}, map[int]int{}}, {"TestFilterOutByValues_NilMap", nil, []int{1, 3}, nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FilterOutByValues(c.input, c.values, func(source, target int) bool {
				return source == target
			})
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the length of map is not equal")
			}
			for k, v := range actual {
				if v != c.expected[k] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the inputV of map is not equal")
				}
			}
		})
	}
}

```


</details>


***
#### func FilterOutByMap\[M ~map[K]V, K comparable, V any\](m M, condition func (k K, v V)  bool) M
<span id="FilterOutByMap"></span>
> 过滤 map 中符合条件的元素，返回过滤后的 map
>   - condition 的返回值为 true 时，将会过滤掉该元素

**示例代码：**

```go

func ExampleFilterOutByMap() {
	var m = map[string]int{"a": 1, "b": 2, "c": 3}
	var result = collection.FilterOutByMap(m, func(k string, v int) bool {
		return k == "a" || v == 3
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFilterOutByMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		filter   map[int]int
		expected map[int]int
	}{{"TestFilterOutByMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]int{1: 1, 3: 3}, map[int]int{2: 2}}, {"TestFilterOutByMap_EmptyMap", map[int]int{}, map[int]int{1: 1, 3: 3}, map[int]int{}}, {"TestFilterOutByMap_NilMap", nil, map[int]int{1: 1, 3: 3}, nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FilterOutByMap(c.input, func(k int, v int) bool {
				return c.filter[k] == v
			})
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the length of map is not equal")
			}
			for k, v := range actual {
				if v != c.expected[k] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "after filter, the inputV of map is not equal")
				}
			}
		})
	}
}

```


</details>


***
#### func FindLoopedNextInSlice\[S ~[]V, V any\](slice S, i int) (next int, value V)
<span id="FindLoopedNextInSlice"></span>
> 返回 i 的下一个数组成员，当 i 达到数组长度时从 0 开始
>   - 当 i 为负数时将返回第一个元素

**示例代码：**

```go

func ExampleFindLoopedNextInSlice() {
	next, v := collection.FindLoopedNextInSlice([]int{1, 2, 3}, 1)
	fmt.Println(next, v)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindLoopedNextInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		i        int
		expected int
	}{{"TestFindLoopedNextInSlice_NonEmptySlice", []int{1, 2, 3}, 1, 2}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, _ := collection.FindLoopedNextInSlice(c.input, c.i)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the next index of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindLoopedPrevInSlice\[S ~[]V, V any\](slice S, i int) (prev int, value V)
<span id="FindLoopedPrevInSlice"></span>
> 返回 i 的上一个数组成员，当 i 为 0 时从数组末尾开始
>   - 当 i 为负数时将返回最后一个元素

**示例代码：**

```go

func ExampleFindLoopedPrevInSlice() {
	prev, v := collection.FindLoopedPrevInSlice([]int{1, 2, 3}, 1)
	fmt.Println(prev, v)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindLoopedPrevInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		i        int
		expected int
	}{{"TestFindLoopedPrevInSlice_NonEmptySlice", []int{1, 2, 3}, 1, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, _ := collection.FindLoopedPrevInSlice(c.input, c.i)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the prev index of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindCombinationsInSliceByRange\[S ~[]V, V any\](s S, minSize int, maxSize int) []S
<span id="FindCombinationsInSliceByRange"></span>
> 获取给定数组的所有组合，且每个组合的成员数量限制在指定范围内

**示例代码：**

```go

func ExampleFindCombinationsInSliceByRange() {
	result := collection.FindCombinationsInSliceByRange([]int{1, 2, 3}, 1, 2)
	fmt.Println(len(result))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindCombinationsInSliceByRange(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		minSize  int
		maxSize  int
		expected [][]int
	}{{"TestFindCombinationsInSliceByRange_NonEmptySlice", []int{1, 2, 3}, 1, 2, [][]int{{1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindCombinationsInSliceByRange(c.input, c.minSize, c.maxSize)
			if len(actual) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindFirstOrDefaultInSlice\[S ~[]V, V any\](slice S, defaultValue V) V
<span id="FindFirstOrDefaultInSlice"></span>
> 判断切片中是否存在元素，返回第一个元素，不存在则返回默认值

**示例代码：**

```go

func ExampleFindFirstOrDefaultInSlice() {
	result := collection.FindFirstOrDefaultInSlice([]int{1, 2, 3}, 0)
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindFirstOrDefaultInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{{"TestFindFirstOrDefaultInSlice_NonEmptySlice", []int{1, 2, 3}, 1}, {"TestFindFirstOrDefaultInSlice_EmptySlice", []int{}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindFirstOrDefaultInSlice(c.input, 0)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the first element of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindOrDefaultInSlice\[S ~[]V, V any\](slice S, defaultValue V, handler func (v V)  bool) (t V)
<span id="FindOrDefaultInSlice"></span>
> 判断切片中是否存在某个元素，返回第一个匹配的索引和元素，不存在则返回默认值

**示例代码：**

```go

func ExampleFindOrDefaultInSlice() {
	result := collection.FindOrDefaultInSlice([]int{1, 2, 3}, 0, func(v int) bool {
		return v == 2
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindOrDefaultInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{{"TestFindOrDefaultInSlice_NonEmptySlice", []int{1, 2, 3}, 2}, {"TestFindOrDefaultInSlice_EmptySlice", []int{}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindOrDefaultInSlice(c.input, 0, func(v int) bool {
				return v == 2
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the element of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindOrDefaultInComparableSlice\[S ~[]V, V comparable\](slice S, v V, defaultValue V) (t V)
<span id="FindOrDefaultInComparableSlice"></span>
> 判断切片中是否存在某个元素，返回第一个匹配的索引和元素，不存在则返回默认值

**示例代码：**

```go

func ExampleFindOrDefaultInComparableSlice() {
	result := collection.FindOrDefaultInComparableSlice([]int{1, 2, 3}, 2, 0)
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindOrDefaultInComparableSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{{"TestFindOrDefaultInComparableSlice_NonEmptySlice", []int{1, 2, 3}, 2}, {"TestFindOrDefaultInComparableSlice_EmptySlice", []int{}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindOrDefaultInComparableSlice(c.input, 2, 0)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the element of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindInSlice\[S ~[]V, V any\](slice S, handler func (v V)  bool) (i int, t V)
<span id="FindInSlice"></span>
> 判断切片中是否存在某个元素，返回第一个匹配的索引和元素，不存在则索引返回 -1

**示例代码：**

```go

func ExampleFindInSlice() {
	_, result := collection.FindInSlice([]int{1, 2, 3}, func(v int) bool {
		return v == 2
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{{"TestFindInSlice_NonEmptySlice", []int{1, 2, 3}, 2}, {"TestFindInSlice_EmptySlice", []int{}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, actual := collection.FindInSlice(c.input, func(v int) bool {
				return v == 2
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the element of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindIndexInSlice\[S ~[]V, V any\](slice S, handler func (v V)  bool) int
<span id="FindIndexInSlice"></span>
> 判断切片中是否存在某个元素，返回第一个匹配的索引，不存在则索引返回 -1

**示例代码：**

```go

func ExampleFindIndexInSlice() {
	result := collection.FindIndexInSlice([]int{1, 2, 3}, func(v int) bool {
		return v == 2
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindIndexInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{{"TestFindIndexInSlice_NonEmptySlice", []int{1, 2, 3}, 1}, {"TestFindIndexInSlice_EmptySlice", []int{}, -1}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindIndexInSlice(c.input, func(v int) bool {
				return v == 2
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the index of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindInComparableSlice\[S ~[]V, V comparable\](slice S, v V) (i int, t V)
<span id="FindInComparableSlice"></span>
> 判断切片中是否存在某个元素，返回第一个匹配的索引和元素，不存在则索引返回 -1

**示例代码：**

```go

func ExampleFindInComparableSlice() {
	index, result := collection.FindInComparableSlice([]int{1, 2, 3}, 2)
	fmt.Println(index, result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindInComparableSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{{"TestFindInComparableSlice_NonEmptySlice", []int{1, 2, 3}, 2}, {"TestFindInComparableSlice_EmptySlice", []int{}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, actual := collection.FindInComparableSlice(c.input, 2)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the element of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindIndexInComparableSlice\[S ~[]V, V comparable\](slice S, v V) int
<span id="FindIndexInComparableSlice"></span>
> 判断切片中是否存在某个元素，返回第一个匹配的索引，不存在则索引返回 -1

**示例代码：**

```go

func ExampleFindIndexInComparableSlice() {
	result := collection.FindIndexInComparableSlice([]int{1, 2, 3}, 2)
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindIndexInComparableSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{{"TestFindIndexInComparableSlice_NonEmptySlice", []int{1, 2, 3}, 1}, {"TestFindIndexInComparableSlice_EmptySlice", []int{}, -1}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindIndexInComparableSlice(c.input, 2)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the index of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindMinimumInComparableSlice\[S ~[]V, V generic.Ordered\](slice S) (result V)
<span id="FindMinimumInComparableSlice"></span>
> 获取切片中的最小值

**示例代码：**

```go

func ExampleFindMinimumInComparableSlice() {
	result := collection.FindMinimumInComparableSlice([]int{1, 2, 3})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindMinimumInComparableSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{{"TestFindMinimumInComparableSlice_NonEmptySlice", []int{1, 2, 3}, 1}, {"TestFindMinimumInComparableSlice_EmptySlice", []int{}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindMinimumInComparableSlice(c.input)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the minimum of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindMinimumInSlice\[S ~[]V, V any, N generic.Ordered\](slice S, handler OrderedValueGetter[V, N]) (result V)
<span id="FindMinimumInSlice"></span>
> 获取切片中的最小值

**示例代码：**

```go

func ExampleFindMinimumInSlice() {
	result := collection.FindMinimumInSlice([]int{1, 2, 3}, func(v int) int {
		return v
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindMinimumInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{{"TestFindMinimumInSlice_NonEmptySlice", []int{1, 2, 3}, 1}, {"TestFindMinimumInSlice_EmptySlice", []int{}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindMinimumInSlice(c.input, func(v int) int {
				return v
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the minimum of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindMaximumInComparableSlice\[S ~[]V, V generic.Ordered\](slice S) (result V)
<span id="FindMaximumInComparableSlice"></span>
> 获取切片中的最大值

**示例代码：**

```go

func ExampleFindMaximumInComparableSlice() {
	result := collection.FindMaximumInComparableSlice([]int{1, 2, 3})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindMaximumInComparableSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{{"TestFindMaximumInComparableSlice_NonEmptySlice", []int{1, 2, 3}, 3}, {"TestFindMaximumInComparableSlice_EmptySlice", []int{}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindMaximumInComparableSlice(c.input)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the maximum of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindMaximumInSlice\[S ~[]V, V any, N generic.Ordered\](slice S, handler OrderedValueGetter[V, N]) (result V)
<span id="FindMaximumInSlice"></span>
> 获取切片中的最大值

**示例代码：**

```go

func ExampleFindMaximumInSlice() {
	result := collection.FindMaximumInSlice([]int{1, 2, 3}, func(v int) int {
		return v
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindMaximumInSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{{"TestFindMaximumInSlice_NonEmptySlice", []int{1, 2, 3}, 3}, {"TestFindMaximumInSlice_EmptySlice", []int{}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindMaximumInSlice(c.input, func(v int) int {
				return v
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the maximum of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindMin2MaxInComparableSlice\[S ~[]V, V generic.Ordered\](slice S) (min V, max V)
<span id="FindMin2MaxInComparableSlice"></span>
> 获取切片中的最小值和最大值

**示例代码：**

```go

func ExampleFindMin2MaxInComparableSlice() {
	minimum, maximum := collection.FindMin2MaxInComparableSlice([]int{1, 2, 3})
	fmt.Println(minimum, maximum)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindMin2MaxInComparableSlice(t *testing.T) {
	var cases = []struct {
		name        string
		input       []int
		expectedMin int
		expectedMax int
	}{{"TestFindMin2MaxInComparableSlice_NonEmptySlice", []int{1, 2, 3}, 1, 3}, {"TestFindMin2MaxInComparableSlice_EmptySlice", []int{}, 0, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			minimum, maximum := collection.FindMin2MaxInComparableSlice(c.input)
			if minimum != c.expectedMin || maximum != c.expectedMax {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expectedMin, minimum, "the minimum of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindMin2MaxInSlice\[S ~[]V, V any, N generic.Ordered\](slice S, handler OrderedValueGetter[V, N]) (min V, max V)
<span id="FindMin2MaxInSlice"></span>
> 获取切片中的最小值和最大值

**示例代码：**

```go

func ExampleFindMin2MaxInSlice() {
	minimum, maximum := collection.FindMin2MaxInSlice([]int{1, 2, 3}, func(v int) int {
		return v
	})
	fmt.Println(minimum, maximum)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindMin2MaxInSlice(t *testing.T) {
	var cases = []struct {
		name        string
		input       []int
		expectedMin int
		expectedMax int
	}{{"TestFindMin2MaxInSlice_NonEmptySlice", []int{1, 2, 3}, 1, 3}, {"TestFindMin2MaxInSlice_EmptySlice", []int{}, 0, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			minimum, maximum := collection.FindMin2MaxInSlice(c.input, func(v int) int {
				return v
			})
			if minimum != c.expectedMin || maximum != c.expectedMax {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expectedMin, minimum, "the minimum of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindMinFromComparableMap\[M ~map[K]V, K comparable, V generic.Ordered\](m M) (result V)
<span id="FindMinFromComparableMap"></span>
> 获取 map 中的最小值

**示例代码：**

```go

func ExampleFindMinFromComparableMap() {
	result := collection.FindMinFromComparableMap(map[int]int{1: 1, 2: 2, 3: 3})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindMinFromComparableMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected int
	}{{"TestFindMinFromComparableMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, 1}, {"TestFindMinFromComparableMap_EmptyMap", map[int]int{}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindMinFromComparableMap(c.input)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the minimum of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindMinFromMap\[M ~map[K]V, K comparable, V any, N generic.Ordered\](m M, handler OrderedValueGetter[V, N]) (result V)
<span id="FindMinFromMap"></span>
> 获取 map 中的最小值

**示例代码：**

```go

func ExampleFindMinFromMap() {
	result := collection.FindMinFromMap(map[int]int{1: 1, 2: 2, 3: 3}, func(v int) int {
		return v
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindMinFromMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected int
	}{{"TestFindMinFromMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, 1}, {"TestFindMinFromMap_EmptyMap", map[int]int{}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindMinFromMap(c.input, func(v int) int {
				return v
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the minimum of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindMaxFromComparableMap\[M ~map[K]V, K comparable, V generic.Ordered\](m M) (result V)
<span id="FindMaxFromComparableMap"></span>
> 获取 map 中的最大值

**示例代码：**

```go

func ExampleFindMaxFromComparableMap() {
	result := collection.FindMaxFromComparableMap(map[int]int{1: 1, 2: 2, 3: 3})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindMaxFromComparableMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected int
	}{{"TestFindMaxFromComparableMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, 3}, {"TestFindMaxFromComparableMap_EmptyMap", map[int]int{}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindMaxFromComparableMap(c.input)
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the maximum of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindMaxFromMap\[M ~map[K]V, K comparable, V any, N generic.Ordered\](m M, handler OrderedValueGetter[V, N]) (result V)
<span id="FindMaxFromMap"></span>
> 获取 map 中的最大值

**示例代码：**

```go

func ExampleFindMaxFromMap() {
	result := collection.FindMaxFromMap(map[int]int{1: 1, 2: 2, 3: 3}, func(v int) int {
		return v
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindMaxFromMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected int
	}{{"TestFindMaxFromMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, 3}, {"TestFindMaxFromMap_EmptyMap", map[int]int{}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.FindMaxFromMap(c.input, func(v int) int {
				return v
			})
			if actual != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, actual, "the maximum of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindMin2MaxFromComparableMap\[M ~map[K]V, K comparable, V generic.Ordered\](m M) (min V, max V)
<span id="FindMin2MaxFromComparableMap"></span>
> 获取 map 中的最小值和最大值

**示例代码：**

```go

func ExampleFindMin2MaxFromComparableMap() {
	minimum, maximum := collection.FindMin2MaxFromComparableMap(map[int]int{1: 1, 2: 2, 3: 3})
	fmt.Println(minimum, maximum)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindMin2MaxFromComparableMap(t *testing.T) {
	var cases = []struct {
		name        string
		input       map[int]int
		expectedMin int
		expectedMax int
	}{{"TestFindMin2MaxFromComparableMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, 1, 3}, {"TestFindMin2MaxFromComparableMap_EmptyMap", map[int]int{}, 0, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			minimum, maximum := collection.FindMin2MaxFromComparableMap(c.input)
			if minimum != c.expectedMin || maximum != c.expectedMax {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expectedMin, minimum, "the minimum of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func FindMin2MaxFromMap\[M ~map[K]V, K comparable, V generic.Ordered\](m M) (min V, max V)
<span id="FindMin2MaxFromMap"></span>
> 获取 map 中的最小值和最大值

**示例代码：**

```go

func ExampleFindMin2MaxFromMap() {
	minimum, maximum := collection.FindMin2MaxFromMap(map[int]int{1: 1, 2: 2, 3: 3})
	fmt.Println(minimum, maximum)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestFindMin2MaxFromMap(t *testing.T) {
	var cases = []struct {
		name        string
		input       map[int]int
		expectedMin int
		expectedMax int
	}{{"TestFindMin2MaxFromMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, 1, 3}, {"TestFindMin2MaxFromMap_EmptyMap", map[int]int{}, 0, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			minimum, maximum := collection.FindMin2MaxFromMap(c.input)
			if minimum != c.expectedMin || maximum != c.expectedMax {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expectedMin, minimum, "the minimum of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func SwapSlice\[S ~[]V, V any\](slice *S, i int, j int)
<span id="SwapSlice"></span>
> 将切片中的两个元素进行交换

**示例代码：**

```go

func ExampleSwapSlice() {
	var s = []int{1, 2, 3}
	collection.SwapSlice(&s, 0, 1)
	fmt.Println(s)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestSwapSlice(t *testing.T) {
	var cases = []struct {
		name   string
		slice  []int
		i      int
		j      int
		expect []int
	}{{"TestSwapSliceNonEmpty", []int{1, 2, 3}, 0, 1, []int{2, 1, 3}}, {"TestSwapSliceEmpty", []int{}, 0, 0, []int{}}, {"TestSwapSliceIndexOutOfBound", []int{1, 2, 3}, 0, 3, []int{1, 2, 3}}, {"TestSwapSliceNil", nil, 0, 0, nil}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.SwapSlice(&c.slice, c.i, c.j)
			for i, v := range c.slice {
				if v != c.expect[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expect, c.slice, "the slice is not equal")
				}
			}
		})
	}
}

```


</details>


***
#### func MappingFromSlice\[S ~[]V, NS []N, V any, N any\](slice S, handler func (value V)  N) NS
<span id="MappingFromSlice"></span>
> 将切片中的元素进行转换

**示例代码：**

```go

func ExampleMappingFromSlice() {
	result := collection.MappingFromSlice([]int{1, 2, 3}, func(value int) int {
		return value + 1
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestMappingFromSlice(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{{"TestMappingFromSlice_NonEmptySlice", []int{1, 2, 3}, []int{2, 3, 4}}, {"TestMappingFromSlice_EmptySlice", []int{}, []int{}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.MappingFromSlice[[]int, []int](c.input, func(value int) int {
				return value + 1
			})
			if len(result) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
			for i := 0; i < len(result); i++ {
				if result[i] != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the value of input is not equal")
				}
			}
		})
	}
}

```


</details>


***
#### func MappingFromMap\[M ~map[K]V, NM map[K]N, K comparable, V any, N any\](m M, handler func (value V)  N) NM
<span id="MappingFromMap"></span>
> 将 map 中的元素进行转换

**示例代码：**

```go

func ExampleMappingFromMap() {
	result := collection.MappingFromMap(map[int]int{1: 1, 2: 2, 3: 3}, func(value int) int {
		return value + 1
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestMappingFromMap(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]int
	}{{"TestMappingFromMap_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]int{1: 2, 2: 3, 3: 4}}, {"TestMappingFromMap_EmptyMap", map[int]int{}, map[int]int{}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.MappingFromMap[map[int]int, map[int]int](c.input, func(value int) int {
				return value + 1
			})
			if len(result) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
			for k, v := range result {
				if v != c.expected[k] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the value of input is not equal")
				}
			}
		})
	}
}

```


</details>


***
#### func MergeSlices\[S ~[]V, V any\](slices ...S) (result S)
<span id="MergeSlices"></span>
> 合并切片

**示例代码：**

```go

func ExampleMergeSlices() {
	fmt.Println(collection.MergeSlices([]int{1, 2, 3}, []int{4, 5, 6}))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestMergeSlices(t *testing.T) {
	var cases = []struct {
		name     string
		input    [][]int
		expected []int
	}{{"TestMergeSlices_NonEmptySlice", [][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2, 3, 4, 5, 6}}, {"TestMergeSlices_EmptySlice", [][]int{{}, {}}, []int{}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.MergeSlices(c.input...)
			if len(result) != len(c.expected) {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
			for i := 0; i < len(result); i++ {
				if result[i] != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the value of input is not equal")
				}
			}
		})
	}
}

```


</details>


***
#### func MergeMaps\[M ~map[K]V, K comparable, V any\](maps ...M) (result M)
<span id="MergeMaps"></span>
> 合并 map，当多个 map 中存在相同的 key 时，后面的 map 中的 key 将会覆盖前面的 map 中的 key

**示例代码：**

```go

func ExampleMergeMaps() {
	m := collection.MergeMaps(map[int]int{1: 1, 2: 2, 3: 3}, map[int]int{4: 4, 5: 5, 6: 6})
	fmt.Println(len(m))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestMergeMaps(t *testing.T) {
	var cases = []struct {
		name     string
		input    []map[int]int
		expected int
	}{{"TestMergeMaps_NonEmptyMap", []map[int]int{{1: 1, 2: 2, 3: 3}, {4: 4, 5: 5, 6: 6}}, 6}, {"TestMergeMaps_EmptyMap", []map[int]int{{}, {}}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.MergeMaps(c.input...)
			if len(result) != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func MergeMapsWithSkip\[M ~map[K]V, K comparable, V any\](maps ...M) (result M)
<span id="MergeMapsWithSkip"></span>
> 合并 map，当多个 map 中存在相同的 key 时，后面的 map 中的 key 将会被跳过

**示例代码：**

```go

func ExampleMergeMapsWithSkip() {
	m := collection.MergeMapsWithSkip(map[int]int{1: 1}, map[int]int{1: 2})
	fmt.Println(m[1])
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestMergeMapsWithSkip(t *testing.T) {
	var cases = []struct {
		name     string
		input    []map[int]int
		expected int
	}{{"TestMergeMapsWithSkip_NonEmptyMap", []map[int]int{{1: 1}, {1: 2}}, 1}, {"TestMergeMapsWithSkip_EmptyMap", []map[int]int{{}, {}}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.MergeMapsWithSkip(c.input...)
			if len(result) != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func ChooseRandomSliceElementRepeatN\[S ~[]V, V any\](slice S, n int) (result []V)
<span id="ChooseRandomSliceElementRepeatN"></span>
> 返回 slice 中的 n 个可重复随机元素
>   - 当 slice 长度为 0 或 n 小于等于 0 时将会返回 nil

**示例代码：**

```go

func ExampleChooseRandomSliceElementRepeatN() {
	result := collection.ChooseRandomSliceElementRepeatN([]int{1}, 10)
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestChooseRandomSliceElementRepeatN(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{{"TestChooseRandomSliceElementRepeatN_NonEmptySlice", []int{1, 2, 3}, 3}, {"TestChooseRandomSliceElementRepeatN_EmptySlice", []int{}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomSliceElementRepeatN(c.input, 3)
			if len(result) != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func ChooseRandomIndexRepeatN\[S ~[]V, V any\](slice S, n int) (result []int)
<span id="ChooseRandomIndexRepeatN"></span>
> 返回 slice 中的 n 个可重复随机元素的索引
>   - 当 slice 长度为 0 或 n 小于等于 0 时将会返回 nil

**示例代码：**

```go

func ExampleChooseRandomIndexRepeatN() {
	result := collection.ChooseRandomIndexRepeatN([]int{1}, 10)
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestChooseRandomIndexRepeatN(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{{"TestChooseRandomIndexRepeatN_NonEmptySlice", []int{1, 2, 3}, 3}, {"TestChooseRandomIndexRepeatN_EmptySlice", []int{}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomIndexRepeatN(c.input, 3)
			if len(result) != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func ChooseRandomSliceElement\[S ~[]V, V any\](slice S) (v V)
<span id="ChooseRandomSliceElement"></span>
> 返回 slice 中随机一个元素，当 slice 长度为 0 时将会得到 V 的零值

**示例代码：**

```go

func ExampleChooseRandomSliceElement() {
	result := collection.ChooseRandomSliceElement([]int{1})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestChooseRandomSliceElement(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected map[int]bool
	}{{"TestChooseRandomSliceElement_NonEmptySlice", []int{1, 2, 3}, map[int]bool{1: true, 2: true, 3: true}}, {"TestChooseRandomSliceElement_EmptySlice", []int{}, map[int]bool{0: true}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomSliceElement(c.input)
			if !c.expected[result] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func ChooseRandomIndex\[S ~[]V, V any\](slice S) (index int)
<span id="ChooseRandomIndex"></span>
> 返回 slice 中随机一个元素的索引，当 slice 长度为 0 时将会得到 -1

**示例代码：**

```go

func ExampleChooseRandomIndex() {
	result := collection.ChooseRandomIndex([]int{1})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestChooseRandomIndex(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected map[int]bool
	}{{"TestChooseRandomIndex_NonEmptySlice", []int{1, 2, 3}, map[int]bool{0: true, 1: true, 2: true}}, {"TestChooseRandomIndex_EmptySlice", []int{}, map[int]bool{-1: true}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomIndex(c.input)
			if !c.expected[result] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func ChooseRandomSliceElementN\[S ~[]V, V any\](slice S, n int) (result []V)
<span id="ChooseRandomSliceElementN"></span>
> 返回 slice 中的 n 个不可重复的随机元素
>   - 当 slice 长度为 0 或 n 大于 slice 长度或小于 0 时将会发生 panic

**示例代码：**

```go

func ExampleChooseRandomSliceElementN() {
	result := collection.ChooseRandomSliceElementN([]int{1}, 1)
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestChooseRandomSliceElementN(t *testing.T) {
	var cases = []struct {
		name  string
		input []int
	}{{"TestChooseRandomSliceElementN_NonEmptySlice", []int{1, 2, 3}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := collection.ChooseRandomSliceElementN(c.input, 3)
			if !collection.AllInComparableSlice(actual, c.input) {
				t.Fatalf("%s failed, actual: %v, error: %s", c.name, actual, "the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func ChooseRandomIndexN\[S ~[]V, V any\](slice S, n int) (result []int)
<span id="ChooseRandomIndexN"></span>
> 获取切片中的 n 个随机元素的索引
>   - 如果 n 大于切片长度或小于 0 时将会发生 panic

**示例代码：**

```go

func ExampleChooseRandomIndexN() {
	result := collection.ChooseRandomIndexN([]int{1}, 1)
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestChooseRandomIndexN(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected map[int]bool
	}{{"TestChooseRandomIndexN_NonEmptySlice", []int{1, 2, 3}, map[int]bool{0: true, 1: true, 2: true}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomIndexN(c.input, 3)
			if !c.expected[result[0]] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func ChooseRandomMapKeyRepeatN\[M ~map[K]V, K comparable, V any\](m M, n int) (result []K)
<span id="ChooseRandomMapKeyRepeatN"></span>
> 获取 map 中的 n 个随机 key，允许重复
>   - 如果 n 大于 map 长度或小于 0 时将会发生 panic

**示例代码：**

```go

func ExampleChooseRandomMapKeyRepeatN() {
	result := collection.ChooseRandomMapKeyRepeatN(map[int]int{1: 1}, 1)
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestChooseRandomMapKeyRepeatN(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]bool
	}{{"TestChooseRandomMapKeyRepeatN_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]bool{1: true, 2: true, 3: true}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomMapKeyRepeatN(c.input, 3)
			if !c.expected[result[0]] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func ChooseRandomMapValueRepeatN\[M ~map[K]V, K comparable, V any\](m M, n int) (result []V)
<span id="ChooseRandomMapValueRepeatN"></span>
> 获取 map 中的 n 个随机 n，允许重复
>   - 如果 n 大于 map 长度或小于 0 时将会发生 panic

**示例代码：**

```go

func ExampleChooseRandomMapValueRepeatN() {
	result := collection.ChooseRandomMapValueRepeatN(map[int]int{1: 1}, 1)
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestChooseRandomMapValueRepeatN(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]bool
	}{{"TestChooseRandomMapValueRepeatN_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]bool{1: true, 2: true, 3: true}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomMapValueRepeatN(c.input, 3)
			if !c.expected[result[0]] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func ChooseRandomMapKeyAndValueRepeatN\[M ~map[K]V, K comparable, V any\](m M, n int) M
<span id="ChooseRandomMapKeyAndValueRepeatN"></span>
> 获取 map 中的 n 个随机 key 和 v，允许重复
>   - 如果 n 大于 map 长度或小于 0 时将会发生 panic

**示例代码：**

```go

func ExampleChooseRandomMapKeyAndValueRepeatN() {
	result := collection.ChooseRandomMapKeyAndValueRepeatN(map[int]int{1: 1}, 1)
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestChooseRandomMapKeyAndValueRepeatN(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]bool
	}{{"TestChooseRandomMapKeyAndValueRepeatN_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]bool{1: true, 2: true, 3: true}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomMapKeyAndValueRepeatN(c.input, 3)
			if !c.expected[result[1]] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func ChooseRandomMapKey\[M ~map[K]V, K comparable, V any\](m M) (k K)
<span id="ChooseRandomMapKey"></span>
> 获取 map 中的随机 key

**示例代码：**

```go

func ExampleChooseRandomMapKey() {
	result := collection.ChooseRandomMapKey(map[int]int{1: 1})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestChooseRandomMapKey(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]bool
	}{{"TestChooseRandomMapKey_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]bool{1: true, 2: true, 3: true}}, {"TestChooseRandomMapKey_EmptyMap", map[int]int{}, map[int]bool{0: true}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomMapKey(c.input)
			if !c.expected[result] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func ChooseRandomMapValue\[M ~map[K]V, K comparable, V any\](m M) (v V)
<span id="ChooseRandomMapValue"></span>
> 获取 map 中的随机 value

**示例代码：**

```go

func ExampleChooseRandomMapValue() {
	result := collection.ChooseRandomMapValue(map[int]int{1: 1})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestChooseRandomMapValue(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]bool
	}{{"TestChooseRandomMapValue_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]bool{1: true, 2: true, 3: true}}, {"TestChooseRandomMapValue_EmptyMap", map[int]int{}, map[int]bool{0: true}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomMapValue(c.input)
			if !c.expected[result] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func ChooseRandomMapKeyN\[M ~map[K]V, K comparable, V any\](m M, n int) (result []K)
<span id="ChooseRandomMapKeyN"></span>
> 获取 map 中的 inputN 个随机 key
>   - 如果 inputN 大于 map 长度或小于 0 时将会发生 panic

**示例代码：**

```go

func ExampleChooseRandomMapKeyN() {
	result := collection.ChooseRandomMapKeyN(map[int]int{1: 1}, 1)
	fmt.Println(result)
}

```

***
#### func ChooseRandomMapValueN\[M ~map[K]V, K comparable, V any\](m M, n int) (result []V)
<span id="ChooseRandomMapValueN"></span>
> 获取 map 中的 n 个随机 value
>   - 如果 n 大于 map 长度或小于 0 时将会发生 panic

**示例代码：**

```go

func ExampleChooseRandomMapValueN() {
	result := collection.ChooseRandomMapValueN(map[int]int{1: 1}, 1)
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestChooseRandomMapValueN(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]bool
	}{{"TestChooseRandomMapValueN_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]bool{1: true, 2: true, 3: true}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ChooseRandomMapValueN(c.input, 3)
			if !c.expected[result[0]] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, result, "the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func ChooseRandomMapKeyAndValue\[M ~map[K]V, K comparable, V any\](m M) (k K, v V)
<span id="ChooseRandomMapKeyAndValue"></span>
> 获取 map 中的随机 key 和 v

**示例代码：**

```go

func ExampleChooseRandomMapKeyAndValue() {
	k, v := collection.ChooseRandomMapKeyAndValue(map[int]int{1: 1})
	fmt.Println(k, v)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestChooseRandomMapKeyAndValue(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]bool
	}{{"TestChooseRandomMapKeyAndValue_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]bool{1: true, 2: true, 3: true}}, {"TestChooseRandomMapKeyAndValue_EmptyMap", map[int]int{}, map[int]bool{0: true}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			k, v := collection.ChooseRandomMapKeyAndValue(c.input)
			if !c.expected[k] || !c.expected[v] {
				t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, k, "the length of input is not equal")
			}
		})
	}
}

```


</details>


***
#### func ChooseRandomMapKeyAndValueN\[M ~map[K]V, K comparable, V any\](m M, n int) M
<span id="ChooseRandomMapKeyAndValueN"></span>
> 获取 map 中的 inputN 个随机 key 和 v
>   - 如果 n 大于 map 长度或小于 0 时将会发生 panic

**示例代码：**

```go

func ExampleChooseRandomMapKeyAndValueN() {
	result := collection.ChooseRandomMapKeyAndValueN(map[int]int{1: 1}, 1)
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestChooseRandomMapKeyAndValueN(t *testing.T) {
	var cases = []struct {
		name     string
		input    map[int]int
		expected map[int]bool
	}{{"TestChooseRandomMapKeyAndValueN_NonEmptyMap", map[int]int{1: 1, 2: 2, 3: 3}, map[int]bool{1: true, 2: true, 3: true}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			kvm := collection.ChooseRandomMapKeyAndValueN(c.input, 1)
			for k := range kvm {
				if !c.expected[k] {
					t.Fatalf("%s failed, expected: %v, actual: %v, error: %s", c.name, c.expected, k, "the length of input is not equal")
				}
			}
		})
	}
}

```


</details>


***
#### func DescBy\[Sort generic.Ordered\](a Sort, b Sort) bool
<span id="DescBy"></span>
> 返回降序比较结果

**示例代码：**

```go

func ExampleDescBy() {
	var slice = []int{1, 2, 3}
	sort.Slice(slice, func(i, j int) bool {
		return collection.DescBy(slice[i], slice[j])
	})
	fmt.Println(slice)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestDescBy(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{{"TestDescBy_NonEmptySlice", []int{1, 2, 3}, []int{3, 2, 1}}, {"TestDescBy_EmptySlice", []int{}, []int{}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sort.Slice(c.input, func(i, j int) bool {
				return collection.DescBy(c.input[i], c.input[j])
			})
			for i, v := range c.input {
				if v != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v", c.name, c.expected, c.input)
				}
			}
		})
	}
}

```


</details>


***
#### func AscBy\[Sort generic.Ordered\](a Sort, b Sort) bool
<span id="AscBy"></span>
> 返回升序比较结果

**示例代码：**

```go

func ExampleAscBy() {
	var slice = []int{1, 2, 3}
	sort.Slice(slice, func(i, j int) bool {
		return collection.AscBy(slice[i], slice[j])
	})
	fmt.Println(slice)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAscBy(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{{"TestAscBy_NonEmptySlice", []int{1, 2, 3}, []int{1, 2, 3}}, {"TestAscBy_EmptySlice", []int{}, []int{}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sort.Slice(c.input, func(i, j int) bool {
				return collection.AscBy(c.input[i], c.input[j])
			})
			for i, v := range c.input {
				if v != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v", c.name, c.expected, c.input)
				}
			}
		})
	}
}

```


</details>


***
#### func Desc\[S ~[]V, V any, Sort generic.Ordered\](slice *S, getter func (index int)  Sort)
<span id="Desc"></span>
> 对切片进行降序排序

**示例代码：**

```go

func ExampleDesc() {
	var slice = []int{1, 2, 3}
	collection.Desc(&slice, func(index int) int {
		return slice[index]
	})
	fmt.Println(slice)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestDesc(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{{"TestDesc_NonEmptySlice", []int{1, 2, 3}, []int{3, 2, 1}}, {"TestDesc_EmptySlice", []int{}, []int{}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.Desc(&c.input, func(index int) int {
				return c.input[index]
			})
			for i, v := range c.input {
				if v != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v", c.name, c.expected, c.input)
				}
			}
		})
	}
}

```


</details>


***
#### func DescByClone\[S ~[]V, V any, Sort generic.Ordered\](slice S, getter func (index int)  Sort) S
<span id="DescByClone"></span>
> 对切片进行降序排序，返回排序后的切片

**示例代码：**

```go

func ExampleDescByClone() {
	var slice = []int{1, 2, 3}
	result := collection.DescByClone(slice, func(index int) int {
		return slice[index]
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestDescByClone(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{{"TestDescByClone_NonEmptySlice", []int{1, 2, 3}, []int{3, 2, 1}}, {"TestDescByClone_EmptySlice", []int{}, []int{}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.DescByClone(c.input, func(index int) int {
				return c.input[index]
			})
			for i, v := range result {
				if v != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v", c.name, c.expected, result)
				}
			}
		})
	}
}

```


</details>


***
#### func Asc\[S ~[]V, V any, Sort generic.Ordered\](slice *S, getter func (index int)  Sort)
<span id="Asc"></span>
> 对切片进行升序排序

**示例代码：**

```go

func ExampleAsc() {
	var slice = []int{1, 2, 3}
	collection.Asc(&slice, func(index int) int {
		return slice[index]
	})
	fmt.Println(slice)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAsc(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{{"TestAsc_NonEmptySlice", []int{1, 2, 3}, []int{1, 2, 3}}, {"TestAsc_EmptySlice", []int{}, []int{}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.Asc(&c.input, func(index int) int {
				return c.input[index]
			})
			for i, v := range c.input {
				if v != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v", c.name, c.expected, c.input)
				}
			}
		})
	}
}

```


</details>


***
#### func AscByClone\[S ~[]V, V any, Sort generic.Ordered\](slice S, getter func (index int)  Sort) S
<span id="AscByClone"></span>
> 对切片进行升序排序，返回排序后的切片

**示例代码：**

```go

func ExampleAscByClone() {
	var slice = []int{1, 2, 3}
	result := collection.AscByClone(slice, func(index int) int {
		return slice[index]
	})
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestAscByClone(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{{"TestAscByClone_NonEmptySlice", []int{1, 2, 3}, []int{1, 2, 3}}, {"TestAscByClone_EmptySlice", []int{}, []int{}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.AscByClone(c.input, func(index int) int {
				return c.input[index]
			})
			for i, v := range result {
				if v != c.expected[i] {
					t.Fatalf("%s failed, expected: %v, actual: %v", c.name, c.expected, result)
				}
			}
		})
	}
}

```


</details>


***
#### func Shuffle\[S ~[]V, V any\](slice *S)
<span id="Shuffle"></span>
> 对切片进行随机排序

**示例代码：**

```go

func ExampleShuffle() {
	var slice = []int{1, 2, 3}
	collection.Shuffle(&slice)
	fmt.Println(len(slice))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestShuffle(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{{"TestShuffle_NonEmptySlice", []int{1, 2, 3}, 3}, {"TestShuffle_EmptySlice", []int{}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collection.Shuffle(&c.input)
			if len(c.input) != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v", c.name, c.expected, c.input)
			}
		})
	}
}

```


</details>


***
#### func ShuffleByClone\[S ~[]V, V any\](slice S) S
<span id="ShuffleByClone"></span>
> 对切片进行随机排序，返回排序后的切片

**示例代码：**

```go

func ExampleShuffleByClone() {
	var slice = []int{1, 2, 3}
	result := collection.ShuffleByClone(slice)
	fmt.Println(len(result))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestShuffleByClone(t *testing.T) {
	var cases = []struct {
		name     string
		input    []int
		expected int
	}{{"TestShuffleByClone_NonEmptySlice", []int{1, 2, 3}, 3}, {"TestShuffleByClone_EmptySlice", []int{}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := collection.ShuffleByClone(c.input)
			if len(result) != c.expected {
				t.Fatalf("%s failed, expected: %v, actual: %v", c.name, c.expected, result)
			}
		})
	}
}

```


</details>


***
<span id="struct_ComparisonHandler"></span>
### ComparisonHandler `STRUCT`
用于比较 `source` 和 `target` 两个值是否相同的比较函数
  - 该函数接受两个参数，分别是源值和目标值，返回 true 的情况下即表示两者相同
```go
type ComparisonHandler[V any] func(source V) bool
```
<span id="struct_OrderedValueGetter"></span>
### OrderedValueGetter `STRUCT`
用于获取 v 的可排序字段值的函数
```go
type OrderedValueGetter[V any, N generic.Ordered] func(v V) N
```
