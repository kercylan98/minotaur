# Changelog

## [0.0.12](https://github.com/kercylan98/minotaur/compare/v0.0.11...v0.0.12) (2023-07-11)


### Code Refactoring | 重构

* server.WithPprof 名称修改为 server.WithPProf ([50ab92e](https://github.com/kercylan98/minotaur/commit/50ab92ef6752bf6a2ea47778680a7d7ab45e7d9c))


### Bug Fixes | 修复

* 修复配置导出 Go 代码注释错误问题 ([9f2242b](https://github.com/kercylan98/minotaur/commit/9f2242b6f7c76384d6a4cc12798331773b0ea5e3))


### Styling | 可读性优化

* 优化 server 包代码可读性 ([74c8f21](https://github.com/kercylan98/minotaur/commit/74c8f215d74c71ebc680e99a4c0c22057554a156))


### Docs | 文档优化

* server 包注释完善 ([9dc73bf](https://github.com/kercylan98/minotaur/commit/9dc73bf281b495434eebb34e28d95284238cc04a))


### Features | 新特性

* server 包 websocket 服务器支持压缩 ([6962cf4](https://github.com/kercylan98/minotaur/commit/6962cf4989edd639c7377c094c08a2dce7316e29))
* server.Server 将记录在线的连接信息，可获取到在线连接和计数等 ([8368fe0](https://github.com/kercylan98/minotaur/commit/8368fe0770fb98a81a9588aef63fd2cc8b0e77c4))

## [0.0.11](https://github.com/kercylan98/minotaur/compare/v0.0.10...v0.0.11) (2023-07-10)


### Bug Fixes | 修复

* 修复 Multiple 模式下启动服务器 listen 有时无法打印的问题 ([d972dc8](https://github.com/kercylan98/minotaur/commit/d972dc864df9024b9e21c109ad1c19ab9b38916a))
* 修复 server.Server 关闭时线程池未释放的问题 ([7228a07](https://github.com/kercylan98/minotaur/commit/7228a07e7e7275a7674ab1568bec7c1b2a8a9105))
* 修复异步慢消息追踪不生效的问题 ([7b8af05](https://github.com/kercylan98/minotaur/commit/7b8af0518e943ded08c6c53c96ca83d263a2af82))
* 修复配置字段描述换行的情况下导出的 Go 代码编译报错问题 ([e676982](https://github.com/kercylan98/minotaur/commit/e676982b9a0fb1f3b64af31e08cb90d6d355fd3b))


### Performance Improvements | 性能优化

* 调整 server.WithBufferSize 默认值 ([1ad6577](https://github.com/kercylan98/minotaur/commit/1ad657799ae09d713a5270076525887cced3c772))


### Features | 新特性

* 增加任务功能 ([bdeaa5a](https://github.com/kercylan98/minotaur/commit/bdeaa5aeb34987564a1184b8d2fc2355deaf25e8))
* 支持对 HTTP 服务器通过 server.WithPprof 开启 pprof ([53e91d1](https://github.com/kercylan98/minotaur/commit/53e91d1fce8fd5aaa365c2d18bdd2175d3f17801))
* 支持对消息增加 mark 标记，可在执行 Message.String() 函数时进行展现 ([1e6974a](https://github.com/kercylan98/minotaur/commit/1e6974ae4be51239e07a0c69091bf45506d2525a))

## [0.0.10](https://github.com/kercylan98/minotaur/compare/v0.0.9...v0.0.10) (2023-07-07)


### Tests | 新增或优化测试用例

* 移除 examples 包 ([f0e3822](https://github.com/kercylan98/minotaur/commit/f0e3822ecfcf514ee928c328ae249c51dcf62352))


### Code Refactoring | 重构

* 优化 server 消息类型，合并 Websocket 数据包监听到统一的 RegConnectionReceivePacketEvent 中 ([8b90307](https://github.com/kercylan98/minotaur/commit/8b903072b12941fd7b39b7e61741909ef31d9b26))
* 服务器支持异步消息类型、死锁阻塞、异步慢消息检测 ([1a2c1df](https://github.com/kercylan98/minotaur/commit/1a2c1df289e927e976cc9db90da557723328a9c5))
* 私有化服务器 PushMessage 函数，移除 PushCrossMessage 函数，改为使用 server.PushXXXMessage 函数 ([6d27433](https://github.com/kercylan98/minotaur/commit/6d27433c4bf933beee48644c1bc8d4d94f783675))
* 移除服务器多核和分流模式的可选项 ([7e67775](https://github.com/kercylan98/minotaur/commit/7e677751577389d675858a48ac5ece3a9fe401ba))

## [0.0.9](https://github.com/kercylan98/minotaur/compare/v0.0.8...v0.0.9) (2023-07-06)


### Bug Fixes | 修复

* 修复导出配置 JSON 特殊字符被转义的问题 ([193763e](https://github.com/kercylan98/minotaur/commit/193763e471d3e63a45e1eee4e2375cf738a9d1aa))
* 修复请求成功 server.Conn 的 callback 不调用的问题 ([8e3325f](https://github.com/kercylan98/minotaur/commit/8e3325fcd8fcaaaaf105d23249ffc5f3fa492108))
* 修复释放定时器后可能造成空指针的问题 ([9f27102](https://github.com/kercylan98/minotaur/commit/9f27102d3ae84cd9034dc8842903264112c63a50))


### Other | 其他更改

* 移除 server.Server.OnConnectionClosedEvent 和 server.Server.OnConnectionOpenedEvent 的日志 ([7065448](https://github.com/kercylan98/minotaur/commit/7065448ddfe9ffc8b09e2133df3c56726bbbdbde))


### Features | 新特性

* 支持通过 hash 包随机的读取 map 数据 ([9a35486](https://github.com/kercylan98/minotaur/commit/9a3548652a13df1bd7e6db3c9a6ebab136fb0c93))
* 支持通过 server.Server.RegStopEvent() 函数注册服务器关闭事件 ([18b9598](https://github.com/kercylan98/minotaur/commit/18b9598f5a807b1b21b380edcdf65b6cb0b88a57))

## [0.0.8](https://github.com/kercylan98/minotaur/compare/v0.0.7...v0.0.8) (2023-07-05)


### Styling | 可读性优化

* 导出日志增加已导出的表信息 ([741da79](https://github.com/kercylan98/minotaur/commit/741da79d6047fce88c19bf50785bb4bde5e66b0b))


### Performance Improvements | 性能优化

* 移除向连接发送数据时的空包处理 ([e0571c7](https://github.com/kercylan98/minotaur/commit/e0571c7ed17eadb89b251944da2c85e347501e97))


### Code Refactoring | 重构

* 由于设计不合理，移除排行榜中的 CompetitorIncrease 函数 ([0f125d4](https://github.com/kercylan98/minotaur/commit/0f125d4de5d29532e79283b3e7d51822a36e1079))


### Tests | 新增或优化测试用例

* 新增 ranking_list 测试用例，调整 aoi2d_test.go 的 packge 为 builtin_test ([b5b428d](https://github.com/kercylan98/minotaur/commit/b5b428ddc106cc1d672789a1a7ff9b1f21f6c2a3))


### Docs | 文档优化

* 排行榜 GetRank 函数增加注释，提示排名从 0 开始 ([1001d50](https://github.com/kercylan98/minotaur/commit/1001d50c04c783b4044aacb014b782f2f1be392e))


### Other | 其他更改

* 在 README.md 中添加 JetBrains OS licenses 信息 ([b234568](https://github.com/kercylan98/minotaur/commit/b234568e5653bbe32eee5149874e4217a542f480))


### Bug Fixes | 修复

* 配置加载后无限刷新修复 ([6634aa6](https://github.com/kercylan98/minotaur/commit/6634aa675ecb69e44851966561f8f4b6f3be01ad))


### Features | 新特性

* server.New 支持通过 server.WithWebsocketReadDeadline 设置超时时间 ([2513714](https://github.com/kercylan98/minotaur/commit/2513714ac44c146dfe73a2875403658a6a83d4e0))
* 可通过 slice.Merge 合并多个切片数据 ([ebfdd7c](https://github.com/kercylan98/minotaur/commit/ebfdd7c28177f15b8c79eb35e9d0c84ffeb1b680))
* 支持在重连等情况时使用 server.Conn.Reuse 函数重用连接数据 ([6144dd6](https://github.com/kercylan98/minotaur/commit/6144dd6bf057d04e94a2244bf2e2933536a069d4))
* 支持对 server.Conn 写入时调用带有 Callback 的写入函数 ([4717566](https://github.com/kercylan98/minotaur/commit/47175660de5645cb06d393f76b3d86a37cd924fe))
* 新增重试函数及两个关于 func 执行的辅助函数 ([ee87612](https://github.com/kercylan98/minotaur/commit/ee87612f60ccecade532a1345e157147597a3540))

## [0.0.7](https://github.com/kercylan98/minotaur/compare/v0.0.6...v0.0.7) (2023-07-05)


### Features | 新特性

* 导表工具导出的 Golang 代码将携带配置名称签名 ([8576d0f](https://github.com/kercylan98/minotaur/commit/8576d0f35229b555fce80110ea681e4c9f09f967))


### Code Refactoring | 重构

* 日志设置生产模式和开发模式写入文件支持开关 ([c6073a9](https://github.com/kercylan98/minotaur/commit/c6073a97a84ff2e118ee349e4ff2b3fffec1c60f))
* 重构 server.ConnectionClosedEventHandle，修复部分问题 ([e0c63d5](https://github.com/kercylan98/minotaur/commit/e0c63d569d13f6a349544bcea00af43c225d84fc))


### Bug Fixes | 修复

* 修复 server.Multiple 关闭服务器空指针异常 ([1136af4](https://github.com/kercylan98/minotaur/commit/1136af4dd87552970eb45594ddcb48ffde0c0a91))
* 配置导表部分未填写的字段导致整个表被截断问题处理 ([65aac67](https://github.com/kercylan98/minotaur/commit/65aac67cf48d4fd73440a0f1acf9fb33d27edd2a))

## [0.0.6](https://github.com/kercylan98/minotaur/compare/v0.0.5...v0.0.6) (2023-07-03)


### Features | 新特性

* 日志 log 包支持更多设置 ([83e0675](https://github.com/kercylan98/minotaur/commit/83e06759a50c7d211932ea95b6d83385f71dd6d3))

## [0.0.5](https://github.com/kercylan98/minotaur/compare/v0.0.4...v0.0.5) (2023-07-03)


### Other | 其他更改

* 删除 net 包中的不合理函数 ([f22bf5b](https://github.com/kercylan98/minotaur/commit/f22bf5bc936e6f709f43c306e38703be498acf01))


### Features | 新特性

* 为 slice 包添加更多的辅助函数 ([d4d11f2](https://github.com/kercylan98/minotaur/commit/d4d11f2a8d1eab2633bc9772c893011d7051706e))
* 配置导出生成的 Go 代码支持获取所有线上配置的集合 ([68cb5f2](https://github.com/kercylan98/minotaur/commit/68cb5f25162302e1c701ad0c81ae719f9426661b))

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
