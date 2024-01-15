# Huge

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
|[NewFloat](#NewFloat)|创建一个 Float
|[NewFloatByString](#NewFloatByString)|通过字符串创建一个 Float
|[NewInt](#NewInt)|创建一个 Int
|[NewIntByString](#NewIntByString)|通过字符串创建一个 Int


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Float](#struct_Float)|暂无描述...
|`STRUCT`|[Int](#struct_Int)|暂无描述...

</details>


***
## 详情信息
#### func NewFloat\[T generic.Number\](x T) *Float
<span id="NewFloat"></span>
> 创建一个 Float

***
#### func NewFloatByString(i string) *Float
<span id="NewFloatByString"></span>
> 通过字符串创建一个 Float
>   - 如果字符串不是一个合法的数字，则返回 0

***
#### func NewInt\[T generic.Number\](x T) *Int
<span id="NewInt"></span>
> 创建一个 Int

***
#### func NewIntByString(i string) *Int
<span id="NewIntByString"></span>
> 通过字符串创建一个 Int
>   - 如果字符串不是一个合法的数字，则返回 0

***
<span id="struct_Float"></span>
### Float `STRUCT`

```go
type Float big.Float
```
#### func (*Float) Copy()  *Float
***
#### func (*Float) Set(i *Float)  *Float
***
#### func (*Float) IsZero()  bool
***
#### func (*Float) ToBigFloat()  *big.Float
***
#### func (*Float) Cmp(i *Float)  int
> 比较，当 slf > i 时返回 1，当 slf < i 时返回 -1，当 slf == i 时返回 0
***
#### func (*Float) GreaterThan(i *Float)  bool
> 大于
***
#### func (*Float) GreaterThanOrEqualTo(i *Float)  bool
> 大于或等于
***
#### func (*Float) LessThan(i *Float)  bool
> 小于
***
#### func (*Float) LessThanOrEqualTo(i *Float)  bool
> 小于或等于
***
#### func (*Float) EqualTo(i *Float)  bool
> 等于
***
#### func (*Float) Float64()  float64
***
#### func (*Float) String()  string
***
#### func (*Float) Add(i *Float)  *Float
***
#### func (*Float) Sub(i *Float)  *Float
***
#### func (*Float) Mul(i *Float)  *Float
***
#### func (*Float) Div(i *Float)  *Float
***
#### func (*Float) Sqrt()  *Float
> 平方根
***
#### func (*Float) Abs()  *Float
> 返回数字的绝对值
***
#### func (*Float) Sign()  int
> 返回数字的符号
>   - 1：正数
>   - 0：零
>   - -1：负数
***
#### func (*Float) IsPositive()  bool
> 是否为正数
***
#### func (*Float) IsNegative()  bool
> 是否为负数
***
<span id="struct_Int"></span>
### Int `STRUCT`

```go
type Int big.Int
```
#### func (*Int) Copy()  *Int
***
#### func (*Int) Set(i *Int)  *Int
***
#### func (*Int) SetInt(i int)  *Int
***
#### func (*Int) SetInt8(i int8)  *Int
***
#### func (*Int) SetInt16(i int16)  *Int
***
#### func (*Int) SetInt32(i int32)  *Int
***
#### func (*Int) SetInt64(i int64)  *Int
***
#### func (*Int) SetUint(i uint)  *Int
***
#### func (*Int) SetUint8(i uint8)  *Int
***
#### func (*Int) SetUint16(i uint16)  *Int
***
#### func (*Int) SetUint32(i uint32)  *Int
***
#### func (*Int) SetUint64(i uint64)  *Int
***
#### func (*Int) IsZero()  bool
***
#### func (*Int) ToBigint()  *big.Int
***
#### func (*Int) Cmp(i *Int)  int
> 比较，当 slf > i 时返回 1，当 slf < i 时返回 -1，当 slf == i 时返回 0
***
#### func (*Int) GreaterThan(i *Int)  bool
> 大于
***
#### func (*Int) GreaterThanOrEqualTo(i *Int)  bool
> 大于或等于
***
#### func (*Int) LessThan(i *Int)  bool
> 小于
***
#### func (*Int) LessThanOrEqualTo(i *Int)  bool
> 小于或等于
***
#### func (*Int) EqualTo(i *Int)  bool
> 等于
***
#### func (*Int) Int64()  int64
***
#### func (*Int) String()  string
***
#### func (*Int) Add(i *Int)  *Int
***
#### func (*Int) AddInt(i int)  *Int
***
#### func (*Int) AddInt8(i int8)  *Int
***
#### func (*Int) AddInt16(i int16)  *Int
***
#### func (*Int) AddInt32(i int32)  *Int
***
#### func (*Int) AddInt64(i int64)  *Int
***
#### func (*Int) AddUint(i uint)  *Int
***
#### func (*Int) AddUint8(i uint8)  *Int
***
#### func (*Int) AddUint16(i uint16)  *Int
***
#### func (*Int) AddUint32(i uint32)  *Int
***
#### func (*Int) AddUint64(i uint64)  *Int
***
#### func (*Int) Mul(i *Int)  *Int
***
#### func (*Int) MulInt(i int)  *Int
***
#### func (*Int) MulInt8(i int8)  *Int
***
#### func (*Int) MulInt16(i int16)  *Int
***
#### func (*Int) MulInt32(i int32)  *Int
***
#### func (*Int) MulInt64(i int64)  *Int
***
#### func (*Int) MulUint(i uint)  *Int
***
#### func (*Int) MulUint8(i uint8)  *Int
***
#### func (*Int) MulUint16(i uint16)  *Int
***
#### func (*Int) MulUint32(i uint32)  *Int
***
#### func (*Int) MulUint64(i uint64)  *Int
***
#### func (*Int) Sub(i *Int)  *Int
***
#### func (*Int) SubInt(i int)  *Int
***
#### func (*Int) SubInt8(i int8)  *Int
***
#### func (*Int) SubInt16(i int16)  *Int
***
#### func (*Int) SubInt32(i int32)  *Int
***
#### func (*Int) SubInt64(i int64)  *Int
***
#### func (*Int) SubUint(i uint)  *Int
***
#### func (*Int) SubUint8(i uint8)  *Int
***
#### func (*Int) SubUint16(i uint16)  *Int
***
#### func (*Int) SubUint32(i uint32)  *Int
***
#### func (*Int) SubUint64(i uint64)  *Int
***
#### func (*Int) Div(i *Int)  *Int
***
#### func (*Int) DivInt(i int)  *Int
***
#### func (*Int) DivInt8(i int8)  *Int
***
#### func (*Int) DivInt16(i int16)  *Int
***
#### func (*Int) DivInt32(i int32)  *Int
***
#### func (*Int) DivInt64(i int64)  *Int
***
#### func (*Int) DivUint(i uint)  *Int
***
#### func (*Int) DivUint8(i uint8)  *Int
***
#### func (*Int) DivUint16(i uint16)  *Int
***
#### func (*Int) DivUint32(i uint32)  *Int
***
#### func (*Int) DivUint64(i uint64)  *Int
***
#### func (*Int) Mod(i *Int)  *Int
***
#### func (*Int) ModInt(i int)  *Int
***
#### func (*Int) ModInt8(i int8)  *Int
***
#### func (*Int) ModInt16(i int16)  *Int
***
#### func (*Int) ModInt32(i int32)  *Int
***
#### func (*Int) ModInt64(i int64)  *Int
***
#### func (*Int) ModUint(i uint)  *Int
***
#### func (*Int) ModUint8(i uint8)  *Int
***
#### func (*Int) ModUint16(i uint16)  *Int
***
#### func (*Int) ModUint32(i uint32)  *Int
***
#### func (*Int) ModUint64(i uint64)  *Int
***
#### func (*Int) Pow(i *Int)  *Int
***
#### func (*Int) PowInt(i int)  *Int
***
#### func (*Int) PowInt8(i int8)  *Int
***
#### func (*Int) PowInt16(i int16)  *Int
***
#### func (*Int) PowInt32(i int32)  *Int
***
#### func (*Int) PowInt64(i int64)  *Int
***
#### func (*Int) PowUint(i uint)  *Int
***
#### func (*Int) PowUint8(i uint8)  *Int
***
#### func (*Int) PowUint16(i uint16)  *Int
***
#### func (*Int) PowUint32(i uint32)  *Int
***
#### func (*Int) PowUint64(i uint64)  *Int
***
#### func (*Int) Lsh(i int)  *Int
> 左移
***
#### func (*Int) Rsh(i int)  *Int
> 右移
***
#### func (*Int) And(i *Int)  *Int
> 与
***
#### func (*Int) AndNot(i *Int)  *Int
> 与非
***
#### func (*Int) Or(i *Int)  *Int
> 或
***
#### func (*Int) Xor(i *Int)  *Int
> 异或
***
#### func (*Int) Not()  *Int
> 非
***
#### func (*Int) Sqrt()  *Int
> 平方根
***
#### func (*Int) GCD(i *Int)  *Int
> 最大公约数
***
#### func (*Int) LCM(i *Int)  *Int
> 最小公倍数
***
#### func (*Int) ModInverse(i *Int)  *Int
> 模反元素
***
#### func (*Int) ModSqrt(i *Int)  *Int
> 模平方根
***
#### func (*Int) BitLen()  int
> 二进制长度
***
#### func (*Int) Bit(i int)  uint
> 二进制位
***
#### func (*Int) SetBit(i int, v uint)  *Int
> 设置二进制位
***
#### func (*Int) Neg()  *Int
> 返回数字的相反数
***
#### func (*Int) Abs()  *Int
> 返回数字的绝对值
***
#### func (*Int) Sign()  int
> 返回数字的符号
>   - 1：正数
>   - 0：零
>   - -1：负数
***
#### func (*Int) IsPositive()  bool
> 是否为正数
***
#### func (*Int) IsNegative()  bool
> 是否为负数
***
#### func (*Int) IsEven()  bool
> 是否为偶数
***
#### func (*Int) IsOdd()  bool
> 是否为奇数
***
#### func (*Int) ProportionalCalc(proportional *Int, formula func (v *Int)  *Int)  *Int
> 比例计算，该函数会再 formula 返回值的基础上除以 proportional
>   - formula 为计算公式，该公式的参数为调用该函数的 Int 的拷贝
***
