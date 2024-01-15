# Random

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
|[Dice](#Dice)|掷骰子
|[DiceN](#DiceN)|掷骰子
|[NetIP](#NetIP)|返回一个随机的IP地址
|[Port](#Port)|返回一个随机的端口号
|[IPv4](#IPv4)|返回一个随机产生的IPv4地址。
|[IPv4Port](#IPv4Port)|返回一个随机产生的IPv4地址和端口。
|[Int64](#Int64)|返回一个介于min和max之间的int64类型的随机数。
|[Int](#Int)|返回一个介于min和max之间的的int类型的随机数。
|[Duration](#Duration)|返回一个介于min和max之间的的Duration类型的随机数。
|[Float64](#Float64)|返回一个0~1的浮点数
|[Float32](#Float32)|返回一个0~1的浮点数
|[IntN](#IntN)|返回一个0~n的整数
|[Bool](#Bool)|返回一个随机的布尔值
|[ProbabilitySlice](#ProbabilitySlice)|按概率随机从切片中产生一个数据并返回命中的对象及是否未命中
|[ProbabilitySliceIndex](#ProbabilitySliceIndex)|按概率随机从切片中产生一个数据并返回命中的对象及对象索引以及是否未命中
|[Probability](#Probability)|输入一个概率，返回是否命中
|[ProbabilityChooseOne](#ProbabilityChooseOne)|输入一组概率，返回命中的索引
|[RefreshSeed](#RefreshSeed)|暂无描述...
|[ChineseName](#ChineseName)|返回一个随机组成的中文姓名。
|[EnglishName](#EnglishName)|返回一个随机组成的英文姓名。
|[Name](#Name)|返回一个随机组成的中文或英文姓名
|[NumberString](#NumberString)|返回一个介于min和max之间的string类型的随机数。
|[NumberStringRepair](#NumberStringRepair)|返回一个介于min和max之间的string类型的随机数
|[HostName](#HostName)|返回一个随机产生的hostname。
|[WeightSlice](#WeightSlice)|按权重随机从切片中产生一个数据并返回
|[WeightSliceIndex](#WeightSliceIndex)|按权重随机从切片中产生一个数据并返回数据和对应索引
|[WeightMap](#WeightMap)|按权重随机从map中产生一个数据并返回
|[WeightMapKey](#WeightMapKey)|按权重随机从map中产生一个数据并返回数据和对应 key



</details>


***
## 详情信息
#### func Dice() int
<span id="Dice"></span>
> 掷骰子
>   - 常规掷骰子将返回 1-6 的随机数

***
#### func DiceN(n int) int
<span id="DiceN"></span>
> 掷骰子
>   - 与 Dice 不同的是，将返回 1-N 的随机数

***
#### func NetIP() net.IP
<span id="NetIP"></span>
> 返回一个随机的IP地址

***
#### func Port() int
<span id="Port"></span>
> 返回一个随机的端口号

***
#### func IPv4() string
<span id="IPv4"></span>
> 返回一个随机产生的IPv4地址。

***
#### func IPv4Port() string
<span id="IPv4Port"></span>
> 返回一个随机产生的IPv4地址和端口。

***
#### func Int64(min int64, max int64) int64
<span id="Int64"></span>
> 返回一个介于min和max之间的int64类型的随机数。

***
#### func Int(min int, max int) int
<span id="Int"></span>
> 返回一个介于min和max之间的的int类型的随机数。

***
#### func Duration(min int64, max int64) time.Duration
<span id="Duration"></span>
> 返回一个介于min和max之间的的Duration类型的随机数。

***
#### func Float64() float64
<span id="Float64"></span>
> 返回一个0~1的浮点数

***
#### func Float32() float32
<span id="Float32"></span>
> 返回一个0~1的浮点数

***
#### func IntN(n int) int
<span id="IntN"></span>
> 返回一个0~n的整数

***
#### func Bool() bool
<span id="Bool"></span>
> 返回一个随机的布尔值

***
#### func ProbabilitySlice\[T any\](getProbabilityHandle func (data T)  float64, data ...T) (hit T, miss bool)
<span id="ProbabilitySlice"></span>
> 按概率随机从切片中产生一个数据并返回命中的对象及是否未命中
>   - 当总概率小于 1 将会发生未命中的情况

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestProbabilitySlice(t *testing.T) {
	var awards = []int{1, 2, 3, 4, 5, 6, 7}
	var probability = []float64{0.1, 2, 0.1, 0.1, 0.1, 0.1, 0.1}
	for i := 0; i < 50; i++ {
		t.Log(random.ProbabilitySlice(func(data int) float64 {
			return probability[data-1]
		}, awards...))
	}
}

```


</details>


***
#### func ProbabilitySliceIndex\[T any\](getProbabilityHandle func (data T)  float64, data ...T) (hit T, index int, miss bool)
<span id="ProbabilitySliceIndex"></span>
> 按概率随机从切片中产生一个数据并返回命中的对象及对象索引以及是否未命中
>   - 当总概率小于 1 将会发生未命中的情况

***
#### func Probability(p int, full ...int) bool
<span id="Probability"></span>
> 输入一个概率，返回是否命中
>   - 当 full 不为空时，将以 full 为基数，p 为分子，计算命中概率

***
#### func ProbabilityChooseOne(ps ...int) int
<span id="ProbabilityChooseOne"></span>
> 输入一组概率，返回命中的索引

***
#### func RefreshSeed(seed ...int64)
<span id="RefreshSeed"></span>

***
#### func ChineseName() string
<span id="ChineseName"></span>
> 返回一个随机组成的中文姓名。

***
#### func EnglishName() string
<span id="EnglishName"></span>
> 返回一个随机组成的英文姓名。

***
#### func Name() string
<span id="Name"></span>
> 返回一个随机组成的中文或英文姓名
>   - 以1/2的概率决定生产的是中文还是英文姓名。

***
#### func NumberString(min int, max int) string
<span id="NumberString"></span>
> 返回一个介于min和max之间的string类型的随机数。

***
#### func NumberStringRepair(min int, max int) string
<span id="NumberStringRepair"></span>
> 返回一个介于min和max之间的string类型的随机数
>   - 通过Int64生成一个随机数，当结果的字符串长度小于max的字符串长度的情况下，使用0在开头补齐。

***
#### func HostName() string
<span id="HostName"></span>
> 返回一个随机产生的hostname。

***
#### func WeightSlice\[T any\](getWeightHandle func (data T)  int64, data ...T) T
<span id="WeightSlice"></span>
> 按权重随机从切片中产生一个数据并返回

***
#### func WeightSliceIndex\[T any\](getWeightHandle func (data T)  int64, data ...T) (item T, index int)
<span id="WeightSliceIndex"></span>
> 按权重随机从切片中产生一个数据并返回数据和对应索引

***
#### func WeightMap\[K comparable, T any\](getWeightHandle func (data T)  int64, data map[K]T) T
<span id="WeightMap"></span>
> 按权重随机从map中产生一个数据并返回

***
#### func WeightMapKey\[K comparable, T any\](getWeightHandle func (data T)  int64, data map[K]T) (item T, key K)
<span id="WeightMapKey"></span>
> 按权重随机从map中产生一个数据并返回数据和对应 key

***
