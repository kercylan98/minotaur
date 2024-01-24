# Runtimes

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
|[GetWorkingDir](#GetWorkingDir)|获取工作目录绝对路径
|[GetTempDir](#GetTempDir)|获取系统临时目录
|[GetExecutablePathByBuild](#GetExecutablePathByBuild)|获取当前执行文件绝对路径（go build）
|[GetExecutablePathByCaller](#GetExecutablePathByCaller)|获取当前执行文件绝对路径（go run）
|[CurrentRunningFuncName](#CurrentRunningFuncName)|获取正在运行的函数名



</details>


***
## 详情信息
#### func GetWorkingDir() string
<span id="GetWorkingDir"></span>
> 获取工作目录绝对路径

***
#### func GetTempDir() string
<span id="GetTempDir"></span>
> 获取系统临时目录

***
#### func GetExecutablePathByBuild() string
<span id="GetExecutablePathByBuild"></span>
> 获取当前执行文件绝对路径（go build）

***
#### func GetExecutablePathByCaller() string
<span id="GetExecutablePathByCaller"></span>
> 获取当前执行文件绝对路径（go run）

***
#### func CurrentRunningFuncName(skip ...int) string
<span id="CurrentRunningFuncName"></span>
> 获取正在运行的函数名

***
