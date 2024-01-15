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
<span id="struct_Float_Copy"></span>

#### func (*Float) Copy()  *Float

***
<span id="struct_Float_Set"></span>

#### func (*Float) Set(i *Float)  *Float

***
<span id="struct_Float_IsZero"></span>

#### func (*Float) IsZero()  bool

***
<span id="struct_Float_ToBigFloat"></span>

#### func (*Float) ToBigFloat()  *big.Float

***
<span id="struct_Float_Cmp"></span>

#### func (*Float) Cmp(i *Float)  int
> 比较，当 slf > i 时返回 1，当 slf < i 时返回 -1，当 slf == i 时返回 0

***
<span id="struct_Float_GreaterThan"></span>

#### func (*Float) GreaterThan(i *Float)  bool
> 大于

***
<span id="struct_Float_GreaterThanOrEqualTo"></span>

#### func (*Float) GreaterThanOrEqualTo(i *Float)  bool
> 大于或等于

***
<span id="struct_Float_LessThan"></span>

#### func (*Float) LessThan(i *Float)  bool
> 小于

***
<span id="struct_Float_LessThanOrEqualTo"></span>

#### func (*Float) LessThanOrEqualTo(i *Float)  bool
> 小于或等于

***
<span id="struct_Float_EqualTo"></span>

#### func (*Float) EqualTo(i *Float)  bool
> 等于

***
<span id="struct_Float_Float64"></span>

#### func (*Float) Float64()  float64

***
<span id="struct_Float_String"></span>

#### func (*Float) String()  string

***
<span id="struct_Float_Add"></span>

#### func (*Float) Add(i *Float)  *Float

***
<span id="struct_Float_Sub"></span>

#### func (*Float) Sub(i *Float)  *Float

***
<span id="struct_Float_Mul"></span>

#### func (*Float) Mul(i *Float)  *Float

***
<span id="struct_Float_Div"></span>

#### func (*Float) Div(i *Float)  *Float

***
<span id="struct_Float_Sqrt"></span>

#### func (*Float) Sqrt()  *Float
> 平方根

***
<span id="struct_Float_Abs"></span>

#### func (*Float) Abs()  *Float
> 返回数字的绝对值

***
<span id="struct_Float_Sign"></span>

#### func (*Float) Sign()  int
> 返回数字的符号
>   - 1：正数
>   - 0：零
>   - -1：负数

***
<span id="struct_Float_IsPositive"></span>

#### func (*Float) IsPositive()  bool
> 是否为正数

***
<span id="struct_Float_IsNegative"></span>

#### func (*Float) IsNegative()  bool
> 是否为负数

***
<span id="struct_Int"></span>
### Int `STRUCT`

```go
type Int big.Int
```
<span id="struct_Int_Copy"></span>

#### func (*Int) Copy()  *Int

***
<span id="struct_Int_Set"></span>

#### func (*Int) Set(i *Int)  *Int

***
<span id="struct_Int_SetInt"></span>

#### func (*Int) SetInt(i int)  *Int

***
<span id="struct_Int_SetInt8"></span>

#### func (*Int) SetInt8(i int8)  *Int

***
<span id="struct_Int_SetInt16"></span>

#### func (*Int) SetInt16(i int16)  *Int

***
<span id="struct_Int_SetInt32"></span>

#### func (*Int) SetInt32(i int32)  *Int

***
<span id="struct_Int_SetInt64"></span>

#### func (*Int) SetInt64(i int64)  *Int

***
<span id="struct_Int_SetUint"></span>

#### func (*Int) SetUint(i uint)  *Int

***
<span id="struct_Int_SetUint8"></span>

#### func (*Int) SetUint8(i uint8)  *Int

***
<span id="struct_Int_SetUint16"></span>

#### func (*Int) SetUint16(i uint16)  *Int

***
<span id="struct_Int_SetUint32"></span>

#### func (*Int) SetUint32(i uint32)  *Int

***
<span id="struct_Int_SetUint64"></span>

#### func (*Int) SetUint64(i uint64)  *Int

***
<span id="struct_Int_IsZero"></span>

#### func (*Int) IsZero()  bool

***
<span id="struct_Int_ToBigint"></span>

#### func (*Int) ToBigint()  *big.Int

***
<span id="struct_Int_Cmp"></span>

#### func (*Int) Cmp(i *Int)  int
> 比较，当 slf > i 时返回 1，当 slf < i 时返回 -1，当 slf == i 时返回 0

***
<span id="struct_Int_GreaterThan"></span>

#### func (*Int) GreaterThan(i *Int)  bool
> 大于

***
<span id="struct_Int_GreaterThanOrEqualTo"></span>

#### func (*Int) GreaterThanOrEqualTo(i *Int)  bool
> 大于或等于

***
<span id="struct_Int_LessThan"></span>

#### func (*Int) LessThan(i *Int)  bool
> 小于

***
<span id="struct_Int_LessThanOrEqualTo"></span>

#### func (*Int) LessThanOrEqualTo(i *Int)  bool
> 小于或等于

***
<span id="struct_Int_EqualTo"></span>

#### func (*Int) EqualTo(i *Int)  bool
> 等于

***
<span id="struct_Int_Int64"></span>

#### func (*Int) Int64()  int64

***
<span id="struct_Int_String"></span>

#### func (*Int) String()  string

***
<span id="struct_Int_Add"></span>

#### func (*Int) Add(i *Int)  *Int

***
<span id="struct_Int_AddInt"></span>

#### func (*Int) AddInt(i int)  *Int

***
<span id="struct_Int_AddInt8"></span>

#### func (*Int) AddInt8(i int8)  *Int

***
<span id="struct_Int_AddInt16"></span>

#### func (*Int) AddInt16(i int16)  *Int

***
<span id="struct_Int_AddInt32"></span>

#### func (*Int) AddInt32(i int32)  *Int

***
<span id="struct_Int_AddInt64"></span>

#### func (*Int) AddInt64(i int64)  *Int

***
<span id="struct_Int_AddUint"></span>

#### func (*Int) AddUint(i uint)  *Int

***
<span id="struct_Int_AddUint8"></span>

#### func (*Int) AddUint8(i uint8)  *Int

***
<span id="struct_Int_AddUint16"></span>

#### func (*Int) AddUint16(i uint16)  *Int

***
<span id="struct_Int_AddUint32"></span>

#### func (*Int) AddUint32(i uint32)  *Int

***
<span id="struct_Int_AddUint64"></span>

#### func (*Int) AddUint64(i uint64)  *Int

***
<span id="struct_Int_Mul"></span>

#### func (*Int) Mul(i *Int)  *Int

***
<span id="struct_Int_MulInt"></span>

#### func (*Int) MulInt(i int)  *Int

***
<span id="struct_Int_MulInt8"></span>

#### func (*Int) MulInt8(i int8)  *Int

***
<span id="struct_Int_MulInt16"></span>

#### func (*Int) MulInt16(i int16)  *Int

***
<span id="struct_Int_MulInt32"></span>

#### func (*Int) MulInt32(i int32)  *Int

***
<span id="struct_Int_MulInt64"></span>

#### func (*Int) MulInt64(i int64)  *Int

***
<span id="struct_Int_MulUint"></span>

#### func (*Int) MulUint(i uint)  *Int

***
<span id="struct_Int_MulUint8"></span>

#### func (*Int) MulUint8(i uint8)  *Int

***
<span id="struct_Int_MulUint16"></span>

#### func (*Int) MulUint16(i uint16)  *Int

***
<span id="struct_Int_MulUint32"></span>

#### func (*Int) MulUint32(i uint32)  *Int

***
<span id="struct_Int_MulUint64"></span>

#### func (*Int) MulUint64(i uint64)  *Int

***
<span id="struct_Int_Sub"></span>

#### func (*Int) Sub(i *Int)  *Int

***
<span id="struct_Int_SubInt"></span>

#### func (*Int) SubInt(i int)  *Int

***
<span id="struct_Int_SubInt8"></span>

#### func (*Int) SubInt8(i int8)  *Int

***
<span id="struct_Int_SubInt16"></span>

#### func (*Int) SubInt16(i int16)  *Int

***
<span id="struct_Int_SubInt32"></span>

#### func (*Int) SubInt32(i int32)  *Int

***
<span id="struct_Int_SubInt64"></span>

#### func (*Int) SubInt64(i int64)  *Int

***
<span id="struct_Int_SubUint"></span>

#### func (*Int) SubUint(i uint)  *Int

***
<span id="struct_Int_SubUint8"></span>

#### func (*Int) SubUint8(i uint8)  *Int

***
<span id="struct_Int_SubUint16"></span>

#### func (*Int) SubUint16(i uint16)  *Int

***
<span id="struct_Int_SubUint32"></span>

#### func (*Int) SubUint32(i uint32)  *Int

***
<span id="struct_Int_SubUint64"></span>

#### func (*Int) SubUint64(i uint64)  *Int

***
<span id="struct_Int_Div"></span>

#### func (*Int) Div(i *Int)  *Int

***
<span id="struct_Int_DivInt"></span>

#### func (*Int) DivInt(i int)  *Int

***
<span id="struct_Int_DivInt8"></span>

#### func (*Int) DivInt8(i int8)  *Int

***
<span id="struct_Int_DivInt16"></span>

#### func (*Int) DivInt16(i int16)  *Int

***
<span id="struct_Int_DivInt32"></span>

#### func (*Int) DivInt32(i int32)  *Int

***
<span id="struct_Int_DivInt64"></span>

#### func (*Int) DivInt64(i int64)  *Int

***
<span id="struct_Int_DivUint"></span>

#### func (*Int) DivUint(i uint)  *Int

***
<span id="struct_Int_DivUint8"></span>

#### func (*Int) DivUint8(i uint8)  *Int

***
<span id="struct_Int_DivUint16"></span>

#### func (*Int) DivUint16(i uint16)  *Int

***
<span id="struct_Int_DivUint32"></span>

#### func (*Int) DivUint32(i uint32)  *Int

***
<span id="struct_Int_DivUint64"></span>

#### func (*Int) DivUint64(i uint64)  *Int

***
<span id="struct_Int_Mod"></span>

#### func (*Int) Mod(i *Int)  *Int

***
<span id="struct_Int_ModInt"></span>

#### func (*Int) ModInt(i int)  *Int

***
<span id="struct_Int_ModInt8"></span>

#### func (*Int) ModInt8(i int8)  *Int

***
<span id="struct_Int_ModInt16"></span>

#### func (*Int) ModInt16(i int16)  *Int

***
<span id="struct_Int_ModInt32"></span>

#### func (*Int) ModInt32(i int32)  *Int

***
<span id="struct_Int_ModInt64"></span>

#### func (*Int) ModInt64(i int64)  *Int

***
<span id="struct_Int_ModUint"></span>

#### func (*Int) ModUint(i uint)  *Int

***
<span id="struct_Int_ModUint8"></span>

#### func (*Int) ModUint8(i uint8)  *Int

***
<span id="struct_Int_ModUint16"></span>

#### func (*Int) ModUint16(i uint16)  *Int

***
<span id="struct_Int_ModUint32"></span>

#### func (*Int) ModUint32(i uint32)  *Int

***
<span id="struct_Int_ModUint64"></span>

#### func (*Int) ModUint64(i uint64)  *Int

***
<span id="struct_Int_Pow"></span>

#### func (*Int) Pow(i *Int)  *Int

***
<span id="struct_Int_PowInt"></span>

#### func (*Int) PowInt(i int)  *Int

***
<span id="struct_Int_PowInt8"></span>

#### func (*Int) PowInt8(i int8)  *Int

***
<span id="struct_Int_PowInt16"></span>

#### func (*Int) PowInt16(i int16)  *Int

***
<span id="struct_Int_PowInt32"></span>

#### func (*Int) PowInt32(i int32)  *Int

***
<span id="struct_Int_PowInt64"></span>

#### func (*Int) PowInt64(i int64)  *Int

***
<span id="struct_Int_PowUint"></span>

#### func (*Int) PowUint(i uint)  *Int

***
<span id="struct_Int_PowUint8"></span>

#### func (*Int) PowUint8(i uint8)  *Int

***
<span id="struct_Int_PowUint16"></span>

#### func (*Int) PowUint16(i uint16)  *Int

***
<span id="struct_Int_PowUint32"></span>

#### func (*Int) PowUint32(i uint32)  *Int

***
<span id="struct_Int_PowUint64"></span>

#### func (*Int) PowUint64(i uint64)  *Int

***
<span id="struct_Int_Lsh"></span>

#### func (*Int) Lsh(i int)  *Int
> 左移

***
<span id="struct_Int_Rsh"></span>

#### func (*Int) Rsh(i int)  *Int
> 右移

***
<span id="struct_Int_And"></span>

#### func (*Int) And(i *Int)  *Int
> 与

***
<span id="struct_Int_AndNot"></span>

#### func (*Int) AndNot(i *Int)  *Int
> 与非

***
<span id="struct_Int_Or"></span>

#### func (*Int) Or(i *Int)  *Int
> 或

***
<span id="struct_Int_Xor"></span>

#### func (*Int) Xor(i *Int)  *Int
> 异或

***
<span id="struct_Int_Not"></span>

#### func (*Int) Not()  *Int
> 非

***
<span id="struct_Int_Sqrt"></span>

#### func (*Int) Sqrt()  *Int
> 平方根

***
<span id="struct_Int_GCD"></span>

#### func (*Int) GCD(i *Int)  *Int
> 最大公约数

***
<span id="struct_Int_LCM"></span>

#### func (*Int) LCM(i *Int)  *Int
> 最小公倍数

***
<span id="struct_Int_ModInverse"></span>

#### func (*Int) ModInverse(i *Int)  *Int
> 模反元素

***
<span id="struct_Int_ModSqrt"></span>

#### func (*Int) ModSqrt(i *Int)  *Int
> 模平方根

***
<span id="struct_Int_BitLen"></span>

#### func (*Int) BitLen()  int
> 二进制长度

***
<span id="struct_Int_Bit"></span>

#### func (*Int) Bit(i int)  uint
> 二进制位

***
<span id="struct_Int_SetBit"></span>

#### func (*Int) SetBit(i int, v uint)  *Int
> 设置二进制位

***
<span id="struct_Int_Neg"></span>

#### func (*Int) Neg()  *Int
> 返回数字的相反数

***
<span id="struct_Int_Abs"></span>

#### func (*Int) Abs()  *Int
> 返回数字的绝对值

***
<span id="struct_Int_Sign"></span>

#### func (*Int) Sign()  int
> 返回数字的符号
>   - 1：正数
>   - 0：零
>   - -1：负数

***
<span id="struct_Int_IsPositive"></span>

#### func (*Int) IsPositive()  bool
> 是否为正数

***
<span id="struct_Int_IsNegative"></span>

#### func (*Int) IsNegative()  bool
> 是否为负数

***
<span id="struct_Int_IsEven"></span>

#### func (*Int) IsEven()  bool
> 是否为偶数

***
<span id="struct_Int_IsOdd"></span>

#### func (*Int) IsOdd()  bool
> 是否为奇数

***
<span id="struct_Int_ProportionalCalc"></span>

#### func (*Int) ProportionalCalc(proportional *Int, formula func (v *Int)  *Int)  *Int
> 比例计算，该函数会再 formula 返回值的基础上除以 proportional
>   - formula 为计算公式，该公式的参数为调用该函数的 Int 的拷贝

***
