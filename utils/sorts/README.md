# Sorts



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/sorts)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[Topological](#Topological)|拓扑排序是一种对有向图进行排序的算法，它可以用来解决一些依赖关系的问题，比如计算字段的依赖关系。拓扑排序会将存在依赖关系的元素进行排序，使得依赖关系的元素总是排在被依赖的元素之前。


> 结构体定义

|结构体|描述
|:--|:--

</details>


#### func Topological(slice S, queryIndexHandler func (item V)  Index, queryDependsHandler func (item V)  []Index)  S,  error
<span id="Topological"></span>
> 拓扑排序是一种对有向图进行排序的算法，它可以用来解决一些依赖关系的问题，比如计算字段的依赖关系。拓扑排序会将存在依赖关系的元素进行排序，使得依赖关系的元素总是排在被依赖的元素之前。
>   - slice: 需要排序的切片
>   - queryIndexHandler: 用于查询切片中每个元素的索引
>   - queryDependsHandler: 用于查询切片中每个元素的依赖关系，返回的是一个索引切片，如果没有依赖关系，那么返回空切片
> 
> 该函数在存在循环依赖的情况下将会返回 ErrCircularDependencyDetected 错误
***
