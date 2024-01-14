# Maths



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/maths)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[Compare](#Compare)|根据特定表达式比较两个值
|[IsContinuity](#IsContinuity)|检查一组值是否连续
|[IsContinuityWithSort](#IsContinuityWithSort)|检查一组值排序后是否连续
|[GetDefaultTolerance](#GetDefaultTolerance)|获取默认误差范围
|[Pow](#Pow)|整数幂运算
|[PowInt64](#PowInt64)|整数幂运算
|[Min](#Min)|返回两个数之中较小的值
|[Max](#Max)|返回两个数之中较大的值
|[MinMax](#MinMax)|将两个数按照较小的和较大的顺序进行返回
|[MaxMin](#MaxMin)|将两个数按照较大的和较小的顺序进行返回
|[Clamp](#Clamp)|将给定值限制在最小值和最大值之间
|[Tolerance](#Tolerance)|检查两个值是否在一个误差范围内
|[Merge](#Merge)|通过一个参考值合并两个数字
|[UnMerge](#UnMerge)|通过一个参考值取消合并的两个数字
|[MergeToInt64](#MergeToInt64)|将两个数字合并为一个 int64 数字
|[UnMergeInt64](#UnMergeInt64)|将一个 int64 数字拆分为两个数字
|[ToContinuous](#ToContinuous)|将一组非连续的数字转换为从1开始的连续数字
|[CountDigits](#CountDigits)|接收一个整数 num 作为输入，并返回该数字的位数
|[GetDigitValue](#GetDigitValue)|接收一个整数 num 和一个表示目标位数的整数 digit 作为输入，并返
|[JoinNumbers](#JoinNumbers)|将一组数字连接起来
|[IsOdd](#IsOdd)|返回 n 是否为奇数
|[IsEven](#IsEven)|返回 n 是否为偶数
|[MakeLastDigitsZero](#MakeLastDigitsZero)|返回一个新的数，其中 num 的最后 digits 位数被设为零。


> 结构体定义

|结构体|描述
|:--|:--
|[CompareExpression](#compareexpression)|比较表达式

</details>


#### func Compare(a V, expression CompareExpression, b V)  bool
<span id="Compare"></span>
> 根据特定表达式比较两个值
***
#### func IsContinuity(values S)  bool
<span id="IsContinuity"></span>
> 检查一组值是否连续
***
#### func IsContinuityWithSort(values S)  bool
<span id="IsContinuityWithSort"></span>
> 检查一组值排序后是否连续
***
#### func GetDefaultTolerance()  float64
<span id="GetDefaultTolerance"></span>
> 获取默认误差范围
***
#### func Pow(a int, n int)  int
<span id="Pow"></span>
> 整数幂运算
***
#### func PowInt64(a int64, n int64)  int64
<span id="PowInt64"></span>
> 整数幂运算
***
#### func Min(a V, b V)  V
<span id="Min"></span>
> 返回两个数之中较小的值
***
#### func Max(a V, b V)  V
<span id="Max"></span>
> 返回两个数之中较大的值
***
#### func MinMax(a V, b V) (min V, max V)
<span id="MinMax"></span>
> 将两个数按照较小的和较大的顺序进行返回
***
#### func MaxMin(a V, b V) (max V, min V)
<span id="MaxMin"></span>
> 将两个数按照较大的和较小的顺序进行返回
***
#### func Clamp(value V, min V, max V)  V
<span id="Clamp"></span>
> 将给定值限制在最小值和最大值之间
***
#### func Tolerance(value1 V, value2 V, tolerance V)  bool
<span id="Tolerance"></span>
> 检查两个值是否在一个误差范围内
***
#### func Merge(refer V, a V, b V)  V
<span id="Merge"></span>
> 通过一个参考值合并两个数字
***
#### func UnMerge(refer V, num V) (a V, b V)
<span id="UnMerge"></span>
> 通过一个参考值取消合并的两个数字
***
#### func MergeToInt64(v1 V, v2 V)  int64
<span id="MergeToInt64"></span>
> 将两个数字合并为一个 int64 数字
***
#### func UnMergeInt64(n int64)  V,  V
<span id="UnMergeInt64"></span>
> 将一个 int64 数字拆分为两个数字
***
#### func ToContinuous(nums S)  map[V]V
<span id="ToContinuous"></span>
> 将一组非连续的数字转换为从1开始的连续数字
>   - 返回值是一个 map，key 是从 1 开始的连续数字，value 是原始数字
***
#### func CountDigits(num V)  int
<span id="CountDigits"></span>
> 接收一个整数 num 作为输入，并返回该数字的位数
***
#### func GetDigitValue(num int64, digit int)  int64
<span id="GetDigitValue"></span>
> 接收一个整数 num 和一个表示目标位数的整数 digit 作为输入，并返
> 回数字 num 在指定位数上的数值。我们使用 math.Abs() 函数获取 num 的绝对值，并通
> 过除以10的操作将 num 移动到目标位数上。然后，通过取余运算得到位数上的数值
***
#### func JoinNumbers(num1 V, n ...V)  V
<span id="JoinNumbers"></span>
> 将一组数字连接起来
***
#### func IsOdd(n V)  bool
<span id="IsOdd"></span>
> 返回 n 是否为奇数
***
#### func IsEven(n V)  bool
<span id="IsEven"></span>
> 返回 n 是否为偶数
***
#### func MakeLastDigitsZero(num T, digits int)  T
<span id="MakeLastDigitsZero"></span>
> 返回一个新的数，其中 num 的最后 digits 位数被设为零。
>   - 函数首先创建一个 10 的 digits 次方的遮罩，然后通过整除和乘以这个遮罩来使 num 的最后 digits 位归零。
>   - 当 T 类型为浮点型时，将被向下取整后再进行转换
***
### CompareExpression
比较表达式
```go
type CompareExpression struct{}
```
