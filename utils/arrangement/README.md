# Arrangement

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/utils/arrangement)

`Arrangement`包提供了一种灵活的方式来管理和操作编排区域。它包含了一些用于处理编排区域和编排选项的函数和类型。

## Area [`编排区域`](https://pkg.go.dev/github.com/kercylan98/minotaur/utils/arrangement#Area)
`Area`类型代表一个编排区域，它包含了一些方法，如`GetAreaInfo`、`GetItems`、`IsAllow`、`IsConflict`、`GetConflictItems`和`GetScore`，这些方法可以用来获取区域信息、获取区域中的所有成员、检查一个成员是否可以被添加到该区域中、检查一个成员是否会造成冲突、获取与一个成员产生冲突的所有其他成员以及获取该区域的评估分数。

## Option [`编排选项`](https://pkg.go.dev/github.com/kercylan98/minotaur/utils/arrangement#Option)
`Option`类型代表一个编排选项，它是一个函数，可以用来修改编排的行为。例如，`WithRetryThreshold`、`WithConstraintHandle`和`WithConflictHandle`函数可以用来设置编排时的重试阈值、约束处理函数和冲突处理函数。

## AreaOption [`编排区域选项`](https://pkg.go.dev/github.com/kercylan98/minotaur/utils/arrangement#AreaOption)
`AreaOption`类型代表一个编排区域选项，它是一个函数，可以用来修改编排区域的行为。例如，`WithAreaConstraint`、`WithAreaConflict`和`WithAreaEvaluate`函数可以用来设置编排区域的约束条件、冲突条件和评估函数。

## 示例代码
[点击查看](./arrangement_test.go)