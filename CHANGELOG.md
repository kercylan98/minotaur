# Changelog

## [0.3.1](https://github.com/kercylan98/minotaur/compare/v0.3.0...v0.3.1) (2023-11-13)


### Other | 其他更改

* 日志调用修改 ([dd3f3ed](https://github.com/kercylan98/minotaur/commit/dd3f3ede078e6d67ea1b3711054196a6f15e1477))


### Features | 新特性

* server 包新增机器人，可通过 server.NewBot 函数进行创建，机器人将模拟普通连接行为，适用于测试等场景 ([4c092c0](https://github.com/kercylan98/minotaur/commit/4c092c04d2151b9764851ab9838ce06c069069f2))
* server 新增 Unique 异步消息，可用于避免相同标识的异步消息在未执行完毕前重复执行 ([e2b7887](https://github.com/kercylan98/minotaur/commit/e2b7887b142be1217572e6f2e487554eedc5010e))
* super 新增 ConditionalRetryByExponentialBackoff 函数，支持可中断的退避指数算法重试 ([274402e](https://github.com/kercylan98/minotaur/commit/274402e721f9b04f7a2ed64c6920f456a3b4df91))


### Bug Fixes | 修复

* 修复配置导出工具无法忽略描述前缀为 # 的字段 ([5c180de](https://github.com/kercylan98/minotaur/commit/5c180de1188692a3e493a89f551d3262ffb52f64))


### Docs | 文档优化

* 优化配置导出工具部分文档描述 ([30c0b3a](https://github.com/kercylan98/minotaur/commit/30c0b3a64bc611885dbca54f190720f51069933e))

## [0.3.0](https://github.com/kercylan98/minotaur/compare/v0.2.9...v0.3.0) (2023-11-11)


### Features | 新特性

* super 包新增 Hostname 函数获取主机名 ([9157c6a](https://github.com/kercylan98/minotaur/commit/9157c6a309d8561b9ed701bb8c1d15383f2d371a))
* super.LossCounter 支持打印 ([01092fe](https://github.com/kercylan98/minotaur/commit/01092fe738c3040a61869693145e437de59cc0da))
* times 包支持设置全局时间偏移 ([f03dd4a](https://github.com/kercylan98/minotaur/commit/f03dd4ac4ff12f0f05f46554ea7c9b785dc5f74f))
* times 包支持重置全局时间偏移量和获取当前全局时间偏移量 ([707fc6c](https://github.com/kercylan98/minotaur/commit/707fc6c5de283af320c46e6d1dc978ad38d86299))
* 修复配置导表工具数组处理异常的问题 ([0f966c0](https://github.com/kercylan98/minotaur/commit/0f966c02f7bf18db5e6b1d4bc49ba7bdecae6c55))
* 增强 server.RegConsoleCommandEvent 函数，支持 url 格式输入命令，并将命令解析为指令和参数 ([d2654cf](https://github.com/kercylan98/minotaur/commit/d2654cfc950bf5ee90a2856dfef368d2bbbc8604))


### Bug Fixes | 修复

* 更新配置导表工具数组处理异常的问题 ([24ba13c](https://github.com/kercylan98/minotaur/commit/24ba13cab215d086b26405d34991b6c9bff2898e))
* 示例代码适配当前版本 ([ab72920](https://github.com/kercylan98/minotaur/commit/ab72920084232fff2306391dde51d00d3f3f1e21))


### Code Refactoring | 重构

* server 包重构及性能优化 ([70f7a79](https://github.com/kercylan98/minotaur/commit/70f7a79c888fe80484b88866cf836f8f4533bb61))

## [0.2.9](https://github.com/kercylan98/minotaur/compare/v0.2.8...v0.2.9) (2023-11-09)


### Other | 其他更改

* xkeys seal encryption used fixed key for all encryption [#6](https://github.com/kercylan98/minotaur/issues/6) ([2079e95](https://github.com/kercylan98/minotaur/commit/2079e9595e1782b28415b36e92c6be7d3dfa1f1c))


### Features | 新特性

* generic 包新增 Basic 类型 ([d405cae](https://github.com/kercylan98/minotaur/commit/d405cae73f527e636f544fffb6e8f9b16965d2ce))
* lockstep 支持获取帧同步客户端数量 ([589a424](https://github.com/kercylan98/minotaur/commit/589a424491dc5b150bed49bc33d1849030dec373))
* server 包支持获取到 HTTP 服务器的 Gin 示例 ([6b2a753](https://github.com/kercylan98/minotaur/commit/6b2a753e67f3605db2cdeb0760fbf30db937037b))
* server.Server 支持使用 PushAsyncMessage 快捷发布异步消息 ([0b77cc9](https://github.com/kercylan98/minotaur/commit/0b77cc9907210ff527d570527071e57cc1804d3c))
* super 包新增规则重试及退避指数重试 ([d191dab](https://github.com/kercylan98/minotaur/commit/d191dabfd3b23c5da171d24d49b69ee646123cf0))
* survey 包的 Analyzer 分析器增加大量辅助函数 ([85176f3](https://github.com/kercylan98/minotaur/commit/85176f32f918b3ad02ace7fa6217b55b6d457d0e))


### Code Refactoring | 重构

* [#60](https://github.com/kercylan98/minotaur/issues/60) 重构 game/task 包，支持更灵活的任务配置方式 ([98c1f39](https://github.com/kercylan98/minotaur/commit/98c1f39ce6a40abb4097842c545216dc180342a4))

## [0.2.8](https://github.com/kercylan98/minotaur/compare/v0.2.7...v0.2.8) (2023-10-31)


### Other | 其他更改

* gRPC-Go HTTP/2 Rapid Reset vulnerability、NATS.io: Adding accounts for just the system account adds auth bypass ([e4d60d7](https://github.com/kercylan98/minotaur/commit/e4d60d7146fcef13c5a729e85faa18eea5debecf))


### Features | 新特性

* server.Conn 支持获取连接打开时间及在线时长 ([18a0b06](https://github.com/kercylan98/minotaur/commit/18a0b06e0ebf8ac0ae5feb841bacfaf2bef9fa66))
* survey 包新增 RecordBytes 函数，支持跳过格式化将数据直接写入，适用于转发至消息队列等场景 ([f475aac](https://github.com/kercylan98/minotaur/commit/f475aac387883ce86d82f00e63bd51cc4bfcdcf8))
* survey.FileFlusher 将会在目录不存在时自行创建 ([d2f982b](https://github.com/kercylan98/minotaur/commit/d2f982bf42aa8c8e026240418ad01ec9cf7ccb5d))


### Bug Fixes | 修复

* lockstep.Lockstep 移除不必要的内容，修复 StartBroadcast 函数锁使用不正确的问题 ([61d41e5](https://github.com/kercylan98/minotaur/commit/61d41e51b5fe49a0ebc787666c1a5f010c573a53))
* 优化 exporter 配置导出器在没有前缀时会默认增加一个 "." 的问题 ([fb5dacb](https://github.com/kercylan98/minotaur/commit/fb5dacb4b4e3bbc871b684b892596f916bac6789))
* 修复 [#58](https://github.com/kercylan98/minotaur/issues/58) taskType 及事件被遗漏的问题 ([9f88265](https://github.com/kercylan98/minotaur/commit/9f882651eb385d6cb328ee857a00980c8076f23e))
* 修复 timer.GetTicker 在获取到定时器后立刻使用造成的竞态问题 ([a4bc828](https://github.com/kercylan98/minotaur/commit/a4bc8280a46a371e59ca656058b7dcb15d545b21))

## [0.2.7](https://github.com/kercylan98/minotaur/compare/v0.2.6...v0.2.7) (2023-10-23)


### Reverts | 回退

* round 并发安全问题回撤（死锁问题） ([6e11c5e](https://github.com/kercylan98/minotaur/commit/6e11c5edec0ebcec8dd4826da195f8db6adf754c))


### Features | 新特性

* concurrent 包 新增 NewMapPool 函数，支持创建 map 对象池 ([74a6b54](https://github.com/kercylan98/minotaur/commit/74a6b545c23034c96d87785d5f209d3569ce29ae))
* fight 包新增 TurnBased 回合制数据结构，用于替代 fight.Round。解决并发安全问题，并且支持按照速度进行回合切换 ([378f855](https://github.com/kercylan98/minotaur/commit/378f855992f9a03eba4853b056f5a0327c669085))
* fight.TurnBased 支持监听回合变更以及刷新当前操作回合超时时间 ([ba2f3af](https://github.com/kercylan98/minotaur/commit/ba2f3af39855b5d860e6483f0281e430742f591b))
* generic 包新增 IDR、IDW、IDRW 的泛型通用接口 ([5259e07](https://github.com/kercylan98/minotaur/commit/5259e07a320c055d15410c96ab35b094d569d19c))
* lockstep 支持丢弃帧缓存 ([803dd4f](https://github.com/kercylan98/minotaur/commit/803dd4f2eb192e268a0fa486f1b83ed269c7e86a))
* server.Conn 支持在 WebSocket 模式下通过 GetWebsocketRequest 函数获取到请求 ([42ab52b](https://github.com/kercylan98/minotaur/commit/42ab52bc668d28194a56f2c9de09833381b36d8f))
* server.Conn 支持通过 GetServer 获取到服务器实例 ([89e9c51](https://github.com/kercylan98/minotaur/commit/89e9c517afb68c405b25c37eead8aafbbd6fbe82))
* super 包新增 LaunchTime 函数，支持获取程序启动时间 ([20f62fe](https://github.com/kercylan98/minotaur/commit/20f62fee87b34b27e31c90ee3f373b795e2fefb1))
* super 包新增 LossCounter，适用于统计代码段时间损耗，可通过 super.StartLossCounter 函数进行使用 ([2b49a36](https://github.com/kercylan98/minotaur/commit/2b49a36e8ef1679d595193a8cd9b4ae7c4164be5))


### Bug Fixes | 修复

* server 包数据竞态问题优化 ([cdbf388](https://github.com/kercylan98/minotaur/commit/cdbf38849824000e224f284dcfe56bae8f986090))
* 修复 concurrent.Pool 可选项无法使用的问题 ([64544e0](https://github.com/kercylan98/minotaur/commit/64544e069d5e91d8da5bbc475fffe18c7fdb7a7a))
* 修复 lockstep.WithInitFrame 不生效的问题 ([859e0a1](https://github.com/kercylan98/minotaur/commit/859e0a1ac14cf7763aa1630814eddf127c1bf266))


### Code Refactoring | 重构

* cross 包服务器 id 更改为 string 类型 ([9e33906](https://github.com/kercylan98/minotaur/commit/9e339065d453635dbc03d088a02634b829e104c6))


### Performance Improvements | 性能优化

* lockstep 包优化同步逻辑，帧 id 由 int 更改为 int64 类型，优化数据竞态问题 ([d3e5632](https://github.com/kercylan98/minotaur/commit/d3e563257f8e72569f76132002c8ea73e5fe39b0))
* lockstep 包优化帧命令逻辑，去除多余字段 ([139fe42](https://github.com/kercylan98/minotaur/commit/139fe4291a229e944f97d1aec3dec001c1612c8f))
* server 包异步消息不再执行额外 defer ([b5b126e](https://github.com/kercylan98/minotaur/commit/b5b126ef073dff627ce533a3bb0ca7f1c0f8c0da))

## [0.2.6](https://github.com/kercylan98/minotaur/compare/v0.2.5...v0.2.6) (2023-10-16)


### Features | 新特性

* super 包新增简单的权限控制器，可通过 super.NewPermission 函数进行创建 ([9e00684](https://github.com/kercylan98/minotaur/commit/9e0068490268aa9ede61832657e3f243b89d24b7))
* 新增 space 包及 space.RoomMananger 结构体，提供了更便于使用的房间结构，用于取代 room 包 ([c3538ab](https://github.com/kercylan98/minotaur/commit/c3538ab530dc70875f91663905bfb6c3d1f32514))


### Bug Fixes | 修复

* 修复 fight.Round 在回合内执行 ActionRefresh 等操作的并发问题 ([2d1e8f1](https://github.com/kercylan98/minotaur/commit/2d1e8f14952171c3aa12cb9ceca7a60b0150f573))

## [0.2.5](https://github.com/kercylan98/minotaur/compare/v0.2.4...v0.2.5) (2023-10-12)


### Features | 新特性

* server.Conn 支持通过 ViewData 函数查看只读的连接数据 ([e60e0a7](https://github.com/kercylan98/minotaur/commit/e60e0a754a4cef3416e8699d40bf95e470b32692))
* 优化 game.Player 的 Send 和 Close 函数与 server.Conn 同步 ([f65a155](https://github.com/kercylan98/minotaur/commit/f65a1555f64f55650b1aac52e00d3f1759c5b97b))


### Bug Fixes | 修复

* HTTP/2 rapid reset can cause excessive work in net/http ([14f542e](https://github.com/kercylan98/minotaur/commit/14f542e5130c9a70da667b4a9f6bb31fa3278fb1))
* random 包按权重产生结果更改为 int64 ([433ba08](https://github.com/kercylan98/minotaur/commit/433ba08c754dcca3d5b4e69e757b8de411207284))
* 修复 poker.CardPile.Reset 函数导致牌组只有大小王的问题 ([fb60065](https://github.com/kercylan98/minotaur/commit/fb60065ec1d1cc0d0890c6c1767dd82b2c28517e))
* 修复 room 包在使用 AddSeat 函数时无法加入空缺位置的问题 ([295aaeb](https://github.com/kercylan98/minotaur/commit/295aaeb4c04d4be3ac6503453265cc136f8a6c3c))
* 修复 room.Manager.GetRoom 函数的空指针问题 ([039500b](https://github.com/kercylan98/minotaur/commit/039500ba87c6706ad84841b00bb7d5d8004f89e7))

## [0.2.4](https://github.com/kercylan98/minotaur/compare/v0.2.3...v0.2.4) (2023-10-08)


### Features | 新特性

* 新增 xlsx 配置导出工具及模板，可手动编译后使用 ([b622175](https://github.com/kercylan98/minotaur/commit/b6221752cacf71aebfe44a389ad31345a1c69274))


### Docs | 文档优化

* README.md 增加配置道具工具相关说明 ([9435ba5](https://github.com/kercylan98/minotaur/commit/9435ba5ecb1330625e0eb331ab5a37ca4648ca52))
* 部分注释优化 ([83ab553](https://github.com/kercylan98/minotaur/commit/83ab55373417b4e4c1940fa7e24c3fbb279e3cb3))

## [0.2.3](https://github.com/kercylan98/minotaur/compare/v0.2.2...v0.2.3) (2023-10-07)


### Features | 新特性

* stream 新增 Maps，以及快捷开启流操作的函数 With... ([cb3bd11](https://github.com/kercylan98/minotaur/commit/cb3bd11248b658294fd76fe40ac8b4fc48a7a524))
* super 包支持通过 MarshalToTargetWithJSON 将对象通过 JSON 序列化为目标对象 ([2e4ab44](https://github.com/kercylan98/minotaur/commit/2e4ab441228d4c9e8940dd2162c0a674f3dc69f3))
* timer 包新增部分获取 分、日、月、年 开始结束时间函数，以及快捷创建时间窗口时间段的函数 ([05f0016](https://github.com/kercylan98/minotaur/commit/05f0016b7ed453f451155b596c32603d6b648313))


### Docs | 文档优化

* README 增加流操作文档 ([ba02fd4](https://github.com/kercylan98/minotaur/commit/ba02fd4accd47dd12ce6319ebbd0ae10e6409adb))

## [0.2.2](https://github.com/kercylan98/minotaur/compare/v0.2.1...v0.2.2) (2023-09-21)


### Reverts | 回退

* 设计不合理原因移除 storage 包 ([d9b9392](https://github.com/kercylan98/minotaur/commit/d9b939295c19fa1910437940a500128a1460b3a1))


### Features | 新特性

* client.Run 支持传入 block 参数指定客户端以阻塞的模式运行 ([534a7e9](https://github.com/kercylan98/minotaur/commit/534a7e962ad6258df277f1f8214f9124975ebcae))
* super 包增加 RetryForever 函数，支持永久重试直到成功 ([13c5483](https://github.com/kercylan98/minotaur/commit/13c5483617223ca6876e31a24deccae6a2d60383))
* 新增 memory 包，适用于游戏数据加载到内存中并周期性持久化 ([ed008cf](https://github.com/kercylan98/minotaur/commit/ed008cf280727f8c40053d6e9b968f66a7ae851a))


### Tests | 新增或优化测试用例

* 新增 times.CalcNextSecWithTime 示例代码 ([149e6a2](https://github.com/kercylan98/minotaur/commit/149e6a2149aedc3b27f049d5940d2727f01a8395))

## [0.2.1](https://github.com/kercylan98/minotaur/compare/v0.2.0...v0.2.1) (2023-09-19)


### Other | 其他更改

* gateway 优化代码逻辑，适配 client.Client 变更 ([0cc8fd8](https://github.com/kercylan98/minotaur/commit/0cc8fd818614a8836f35255c185975872bce797f))


### Features | 新特性

* buffer.Unbounded 增加新的构造函数，支持省略 generateNil 函数，新增 IsClosed 函数检查无界缓冲区是否已经关闭 ([e9bc9fb](https://github.com/kercylan98/minotaur/commit/e9bc9fb48175dc6544a570fd82d71af66ca8f801))
* concurrent.Pool 新增静默模式可选项 WithPoolSilent，在该模式下当缓冲区大小不足时，将不再输出警告日志 ([3ad1330](https://github.com/kercylan98/minotaur/commit/3ad1330cd937635ce56f1ca70365836f19c97fc8))
* random 包通过权重和概率随机产生一个成员支持返回产生成员的索引或 Key ([782a1ad](https://github.com/kercylan98/minotaur/commit/782a1adb37da396028987c3ff35917c0ddf8b4e2))
* 新增 writeloop 包，内置了一个写循环的实现 ([dd1acfd](https://github.com/kercylan98/minotaur/commit/dd1acfd017e9f0eccbc23663fa7f871a6b2b7de4))


### Bug Fixes | 修复

* super 包优化 GetError 函数的空指针问题 ([ab3926e](https://github.com/kercylan98/minotaur/commit/ab3926e307bed9d665010f6613409ed093e256fe))


### Docs | 文档优化

* 修复 server 在 WebSocket 模式下超时时间无效的问题 ([1bc32e2](https://github.com/kercylan98/minotaur/commit/1bc32e2026da59283b9302ec753699919a994cce))


### Styling | 可读性优化

* server 包为服务器启动添加 IP 信息，死锁检测的日志内容优化 ([42465a8](https://github.com/kercylan98/minotaur/commit/42465a8f42cb3a3515263318ac41403473476cb4))


### Code Refactoring | 重构

* client 包采用无界缓冲区替代通过 chan 实现的写通道，移除消息堆积功能，优化代码逻辑 ([2d9ffad](https://github.com/kercylan98/minotaur/commit/2d9ffad2ab0277c0a83842c3d27ca31b820a51de))
* server 移除 WithConnMessageChannelSize 可选项 ([31c0e1b](https://github.com/kercylan98/minotaur/commit/31c0e1b7356e062ee741cd4aeacca8f96b62953d))


### Performance Improvements | 性能优化

* server.Conn 写循环更改为采用无界缓冲区的写入，优化整体逻辑 ([551a3e5](https://github.com/kercylan98/minotaur/commit/551a3e5c51c048eac13bf31c0fc1665d2b7b8431))

## [0.2.0](https://github.com/kercylan98/minotaur/compare/v0.1.7...v0.2.0) (2023-09-18)


### Features | 新特性

* buffer 包新增 Unbounded 实现 ([d56c1df](https://github.com/kercylan98/minotaur/commit/d56c1df6e1d9f6f1f2f01d8f461858ec9c9139c2))
* random 包新增 ProbabilitySlice 函数，用于基于概率产生一个结果，当概率总和小于 1 会发生未命中的情况，概率总和大于 1 将等比缩放至 1 ([7c9bc46](https://github.com/kercylan98/minotaur/commit/7c9bc46a3506d5722da7b0062b9b493f709fbb97))
* 新增 buffer 包，内置了一个环形缓冲区的实现 ([12d1aba](https://github.com/kercylan98/minotaur/commit/12d1abab9aa4f09bf9c0f7b5628c50155f2aaf4e))


### Bug Fixes | 修复

* server 修复慢消息导致堆栈溢出的问题 ([e95e1ba](https://github.com/kercylan98/minotaur/commit/e95e1ba3997d57dc5eb7f8bd99e1d6adb7da19c2))
* 修复 gnet 作为服务器核心关闭时导致的空指针问题 ([2712f3b](https://github.com/kercylan98/minotaur/commit/2712f3b98e2d871e5bf2101e338c50b343fe05e0))


### Docs | 文档优化

* 优化文档内容兼容 WithShunt ([00eaa36](https://github.com/kercylan98/minotaur/commit/00eaa362262f158fd71bc72c8c0522a5b6bb0d0e))


### Code Refactoring | 重构

* server 包重构消息通道，采用无界缓冲区替代原本的 chan，解决消息通道的缓冲区达到上限时造成永久阻塞的问题，移除 WithMessageChannelSize 可选项，修改 WithShunt 可选项不再需要 channelGenerator 参数 ([810a9fd](https://github.com/kercylan98/minotaur/commit/810a9fdb73f4460f181ff7a60937615a6c926db8))

## [0.1.7](https://github.com/kercylan98/minotaur/compare/v0.1.6...v0.1.7) (2023-09-12)


### Features | 新特性

* router 包新增 Multistage 多级分类路由器，用于替代原有的 1~3 级路由器 ([10cc443](https://github.com/kercylan98/minotaur/commit/10cc443b3af307111af92ec850d8a3f1f277355c))
* router.Multistage 支持通过可选项创建 ([637ae27](https://github.com/kercylan98/minotaur/commit/637ae2788219d9cb23698953b8ca6967f59087ad))
* server 包新增 WithLimitLift 和 WithConnectionMessageChannelSize 函数，用于限制服务器最大生命周期及连接的消息写入通道大小 ([064d434](https://github.com/kercylan98/minotaur/commit/064d434a0cb777665248b825662db32f58a339a9))
* slice.Priority 优先级队列新增 Slice 函数，获取队列成员为切片类型 ([30dbb14](https://github.com/kercylan98/minotaur/commit/30dbb14addd714081b9a2e98880d96d7b9bf4229))
* utils 下新增 sorts.Topological 拓扑排序函数 ([7a5e2c1](https://github.com/kercylan98/minotaur/commit/7a5e2c1e7e5e14c7820871adba269985d01bd129))


### Bug Fixes | 修复

* 修复 super.RegError 和 RegErrorRef 空指针问题 ([82973dd](https://github.com/kercylan98/minotaur/commit/82973dd11adc4648d30bbe36b8f272fe77a6031f))


### Docs | 文档优化

* gateway 包注释优化 ([5103103](https://github.com/kercylan98/minotaur/commit/5103103fb5f1e480c1b0204d5dc98e149c7c36c7))
* 弃用文档优化 ([a0d5fc8](https://github.com/kercylan98/minotaur/commit/a0d5fc860ae402c5504994c9d78110782dd2c2c0))


### Performance Improvements | 性能优化

* server 包 websocket SetReadDeadline 优化 ([dc3c7d2](https://github.com/kercylan98/minotaur/commit/dc3c7d2eeaa9743400d156758d69f2bab87858a8))


### Tests | 新增或优化测试用例

* slice 包新增部分单元测试 ([4982e6d](https://github.com/kercylan98/minotaur/commit/4982e6d7b691c16b634d6d79c1cf5119eaf89524))

## [0.1.6](https://github.com/kercylan98/minotaur/compare/v0.1.5...v0.1.6) (2023-09-09)


### Features | 新特性

* survey 包 Report 新增 Avg、Count、Sum 等辅助计算函数 ([8fd4e8f](https://github.com/kercylan98/minotaur/commit/8fd4e8f722dc7b6813912d9852428971ee4ddfe8))


### Bug Fixes | 修复

* 修复 server.Conn 和 client.Client 连接关闭时发生的竞态问题 ([0215c54](https://github.com/kercylan98/minotaur/commit/0215c5449ae5c46890fca362279f94debb72aa29))
* 修复 server.Conn 连接关闭时发生的竞态问题 ([674c38a](https://github.com/kercylan98/minotaur/commit/674c38a066ac17719019a70c02a6de11a3847208))


### Code Refactoring | 重构

* 重构 super 包中的 error 部分，优化设计不合理的地方，支持动态注册错误码，支持并发读写 ([de7b085](https://github.com/kercylan98/minotaur/commit/de7b085cf7bed61637a55a4d4f7010de581ee244))


### Performance Improvements | 性能优化

* 调整 super 包 error 部分为使用后再申请内存 ([83b2800](https://github.com/kercylan98/minotaur/commit/83b28003c85f0caa87f5d48697bef7f3072ee58a))

## [0.1.5](https://github.com/kercylan98/minotaur/compare/v0.1.4...v0.1.5) (2023-09-08)


### Features | 新特性

* slice 包新增 Zoom 函数， stream 包支持 Zoom 函数 ([62ef35a](https://github.com/kercylan98/minotaur/commit/62ef35a518c259142679d171f53060d0cef79d13))
* stream.Slice 新增 Indexes 和 Map 函数 ([5024022](https://github.com/kercylan98/minotaur/commit/5024022366aaa52cfdd36afc5440266baa633021))
* survey 分析记录支持通过 GetTime 函数获取记录时间 ([3c3dc83](https://github.com/kercylan98/minotaur/commit/3c3dc83830e7843ba09fdc3ed2a9ad9d7e099d95))
* 优化 slice 包 Filter 和 Map 函数，新增 Reduce 函数 ([5ab9902](https://github.com/kercylan98/minotaur/commit/5ab990246ddb7059bc83ec65f485cb7bbb1ded22))
* 新增大量 slice 包和 hash 包的辅助函数 ([d772409](https://github.com/kercylan98/minotaur/commit/d7724094d19943303b9bbe2b61fa8cb3e595c7c8))


### Bug Fixes | 修复

* server 修复消息计数始终为1的问题 ([6c882ed](https://github.com/kercylan98/minotaur/commit/6c882edb09dcd3d7979da42d951eddb63bc6f240))
* 修复 server 关闭服务器后，如果等待消息结束过程中，新消息将阻塞的问题 ([19df61b](https://github.com/kercylan98/minotaur/commit/19df61b97fc17f5dc7fdcf04d6d23cb72aaa1500))
* 修复 survey.Analyzer 去重 BUG ([a4ba3f1](https://github.com/kercylan98/minotaur/commit/a4ba3f1fa86ab2ad682c28f6e3ab0258099b4ac6))


### Styling | 可读性优化

* 常量调整为从 1 开始 ([5fb1dcb](https://github.com/kercylan98/minotaur/commit/5fb1dcbcea0c56aeafd271e3d7ff3c8cd1eece9b))


### Code Refactoring | 重构

* 优化 server 包消息日志，移除 server.Conn.Reuse 函数（不合理） ([376ff77](https://github.com/kercylan98/minotaur/commit/376ff779e129f2ced628f48e4cffdad507def19d))
* 重构 stream 包，提供更便捷的使用方式 ([d72f185](https://github.com/kercylan98/minotaur/commit/d72f18590bec72f6321fb990f1428a12c30c00e6))


### Performance Improvements | 性能优化

* server 包连接关闭逻辑优化 ([483ace2](https://github.com/kercylan98/minotaur/commit/483ace2fa9e1d60069fb6dff234505efd0fc4cd6))


### Tests | 新增或优化测试用例

* 新增 stream.Slice 测试用例 ([d9b68fc](https://github.com/kercylan98/minotaur/commit/d9b68fc037a5fdf068c9d3f3d42785ccf12a8928))

## [0.1.4](https://github.com/kercylan98/minotaur/compare/v0.1.3...v0.1.4) (2023-09-06)


### Features | 新特性

* server 包 Server 新增 RegMessageReadyEvent 函数 ([04c40bf](https://github.com/kercylan98/minotaur/commit/04c40bf87379f3216c1eb6dcb36b44f4b1fd0ee0))
* slice 包新增 Mapping 函数，支持将切片中的元素进行转换 ([da68945](https://github.com/kercylan98/minotaur/commit/da68945f7eea9806bf1e3c3fe3c015b997f11596))


### Bug Fixes | 修复

* client 包错误类型转换错误问题处理 ([034ca17](https://github.com/kercylan98/minotaur/commit/034ca174b6461e15f078420d5dcc7172113ee477))
* 修复 server 包 Server.RegConsoleCommandEvent 函数在无终端环境下导致 CPU 飙升的问题 ([3e35e73](https://github.com/kercylan98/minotaur/commit/3e35e73c9094fba66c61853a8c41cfa36bba10cf))


### Docs | 文档优化

* README.md 增加部分示例 ([e5bf7f3](https://github.com/kercylan98/minotaur/commit/e5bf7f31207831153ba2f36d743ca18ca1331fc4))


### Code Refactoring | 重构

* survey 包 AllWithPath 函数更改为 Analyze，新增分析报告，及分析器，提供方便的统计功能 ([ac11e9e](https://github.com/kercylan98/minotaur/commit/ac11e9e9727990a831296f20ffc306a2408cbef1))


### Performance Improvements | 性能优化

* 优化 concurrent.Pool 池对象不够用的日志打印为 1 秒一次，而不是频繁打印 ([989b9da](https://github.com/kercylan98/minotaur/commit/989b9da33d282369b5771621b0eb7d6fe03dd6c0))


### Tests | 新增或优化测试用例

* 优化 server 服务器并发测试用例 ([4214ea4](https://github.com/kercylan98/minotaur/commit/4214ea4c2b57858cffb2da6fd10190140f1fd0d3))

## [0.1.3](https://github.com/kercylan98/minotaur/compare/v0.1.2...v0.1.3) (2023-09-05)


### Features | 新特性

* server 包新增 Server.RegMessageExecBeforeEvent 函数，支持在消息执行前进行处理，适用于限流等场景 ([0297c44](https://github.com/kercylan98/minotaur/commit/0297c4444aba9f13c7b60276c3b54f83d3ab8174))


### Bug Fixes | 修复

* 修复 server 包 RegMessageExecBeforeEvent 函数导致内存泄露的问题 ([15a4918](https://github.com/kercylan98/minotaur/commit/15a491816a26afdcda9f241de20740ccc8e27d83))
* 修复宕机问题 ([add1e4b](https://github.com/kercylan98/minotaur/commit/add1e4bc8c07f934a419da16c5f9edbf4bf88535))

## [0.1.2](https://github.com/kercylan98/minotaur/compare/v0.1.1...v0.1.2) (2023-09-01)


### Features | 新特性

* counter 包增加简单去重计数器 ([0d76507](https://github.com/kercylan98/minotaur/commit/0d765075e07b414a3940d643db273332ae79b404))
* gateway 支持连接与某一端点保持持久通讯，支持将端点的所有连接切换到另一端点 ([6d5aa59](https://github.com/kercylan98/minotaur/commit/6d5aa599d76ac3e297077781401e039df6562ec7))
* gateway 新增 WithEndpointConnectionPoolSize 支持配置与端点建立连接的数量 ([3ca6ed0](https://github.com/kercylan98/minotaur/commit/3ca6ed00ec91c34a4a61a61dcfd5731da8faba66))
* super 包新增函数 IsNumber，用于判断一个值是否为数字类型 ([518d47a](https://github.com/kercylan98/minotaur/commit/518d47ae6a13eda45cd7d650d5e07477869c2eff))


### Bug Fixes | 修复

* server 包中 RegConsoleCommandEvent 函数空指针问题处理 ([2ed52fc](https://github.com/kercylan98/minotaur/commit/2ed52fc814130a57b3c356214052069c094d7bed))
* server 包修复服务器关闭后发送消息引起的空指针问题 ([3062428](https://github.com/kercylan98/minotaur/commit/3062428051b075ccb53f1758d5f905b047401af1))
* survey 包修复 All 函数无用的返回参 ([c6f8c19](https://github.com/kercylan98/minotaur/commit/c6f8c190862e9af516dda9338225de2c960c3b2f))
* 修复 server 中 HTTP 服务器关闭时会引发空指针的问题 ([8cd9979](https://github.com/kercylan98/minotaur/commit/8cd9979e2be17e6043959f7206bf254071397b72))


### Code Refactoring | 重构

* 优化 survey，移除 All 函数，新增 Flusher 接口，可自行实现其他持久化方式 ([d9ba1bc](https://github.com/kercylan98/minotaur/commit/d9ba1bc85caa0c3a7453515b2e5452cbfb31c640))


### Performance Improvements | 性能优化

* 提高消息吞吐量，降低消息延迟 ([1cbe8ec](https://github.com/kercylan98/minotaur/commit/1cbe8ecf56430318ca1f5a190e311cbb1bcbb2a4))

## [0.1.1](https://github.com/kercylan98/minotaur/compare/v0.1.0...v0.1.1) (2023-08-24)


### Other | 其他更改

* 关闭 release-as ([75a8608](https://github.com/kercylan98/minotaur/commit/75a8608bf4143183a2525be6d76e0fde0ccdc4c7))


### Features | 新特性

* client 包增加 tcp 客户端 ([a3bb100](https://github.com/kercylan98/minotaur/commit/a3bb10012ed266b2fd1800154da4dcb960dd558e))
* gateway.Endpoint 支持设置重连间隔 ([cdfecb4](https://github.com/kercylan98/minotaur/commit/cdfecb41e84632f7d0ca429e39f45e49355f8368))
* survey.Reg 支持每次写入就持久化的策略 ([3fc282a](https://github.com/kercylan98/minotaur/commit/3fc282afabcff881df5e98ed4883fbf438e46156))


### Bug Fixes | 修复

* client 包内存溢出、死循环等问题处理 ([08559d8](https://github.com/kercylan98/minotaur/commit/08559d822506bc5695fafae531260ee69447d9bf))
* 修复 tcp、udp、uds 等类型服务器数据包会额外增加一个长度的问题 ([124635c](https://github.com/kercylan98/minotaur/commit/124635c72c64a870be0b05dd672f76a7343ff361))


### Styling | 可读性优化

* 错误的命名更正 ([1558b28](https://github.com/kercylan98/minotaur/commit/1558b2868d0b1d5d475987442268bac21b57f3e0))


### Code Refactoring | 重构

* gateway 整体优化重构 ([30e7894](https://github.com/kercylan98/minotaur/commit/30e7894a376ff66ca33faa74bcbdfb77576715b8))
* storage 包重构，优化整体设计 ([0ad8a5c](https://github.com/kercylan98/minotaur/commit/0ad8a5c7d54513af612b4056c316277fe1cf6bd0))

## [0.1.0](https://github.com/kercylan98/minotaur/compare/v0.0.31...v0.1.0) (2023-08-23)


### Other | 其他更改

* 版本调整至 0.1.0 ([74899af](https://github.com/kercylan98/minotaur/commit/74899af44443c90201d03400f24557314bbbf612))
* 移除 qodana workflow ([7fa369f](https://github.com/kercylan98/minotaur/commit/7fa369fd8b3b1694167cccc4f7d810510cfa7e1c))


### Features | 新特性

* survey.R 增加更多的辅助函数 ([4785c60](https://github.com/kercylan98/minotaur/commit/4785c60c5e93804c76988fffba6f21b06414a738))
* 新增 survey 包，包含了运营日志的基本功能实现 ([e962009](https://github.com/kercylan98/minotaur/commit/e962009efffcdbda4abb433761349704f0296d17))


### Bug Fixes | 修复

* [#40](https://github.com/kercylan98/minotaur/issues/40) uint 类型溢出问题处理 ([ed45d1a](https://github.com/kercylan98/minotaur/commit/ed45d1a643647d60c4d430a4c5710719d8f7a17b))


### Code Refactoring | 重构

* 调整 server 数据包相关处理函数的接收参数不再为 server.Packet，同时移除 server.Packet ([4850dd4](https://github.com/kercylan98/minotaur/commit/4850dd4aa3bccacb92bf4d866db236c7930635e6))


### Performance Improvements | 性能优化

* survey 包整体优化 ([50f6b1b](https://github.com/kercylan98/minotaur/commit/50f6b1b085887bfc985b33d384cd3a7c3248ef09))


### Build System | 影响构建的修改

* 更新依赖版本 ([c9ff457](https://github.com/kercylan98/minotaur/commit/c9ff4570fe786ca17c82bf32d75846d74c51911c))

## [0.0.31](https://github.com/kercylan98/minotaur/compare/v0.0.30...v0.0.31) (2023-08-22)


### Other | 其他更改

* server 异步消息回调将不再使用 MessageTypeSystem，更改为 MessageTypeAsyncCallback ([811e1bd](https://github.com/kercylan98/minotaur/commit/811e1bd29ec4c4859a439c7bdc9655cd8abea635))
* 调整 log.Duration 和 log.DurationP 函数为 String 调用 ([a1c15a2](https://github.com/kercylan98/minotaur/commit/a1c15a2c26d22babe27d9b64fae7bf52bfdddbd9))


### Reverts | 回退

* 设计原因移除 report 包，采用 utils/counter ([7cbe5c4](https://github.com/kercylan98/minotaur/commit/7cbe5c4805585ab9d06ad0e6ae3d553e57b77f06))


### Features | 新特性

* concurrent.Pool 新增 EAC 函数，用于动态调整缓冲区大小。优化超出缓冲区大小警告日志，增加堆栈信息，用于定位高频点 ([64ecd45](https://github.com/kercylan98/minotaur/commit/64ecd459a1b29a4dceadf9b09fad265e1043b5cf))
* hash 包增加 Clear 函数，用于清空 map ([7f316d4](https://github.com/kercylan98/minotaur/commit/7f316d4a7a918d7fdf6bcb28e9a3fec2e3772efe))
* server/client 新增 Unix Domain Socket 客户端 ([3de1f6b](https://github.com/kercylan98/minotaur/commit/3de1f6b9d3ece2bc33c162886da1ec562f8f5512))
* slice 包新增 Filter 函数用于过滤切片 ([ab19bd6](https://github.com/kercylan98/minotaur/commit/ab19bd6f6ac95c9ee0ee1ea656cba50c8b56a830))
* super 包新增 StringToFloat64 函数 ([89c32b4](https://github.com/kercylan98/minotaur/commit/89c32b4ce3f187c78e5673a4bdca885e0ca3563d))
* super 包新增大量 string 和 number 基本类型转换的辅助函数 ([d37fbb7](https://github.com/kercylan98/minotaur/commit/d37fbb7aa2dfd3839b3909c9be5ea3bb76e7da65))


### Bug Fixes | 修复

* counter 包修复 mark key 无法被清理、重置的问题 ([a31369a](https://github.com/kercylan98/minotaur/commit/a31369abbe4f47f7630e2f11071edba9ec9e6376))
* 优化 go1.21 以下项目的兼容性 ([ab90fa8](https://github.com/kercylan98/minotaur/commit/ab90fa8928151554324f07f27126ddb104682bb2))
* 优化 uds 客户端无法正常接收数据包的问题 ([6792e22](https://github.com/kercylan98/minotaur/commit/6792e227c010a2e543e6dd402da2e733a3ab7ffe))
* 修复非 gateway 数据包导致数组越界的问题 ([5096e6f](https://github.com/kercylan98/minotaur/commit/5096e6f88458b108887815e05693d3ff6292d305))


### Code Refactoring | 重构

* 调整事件函数名称 ([dc76196](https://github.com/kercylan98/minotaur/commit/dc761964b9923a10b519b1fb04cc2330689629ed))

## [0.0.30](https://github.com/kercylan98/minotaur/compare/v0.0.29...v0.0.30) (2023-08-21)


### Features | 新特性

* gateway 数据包支持像普通数据包一样处理，并且支持自定义端点健康评估函数 ([3512570](https://github.com/kercylan98/minotaur/commit/351257033ec6cbf13a6e172eb36800e7816b0dc5))
* server 包支持设置注册事件的优先级 ([3c6ce9c](https://github.com/kercylan98/minotaur/commit/3c6ce9cfdf5a40c00edc9f132d3209d81802f816))
* slice 包新增 GetValue 和 GetValueHandle 函数，用于获取特定索引的元素，如果索引超出范围将返回零值 ([2dd5dd5](https://github.com/kercylan98/minotaur/commit/2dd5dd5c6c96809dc8623863be6dba8adaabe25c))
* slice 包新增优先级切片 ([93e63b1](https://github.com/kercylan98/minotaur/commit/93e63b1ace526018c767c24dd1c2fa75716e3c79))
* 新增 counter 包，用于创建支持特定时间内去重的计数器 ([1005d74](https://github.com/kercylan98/minotaur/commit/1005d7458d2e32c45a170449b0fe5b276d6d39ea))


### Bug Fixes | 修复

* 修复 count.Shadow 函数死锁问题 ([34ca7f0](https://github.com/kercylan98/minotaur/commit/34ca7f07d25ad4fa3896e5ba7515c3c0ddeef5e8))
* 修复 websocket 客户端死锁问题 ([7bf4e82](https://github.com/kercylan98/minotaur/commit/7bf4e82183c7f1b259e12cf329796812d5da296f))

## [0.0.29](https://github.com/kercylan98/minotaur/compare/v0.0.28...v0.0.29) (2023-08-17)


### Features | 新特性

* server 新增 NetworkNone 网络类型，该模式下不监听任何网络端口，仅开启消息队列，适用于纯粹的跨服服务器等情况 ([dcfb3da](https://github.com/kercylan98/minotaur/commit/dcfb3da534b15ee2d3792ea5fcff61e669be058d))
* server.Server 新增 HttpServer 函数，用于替代 HttpRouter ([b87df07](https://github.com/kercylan98/minotaur/commit/b87df072fc0b3982ce8c144ad32ad1ff5b77a414))
* server.Server 的 HTTP 请求支持慢消息检测 ([36a3333](https://github.com/kercylan98/minotaur/commit/36a333379e1d5f7906003089d994fa04575b882b))


### Bug Fixes | 修复

* server 包优化 Shutdown 逻辑，修复服务器关闭时不会等待消息执行完毕的问题 ([93c5f36](https://github.com/kercylan98/minotaur/commit/93c5f3695f8e43e84aa7b94d6bedb4d9f4bf0a9b))

## [0.0.28](https://github.com/kercylan98/minotaur/compare/v0.0.27...v0.0.28) (2023-08-16)


### Features | 新特性

* gateway 网关支持通过可选项自定义端点选择器 ([e0f43c5](https://github.com/kercylan98/minotaur/commit/e0f43c5bfb96654fb682de22ad07af91c8c40958))
* server 目录中新增 client 包，提供了 Websocket 客户端实例 ([322938a](https://github.com/kercylan98/minotaur/commit/322938accf969509967c81e29f026aeca3af1d33))
* server 目录中新增 gateway 包，提供了基本的 Websocket 网关实现 ([5ff74b6](https://github.com/kercylan98/minotaur/commit/5ff74b623d13734eb65ad9f7d72a749297427e6a))
* server.Server 新增 RegConnectionPacketPreprocessEvent 函数用于对数据包进行预处理操作 ([b3e4bb6](https://github.com/kercylan98/minotaur/commit/b3e4bb6166c3abd3faffb49cddfd19fce5afc0e1))
* times 包增加部分时间处理函数 ([157b6b5](https://github.com/kercylan98/minotaur/commit/157b6b5aafb645ee61336b47c046b96f746d4e11))


### Bug Fixes | 修复

* 修复 timer.Ticker.Loop 函数首次触发时会触发两次的问题 ([2bd6aa5](https://github.com/kercylan98/minotaur/commit/2bd6aa50cbbbc370b93b32c867fadaaf1a18bb02))

## [0.0.27](https://github.com/kercylan98/minotaur/compare/v0.0.26...v0.0.27) (2023-08-14)


### Features | 新特性

* ranking.List 新增支持默认值的获取排名和分数的函数 ([57ee7ff](https://github.com/kercylan98/minotaur/commit/57ee7ff3ef634f754945c701fd4ae7336290d53b))
* sole 包新增 Once 结构体，用于数据取值去重 ([0f31173](https://github.com/kercylan98/minotaur/commit/0f31173291efafd4cb5594f56d54fb68903179a6))


### Bug Fixes | 修复

* 修复配置显示声明的字符串时，导出的数据包含双引号的问题 ([31cd79c](https://github.com/kercylan98/minotaur/commit/31cd79c2218674fee5431d6b25586f937f76d716))

## [0.0.26](https://github.com/kercylan98/minotaur/compare/v0.0.25...v0.0.26) (2023-08-10)


### Features | 新特性

* arrangement 新增冲突、冲突处理函数、约束处理函数 ([84f36ea](https://github.com/kercylan98/minotaur/commit/84f36eaabaaafa072777c393eafd27d03e6ebf2a))
* arrangement.Engine 新增更多的辅助函数 ([822ffc7](https://github.com/kercylan98/minotaur/commit/822ffc7041c3a4d97c457cff0a6fc0da5183f17a))
* server 包新增 HTTP 包装器 ([cec7e5b](https://github.com/kercylan98/minotaur/commit/cec7e5b341d508aaead33a471ec86b665dd8a8c5))
* 新增 reflects 包，包含反射相关辅助函数 ([340b00e](https://github.com/kercylan98/minotaur/commit/340b00eb76135bb5323d9736906f4a19ea4a82f2))


### Bug Fixes | 修复

* http 包装器 group 修复 ([dbf7ed7](https://github.com/kercylan98/minotaur/commit/dbf7ed717ab6b6305013eeb3d5bf515d73b8acb0))


### Build System | 影响构建的修改

* 升级 go 至 1.21 版本 ([9596320](https://github.com/kercylan98/minotaur/commit/9596320e6508c87616a8e202aab3e3db64252a50))

## [0.0.25](https://github.com/kercylan98/minotaur/compare/v0.0.24...v0.0.25) (2023-08-03)


### Features | 新特性

* combination 包新增 Validator 校验器，用于校验组合是否匹配，取代 poker.Rule ([f6873bd](https://github.com/kercylan98/minotaur/commit/f6873bd5dc59af9ec2997029ba30063d18b15238))
* combination 包新增 WithValidatorHandleNCarryM、WithValidatorHandleNCarryIndependentM 函数 ([87a1ca9](https://github.com/kercylan98/minotaur/commit/87a1ca90bd80f9a2ff4ef06d56d3e7c0ce77a4b3))
* room.Helper 支持通过 BroadcastExcept 向被排除表达式命中外的玩家广播消息 ([0804508](https://github.com/kercylan98/minotaur/commit/08045088e612009ccad91eb120c835365e552b06))
* 新增 arrangement 包，用于针对多条数据进行合理编排的数据结构 ([1f5f95a](https://github.com/kercylan98/minotaur/commit/1f5f95ae6de5df7849318ae2b27c79689f240d77))


### Bug Fixes | 修复

* combination.WithValidatorHandleNCarryM 修复 M 允许类型不同的问题 ([0db1e5c](https://github.com/kercylan98/minotaur/commit/0db1e5c30b768a4918f4b2be9e3f03cfe29d9f8e))
* room.Helper.BroadcastExcept 函数返回值修复 ([faac7b2](https://github.com/kercylan98/minotaur/commit/faac7b27bbd57e7602a62acd699bc556e748d9c5))


### Docs | 文档优化

* poker 包过时标记 ([553c436](https://github.com/kercylan98/minotaur/commit/553c4362e3160e76bf0866092802fd2001e3118f))
* README.md 及 CONTRIBUTING.md 完善 ([7cfdbb1](https://github.com/kercylan98/minotaur/commit/7cfdbb12a4b62d4c618261a70ee567055dae80ff))

## [0.0.24](https://github.com/kercylan98/minotaur/compare/v0.0.23...v0.0.24) (2023-08-02)


### Features | 新特性

* fight.Round 新增操作刷新事件 ([d96ed58](https://github.com/kercylan98/minotaur/commit/d96ed58548ed87ec0e2730ed90aa32b11a2c3394))
* fight.Round 新增获取当前操作超时时间的函数 ([060fb05](https://github.com/kercylan98/minotaur/commit/060fb05fb8cdeff4008706527806193f808d48f4))
* random 包新增 Dice 掷骰子和 Probability 概率函数 ([d9d0392](https://github.com/kercylan98/minotaur/commit/d9d0392db39cff582d1af1e78286d62570bb1979))
* room.Helper 新增获取玩家切片、广播所有玩家、广播在座玩家的函数 ([ab180f3](https://github.com/kercylan98/minotaur/commit/ab180f384b8ee6272e2d8abe21dd73e802007bed))
* server.Server 支持通过 WithShunt 函数对服务器消息进行分流 ([c92f16c](https://github.com/kercylan98/minotaur/commit/c92f16c17060d940346b17000d5d59fd660269e7))
* server.Server 新增分流通道创建和关闭事件 ([b9d9533](https://github.com/kercylan98/minotaur/commit/b9d953338f7efdac1d9ca97c7494a3ff0718adcd))
* 新增 deck 包，用于对牌堆、麻将牌堆、一组数据等情况的管理 ([ace17a6](https://github.com/kercylan98/minotaur/commit/ace17a6a76b7b4324135f1e5a476dead6a7281e3))


### Bug Fixes | 修复

* configuration 包字段类型转换修复 ([aef7740](https://github.com/kercylan98/minotaur/commit/aef7740f5c0d44325aadfc17edf1b565c5d16fa5))
* 修复 room 包中通过 Manager 获取 Helper 时，当传入的 room 为空依旧会返回不为空指针的 Helper 问题 ([e8c2cf2](https://github.com/kercylan98/minotaur/commit/e8c2cf28357dbff94293b8a9247ba6de084467b8))


### Code Refactoring | 重构

* moving2d 移动到 game 包中 ([e3224d0](https://github.com/kercylan98/minotaur/commit/e3224d010b0017a3e1eb80f0c15f002778e0b9f9))
* 移除 component 包，lockstep 迁移至 server/lockstep ([1b8d041](https://github.com/kercylan98/minotaur/commit/1b8d041ae0b5400c008a1e255f80a096a56bb425))


### Tests | 新增或优化测试用例

* fight.Round 单元测试函数名变更 ([ffd8d04](https://github.com/kercylan98/minotaur/commit/ffd8d047f9cd101d52a25cffd9f35dce9a25144a))

## [0.0.23](https://github.com/kercylan98/minotaur/compare/v0.0.22...v0.0.23) (2023-08-01)


### Other | 其他更改

* 优化 combination 包命名，删除无用文件 ([57936b2](https://github.com/kercylan98/minotaur/commit/57936b2b25426055de409659b5f5a2a018f9031e))


### Reverts | 回退

* 移除 poker 包的 matcher，改为使用 combination 包 ([8b92921](https://github.com/kercylan98/minotaur/commit/8b929212303e020db4476842566449f7a3b605fc))


### Features | 新特性

* fight 包的 Round 新增操作超时事件，优化事件逻辑 ([9198faa](https://github.com/kercylan98/minotaur/commit/9198faa06140404f947bf954e36d3d94975ee46a))
* maths 包支持奇偶数判断 ([ac43963](https://github.com/kercylan98/minotaur/commit/ac43963a864a74c499450838ed7f1d8c53700826))
* room 包新增房间创建事件 ([87c6695](https://github.com/kercylan98/minotaur/commit/87c66954a3ea1215b587aa3a22b464e6d2066321))
* 新增 combination 包，用于数组组合筛选（抽离自 poker 包） ([48d9c11](https://github.com/kercylan98/minotaur/commit/48d9c1131627087b39456b9f376d3148942ad259))
* 新增 fight 包，提供了回合制战斗的功能实现 ([df8f6fc](https://github.com/kercylan98/minotaur/commit/df8f6fc53e5bfdc481351c962f4f10a3585d3796))


### Bug Fixes | 修复

* 修复 server 异步消息的 callback 的并发问题 ([1297ae7](https://github.com/kercylan98/minotaur/commit/1297ae7a8f246f8929131b299ba6cfcffc585c4e))
* 修复泛型对象 player 不能判断 nil 的表达式错误 ([4dddd14](https://github.com/kercylan98/minotaur/commit/4dddd1422bc00f40be43050f53cd7525f9a73341))
* 修复牌堆重置时不会重置 guid 的问题 ([39ccad4](https://github.com/kercylan98/minotaur/commit/39ccad42411774058e14a520fcfc16960e22a9f5))
* 状态机 fsm 包名修复，优化注释 ([cee067e](https://github.com/kercylan98/minotaur/commit/cee067e246942024acf44261a3e7d549b4b85b7a))
* 状态机 State 名称修复 ([de76411](https://github.com/kercylan98/minotaur/commit/de76411726854f0f11ffad405ded8dc5e1b89ec4))


### Docs | 文档优化

* server.PushAsyncMessage 注意事项补全 ([2482d2e](https://github.com/kercylan98/minotaur/commit/2482d2e7f0dcfd3bea2be2474102dcd7b10d6da5))


### Code Refactoring | 重构

* fsm 包状态机事件优化，新增部分获取状态机信息的函数 ([0fad041](https://github.com/kercylan98/minotaur/commit/0fad0417c7cbd27a228b199c58c209de71ebbb0f))


### Performance Improvements | 性能优化

* 优化 combination 包 NCarryM 性能 ([abd1db5](https://github.com/kercylan98/minotaur/commit/abd1db55860a26d3cad0d12e7cf7aa66304e852a))
* 优化 slice.Combinations 效率 ([03028b1](https://github.com/kercylan98/minotaur/commit/03028b1a41567b2a9bfa1f4c4f8d4d5e6cc4264c))

## [0.0.22](https://github.com/kercylan98/minotaur/compare/v0.0.21...v0.0.22) (2023-07-28)


### Features | 新特性

* maths 包新增支持 int64 的数字合并函数 ([a6fb7fb](https://github.com/kercylan98/minotaur/commit/a6fb7fb8dc69962ef3c969c66775574ca1a3f081))
* room 支持获取座位上的玩家数量 ([24f54a1](https://github.com/kercylan98/minotaur/commit/24f54a1536620d88de5319c1dee14b85cdfe3a61))
* super 包支持使用 Convert 强制转换数据类型 ([867d1ec](https://github.com/kercylan98/minotaur/commit/867d1ecf82d95cc1f468bf190d1018367c1362ef))
* times 包新增 SystemNewDay 和 OffsetTimeNewDay 事件 ([2a0c5b8](https://github.com/kercylan98/minotaur/commit/2a0c5b84a83e8bf6db6263dcd329cf225ed8d79f))


### Bug Fixes | 修复

* fms 包迁移问题处理 ([996f5af](https://github.com/kercylan98/minotaur/commit/996f5af8bd48e998987146a8615f6a795692381f))


### Code Refactoring | 重构

* room 包移除大量 error 返回，增加易于房间操作 Helper 数据结构，可通过 Manager.GetHelper 和 room.NewHelper 获取 ([3dec407](https://github.com/kercylan98/minotaur/commit/3dec4075d5929dcd4a064350dcdfbe8e3287b7e4))


### Tests | 新增或优化测试用例

* test:  ([930fe15](https://github.com/kercylan98/minotaur/commit/930fe159bffc22e15f462febcadf89ce7a4648ff))

## [0.0.21](https://github.com/kercylan98/minotaur/compare/v0.0.20...v0.0.21) (2023-07-27)


### Reverts | 回退

* 移除 attrs，设计不合理 ([87f26dd](https://github.com/kercylan98/minotaur/commit/87f26dd394ad99f48cd75dda61cbce6e946ab733))
* 移除 gameplay，设计不合理 ([41ea022](https://github.com/kercylan98/minotaur/commit/41ea0222612972f925746bf06bc1f4441176a11d))
* 移除 terrain 和 world，设计不合理 ([361e269](https://github.com/kercylan98/minotaur/commit/361e269f125eb81176c36d3f816495dddd75c667))


### Features | 新特性

* generic 包支持更多的空指针判断函数 ([d06c840](https://github.com/kercylan98/minotaur/commit/d06c840c463810f56b2023751ea15261c5298b85))
* hash 包新增 Set 数据结构 ([9fcc75e](https://github.com/kercylan98/minotaur/commit/9fcc75e0d7545fcbdb65f87ec9e1a12b03b7bce0))
* maths 包新增 CountDigits 和 GetDigitValue 函数，用于计算一个数字的位数和获取特定位数上的值 ([3f94f38](https://github.com/kercylan98/minotaur/commit/3f94f38e99d304eaa02a74d5bd8063c75919bbf0))
* room 包添加更多的事件，添加座位号支持 ([c8f181f](https://github.com/kercylan98/minotaur/commit/c8f181f63eaad5310d263621f222985baad35fd1))
* server 异步消息支持将 callback 设置为 nil ([b63975e](https://github.com/kercylan98/minotaur/commit/b63975ea09cff8510118b0772cca66452168a6ff))
* server.Server 事件消息添加 mark 标记，方便问题定位 ([471ee48](https://github.com/kercylan98/minotaur/commit/471ee48644eee5e5b527c5ad8e24761498bfdce5))
* server.Server 新增 ConnectionOpenedAfterEvent ([8dde18a](https://github.com/kercylan98/minotaur/commit/8dde18a36ed99e02a45c7b63e8c0d8887447ea78))
* server.Server 新增对连接写入事件前的处理函数 ([5e26467](https://github.com/kercylan98/minotaur/commit/5e26467deef2e2dcf6d0b04c918e59193942d432))
* slice 包新增 CombinationsPari 函数，用于从给定的两个数组中按照特定数量得到所有组合后，再将两个数组的组合进行组合 ([d26ef3a](https://github.com/kercylan98/minotaur/commit/d26ef3aca6ded00f91bc912488453948dbe3d9c2))
* super 包支持无错的 json 序列化 ([11ad997](https://github.com/kercylan98/minotaur/commit/11ad997eaa4bb16e0a1e64f967761ed5e1c6a7c6))
* 房间管理器实现 ([45c855a](https://github.com/kercylan98/minotaur/commit/45c855a5160e1918707c2a6bef422b261486af72))


### Bug Fixes | 修复

* 修复 room.NewManager 没有初始化 rp 字段的问题 ([5c3c959](https://github.com/kercylan98/minotaur/commit/5c3c9592c538ec4d3c1b757d9f0482ee6b266abb))


### Docs | 文档优化

* game 包文档优化 ([054b3a7](https://github.com/kercylan98/minotaur/commit/054b3a7ec9f2e30adf61f1c0db77778b790608c7))


### Code Refactoring | 重构

* kcrypto 包更名为 crypto，与目录名对应 ([1ae14f0](https://github.com/kercylan98/minotaur/commit/1ae14f0d7be64ae3c8eb8d522f29a26442e50f7d))
* RankingList 更名为 List，并且移动至 ranking 包中 ([ed8ee4a](https://github.com/kercylan98/minotaur/commit/ed8ee4a542228278376d5592a18775fa8b5bd6d6))
* 从 builtin 包中单独抽离到 aoi 包，更名为 TwoDimensional ([bca8a98](https://github.com/kercylan98/minotaur/commit/bca8a98463ba19fa9722e486fd612757123cfe78))
* 状态机从 builtin 包中单独抽离到 fsm 包 ([6fb24da](https://github.com/kercylan98/minotaur/commit/6fb24da8c186db0a567cb4527ed7ba3610bc3f79))
* 移除原有的 builtin 中的各类 room 实现 ([ee18934](https://github.com/kercylan98/minotaur/commit/ee18934768507a621406399d9b4c2e4f5d5ccfa7))

## [0.0.20](https://github.com/kercylan98/minotaur/compare/v0.0.19...v0.0.20) (2023-07-25)


### Reverts | 回退

* 移除 storage 包，不合理的设计 ([3e956b6](https://github.com/kercylan98/minotaur/commit/3e956b64cf097894fac6aba8c4bc0f103bd705c7))


### Features | 新特性

* super 包支持注册第三方错误，将第三方错误转换为特定错误代码和信息 ([2cbffbf](https://github.com/kercylan98/minotaur/commit/2cbffbf967aef46b58596ea89924c09ce54470d9))
* super 包添加 []byte、string 零拷贝转换函数 ([506e0f2](https://github.com/kercylan98/minotaur/commit/506e0f2ee411e91d96695880cd81d2acc41464af))


### Code Refactoring | 重构

* map 移除适配 ([d446ff1](https://github.com/kercylan98/minotaur/commit/d446ff18b97aa2534303049396257dcca6b22b48))
* storage 中的 Delete 要求返回 error ([a43fb4f](https://github.com/kercylan98/minotaur/commit/a43fb4faea167b7a94ef7714ce9e69c18ca06b01))
* storage 包重新实现 ([b6f28dd](https://github.com/kercylan98/minotaur/commit/b6f28dd7431ca0d59292b9c3f993ae23320db63b))
* storage 要求 Load 等函数返回错误信息 ([0d1a985](https://github.com/kercylan98/minotaur/commit/0d1a985e691fdc4f6af7bc4c23fab7687fc86238))
* 优化 solo.guid 的使用，命名空间需要注册 ([6238883](https://github.com/kercylan98/minotaur/commit/6238883dc97839b089fb36544252614c4d5860ff))
* 去除 storage 中的 errHandle 参数 ([3befe64](https://github.com/kercylan98/minotaur/commit/3befe645b71473799e73127f76a2e91c7e67fa5e))
* 移除分段锁map实现及 hash.Map、hash.ReadonlyMap 接口，移除 asynchronous 包，同步包更名为 concurrent ([d0d2087](https://github.com/kercylan98/minotaur/commit/d0d2087fee823e5821fe6c88c871bb94e5fa69cc))
* 重构 poker 包为全泛型包，支持通过 poker.Matcher 根据一组扑克牌选出最佳组合 ([d71d843](https://github.com/kercylan98/minotaur/commit/d71d8434b6c431327fd405535843ca52c65c9973))

## [0.0.19](https://github.com/kercylan98/minotaur/compare/v0.0.18...v0.0.19) (2023-07-20)


### Bug Fixes | 修复

* 修复 onStop 无法等待逻辑执行完成的问题 ([037c9b7](https://github.com/kercylan98/minotaur/commit/037c9b7bbd43f6b893b528df693855b12942cb4e))

## [0.0.18](https://github.com/kercylan98/minotaur/compare/v0.0.17...v0.0.18) (2023-07-19)


### Features | 新特性

* builtin.Player 可以通过 GetConn 函数获取到网络连接 ([31ad0ee](https://github.com/kercylan98/minotaur/commit/31ad0ee4fbfe8c0fe1d4225c11b250559154d21c))
* storage 添加内置实现的文件存储器，可以通过 storages 包进行使用 ([c447c8a](https://github.com/kercylan98/minotaur/commit/c447c8afb395558a2dd85117b3fac8e093a8cfa7))
* 支持使用 super.RegError 函数为错误注册全局错误码，使用 super.GetErrorCode 根据错误获取全局错误码 ([1dcbd0a](https://github.com/kercylan98/minotaur/commit/1dcbd0a2203c5b0384969836bf4083f3fedce418))
* 支持通过 timer.CalcNextTimeWithRefer 计算下一个整点时间 ([8835e4a](https://github.com/kercylan98/minotaur/commit/8835e4a88bd80bb795a93dfe2494445d8acf0d95))
* 新增 storage 支持数据持久化 ([f59354d](https://github.com/kercylan98/minotaur/commit/f59354db3f244e76faf3590f6865088e5ed6e226))


### Tests | 新增或优化测试用例

* 新增 GlobalDataFileStorage 和 IndexDataFileStorage 的测试用例 ([4378aa0](https://github.com/kercylan98/minotaur/commit/4378aa0eb79f08052121c7ec6f3648aa2248d3dd))

## [0.0.17](https://github.com/kercylan98/minotaur/compare/v0.0.16...v0.0.17) (2023-07-18)


### Bug Fixes | 修复

* 修复主键为空的数据被导出的问题 ([ab0a7cb](https://github.com/kercylan98/minotaur/commit/ab0a7cbbbc786f198ab71aed4515ed8422a78cdf))


### Features | 新特性

* 增加部分字符串转换函数 ([28c6097](https://github.com/kercylan98/minotaur/commit/28c60970447da65180a74c69cc47dfa25cce4cac))
* 通过 golang 模板生成的配置结构代码支持通过 Sync 函数执行安全的配置操作，避免配置被刷新造成的异常 ([8bbd495](https://github.com/kercylan98/minotaur/commit/8bbd49554f6df0b9b6e0eacbe8d0eb9ba9f839bf))

## [0.0.16](https://github.com/kercylan98/minotaur/compare/v0.0.15...v0.0.16) (2023-07-17)


### Bug Fixes | 修复

* 修复 server.Server 部分事件中发生 panic 导致程序退出的问题 ([0215d9f](https://github.com/kercylan98/minotaur/commit/0215d9ff8c6771bc398149fbaca35ae3862aa329))


### Styling | 可读性优化

* 去除部分无用字段，优化整体可读性 ([c1e3c65](https://github.com/kercylan98/minotaur/commit/c1e3c65c1cba9edd91268c99943bce64b904b428))


### Other | 其他更改

* pce.ce 包提供内置的 xlsx 配置表 ([91b2b52](https://github.com/kercylan98/minotaur/commit/91b2b52fc8229959c18d048c0c33d49da8b7b4ae))
* 日志字段调用由 zap.Field 更改为 log.Field ([8e2b4eb](https://github.com/kercylan98/minotaur/commit/8e2b4ebc89ed56a3e1e091a8905641ee3461f1c2))
* 配置导出 Golang 结构体注释优化 ([9349e3c](https://github.com/kercylan98/minotaur/commit/9349e3cdbedfdc7d9a4e3a68294afce8ca63da1d))
* 配置导表优化 ([130869a](https://github.com/kercylan98/minotaur/commit/130869af4eb0cec045730d7bc85cf11c0137a236))


### Features | 新特性

* super 包支持 match 控制函数 ([25ed712](https://github.com/kercylan98/minotaur/commit/25ed712fc9ba1f18fe2c1ce5524e2917160ae295))
* super 包支持使用 super.GoFormat 函数格式化 go 文件 ([3ee638f](https://github.com/kercylan98/minotaur/commit/3ee638f4df459f36a05190b3874502f68e815fd2))
* 修复 server.PushAsyncMessage 无法正确调用回调函数的问题 ([1b9ec9f](https://github.com/kercylan98/minotaur/commit/1b9ec9f2b69b3d2eadbca05447b4d69d1a97a232))
* 重构 config 和 configexport 包 ([7e7a504](https://github.com/kercylan98/minotaur/commit/7e7a504421ba430537f3b70e78334fc30a4a1681))

## [0.0.15](https://github.com/kercylan98/minotaur/compare/v0.0.14...v0.0.15) (2023-07-14)


### Bug Fixes | 修复

* 修复 log 无法正确打印 Caller 的问题 ([349ec42](https://github.com/kercylan98/minotaur/commit/349ec42a7289879d830e592328970bd0ba77e817))


### Other | 其他更改

* mod 优化 ([8f9589d](https://github.com/kercylan98/minotaur/commit/8f9589df4270aa5248ee78a005bccf0ffe38c6f7))
* 移除 tools 包 ([3faca36](https://github.com/kercylan98/minotaur/commit/3faca36d5173c79b2dd5f17b08907a44b127a3d0))


### Features | 新特性

* 新增 steram 包，支持 map 和 slice 的链式操作 ([10fcb54](https://github.com/kercylan98/minotaur/commit/10fcb54322e9afeb9eac61175780126bf753967e))

## [0.0.14](https://github.com/kercylan98/minotaur/compare/v0.0.13...v0.0.14) (2023-07-13)


### Features | 新特性

* slice 包支持获取数组的部分数据 ([c211d62](https://github.com/kercylan98/minotaur/commit/c211d626203b9355fb72e6796faa7aba8728ec0c))
* 支持通过 file.FilePaths 获取目录下所有文件，通过 file.LineCount 统计文件行数 ([0c5ff89](https://github.com/kercylan98/minotaur/commit/0c5ff894f8c731a84b46d2c2b3bee91588d84efd))
* 支持通过 server.NewPacket、 server.NewWSPacket、server.NewPacketString、server.NewWSPacketString 函数快捷创建数据包 ([26993d9](https://github.com/kercylan98/minotaur/commit/26993d94d90a5664c003cd8893a214e747d528df))
* 支持通过 server.SetMessagePacketVisualizer 函数设置服务器数据包消息可视化函数 ([676b542](https://github.com/kercylan98/minotaur/commit/676b5429433cf73bb2bc9fe8b494de4906ade88a))


### Performance Improvements | 性能优化

* 调整 server.DefaultMessageChannelSize 为 65535，优化默认内存占用 ([3e9d56e](https://github.com/kercylan98/minotaur/commit/3e9d56ec5b49fde53e8e8ddf9577f619eed98922))

## [0.0.13](https://github.com/kercylan98/minotaur/compare/v0.0.12...v0.0.13) (2023-07-12)


### Performance Improvements | 性能优化

* 优化代码结构，去除无用代码，去除重复代码 ([47b8a33](https://github.com/kercylan98/minotaur/commit/47b8a333eb22c6bbfe1c6e533681c7dcb5ca34fd))


### Other | 其他更改

* 修改 server.Server 慢消息检测的异步消息判定条件为 1 秒 ([8917326](https://github.com/kercylan98/minotaur/commit/8917326a246d52599ef5de18f939a3e035c245db))


### Code Refactoring | 重构

* log 包重构，优化使用方式 ([98234e5](https://github.com/kercylan98/minotaur/commit/98234e5f861cecfe4fd30f3db51713201d19c725))
* 任务 task 包重构 ([a23e48b](https://github.com/kercylan98/minotaur/commit/a23e48b087252995f8a42212bd77f3d0d8126578))


### Features | 新特性

* str 包增加内置字符 Dunno、CenterDot、Dot、Slash 和其 []byte 形式 ([94147e8](https://github.com/kercylan98/minotaur/commit/94147e8b9c99de298cc6a6d8957f286f9409f54f))
* 可使用 super.NewStackGo 创建用于对上一个协程堆栈进行收集的收集器 ([a4a27ea](https://github.com/kercylan98/minotaur/commit/a4a27ea9da7d1d61ad1a5972077f888f309e8f4d))
* 支持通过 super.StackGO 进行跨协程同步运行堆栈抓取 ([b5a4bc9](https://github.com/kercylan98/minotaur/commit/b5a4bc959df8aee316bedd4050ac34d77a858162))


### Bug Fixes | 修复

* 修复服务器消息报错不打印堆栈信息的问题 ([aa39d39](https://github.com/kercylan98/minotaur/commit/aa39d391606b0a1817b16886616d8803925c90cf))

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
