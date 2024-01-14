# Compress

compress 提供了一些用于压缩和解压缩数据的函数。

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/compress)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[GZipCompress](#GZipCompress)|对数据进行GZip压缩，返回bytes.Buffer和错误信息
|[GZipUnCompress](#GZipUnCompress)|对已进行GZip压缩的数据进行解压缩，返回字节数组及错误信息
|[TARCompress](#TARCompress)|对数据进行TAR压缩，返回bytes.Buffer和错误信息
|[TARUnCompress](#TARUnCompress)|对已进行TAR压缩的数据进行解压缩，返回字节数组及错误信息
|[ZIPCompress](#ZIPCompress)|对数据进行ZIP压缩，返回bytes.Buffer和错误信息
|[ZIPUnCompress](#ZIPUnCompress)|对已进行ZIP压缩的数据进行解压缩，返回字节数组及错误信息


> 结构体定义

|结构体|描述
|:--|:--

</details>


#### func GZipCompress(data []byte)  bytes.Buffer,  error
<span id="GZipCompress"></span>
> 对数据进行GZip压缩，返回bytes.Buffer和错误信息
***
#### func GZipUnCompress(dataByte []byte)  []byte,  error
<span id="GZipUnCompress"></span>
> 对已进行GZip压缩的数据进行解压缩，返回字节数组及错误信息
***
#### func TARCompress(data []byte)  bytes.Buffer,  error
<span id="TARCompress"></span>
> 对数据进行TAR压缩，返回bytes.Buffer和错误信息
***
#### func TARUnCompress(dataByte []byte)  []byte,  error
<span id="TARUnCompress"></span>
> 对已进行TAR压缩的数据进行解压缩，返回字节数组及错误信息
***
#### func ZIPCompress(data []byte)  bytes.Buffer,  error
<span id="ZIPCompress"></span>
> 对数据进行ZIP压缩，返回bytes.Buffer和错误信息
***
#### func ZIPUnCompress(dataByte []byte)  []byte,  error
<span id="ZIPUnCompress"></span>
> 对已进行ZIP压缩的数据进行解压缩，返回字节数组及错误信息
***
