# File

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/file)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)




## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[PathExist](#PathExist)|路径是否存在
|[IsDir](#IsDir)|路径是否是文件夹
|[WriterFile](#WriterFile)|向特定文件写入内容
|[ReadOnce](#ReadOnce)|单次读取文件
|[ReadBlockHook](#ReadBlockHook)|分块读取文件
|[ReadLine](#ReadLine)|分行读取文件
|[LineCount](#LineCount)|统计文件行数
|[Paths](#Paths)|获取指定目录下的所有文件路径
|[ReadLineWithParallel](#ReadLineWithParallel)|并行的分行读取文件并行处理，处理过程中会将每一行的内容传入 handlerFunc 中进行处理
|[FindLineChunks](#FindLineChunks)|查找文件按照每行划分的分块，每个分块的大小将在 chunkSize 和分割后的分块距离行首及行尾的距离中范围内
|[FindLineChunksByOffset](#FindLineChunksByOffset)|该函数与 FindLineChunks 类似，不同的是该函数可以指定 offset 从指定位置开始读取文件


***
## 详情信息
#### func PathExist(path string)  bool,  error
<span id="PathExist"></span>
> 路径是否存在

***
#### func IsDir(path string)  bool,  error
<span id="IsDir"></span>
> 路径是否是文件夹

***
#### func WriterFile(filePath string, content []byte)  error
<span id="WriterFile"></span>
> 向特定文件写入内容

***
#### func ReadOnce(filePath string)  []byte,  error
<span id="ReadOnce"></span>
> 单次读取文件
>   - 一次性对整个文件进行读取，小文件读取可以很方便的一次性将文件内容读取出来，而大文件读取会造成性能影响。

***
#### func ReadBlockHook(filePath string, bufferSize int, hook func (data []byte))  error
<span id="ReadBlockHook"></span>
> 分块读取文件
>   - 将filePath路径对应的文件数据并将读到的每一部分传入hook函数中，当过程中如果产生错误则会返回error。
>   - 分块读取可以在读取速度和内存消耗之间有一个很好的平衡。

***
#### func ReadLine(filePath string, hook func (line string))  error
<span id="ReadLine"></span>
> 分行读取文件
>   - 将filePath路径对应的文件数据并将读到的每一行传入hook函数中，当过程中如果产生错误则会返回error。

***
#### func LineCount(filePath string)  int
<span id="LineCount"></span>
> 统计文件行数

***
#### func Paths(dir string)  []string
<span id="Paths"></span>
> 获取指定目录下的所有文件路径
>   - 包括了子目录下的文件
>   - 不包含目录

***
#### func ReadLineWithParallel(filename string, chunkSize int64, handlerFunc func ( string), start ...int64) (n int64, err error)
<span id="ReadLineWithParallel"></span>
> 并行的分行读取文件并行处理，处理过程中会将每一行的内容传入 handlerFunc 中进行处理
>   - 由于是并行处理，所以处理过程中的顺序是不确定的。
>   - 可通过 start 参数指定开始读取的位置，如果不指定则从文件开头开始读取。

***
#### func FindLineChunks(file *os.File, chunkSize int64)  [][2]int64
<span id="FindLineChunks"></span>
> 查找文件按照每行划分的分块，每个分块的大小将在 chunkSize 和分割后的分块距离行首及行尾的距离中范围内
>   - 使用该函数得到的分块是完整的行，不会出现行被分割的情况
>   - 当过程中发生错误将会发生 panic
>   - 返回值的成员是一个长度为 2 的数组，第一个元素是分块的起始位置，第二个元素是分块的结束位置

***
#### func FindLineChunksByOffset(file *os.File, offset int64, chunkSize int64)  [][2]int64
<span id="FindLineChunksByOffset"></span>
> 该函数与 FindLineChunks 类似，不同的是该函数可以指定 offset 从指定位置开始读取文件

***
