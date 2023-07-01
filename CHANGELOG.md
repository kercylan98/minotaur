# Changelog

## [0.0.4](https://github.com/kercylan98/minotaur/compare/v0.0.3...v0.0.4) (2023-07-01)


### Bug Fixes | 修复

* 多服务器情况下日志错乱及无法正常 Shuntdown 问题修复 ([67616b2](https://github.com/kercylan98/minotaur/commit/67616b29632e62a00107f3c0d80dc4a4609afe11))


### Tests | 新增或优化测试用例

* components.Moving2D 增加示例 ([01bafe6](https://github.com/kercylan98/minotaur/commit/01bafe6fc030a00fd9ec9a4282754aeb7e9e00bc))
* components.Moving2D 测试用例优化 ([49bc143](https://github.com/kercylan98/minotaur/commit/49bc143a72bce85e84e717894c0dfd693a945691))


### Features | 新特性

* components.Moving2D 支持停止移动事件注册 ([f67a66d](https://github.com/kercylan98/minotaur/commit/f67a66d2d0b5543635f075a8a30e534ea76d99cf))
* 对 poker.Rule 提供功能的辅助函数 ([0172c67](https://github.com/kercylan98/minotaur/commit/0172c67a44f0bb969abbec1f3d4e15b785a1d484))
* 服务器支持通过 server.WithDiversion 可选项对数据包消息进行分流处理 ([73cefc9](https://github.com/kercylan98/minotaur/commit/73cefc9b48be3c0f537b4d0ed93b5b73087701da))


### Code Refactoring | 重构

* 导表工具重构，增加部分特性，修复部分问题 ([afdda79](https://github.com/kercylan98/minotaur/commit/afdda793bc46a496dafd8ac493e275a462b6ee74))

## [0.0.3](https://github.com/kercylan98/minotaur/compare/v0.0.2...v0.0.3) (2023-06-30)


### Bug Fixes | 修复

* 修复 file.ReadOnce 读文件错误 ([b0ae569](https://github.com/kercylan98/minotaur/commit/b0ae56991be4ad584550edd64985207b005ed0d5))


### Features | 新特性

* generic 包支持检查泛型类型是否为空指针 ([6023f59](https://github.com/kercylan98/minotaur/commit/6023f591608efa64ee543884917b6f3fc72f1d05))
* maths 包支持比较一组数是否连续 ([0ab38c7](https://github.com/kercylan98/minotaur/commit/0ab38c7023d37da81967d03392d8dfdb8d715c89))
* timer.Ticker 支持附加标记信息 ([db51edf](https://github.com/kercylan98/minotaur/commit/db51edfa1cc44932a357d0d2b7d7dc2a934938f9))
* 增加时间段 times.Period 数据结构 ([a6ca8a9](https://github.com/kercylan98/minotaur/commit/a6ca8a9f9ee00f599879800ae3d5ce259605848d))


### Code Refactoring | 重构

* 重构 poker 包设计，移除 Poker 结构体，以 Rule 结构体进行取代 ([d1b7699](https://github.com/kercylan98/minotaur/commit/d1b7699cb4790098e3eb7bf093b1c6d1a1f0242e))
* 重构游戏活动实现 ([390e8e7](https://github.com/kercylan98/minotaur/commit/390e8e75efe9b13abba3e5215de780f05e83a5aa))


### Tests | 新增或优化测试用例

* 完善测试用例 ([741a25c](https://github.com/kercylan98/minotaur/commit/741a25cf42a2b76d14e4b72283c1e699c83e48df))

## [0.0.2](https://github.com/kercylan98/minotaur/compare/v0.0.1...v0.0.2) (2023-06-27)


### Features | 新特性

* 增加时间转换辅助函数 ([05a328e](https://github.com/kercylan98/minotaur/commit/05a328e34493b5c96c88b8ff285e2f3107aae6e0))
* 增加更多的时间处理函数 ([2127978](https://github.com/kercylan98/minotaur/commit/2127978093ec53a07d704857dee83c3df3137038))
* 支持获取全局偏移时间 ([77e7d46](https://github.com/kercylan98/minotaur/commit/77e7d468838fb405a387cd9200b74cc970ca02b1))
* 新增全局偏移时间 ([6c4f59f](https://github.com/kercylan98/minotaur/commit/6c4f59f1a0baf54e4bcd7ac13d3ecad06d9e3792))
* 新增游戏活动功能支持 ([83531b6](https://github.com/kercylan98/minotaur/commit/83531b65c6d9b9ffc23247ea0dc86ce6a1214aae))


### Bug Fixes | 修复

* 修复使用 int 和 math.MaxUint 比较导致溢出的问题 ([a4e9b5f](https://github.com/kercylan98/minotaur/commit/a4e9b5f14397e095c20c9f63e33d88a8cd87bfa5))

## 0.0.1 (2023-06-26)

### Features | 新特性

* 支持通过 server 包支持快速创建 TCP、UDP、Websocket、UnixSock、HTTP、GRPC、KCP 服务器
* 支持通过 router 包创建最多支持三级的路由器
* 支持通过 cross 对 server 创建的服务器提供跨服支持
* 通过 configexport 包提供了针对策划及开发人员的配置表模板及导表工具，支持导出 json 和 go 配置文件
* 支持通过 notify 包快速实现通知功能，默认支持飞书群聊机器人通知
* 组件 component 包中提供了帧同步组件的实现及 2D 移动组件的实现
* 支持通过 report 包实现快捷的数据上报功能
* utils 包中提供了大量常用的辅助函数
  * asynchronization 包中提供了实现了 hash.Map 的非并发安全 map 数据结构
  * compress 包中提供了 gzip 压缩与解压缩的算法
  * crypto 包中支持对数据进行 base64、crc、md5、sha1、sha256 的编码解码函数
  * file 包中提供了常用的文件操作函数
  * generic 包中提供了常见的泛型约束
  * geometry 包中提供了几何相关的处理函数，包括线、形状、点等内容
    * astar 包中提供了 A* 算法的实现
    * dp 包中提供了基于二维数组的分布链接的机制，可以快速查找与给定成员具有相同特征且位置紧邻的其他成员
    * matrix 包中提供了一个简单的二维矩阵实现
    * navmesh 包提供了基于 astar 的网格寻路功能
  * hash 包提供了常用了 hashmap 转换、接口等功能
  * huge 包提供了 int 类型的大整数实现
  * log 包中提供了基于 zap 的默认日志组件
  * maths 包中提供了常用的数学处理函数
  * network 包中提供了常用的网络辅助函数
  * offset 包中提供了带偏移的时间实现
  * random 包中提供了常用的随机函数，包括随机 hash、名称等
  * runtimes 包中提供了常用的运行时辅助函数
  * slice 包中提供了基于切片的辅助函数
  * sole 包中提供了 guid 和 雪花id 的实现
  * str 包中提供了常用的字符串处理函数
  * super 包中提供了 if 的三目表达式函数
  * synchronization 包中提供了并发安全的数据结构
  * timer 包中提供了定时器组件
  * times 包中提供了常用的时间处理函数
