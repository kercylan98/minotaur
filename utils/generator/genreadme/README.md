# Genreadme

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/genreadme)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)




## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[New](#New)|暂无描述...


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Builder](#builder)|暂无描述...

</details>


***
## 详情信息
#### func New(pkgDirPath string, output string)  *Builder,  error
<span id="New"></span>

***
### Builder `STRUCT`

```go
type Builder struct {
	p *astgo.Package
	b *strings.Builder
	o string
}
```
#### func (*Builder) Generate()  error
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestBuilder_Generate(t *testing.T) {
	filepath.Walk("/Users/kercylan/Coding.localized/Go/minotaur", func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}
		if strings.Contains(strings.TrimPrefix(path, "/Users/kercylan/Coding.localized/Go/minotaur"), ".") {
			return nil
		}
		b, err := New(path, filepath.Join(path, "README.md"))
		if err != nil {
			return nil
		}
		if err = b.Generate(); err != nil {
			panic(err)
		}
		return nil
	})
}

```


</details>


***
