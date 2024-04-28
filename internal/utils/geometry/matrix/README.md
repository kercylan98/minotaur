# Matrix

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

matrix 提供了一个简单的二维数组的实现


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[NewMatrix](#NewMatrix)|生成特定宽高的二维矩阵


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Matrix](#struct_Matrix)|二维矩阵

</details>


***
## 详情信息
#### func NewMatrix\[T any\](width int, height int) *Matrix[T]
<span id="NewMatrix"></span>
> 生成特定宽高的二维矩阵
>   - 虽然提供了通过x、y坐标的操作函数，但是建议无论如何使用pos进行处理
>   - 该矩阵为XY，而非YX

***
<span id="struct_Matrix"></span>
### Matrix `STRUCT`
二维矩阵
```go
type Matrix[T any] struct {
	w int
	h int
	m []T
}
```
<span id="struct_Matrix_GetWidth"></span>

#### func (*Matrix) GetWidth()  int
> 获取二维矩阵宽度

***
<span id="struct_Matrix_GetHeight"></span>

#### func (*Matrix) GetHeight()  int
> 获取二维矩阵高度

***
<span id="struct_Matrix_GetWidth2Height"></span>

#### func (*Matrix) GetWidth2Height() (width int, height int)
> 获取二维矩阵的宽度和高度

***
<span id="struct_Matrix_GetMatrix"></span>

#### func (*Matrix) GetMatrix()  [][]T
> 获取二维矩阵
>   - 通常建议使用 GetMatrixWithPos 进行处理这样将拥有更高的效率

***
<span id="struct_Matrix_GetMatrixWithPos"></span>

#### func (*Matrix) GetMatrixWithPos()  []T
> 获取顺序的矩阵

***
<span id="struct_Matrix_Get"></span>

#### func (*Matrix) Get(x int, y int) (value T)
> 获取特定坐标的内容

***
<span id="struct_Matrix_GetExist"></span>

#### func (*Matrix) GetExist(x int, y int) (value T, exist bool)
> 获取特定坐标的内容，如果不存在则返回 false

***
<span id="struct_Matrix_GetWithPos"></span>

#### func (*Matrix) GetWithPos(pos int) (value T)
> 获取特定坐标的内容

***
<span id="struct_Matrix_Set"></span>

#### func (*Matrix) Set(x int, y int, data T)
> 设置特定坐标的内容

***
<span id="struct_Matrix_SetWithPos"></span>

#### func (*Matrix) SetWithPos(pos int, data T)
> 设置特定坐标的内容

***
<span id="struct_Matrix_Swap"></span>

#### func (*Matrix) Swap(x1 int, y1 int, x2 int, y2 int)
> 交换两个位置的内容

***
<span id="struct_Matrix_SwapWithPos"></span>

#### func (*Matrix) SwapWithPos(pos1 int, pos2 int)
> 交换两个位置的内容

***
<span id="struct_Matrix_TrySwap"></span>

#### func (*Matrix) TrySwap(x1 int, y1 int, x2 int, y2 int, expressionHandle func (matrix *Matrix[T])  bool)
> 尝试交换两个位置的内容，交换后不满足表达式时进行撤销

***
<span id="struct_Matrix_TrySwapWithPos"></span>

#### func (*Matrix) TrySwapWithPos(pos1 int, pos2 int, expressionHandle func (matrix *Matrix[T])  bool)
> 尝试交换两个位置的内容，交换后不满足表达式时进行撤销

***
<span id="struct_Matrix_FillFull"></span>

#### func (*Matrix) FillFull(generateHandle func (x int)  T)
> 根据提供的生成器填充整个矩阵

***
<span id="struct_Matrix_FillFullWithPos"></span>

#### func (*Matrix) FillFullWithPos(generateHandle func (pos int)  T)
> 根据提供的生成器填充整个矩阵

***
