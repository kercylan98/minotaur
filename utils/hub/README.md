# Hub

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
|[NewObjectPool](#NewObjectPool)|创建一个 ObjectPool


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[ObjectPool](#struct_ObjectPool)|基于 sync.Pool 实现的线程安全的对象池

</details>


***
## 详情信息
#### func NewObjectPool\[T any\](generator func ()  *T, releaser func (data *T)) *ObjectPool[*T]
<span id="NewObjectPool"></span>
> 创建一个 ObjectPool

**示例代码：**

```go

func ExampleNewObjectPool() {
	var p = hub.NewObjectPool[map[int]int](func() *map[int]int {
		return &map[int]int{}
	}, func(data *map[int]int) {
		collection.ClearMap(*data)
	})
	m := *p.Get()
	m[1] = 1
	p.Release(&m)
	fmt.Println(m)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewObjectPool(t *testing.T) {
	var cases = []struct {
		name        string
		generator   func() *map[string]int
		releaser    func(data *map[string]int)
		shouldPanic bool
	}{{name: "TestNewObjectPool_NilGenerator", generator: nil, releaser: func(data *map[string]int) {
	}, shouldPanic: true}, {name: "TestNewObjectPool_NilReleaser", generator: func() *map[string]int {
		return &map[string]int{}
	}, releaser: nil, shouldPanic: true}, {name: "TestNewObjectPool_NilGeneratorAndReleaser", generator: nil, releaser: nil, shouldPanic: true}, {name: "TestNewObjectPool_Normal", generator: func() *map[string]int {
		return &map[string]int{}
	}, releaser: func(data *map[string]int) {
	}, shouldPanic: false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer func() {
				if err := recover(); c.shouldPanic && err == nil {
					t.Error("TestNewObjectPool should panic")
				}
			}()
			_ = hub.NewObjectPool[map[string]int](c.generator, c.releaser)
		})
	}
}

```


</details>


***
<span id="struct_ObjectPool"></span>
### ObjectPool `STRUCT`
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
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestObjectPool_Get(t *testing.T) {
	var cases = []struct {
		name      string
		generator func() *map[string]int
		releaser  func(data *map[string]int)
	}{{name: "TestObjectPool_Get_Normal", generator: func() *map[string]int {
		return &map[string]int{}
	}, releaser: func(data *map[string]int) {
		for k := range *data {
			delete(*data, k)
		}
	}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			pool := hub.NewObjectPool[map[string]int](c.generator, c.releaser)
			if actual := pool.Get(); len(*actual) != 0 {
				t.Error("TestObjectPool_Get failed")
			}
		})
	}
}

```


</details>


***
#### func (*ObjectPool) Release(data T)
> 将使用完成的对象放回缓冲区
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestObjectPool_Release(t *testing.T) {
	var cases = []struct {
		name      string
		generator func() *map[string]int
		releaser  func(data *map[string]int)
	}{{name: "TestObjectPool_Release_Normal", generator: func() *map[string]int {
		return &map[string]int{}
	}, releaser: func(data *map[string]int) {
		for k := range *data {
			delete(*data, k)
		}
	}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			pool := hub.NewObjectPool[map[string]int](c.generator, c.releaser)
			msg := pool.Get()
			m := *msg
			m["test"] = 1
			pool.Release(msg)
			if len(m) != 0 {
				t.Error("TestObjectPool_Release failed")
			}
		})
	}
}

```


</details>


***
