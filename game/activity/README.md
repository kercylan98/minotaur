# Activity

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/game/activity)

`activity` 是为不同类型的活动提供灵活的状态管理实现，支持活动的预告、开始、结束、延长展示等状态。

## 设计思路
- 为不同类型的活动提供灵活的状态管理框架，支持活动的预告、开始、结束、延长展示等状态。
- 支持事件驱动，根据活动状态变化和时间触发事件。
- 允许活动循环，并支持配置延长展示时间。
- 在多线程环境下使用互斥锁进行同步。
- 使用反射处理不同类型的活动数据。