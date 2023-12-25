# Lockstep [`锁步（帧）同步`](https://pkg.go.dev/github.com/kercylan98/minotaur/server/lockstep#Lockstep)


[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/server/lockstep)
> 它是一个不限制网络类型的实现，仅需要对应连接实现 [`lockstep.Client`](https://pkg.go.dev/github.com/kercylan98/minotaur/server/lockstep#Client) 接口即可。

该包提供了一个并发安全的锁步（帧）同步实现，其中内置了频率设置、帧上限、序列化、初始帧、追帧等功能。可使用其来快速构建和管理锁步（帧）同步。

锁步（帧）同步是一种特殊的同步，它可以并发安全地将数据同步到底层连接。