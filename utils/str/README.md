# Str

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/str)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)




## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[RangeLine](#RangeLine)|对传入的eachString进行按行切片后再进行遍历
|[SplitTrimSpace](#SplitTrimSpace)|按照空格分割字符串并去除空格
|[FirstUpper](#FirstUpper)|首字母大写
|[FirstLower](#FirstLower)|首字母小写
|[FirstUpperBytes](#FirstUpperBytes)|首字母大写
|[FirstLowerBytes](#FirstLowerBytes)|首字母小写
|[IsEmpty](#IsEmpty)|判断字符串是否为空
|[IsEmptyBytes](#IsEmptyBytes)|判断字符串是否为空
|[IsNotEmpty](#IsNotEmpty)|判断字符串是否不为空
|[IsNotEmptyBytes](#IsNotEmptyBytes)|判断字符串是否不为空
|[SnakeString](#SnakeString)|蛇形字符串
|[SnakeStringBytes](#SnakeStringBytes)|蛇形字符串
|[CamelString](#CamelString)|驼峰字符串
|[CamelStringBytes](#CamelStringBytes)|驼峰字符串
|[SortJoin](#SortJoin)|将多个字符串排序后拼接
|[HideSensitivity](#HideSensitivity)|返回防敏感化后的字符串
|[ThousandsSeparator](#ThousandsSeparator)|返回将str进行千位分隔符处理后的字符串。
|[KV](#KV)|返回str经过转换后形成的key、value
|[FormatSpeedyInt](#FormatSpeedyInt)|返回numberStr经过格式化后去除空格和“,”分隔符的结果
|[FormatSpeedyInt64](#FormatSpeedyInt64)|返回numberStr经过格式化后去除空格和“,”分隔符的结果
|[FormatSpeedyFloat32](#FormatSpeedyFloat32)|返回numberStr经过格式化后去除空格和“,”分隔符的结果
|[FormatSpeedyFloat64](#FormatSpeedyFloat64)|返回numberStr经过格式化后去除空格和“,”分隔符的结果


***
## 详情信息
#### func RangeLine(eachString string, eachFunc func (index int, line string)  error)  error
<span id="RangeLine"></span>
> 对传入的eachString进行按行切片后再进行遍历
>   - 该函数会预先对“\r\n”进行处理替换为“\n”。
>   - 在遍历到每一行的时候会将结果index和line作为入参传入eachFunc中进行调用。
>   - index表示了当前行的行号（由0开始），line表示了当前行的内容。

***
#### func SplitTrimSpace(str string, sep string)  []string
<span id="SplitTrimSpace"></span>
> 按照空格分割字符串并去除空格

***
#### func FirstUpper(str string)  string
<span id="FirstUpper"></span>
> 首字母大写

***
#### func FirstLower(str string)  string
<span id="FirstLower"></span>
> 首字母小写

***
#### func FirstUpperBytes(str []byte)  []byte
<span id="FirstUpperBytes"></span>
> 首字母大写

***
#### func FirstLowerBytes(str []byte)  []byte
<span id="FirstLowerBytes"></span>
> 首字母小写

***
#### func IsEmpty(str string)  bool
<span id="IsEmpty"></span>
> 判断字符串是否为空

***
#### func IsEmptyBytes(str []byte)  bool
<span id="IsEmptyBytes"></span>
> 判断字符串是否为空

***
#### func IsNotEmpty(str string)  bool
<span id="IsNotEmpty"></span>
> 判断字符串是否不为空

***
#### func IsNotEmptyBytes(str []byte)  bool
<span id="IsNotEmptyBytes"></span>
> 判断字符串是否不为空

***
#### func SnakeString(str string)  string
<span id="SnakeString"></span>
> 蛇形字符串

***
#### func SnakeStringBytes(str []byte)  []byte
<span id="SnakeStringBytes"></span>
> 蛇形字符串

***
#### func CamelString(str string)  string
<span id="CamelString"></span>
> 驼峰字符串

***
#### func CamelStringBytes(str []byte)  []byte
<span id="CamelStringBytes"></span>
> 驼峰字符串

***
#### func SortJoin(delimiter string, s ...string)  string
<span id="SortJoin"></span>
> 将多个字符串排序后拼接

***
#### func HideSensitivity(str string) (result string)
<span id="HideSensitivity"></span>
> 返回防敏感化后的字符串
>   - 隐藏身份证、邮箱、手机号等敏感信息用*号替代

***
#### func ThousandsSeparator(str string)  string
<span id="ThousandsSeparator"></span>
> 返回将str进行千位分隔符处理后的字符串。

***
#### func KV(str string, tag ...string)  string,  string
<span id="KV"></span>
> 返回str经过转换后形成的key、value
>   - 这里tag表示使用什么字符串来区分key和value的分隔符。
>   - 默认情况即不传入tag的情况下分隔符为“=”。

***
#### func FormatSpeedyInt(numberStr string)  int,  error
<span id="FormatSpeedyInt"></span>
> 返回numberStr经过格式化后去除空格和“,”分隔符的结果
>   - 当字符串为“123,456,789”的时候，返回结果为“123456789”。
>   - 当字符串为“123 456 789”的时候，返回结果为“123456789”。
>   - 当字符串为“1 23, 45 6, 789”的时候，返回结果为“123456789”。

***
#### func FormatSpeedyInt64(numberStr string)  int64,  error
<span id="FormatSpeedyInt64"></span>
> 返回numberStr经过格式化后去除空格和“,”分隔符的结果
>   - 当字符串为“123,456,789”的时候，返回结果为“123456789”。
>   - 当字符串为“123 456 789”的时候，返回结果为“123456789”。
>   - 当字符串为“1 23, 45 6, 789”的时候，返回结果为“123456789”。

***
#### func FormatSpeedyFloat32(numberStr string)  float64,  error
<span id="FormatSpeedyFloat32"></span>
> 返回numberStr经过格式化后去除空格和“,”分隔符的结果
>   - 当字符串为“123,456,789.123”的时候，返回结果为“123456789.123”。
>   - 当字符串为“123 456 789.123”的时候，返回结果为“123456789.123”。
>   - 当字符串为“1 23, 45 6, 789.123”的时候，返回结果为“123456789.123”。

***
#### func FormatSpeedyFloat64(numberStr string)  float64,  error
<span id="FormatSpeedyFloat64"></span>
> 返回numberStr经过格式化后去除空格和“,”分隔符的结果
>   - 当字符串为“123,456,789.123”的时候，返回结果为“123456789.123”。
>   - 当字符串为“123 456 789.123”的时候，返回结果为“123456789.123”。
>   - 当字符串为“1 23, 45 6, 789.123”的时候，返回结果为“123456789.123”。

***
