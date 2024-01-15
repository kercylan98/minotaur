# Combination

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

combination 包提供了一些实用的组合函数。


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[NewCombination](#NewCombination)|创建一个新的组合器
|[WithEvaluation](#WithEvaluation)|设置组合评估函数
|[NewMatcher](#NewMatcher)|创建一个新的匹配器
|[WithMatcherEvaluation](#WithMatcherEvaluation)|设置匹配器评估函数
|[WithMatcherLeastLength](#WithMatcherLeastLength)|通过匹配最小长度的组合创建匹配器
|[WithMatcherLength](#WithMatcherLength)|通过匹配长度的组合创建匹配器
|[WithMatcherMostLength](#WithMatcherMostLength)|通过匹配最大长度的组合创建匹配器
|[WithMatcherIntervalLength](#WithMatcherIntervalLength)|通过匹配长度区间的组合创建匹配器
|[WithMatcherContinuity](#WithMatcherContinuity)|通过匹配连续的组合创建匹配器
|[WithMatcherSame](#WithMatcherSame)|通过匹配相同的组合创建匹配器
|[WithMatcherNCarryM](#WithMatcherNCarryM)|通过匹配 N 携带 M 的组合创建匹配器
|[WithMatcherNCarryIndependentM](#WithMatcherNCarryIndependentM)|通过匹配 N 携带独立 M 的组合创建匹配器
|[NewValidator](#NewValidator)|创建一个新的校验器
|[WithValidatorHandle](#WithValidatorHandle)|通过特定的验证函数对组合进行验证
|[WithValidatorHandleLength](#WithValidatorHandleLength)|校验组合的长度是否符合要求
|[WithValidatorHandleLengthRange](#WithValidatorHandleLengthRange)|校验组合的长度是否在指定的范围内
|[WithValidatorHandleLengthMin](#WithValidatorHandleLengthMin)|校验组合的长度是否大于等于指定的最小值
|[WithValidatorHandleLengthMax](#WithValidatorHandleLengthMax)|校验组合的长度是否小于等于指定的最大值
|[WithValidatorHandleLengthNot](#WithValidatorHandleLengthNot)|校验组合的长度是否不等于指定的值
|[WithValidatorHandleTypeLength](#WithValidatorHandleTypeLength)|校验组合成员类型数量是否为指定的值
|[WithValidatorHandleTypeLengthRange](#WithValidatorHandleTypeLengthRange)|校验组合成员类型数量是否在指定的范围内
|[WithValidatorHandleTypeLengthMin](#WithValidatorHandleTypeLengthMin)|校验组合成员类型数量是否大于等于指定的最小值
|[WithValidatorHandleTypeLengthMax](#WithValidatorHandleTypeLengthMax)|校验组合成员类型数量是否小于等于指定的最大值
|[WithValidatorHandleTypeLengthNot](#WithValidatorHandleTypeLengthNot)|校验组合成员类型数量是否不等于指定的值
|[WithValidatorHandleContinuous](#WithValidatorHandleContinuous)|校验组合成员是否连续
|[WithValidatorHandleContinuousNot](#WithValidatorHandleContinuousNot)|校验组合成员是否不连续
|[WithValidatorHandleGroupContinuous](#WithValidatorHandleGroupContinuous)|校验组合成员是否能够按类型分组并且连续
|[WithValidatorHandleGroupContinuousN](#WithValidatorHandleGroupContinuousN)|校验组合成员是否能够按分组为 n 组类型并且连续
|[WithValidatorHandleNCarryM](#WithValidatorHandleNCarryM)| 校验组合成员是否匹配 N 携带相同的 M 的组合
|[WithValidatorHandleNCarryIndependentM](#WithValidatorHandleNCarryIndependentM)|校验组合成员是否匹配 N 携带独立的 M 的组合


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Combination](#combination)|用于从多个匹配器内提取组合的数据结构
|`STRUCT`|[Option](#option)|组合器选项
|`INTERFACE`|[Item](#item)|暂无描述...
|`STRUCT`|[Matcher](#matcher)|用于从一组数据内提取组合的数据结构
|`STRUCT`|[MatcherOption](#matcheroption)|匹配器选项
|`STRUCT`|[Validator](#validator)|用于对组合进行验证的校验器
|`STRUCT`|[ValidatorOption](#validatoroption)|暂无描述...

</details>


***
## 详情信息
#### func NewCombination(options ...Option[T]) *Combination[T]
<span id="NewCombination"></span>
> 创建一个新的组合器

***
#### func WithEvaluation(evaluate func (items []T)  float64) Option[T]
<span id="WithEvaluation"></span>
> 设置组合评估函数
>   - 用于对组合进行评估，返回一个分值的评价函数
>   - 通过该选项将设置所有匹配器的默认评估函数为该函数
>   - 通过匹配器选项 WithMatcherEvaluation 可以覆盖该默认评估函数
>   - 默认的评估函数将返回一个随机数

***
#### func NewMatcher(options ...MatcherOption[T]) *Matcher[T]
<span id="NewMatcher"></span>
> 创建一个新的匹配器

***
#### func WithMatcherEvaluation(evaluate func (items []T)  float64) MatcherOption[T]
<span id="WithMatcherEvaluation"></span>
> 设置匹配器评估函数
>   - 用于对组合进行评估，返回一个分值的评价函数
>   - 通过该选项将覆盖匹配器的默认(WithEvaluation)评估函数

***
#### func WithMatcherLeastLength(length int) MatcherOption[T]
<span id="WithMatcherLeastLength"></span>
> 通过匹配最小长度的组合创建匹配器
>   - length: 组合的长度，表示需要匹配的组合最小数量

***
#### func WithMatcherLength(length int) MatcherOption[T]
<span id="WithMatcherLength"></span>
> 通过匹配长度的组合创建匹配器
>   - length: 组合的长度，表示需要匹配的组合数量

***
#### func WithMatcherMostLength(length int) MatcherOption[T]
<span id="WithMatcherMostLength"></span>
> 通过匹配最大长度的组合创建匹配器
>   - length: 组合的长度，表示需要匹配的组合最大数量

***
#### func WithMatcherIntervalLength(min int, max int) MatcherOption[T]
<span id="WithMatcherIntervalLength"></span>
> 通过匹配长度区间的组合创建匹配器
>   - min: 组合的最小长度，表示需要匹配的组合最小数量
>   - max: 组合的最大长度，表示需要匹配的组合最大数量

***
#### func WithMatcherContinuity(getIndex func (item T)  Index) MatcherOption[T]
<span id="WithMatcherContinuity"></span>
> 通过匹配连续的组合创建匹配器
>   - index: 用于获取组合中元素的索引值，用于判断是否连续

***
#### func WithMatcherSame(count int, getType func (item T)  E) MatcherOption[T]
<span id="WithMatcherSame"></span>
> 通过匹配相同的组合创建匹配器
>   - count: 组合中相同元素的数量，当 count <= 0 时，表示相同元素的数量不限
>   - getType: 用于获取组合中元素的类型，用于判断是否相同

***
#### func WithMatcherNCarryM(n int, m int, getType func (item T)  E) MatcherOption[T]
<span id="WithMatcherNCarryM"></span>
> 通过匹配 N 携带 M 的组合创建匹配器
>   - n: 组合中元素的数量，表示需要匹配的组合数量，n 的类型需要全部相同
>   - m: 组合中元素的数量，表示需要匹配的组合数量，m 的类型需要全部相同
>   - getType: 用于获取组合中元素的类型，用于判断是否相同

***
#### func WithMatcherNCarryIndependentM(n int, m int, getType func (item T)  E) MatcherOption[T]
<span id="WithMatcherNCarryIndependentM"></span>
> 通过匹配 N 携带独立 M 的组合创建匹配器
>   - n: 组合中元素的数量，表示需要匹配的组合数量，n 的类型需要全部相同
>   - m: 组合中元素的数量，表示需要匹配的组合数量，m 的类型无需全部相同
>   - getType: 用于获取组合中元素的类型，用于判断是否相同

***
#### func NewValidator(options ...ValidatorOption[T]) *Validator[T]
<span id="NewValidator"></span>
> 创建一个新的校验器

***
#### func WithValidatorHandle(handle func (items []T)  bool) ValidatorOption[T]
<span id="WithValidatorHandle"></span>
> 通过特定的验证函数对组合进行验证

***
#### func WithValidatorHandleLength(length int) ValidatorOption[T]
<span id="WithValidatorHandleLength"></span>
> 校验组合的长度是否符合要求

***
#### func WithValidatorHandleLengthRange(min int, max int) ValidatorOption[T]
<span id="WithValidatorHandleLengthRange"></span>
> 校验组合的长度是否在指定的范围内

***
#### func WithValidatorHandleLengthMin(min int) ValidatorOption[T]
<span id="WithValidatorHandleLengthMin"></span>
> 校验组合的长度是否大于等于指定的最小值

***
#### func WithValidatorHandleLengthMax(max int) ValidatorOption[T]
<span id="WithValidatorHandleLengthMax"></span>
> 校验组合的长度是否小于等于指定的最大值

***
#### func WithValidatorHandleLengthNot(length int) ValidatorOption[T]
<span id="WithValidatorHandleLengthNot"></span>
> 校验组合的长度是否不等于指定的值

***
#### func WithValidatorHandleTypeLength(length int, getType func (item T)  E) ValidatorOption[T]
<span id="WithValidatorHandleTypeLength"></span>
> 校验组合成员类型数量是否为指定的值

***
#### func WithValidatorHandleTypeLengthRange(min int, max int, getType func (item T)  E) ValidatorOption[T]
<span id="WithValidatorHandleTypeLengthRange"></span>
> 校验组合成员类型数量是否在指定的范围内

***
#### func WithValidatorHandleTypeLengthMin(min int, getType func (item T)  E) ValidatorOption[T]
<span id="WithValidatorHandleTypeLengthMin"></span>
> 校验组合成员类型数量是否大于等于指定的最小值

***
#### func WithValidatorHandleTypeLengthMax(max int, getType func (item T)  E) ValidatorOption[T]
<span id="WithValidatorHandleTypeLengthMax"></span>
> 校验组合成员类型数量是否小于等于指定的最大值

***
#### func WithValidatorHandleTypeLengthNot(length int, getType func (item T)  E) ValidatorOption[T]
<span id="WithValidatorHandleTypeLengthNot"></span>
> 校验组合成员类型数量是否不等于指定的值

***
#### func WithValidatorHandleContinuous(getIndex func (item T)  Index) ValidatorOption[T]
<span id="WithValidatorHandleContinuous"></span>
> 校验组合成员是否连续

***
#### func WithValidatorHandleContinuousNot(getIndex func (item T)  Index) ValidatorOption[T]
<span id="WithValidatorHandleContinuousNot"></span>
> 校验组合成员是否不连续

***
#### func WithValidatorHandleGroupContinuous(getType func (item T)  E, getIndex func (item T)  Index) ValidatorOption[T]
<span id="WithValidatorHandleGroupContinuous"></span>
> 校验组合成员是否能够按类型分组并且连续

***
#### func WithValidatorHandleGroupContinuousN(n int, getType func (item T)  E, getIndex func (item T)  Index) ValidatorOption[T]
<span id="WithValidatorHandleGroupContinuousN"></span>
> 校验组合成员是否能够按分组为 n 组类型并且连续

***
#### func WithValidatorHandleNCarryM(n int, m int, getType func (item T)  E) ValidatorOption[T]
<span id="WithValidatorHandleNCarryM"></span>
>  校验组合成员是否匹配 N 携带相同的 M 的组合
>   - n: 组合中元素的数量，表示需要匹配的组合数量，n 的类型需要全部相同
>   - m: 组合中元素的数量，表示需要匹配的组合数量，m 的类型需要全部相同
>   - getType: 用于获取组合中元素的类型，用于判断是否相同

***
#### func WithValidatorHandleNCarryIndependentM(n int, m int, getType func (item T)  E) ValidatorOption[T]
<span id="WithValidatorHandleNCarryIndependentM"></span>
> 校验组合成员是否匹配 N 携带独立的 M 的组合
>   - n: 组合中元素的数量，表示需要匹配的组合数量，n 的类型需要全部相同
>   - m: 组合中元素的数量，表示需要匹配的组合数量，m 的类型无需全部相同
>   - getType: 用于获取组合中元素的类型，用于判断是否相同

***
### Combination `STRUCT`
用于从多个匹配器内提取组合的数据结构
```go
type Combination[T Item] struct {
	evaluate func([]T) float64
	matchers map[string]*Matcher[T]
	priority []string
}
```
#### func (*Combination) NewMatcher(name string, options ...MatcherOption[T])  *Combination[T]
> 添加一个新的匹配器
***
#### func (*Combination) AddMatcher(name string, matcher *Matcher[T])  *Combination[T]
> 添加一个匹配器
***
#### func (*Combination) RemoveMatcher(name string)  *Combination[T]
> 移除一个匹配器
***
#### func (*Combination) Combinations(items []T) (result [][]T)
> 从一组数据中提取所有符合匹配器规则的组合
***
#### func (*Combination) CombinationsToName(items []T) (result map[string][][]T)
> 从一组数据中提取所有符合匹配器规则的组合，并返回匹配器名称
***
#### func (*Combination) Best(items []T) (name string, result []T)
> 从一组数据中提取符合匹配器规则的最佳组合
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestCombination_Best(t *testing.T) {
	combine := combination.NewCombination(combination.WithEvaluation(func(items []*Poker) float64 {
		var total float64
		for _, item := range items {
			total += float64(item.Point)
		}
		return total
	}))
	combine.NewMatcher("炸弹", combination.WithMatcherSame[*Poker, int](4, func(item *Poker) int {
		return item.Point
	})).NewMatcher("三带一", combination.WithMatcherNCarryM[*Poker, int](3, 1, func(item *Poker) int {
		return item.Point
	}))
	var cards = []*Poker{{Point: 2, Color: 1}, {Point: 2, Color: 2}, {Point: 2, Color: 3}, {Point: 3, Color: 4}, {Point: 4, Color: 1}, {Point: 4, Color: 2}, {Point: 5, Color: 3}, {Point: 6, Color: 4}, {Point: 7, Color: 1}, {Point: 8, Color: 2}, {Point: 9, Color: 3}, {Point: 10, Color: 4}, {Point: 11, Color: 1}, {Point: 12, Color: 2}, {Point: 13, Color: 3}, {Point: 10, Color: 3}, {Point: 11, Color: 2}, {Point: 12, Color: 1}, {Point: 13, Color: 4}, {Point: 10, Color: 2}}
	name, result := combine.Worst(cards)
	fmt.Println("best:", name)
	for _, item := range result {
		fmt.Println(item)
	}
}

```


</details>


***
#### func (*Combination) Worst(items []T) (name string, result []T)
> 从一组数据中提取符合匹配器规则的最差组合
***
### Option `STRUCT`
组合器选项
```go
type Option[T Item] func(*Combination[T])
```
### Item `INTERFACE`

```go
type Item interface{}
```
### Matcher `STRUCT`
用于从一组数据内提取组合的数据结构
```go
type Matcher[T Item] struct {
	evaluate func(items []T) float64
	filter   []func(items []T) [][]T
}
```
#### func (*Matcher) AddFilter(filter func (items []T)  [][]T)
> 添加一个筛选器
>   - 筛选器用于对组合进行筛选，返回一个二维数组，每个数组内的元素都是一个组合
***
#### func (*Matcher) Combinations(items []T)  [][]T
> 从一组数据中提取所有符合筛选器规则的组合
***
#### func (*Matcher) Best(items []T)  []T
> 从一组数据中提取符筛选器规则的最佳组合
***
#### func (*Matcher) Worst(items []T)  []T
> 从一组数据中提取符筛选器规则的最差组合
***
### MatcherOption `STRUCT`
匹配器选项
```go
type MatcherOption[T Item] func(matcher *Matcher[T])
```
### Validator `STRUCT`
用于对组合进行验证的校验器
```go
type Validator[T Item] struct {
	vh []func(items []T) bool
}
```
#### func (*Validator) Validate(items []T)  bool
> 校验组合是否符合要求
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestValidator_Validate(t *testing.T) {
	v := combination.NewValidator[*Card](combination.WithValidatorHandleContinuous[*Card, int](func(item *Card) int {
		switch item.Point {
		case "A":
			return 1
		case "2", "3", "4", "5", "6", "7", "8", "9", "10":
			return super.StringToInt(item.Point)
		case "J":
			return 11
		case "Q":
			return 12
		case "K":
			return 13
		}
		return -1
	}), combination.WithValidatorHandleLength[*Card](3))
	cards := []*Card{{Point: "2", Color: "Spade"}, {Point: "4", Color: "Heart"}, {Point: "3", Color: "Diamond"}}
	fmt.Println(v.Validate(cards))
}

```


</details>


***
### ValidatorOption `STRUCT`

```go
type ValidatorOption[T Item] func(validator *Validator[T])
```
