# Survey



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/survey)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewFileFlusher](#NewFileFlusher)|创建一个文件刷新器
|[WithFlushInterval](#WithFlushInterval)|设置日志文件刷新间隔
|[Reg](#Reg)|注册一个运营日志记录器
|[Record](#Record)|记录一条运营日志
|[RecordBytes](#RecordBytes)|记录一条运营日志
|[Flush](#Flush)|将运营日志记录器的缓冲区数据写入到文件
|[Close](#Close)|关闭运营日志记录器
|[Analyze](#Analyze)|分析特定文件的记录，当发生错误时，会发生 panic
|[AnalyzeMulti](#AnalyzeMulti)|与 Analyze 类似，但是可以分析多个文件
|[IncrementAnalyze](#IncrementAnalyze)|增量分析，返回一个函数，每次调用该函数都会分析文件中新增的内容


> 结构体定义

|结构体|描述
|:--|:--
|[Analyzer](#analyzer)|分析器
|[Flusher](#flusher)|用于刷新缓冲区的接口
|[FileFlusher](#fileflusher)|暂无描述...
|[Option](#option)|选项
|[Result](#result)|暂无描述...
|[R](#r)|记录器所记录的一条数据
|[Report](#report)|分析报告

</details>


#### func NewFileFlusher(filePath string, layout ...string)  *FileFlusher
<span id="NewFileFlusher"></span>
> 创建一个文件刷新器
>   - layout 为日志文件名的时间戳格式 (默认为 time.DateOnly)
***
#### func WithFlushInterval(interval time.Duration)  Option
<span id="WithFlushInterval"></span>
> 设置日志文件刷新间隔
>   - 默认为 3s，当日志文件刷新间隔 <= 0 时，将会在每次写入日志时刷新日志文件
***
#### func Reg(name string, flusher Flusher, options ...Option)
<span id="Reg"></span>
> 注册一个运营日志记录器
***
#### func Record(name string, data map[string]any)
<span id="Record"></span>
> 记录一条运营日志
***
#### func RecordBytes(name string, data []byte)
<span id="RecordBytes"></span>
> 记录一条运营日志
***
#### func Flush(names ...string)
<span id="Flush"></span>
> 将运营日志记录器的缓冲区数据写入到文件
>   - name 为空时，将所有记录器的缓冲区数据写入到文件
***
#### func Close(names ...string)
<span id="Close"></span>
> 关闭运营日志记录器
***
#### func Analyze(filePath string, handle func (analyzer *Analyzer, record R))  *Report
<span id="Analyze"></span>
> 分析特定文件的记录，当发生错误时，会发生 panic
>   - handle 为并行执行的，需要自行处理并发安全
>   - 适用于外部进程对于日志文件的读取，但是需要注意的是，此时日志文件可能正在被写入，所以可能会读取到错误的数据
***
#### func AnalyzeMulti(filePaths []string, handle func (analyzer *Analyzer, record R))  *Report
<span id="AnalyzeMulti"></span>
> 与 Analyze 类似，但是可以分析多个文件
***
#### func IncrementAnalyze(filePath string, handle func (analyzer *Analyzer, record R))  func ()  *Report,  error
<span id="IncrementAnalyze"></span>
> 增量分析，返回一个函数，每次调用该函数都会分析文件中新增的内容
***
### Analyzer
分析器
```go
type Analyzer struct {
	v      map[string]any
	vc     map[string]int64
	repeat map[string]struct{}
	subs   map[string]*Analyzer
	format map[string]func(v any) any
	m      sync.Mutex
}
```
#### func (*Analyzer) Sub(key string)  *Analyzer
> 获取子分析器
***
#### func (*Analyzer) SetFormat(key string, format func (v any)  any)
> 设置格式化函数
***
#### func (*Analyzer) SetValueIfGreaterThan(key string, value float64)
> 设置指定 key 的值，当新值大于旧值时
>   - 当已有值不为 float64 时，将会被忽略
***
#### func (*Analyzer) SetValueIfLessThan(key string, value float64)
> 设置指定 key 的值，当新值小于旧值时
>   - 当已有值不为 float64 时，将会被忽略
***
#### func (*Analyzer) SetValueIf(key string, expression bool, value float64)
> 当表达式满足的时候将设置指定 key 的值为 value
***
#### func (*Analyzer) SetValueStringIf(key string, expression bool, value string)
> 当表达式满足的时候将设置指定 key 的值为 value
***
#### func (*Analyzer) SetValue(key string, value float64)
> 设置指定 key 的值
***
#### func (*Analyzer) SetValueString(key string, value string)
> 设置指定 key 的值
***
#### func (*Analyzer) Increase(key string, record R, recordKey string)
> 在指定 key 现有值的基础上增加 recordKey 的值
>   - 当分析器已经记录过相同 key 的值时，会根据已有的值类型进行不同处理
> 
> 处理方式：
>   - 当已有值类型为 string 时，将会使用新的值的 string 类型进行覆盖
>   - 当已有值类型为 float64 时，当新的值类型不为 float64 时，将会被忽略
***
#### func (*Analyzer) IncreaseValue(key string, value float64)
> 在指定 key 现有值的基础上增加 value
***
#### func (*Analyzer) IncreaseNonRepeat(key string, record R, recordKey string, dimension ...string)
> 在指定 key 现有值的基础上增加 recordKey 的值，但是当去重维度 dimension 相同时，不会增加
***
#### func (*Analyzer) IncreaseValueNonRepeat(key string, record R, value float64, dimension ...string)
> 在指定 key 现有值的基础上增加 value，但是当去重维度 dimension 相同时，不会增加
***
#### func (*Analyzer) GetValue(key string)  float64
> 获取当前记录的值
***
#### func (*Analyzer) GetValueString(key string)  string
> 获取当前记录的值
***
### Flusher
用于刷新缓冲区的接口
```go
type Flusher struct{}
```
### FileFlusher

```go
type FileFlusher struct {
	dir       string
	fn        string
	fe        string
	layout    string
	layoutLen int
}
```
#### func (*FileFlusher) Flush(records []string)
***
#### func (*FileFlusher) Info()  string
***
### Option
选项
```go
type Option struct{}
```
### Result

```go
type Result struct{}
```
### R
记录器所记录的一条数据
```go
type R struct{}
```
#### func (R) GetTime(layout string)  time.Time
> 获取该记录的时间
***
#### func (R) Get(key string)  Result
> 获取指定 key 的值
>   - 当 key 为嵌套 key 时，使用 . 进行分割，例如：a.b.c
>   - 更多用法参考：https://github.com/tidwall/gjson
***
#### func (R) Exist(key string)  bool
> 判断指定 key 是否存在
***
#### func (R) GetString(key string)  string
> 该函数为 Get(key).String() 的简写
***
#### func (R) GetInt64(key string)  int64
> 该函数为 Get(key).Int() 的简写
***
#### func (R) GetInt(key string)  int
> 该函数为 Get(key).Int() 的简写，但是返回值为 int 类型
***
#### func (R) GetFloat64(key string)  float64
> 该函数为 Get(key).Float() 的简写
***
#### func (R) GetBool(key string)  bool
> 该函数为 Get(key).Bool() 的简写
***
#### func (R) String()  string
***
### Report
分析报告
```go
type Report struct {
	analyzer *Analyzer
	Name     string
	Values   map[string]any
	Counter  map[string]int64
	Subs     []*Report
}
```
#### func (*Report) Avg(key string)  float64
> 计算平均值
***
#### func (*Report) Count(key string)  int64
> 获取特定 key 的计数次数
***
#### func (*Report) Sum(keys ...string)  float64
> 获取特定 key 的总和
***
#### func (*Report) Sub(name string)  *Report
> 获取特定名称的子报告
***
#### func (*Report) ReserveSubByPrefix(prefix string)  *Report
> 仅保留特定前缀的子报告
***
#### func (*Report) ReserveSub(names ...string)  *Report
> 仅保留特定名称子报告
***
#### func (*Report) FilterSub(names ...string)  *Report
> 将特定名称的子报告过滤掉
***
#### func (*Report) String()  string
***
