# Reflects

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
|[WrapperFunc](#WrapperFunc)|包装函数
|[WrapperFuncBefore2After](#WrapperFuncBefore2After)|包装函数，前置函数执行前，后置函数执行后
|[WrapperFuncBefore](#WrapperFuncBefore)|包装函数，前置函数执行前
|[WrapperFuncAfter](#WrapperFuncAfter)|包装函数，后置函数执行后
|[GetPtrUnExportFiled](#GetPtrUnExportFiled)|获取指针类型的未导出字段
|[SetPtrUnExportFiled](#SetPtrUnExportFiled)|设置指针类型的未导出字段
|[Copy](#Copy)|拷贝
|[GetPointer](#GetPointer)|获取指针



</details>


***
## 详情信息
#### func WrapperFunc(f any, wrapper func (call func ( []reflect.Value)  []reflect.Value)  func (args []reflect.Value)  []reflect.Value) (wf Func, err error)
<span id="WrapperFunc"></span>
> 包装函数

***
#### func WrapperFuncBefore2After(f Func, before func (), after func ()) (wf Func, err error)
<span id="WrapperFuncBefore2After"></span>
> 包装函数，前置函数执行前，后置函数执行后

***
#### func WrapperFuncBefore(f Func, before func ()) (wf Func, err error)
<span id="WrapperFuncBefore"></span>
> 包装函数，前置函数执行前

***
#### func WrapperFuncAfter(f Func, after func ()) (wf Func, err error)
<span id="WrapperFuncAfter"></span>
> 包装函数，后置函数执行后

***
#### func GetPtrUnExportFiled(s reflect.Value, filedIndex int) reflect.Value
<span id="GetPtrUnExportFiled"></span>
> 获取指针类型的未导出字段

***
#### func SetPtrUnExportFiled(s reflect.Value, filedIndex int, val reflect.Value)
<span id="SetPtrUnExportFiled"></span>
> 设置指针类型的未导出字段

***
#### func Copy(s reflect.Value) reflect.Value
<span id="Copy"></span>
> 拷贝

***
#### func GetPointer(src T) reflect.Value
<span id="GetPointer"></span>
> 获取指针

***
