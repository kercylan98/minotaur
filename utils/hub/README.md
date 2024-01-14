# Hub



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/hub)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewObjectPool](#NewObjectPool)|创建一个 ObjectPool


> 结构体定义

|结构体|描述
|:--|:--
|[ObjectPool](#objectpool)|基于 sync.Pool 实现的线程安全的对象池

</details>


#### func NewObjectPool(generator func ()  *T, releaser func (data *T))  *ObjectPool[*T]
<span id="NewObjectPool"></span>
> 创建一个 ObjectPool
***
### ObjectPool
基于 sync.Pool 实现的线程安全的对象池
  - 一些高频临时生成使用的对象可以通过 ObjectPool 进行管理，例如属性计算等
```go
type ObjectPool[T any] struct {
	p        sync.Pool
	releaser func(data T)
}
```
#### func (*ObjectPool) Get()  T
> 获取一个对象
***
#### func (*ObjectPool) Release(data T)
> 将使用完成的对象放回缓冲区
***
