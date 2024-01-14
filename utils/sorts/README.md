# Sorts

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

暂无介绍...


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[Topological](#Topological)|拓扑排序是一种对有向图进行排序的算法，它可以用来解决一些依赖关系的问题，比如计算字段的依赖关系。拓扑排序会将存在依赖关系的元素进行排序，使得依赖关系的元素总是排在被依赖的元素之前。



</details>


***
## 详情信息
#### func Topological(slice S, queryIndexHandler func (item V)  Index, queryDependsHandler func (item V)  []Index)  S,  error
<span id="Topological"></span>
> 拓扑排序是一种对有向图进行排序的算法，它可以用来解决一些依赖关系的问题，比如计算字段的依赖关系。拓扑排序会将存在依赖关系的元素进行排序，使得依赖关系的元素总是排在被依赖的元素之前。
>   - slice: 需要排序的切片
>   - queryIndexHandler: 用于查询切片中每个元素的索引
>   - queryDependsHandler: 用于查询切片中每个元素的依赖关系，返回的是一个索引切片，如果没有依赖关系，那么返回空切片
> 
> 该函数在存在循环依赖的情况下将会返回 ErrCircularDependencyDetected 错误

示例代码：
```go

func ExampleTopological() {
	type Item struct {
		ID      int
		Depends []int
	}
	var items = []Item{{ID: 2, Depends: []int{4}}, {ID: 1, Depends: []int{2, 3}}, {ID: 3, Depends: []int{4}}, {ID: 4, Depends: []int{5}}, {ID: 5, Depends: []int{}}}
	var sorted, err = sorts.Topological(items, func(item Item) int {
		return item.ID
	}, func(item Item) []int {
		return item.Depends
	})
	if err != nil {
		return
	}
	for _, item := range sorted {
		fmt.Println(item.ID, "|", item.Depends)
	}
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestTopological(t *testing.T) {
	type Item struct {
		ID      int
		Depends []int
	}
	var items = []Item{{ID: 2, Depends: []int{4}}, {ID: 1, Depends: []int{2, 3}}, {ID: 3, Depends: []int{4}}, {ID: 4, Depends: []int{5}}, {ID: 5, Depends: []int{}}}
	var sorted, err = sorts.Topological(items, func(item Item) int {
		return item.ID
	}, func(item Item) []int {
		return item.Depends
	})
	if err != nil {
		t.Error(err)
		return
	}
	for _, item := range sorted {
		t.Log(item.ID, "|", item.Depends)
	}
}

```


</details>


<details>
<summary>查看 / 收起基准测试</summary>


```go

func BenchmarkTopological(b *testing.B) {
	type Item struct {
		ID      int
		Depends []int
	}
	var items = []Item{{ID: 2, Depends: []int{4}}, {ID: 1, Depends: []int{2, 3}}, {ID: 3, Depends: []int{4}}, {ID: 4, Depends: []int{5}}, {ID: 5, Depends: []int{}}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Topological(items, func(item Item) int {
			return item.ID
		}, func(item Item) []int {
			return item.Depends
		})
		if err != nil {
			b.Error(err)
			return
		}
	}
}

```


</details>


***
