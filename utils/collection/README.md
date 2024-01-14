# Collection

collection 用于对 input 和 map 操作的工具函数

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/collection)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[CloneSlice](#CloneSlice)|克隆切片，该函数是 slices.Clone 的快捷方式
|[CloneMap](#CloneMap)|克隆 map
|[CloneSliceN](#CloneSliceN)|克隆 slice 为 n 个切片进行返回
|[CloneMapN](#CloneMapN)|克隆 map 为 n 个 map 进行返回
|[CloneSlices](#CloneSlices)|克隆多个切片
|[CloneMaps](#CloneMaps)|克隆多个 map
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


> 结构体定义

|结构体|描述
|:--|:--
|[ComparisonHandler](#comparisonhandler)|暂无描述...
|[OrderedValueGetter](#orderedvaluegetter)|暂无描述...

</details>


#### func CloneSlice(slice S)  S
<span id="CloneSlice"></span>
> 克隆切片，该函数是 slices.Clone 的快捷方式
***
#### func CloneMap(m M)  M
<span id="CloneMap"></span>
> 克隆 map
***
#### func CloneSliceN(slice S, n int)  []S
<span id="CloneSliceN"></span>
> 克隆 slice 为 n 个切片进行返回
***
#### func CloneMapN(m M, n int)  []M
<span id="CloneMapN"></span>
> 克隆 map 为 n 个 map 进行返回
***
#### func CloneSlices(slices ...S)  []S
<span id="CloneSlices"></span>
> 克隆多个切片
***
#### func CloneMaps(maps ...M)  []M
<span id="CloneMaps"></span>
> 克隆多个 map
***
#### func InSlice(slice S, v V, handler ComparisonHandler[V])  bool
<span id="InSlice"></span>
> 检查 v 是否被包含在 slice 中，当 handler 返回 true 时，表示 v 和 slice 中的某个元素相匹配
***
#### func InComparableSlice(slice S, v V)  bool
<span id="InComparableSlice"></span>
> 检查 v 是否被包含在 slice 中
***
#### func AllInSlice(slice S, values []V, handler ComparisonHandler[V])  bool
<span id="AllInSlice"></span>
> 检查 values 中的所有元素是否均被包含在 slice 中，当 handler 返回 true 时，表示 values 中的某个元素和 slice 中的某个元素相匹配
>   - 在所有 values 中的元素都被包含在 slice 中时，返回 true
>   - 当 values 长度为 0 或为 nil 时，将返回 true
***
#### func AllInComparableSlice(slice S, values []V)  bool
<span id="AllInComparableSlice"></span>
> 检查 values 中的所有元素是否均被包含在 slice 中
>   - 在所有 values 中的元素都被包含在 slice 中时，返回 true
>   - 当 values 长度为 0 或为 nil 时，将返回 true
***
#### func AnyInSlice(slice S, values []V, handler ComparisonHandler[V])  bool
<span id="AnyInSlice"></span>
> 检查 values 中的任意一个元素是否被包含在 slice 中，当 handler 返回 true 时，表示 value 中的某个元素和 slice 中的某个元素相匹配
>   - 当 values 中的任意一个元素被包含在 slice 中时，返回 true
***
#### func AnyInComparableSlice(slice S, values []V)  bool
<span id="AnyInComparableSlice"></span>
> 检查 values 中的任意一个元素是否被包含在 slice 中
>   - 当 values 中的任意一个元素被包含在 slice 中时，返回 true
***
#### func InSlices(slices []S, v V, handler ComparisonHandler[V])  bool
<span id="InSlices"></span>
> 通过将多个切片合并后检查 v 是否被包含在 slices 中，当 handler 返回 true 时，表示 v 和 slices 中的某个元素相匹配
>   - 当传入的 v 被包含在 slices 的任一成员中时，返回 true
***
#### func InComparableSlices(slices []S, v V)  bool
<span id="InComparableSlices"></span>
> 通过将多个切片合并后检查 v 是否被包含在 slices 中
>   - 当传入的 v 被包含在 slices 的任一成员中时，返回 true
***
#### func AllInSlices(slices []S, values []V, handler ComparisonHandler[V])  bool
<span id="AllInSlices"></span>
> 通过将多个切片合并后检查 values 中的所有元素是否被包含在 slices 中，当 handler 返回 true 时，表示 values 中的某个元素和 slices 中的某个元素相匹配
>   - 当 values 中的所有元素都被包含在 slices 的任一成员中时，返回 true
***
#### func AllInComparableSlices(slices []S, values []V)  bool
<span id="AllInComparableSlices"></span>
> 通过将多个切片合并后检查 values 中的所有元素是否被包含在 slices 中
>   - 当 values 中的所有元素都被包含在 slices 的任一成员中时，返回 true
***
#### func AnyInSlices(slices []S, values []V, handler ComparisonHandler[V])  bool
<span id="AnyInSlices"></span>
> 通过将多个切片合并后检查 values 中的任意一个元素是否被包含在 slices 中，当 handler 返回 true 时，表示 values 中的某个元素和 slices 中的某个元素相匹配
>   - 当 values 中的任意一个元素被包含在 slices 的任一成员中时，返回 true
***
#### func AnyInComparableSlices(slices []S, values []V)  bool
<span id="AnyInComparableSlices"></span>
> 通过将多个切片合并后检查 values 中的任意一个元素是否被包含在 slices 中
>   - 当 values 中的任意一个元素被包含在 slices 的任一成员中时，返回 true
***
#### func InAllSlices(slices []S, v V, handler ComparisonHandler[V])  bool
<span id="InAllSlices"></span>
> 检查 v 是否被包含在 slices 的每一项元素中，当 handler 返回 true 时，表示 v 和 slices 中的某个元素相匹配
>   - 当 v 被包含在 slices 的每一项元素中时，返回 true
***
#### func InAllComparableSlices(slices []S, v V)  bool
<span id="InAllComparableSlices"></span>
> 检查 v 是否被包含在 slices 的每一项元素中
>   - 当 v 被包含在 slices 的每一项元素中时，返回 true
***
#### func AnyInAllSlices(slices []S, values []V, handler ComparisonHandler[V])  bool
<span id="AnyInAllSlices"></span>
> 检查 slices 中的每一个元素是否均包含至少任意一个 values 中的元素，当 handler 返回 true 时，表示 value 中的某个元素和 slices 中的某个元素相匹配
>   - 当 slices 中的每一个元素均包含至少任意一个 values 中的元素时，返回 true
***
#### func AnyInAllComparableSlices(slices []S, values []V)  bool
<span id="AnyInAllComparableSlices"></span>
> 检查 slices 中的每一个元素是否均包含至少任意一个 values 中的元素
>   - 当 slices 中的每一个元素均包含至少任意一个 values 中的元素时，返回 true
***
#### func KeyInMap(m M, key K)  bool
<span id="KeyInMap"></span>
> 检查 m 中是否包含特定 key
***
#### func ValueInMap(m M, value V, handler ComparisonHandler[V])  bool
<span id="ValueInMap"></span>
> 检查 m 中是否包含特定 value，当 handler 返回 true 时，表示 value 和 m 中的某个元素相匹配
***
#### func AllKeyInMap(m M, keys ...K)  bool
<span id="AllKeyInMap"></span>
> 检查 m 中是否包含 keys 中所有的元素
***
#### func AllValueInMap(m M, values []V, handler ComparisonHandler[V])  bool
<span id="AllValueInMap"></span>
> 检查 m 中是否包含 values 中所有的元素，当 handler 返回 true 时，表示 values 中的某个元素和 m 中的某个元素相匹配
***
#### func AnyKeyInMap(m M, keys ...K)  bool
<span id="AnyKeyInMap"></span>
> 检查 m 中是否包含 keys 中任意一个元素
***
#### func AnyValueInMap(m M, values []V, handler ComparisonHandler[V])  bool
<span id="AnyValueInMap"></span>
> 检查 m 中是否包含 values 中任意一个元素，当 handler 返回 true 时，表示 values 中的某个元素和 m 中的某个元素相匹配
***
#### func AllKeyInMaps(maps []M, keys ...K)  bool
<span id="AllKeyInMaps"></span>
> 检查 maps 中的每一个元素是否均包含 keys 中所有的元素
***
#### func AllValueInMaps(maps []M, values []V, handler ComparisonHandler[V])  bool
<span id="AllValueInMaps"></span>
> 检查 maps 中的每一个元素是否均包含 value 中所有的元素，当 handler 返回 true 时，表示 value 中的某个元素和 maps 中的某个元素相匹配
***
#### func AnyKeyInMaps(maps []M, keys ...K)  bool
<span id="AnyKeyInMaps"></span>
> 检查 keys 中的任意一个元素是否被包含在 maps 中的任意一个元素中
>   - 当 keys 中的任意一个元素被包含在 maps 中的任意一个元素中时，返回 true
***
#### func AnyValueInMaps(maps []M, values []V, handler ComparisonHandler[V])  bool
<span id="AnyValueInMaps"></span>
> 检查 maps 中的任意一个元素是否包含 value 中的任意一个元素，当 handler 返回 true 时，表示 value 中的某个元素和 maps 中的某个元素相匹配
>   - 当 maps 中的任意一个元素包含 value 中的任意一个元素时，返回 true
***
#### func KeyInAllMaps(maps []M, key K)  bool
<span id="KeyInAllMaps"></span>
> 检查 key 是否被包含在 maps 的每一个元素中
***
#### func AnyKeyInAllMaps(maps []M, keys []K)  bool
<span id="AnyKeyInAllMaps"></span>
> 检查 maps 中的每一个元素是否均包含 keys 中任意一个元素
>   - 当 maps 中的每一个元素均包含 keys 中任意一个元素时，返回 true
***
#### func ConvertSliceToAny(s S)  []any
<span id="ConvertSliceToAny"></span>
> 将切片转换为任意类型的切片
***
#### func ConvertSliceToIndexMap(s S)  map[int]V
<span id="ConvertSliceToIndexMap"></span>
> 将切片转换为索引为键的映射
***
#### func ConvertSliceToIndexOnlyMap(s S)  map[int]struct {}
<span id="ConvertSliceToIndexOnlyMap"></span>
> 将切片转换为索引为键的映射
***
#### func ConvertSliceToMap(s S)  map[V]struct {}
<span id="ConvertSliceToMap"></span>
> 将切片转换为值为键的映射
***
#### func ConvertSliceToBoolMap(s S)  map[V]bool
<span id="ConvertSliceToBoolMap"></span>
> 将切片转换为值为键的映射
***
#### func ConvertMapKeysToSlice(m M)  []K
<span id="ConvertMapKeysToSlice"></span>
> 将映射的键转换为切片
***
#### func ConvertMapValuesToSlice(m M)  []V
<span id="ConvertMapValuesToSlice"></span>
> 将映射的值转换为切片
***
#### func InvertMap(m M)  N
<span id="InvertMap"></span>
> 将映射的键和值互换
***
#### func ConvertMapValuesToBool(m M)  N
<span id="ConvertMapValuesToBool"></span>
> 将映射的值转换为布尔值
***
#### func ReverseSlice(s *S)
<span id="ReverseSlice"></span>
> 将切片反转
***
#### func ClearSlice(slice *S)
<span id="ClearSlice"></span>
> 清空切片
***
#### func ClearMap(m M)
<span id="ClearMap"></span>
> 清空 map
***
#### func DropSliceByIndices(slice *S, indices ...int)
<span id="DropSliceByIndices"></span>
> 删除切片中特定索引的元素
***
#### func DropSliceByCondition(slice *S, condition func (v V)  bool)
<span id="DropSliceByCondition"></span>
> 删除切片中符合条件的元素
>   - condition 的返回值为 true 时，将会删除该元素
***
#### func DropSliceOverlappingElements(slice *S, anotherSlice []V, comparisonHandler ComparisonHandler[V])
<span id="DropSliceOverlappingElements"></span>
> 删除切片中与另一个切片重叠的元素
***
#### func DeduplicateSliceInPlace(s *S)
<span id="DeduplicateSliceInPlace"></span>
> 去除切片中的重复元素
***
#### func DeduplicateSlice(s S)  S
<span id="DeduplicateSlice"></span>
> 去除切片中的重复元素，返回新切片
***
#### func DeduplicateSliceInPlaceWithCompare(s *S, compare func (a V)  bool)
<span id="DeduplicateSliceInPlaceWithCompare"></span>
> 去除切片中的重复元素，使用自定义的比较函数
***
#### func DeduplicateSliceWithCompare(s S, compare func (a V)  bool)  S
<span id="DeduplicateSliceWithCompare"></span>
> 去除切片中的重复元素，使用自定义的比较函数，返回新的切片
***
#### func FilterOutByIndices(slice S, indices ...int)  S
<span id="FilterOutByIndices"></span>
> 过滤切片中特定索引的元素，返回过滤后的切片
***
#### func FilterOutByCondition(slice S, condition func (v V)  bool)  S
<span id="FilterOutByCondition"></span>
> 过滤切片中符合条件的元素，返回过滤后的切片
>   - condition 的返回值为 true 时，将会过滤掉该元素
***
#### func FilterOutByKey(m M, key K)  M
<span id="FilterOutByKey"></span>
> 过滤 map 中特定的 key，返回过滤后的 map
***
#### func FilterOutByValue(m M, value V, handler ComparisonHandler[V])  M
<span id="FilterOutByValue"></span>
> 过滤 map 中特定的 value，返回过滤后的 map
***
#### func FilterOutByKeys(m M, keys ...K)  M
<span id="FilterOutByKeys"></span>
> 过滤 map 中多个 key，返回过滤后的 map
***
#### func FilterOutByValues(m M, values []V, handler ComparisonHandler[V])  M
<span id="FilterOutByValues"></span>
> 过滤 map 中多个 values，返回过滤后的 map
***
#### func FilterOutByMap(m M, condition func (k K, v V)  bool)  M
<span id="FilterOutByMap"></span>
> 过滤 map 中符合条件的元素，返回过滤后的 map
>   - condition 的返回值为 true 时，将会过滤掉该元素
***
#### func FindLoopedNextInSlice(slice S, i int) (next int, value V)
<span id="FindLoopedNextInSlice"></span>
> 返回 i 的下一个数组成员，当 i 达到数组长度时从 0 开始
>   - 当 i 为负数时将返回第一个元素
***
#### func FindLoopedPrevInSlice(slice S, i int) (prev int, value V)
<span id="FindLoopedPrevInSlice"></span>
> 返回 i 的上一个数组成员，当 i 为 0 时从数组末尾开始
>   - 当 i 为负数时将返回最后一个元素
***
#### func FindCombinationsInSliceByRange(s S, minSize int, maxSize int)  []S
<span id="FindCombinationsInSliceByRange"></span>
> 获取给定数组的所有组合，且每个组合的成员数量限制在指定范围内
***
#### func FindFirstOrDefaultInSlice(slice S, defaultValue V)  V
<span id="FindFirstOrDefaultInSlice"></span>
> 判断切片中是否存在元素，返回第一个元素，不存在则返回默认值
***
#### func FindOrDefaultInSlice(slice S, defaultValue V, handler func (v V)  bool) (t V)
<span id="FindOrDefaultInSlice"></span>
> 判断切片中是否存在某个元素，返回第一个匹配的索引和元素，不存在则返回默认值
***
#### func FindOrDefaultInComparableSlice(slice S, v V, defaultValue V) (t V)
<span id="FindOrDefaultInComparableSlice"></span>
> 判断切片中是否存在某个元素，返回第一个匹配的索引和元素，不存在则返回默认值
***
#### func FindInSlice(slice S, handler func (v V)  bool) (i int, t V)
<span id="FindInSlice"></span>
> 判断切片中是否存在某个元素，返回第一个匹配的索引和元素，不存在则索引返回 -1
***
#### func FindIndexInSlice(slice S, handler func (v V)  bool)  int
<span id="FindIndexInSlice"></span>
> 判断切片中是否存在某个元素，返回第一个匹配的索引，不存在则索引返回 -1
***
#### func FindInComparableSlice(slice S, v V) (i int, t V)
<span id="FindInComparableSlice"></span>
> 判断切片中是否存在某个元素，返回第一个匹配的索引和元素，不存在则索引返回 -1
***
#### func FindIndexInComparableSlice(slice S, v V)  int
<span id="FindIndexInComparableSlice"></span>
> 判断切片中是否存在某个元素，返回第一个匹配的索引，不存在则索引返回 -1
***
#### func FindMinimumInComparableSlice(slice S) (result V)
<span id="FindMinimumInComparableSlice"></span>
> 获取切片中的最小值
***
#### func FindMinimumInSlice(slice S, handler OrderedValueGetter[V, N]) (result V)
<span id="FindMinimumInSlice"></span>
> 获取切片中的最小值
***
#### func FindMaximumInComparableSlice(slice S) (result V)
<span id="FindMaximumInComparableSlice"></span>
> 获取切片中的最大值
***
#### func FindMaximumInSlice(slice S, handler OrderedValueGetter[V, N]) (result V)
<span id="FindMaximumInSlice"></span>
> 获取切片中的最大值
***
#### func FindMin2MaxInComparableSlice(slice S) (min V, max V)
<span id="FindMin2MaxInComparableSlice"></span>
> 获取切片中的最小值和最大值
***
#### func FindMin2MaxInSlice(slice S, handler OrderedValueGetter[V, N]) (min V, max V)
<span id="FindMin2MaxInSlice"></span>
> 获取切片中的最小值和最大值
***
#### func FindMinFromComparableMap(m M) (result V)
<span id="FindMinFromComparableMap"></span>
> 获取 map 中的最小值
***
#### func FindMinFromMap(m M, handler OrderedValueGetter[V, N]) (result V)
<span id="FindMinFromMap"></span>
> 获取 map 中的最小值
***
#### func FindMaxFromComparableMap(m M) (result V)
<span id="FindMaxFromComparableMap"></span>
> 获取 map 中的最大值
***
#### func FindMaxFromMap(m M, handler OrderedValueGetter[V, N]) (result V)
<span id="FindMaxFromMap"></span>
> 获取 map 中的最大值
***
#### func FindMin2MaxFromComparableMap(m M) (min V, max V)
<span id="FindMin2MaxFromComparableMap"></span>
> 获取 map 中的最小值和最大值
***
#### func FindMin2MaxFromMap(m M) (min V, max V)
<span id="FindMin2MaxFromMap"></span>
> 获取 map 中的最小值和最大值
***
#### func SwapSlice(slice *S, i int, j int)
<span id="SwapSlice"></span>
> 将切片中的两个元素进行交换
***
#### func MappingFromSlice(slice S, handler func (value V)  N)  NS
<span id="MappingFromSlice"></span>
> 将切片中的元素进行转换
***
#### func MappingFromMap(m M, handler func (value V)  N)  NM
<span id="MappingFromMap"></span>
> 将 map 中的元素进行转换
***
#### func MergeSlices(slices ...S) (result S)
<span id="MergeSlices"></span>
> 合并切片
***
#### func MergeMaps(maps ...M) (result M)
<span id="MergeMaps"></span>
> 合并 map，当多个 map 中存在相同的 key 时，后面的 map 中的 key 将会覆盖前面的 map 中的 key
***
#### func MergeMapsWithSkip(maps ...M) (result M)
<span id="MergeMapsWithSkip"></span>
> 合并 map，当多个 map 中存在相同的 key 时，后面的 map 中的 key 将会被跳过
***
#### func ChooseRandomSliceElementRepeatN(slice S, n int) (result []V)
<span id="ChooseRandomSliceElementRepeatN"></span>
> 返回 slice 中的 n 个可重复随机元素
>   - 当 slice 长度为 0 或 n 小于等于 0 时将会返回 nil
***
#### func ChooseRandomIndexRepeatN(slice S, n int) (result []int)
<span id="ChooseRandomIndexRepeatN"></span>
> 返回 slice 中的 n 个可重复随机元素的索引
>   - 当 slice 长度为 0 或 n 小于等于 0 时将会返回 nil
***
#### func ChooseRandomSliceElement(slice S) (v V)
<span id="ChooseRandomSliceElement"></span>
> 返回 slice 中随机一个元素，当 slice 长度为 0 时将会得到 V 的零值
***
#### func ChooseRandomIndex(slice S) (index int)
<span id="ChooseRandomIndex"></span>
> 返回 slice 中随机一个元素的索引，当 slice 长度为 0 时将会得到 -1
***
#### func ChooseRandomSliceElementN(slice S, n int) (result []V)
<span id="ChooseRandomSliceElementN"></span>
> 返回 slice 中的 n 个不可重复的随机元素
>   - 当 slice 长度为 0 或 n 大于 slice 长度或小于 0 时将会发生 panic
***
#### func ChooseRandomIndexN(slice S, n int) (result []int)
<span id="ChooseRandomIndexN"></span>
> 获取切片中的 n 个随机元素的索引
>   - 如果 n 大于切片长度或小于 0 时将会发生 panic
***
#### func ChooseRandomMapKeyRepeatN(m M, n int) (result []K)
<span id="ChooseRandomMapKeyRepeatN"></span>
> 获取 map 中的 n 个随机 key，允许重复
>   - 如果 n 大于 map 长度或小于 0 时将会发生 panic
***
#### func ChooseRandomMapValueRepeatN(m M, n int) (result []V)
<span id="ChooseRandomMapValueRepeatN"></span>
> 获取 map 中的 n 个随机 n，允许重复
>   - 如果 n 大于 map 长度或小于 0 时将会发生 panic
***
#### func ChooseRandomMapKeyAndValueRepeatN(m M, n int)  M
<span id="ChooseRandomMapKeyAndValueRepeatN"></span>
> 获取 map 中的 n 个随机 key 和 v，允许重复
>   - 如果 n 大于 map 长度或小于 0 时将会发生 panic
***
#### func ChooseRandomMapKey(m M) (k K)
<span id="ChooseRandomMapKey"></span>
> 获取 map 中的随机 key
***
#### func ChooseRandomMapValue(m M) (v V)
<span id="ChooseRandomMapValue"></span>
> 获取 map 中的随机 value
***
#### func ChooseRandomMapKeyN(m M, n int) (result []K)
<span id="ChooseRandomMapKeyN"></span>
> 获取 map 中的 inputN 个随机 key
>   - 如果 inputN 大于 map 长度或小于 0 时将会发生 panic
***
#### func ChooseRandomMapValueN(m M, n int) (result []V)
<span id="ChooseRandomMapValueN"></span>
> 获取 map 中的 n 个随机 value
>   - 如果 n 大于 map 长度或小于 0 时将会发生 panic
***
#### func ChooseRandomMapKeyAndValue(m M) (k K, v V)
<span id="ChooseRandomMapKeyAndValue"></span>
> 获取 map 中的随机 key 和 v
***
#### func ChooseRandomMapKeyAndValueN(m M, n int)  M
<span id="ChooseRandomMapKeyAndValueN"></span>
> 获取 map 中的 inputN 个随机 key 和 v
>   - 如果 n 大于 map 长度或小于 0 时将会发生 panic
***
#### func DescBy(a Sort, b Sort)  bool
<span id="DescBy"></span>
> 返回降序比较结果
***
#### func AscBy(a Sort, b Sort)  bool
<span id="AscBy"></span>
> 返回升序比较结果
***
#### func Desc(slice *S, getter func (index int)  Sort)
<span id="Desc"></span>
> 对切片进行降序排序
***
#### func DescByClone(slice S, getter func (index int)  Sort)  S
<span id="DescByClone"></span>
> 对切片进行降序排序，返回排序后的切片
***
#### func Asc(slice *S, getter func (index int)  Sort)
<span id="Asc"></span>
> 对切片进行升序排序
***
#### func AscByClone(slice S, getter func (index int)  Sort)  S
<span id="AscByClone"></span>
> 对切片进行升序排序，返回排序后的切片
***
#### func Shuffle(slice *S)
<span id="Shuffle"></span>
> 对切片进行随机排序
***
#### func ShuffleByClone(slice S)  S
<span id="ShuffleByClone"></span>
> 对切片进行随机排序，返回排序后的切片
***
### ComparisonHandler

```go
type ComparisonHandler[V any] struct{}
```
### OrderedValueGetter

```go
type OrderedValueGetter[V any, N generic.Ordered] struct{}
```
