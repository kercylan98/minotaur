# Crypto

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/crypto)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)




## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[EncryptBase64](#EncryptBase64)|对数据进行Base64编码
|[DecodedBase64](#DecodedBase64)|对数据进行Base64解码
|[EncryptCRC32](#EncryptCRC32)|对字符串进行CRC加密并返回其结果。
|[DecodedCRC32](#DecodedCRC32)|对字节数组进行CRC加密并返回其结果。
|[EncryptMD5](#EncryptMD5)|对字符串进行MD5加密并返回其结果。
|[DecodedMD5](#DecodedMD5)|对字节数组进行MD5加密并返回其结果。
|[EncryptSHA1](#EncryptSHA1)|对字符串进行SHA1加密并返回其结果。
|[DecodedSHA1](#DecodedSHA1)|对字节数组进行SHA1加密并返回其结果。
|[EncryptSHA256](#EncryptSHA256)|对字符串进行SHA256加密并返回其结果。
|[DecodedSHA256](#DecodedSHA256)|对字节数组进行SHA256加密并返回其结果。


***
## 详情信息
#### func EncryptBase64(data []byte)  string
<span id="EncryptBase64"></span>
> 对数据进行Base64编码

***
#### func DecodedBase64(data string)  []byte,  error
<span id="DecodedBase64"></span>
> 对数据进行Base64解码

***
#### func EncryptCRC32(str string)  uint32
<span id="EncryptCRC32"></span>
> 对字符串进行CRC加密并返回其结果。

***
#### func DecodedCRC32(data []byte)  uint32
<span id="DecodedCRC32"></span>
> 对字节数组进行CRC加密并返回其结果。

***
#### func EncryptMD5(str string)  string
<span id="EncryptMD5"></span>
> 对字符串进行MD5加密并返回其结果。

***
#### func DecodedMD5(data []byte)  string
<span id="DecodedMD5"></span>
> 对字节数组进行MD5加密并返回其结果。

***
#### func EncryptSHA1(str string)  string
<span id="EncryptSHA1"></span>
> 对字符串进行SHA1加密并返回其结果。

***
#### func DecodedSHA1(data []byte)  string
<span id="DecodedSHA1"></span>
> 对字节数组进行SHA1加密并返回其结果。

***
#### func EncryptSHA256(str string)  string
<span id="EncryptSHA256"></span>
> 对字符串进行SHA256加密并返回其结果。

***
#### func DecodedSHA256(data []byte)  string
<span id="DecodedSHA256"></span>
> 对字节数组进行SHA256加密并返回其结果。

***
