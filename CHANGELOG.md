# Changelog

## [0.6.1](https://github.com/kercylan98/minotaur/compare/v0.6.1...v0.6.1) (2024-08-17)

### ⚠ BREAKING CHANGES

- **configuration:** configuration 包不再适用，它之前提供的配置表的导出功能目前已转移到 AX CLI 中。

### Features

- **ax:** 添加xlsx转换支持和配置生成 ([818b80c](https://github.com/kercylan98/minotaur/commit/818b80c90883f85b35ebec5515dc5cbfac72fa85))
- **collection:** 添加 SliceSum 和 MapSum 函数 ([b448e0a](https://github.com/kercylan98/minotaur/commit/b448e0a9035a242df81f9c6ba0c2005e53418e72))

### Code Refactoring

- **configuration:** 移除不再适用的 configuration 包 ([e389b55](https://github.com/kercylan98/minotaur/commit/e389b5591f2eb46500d22a04b76900ded12331bb))

## [0.6.0](https://github.com/kercylan98/minotaur/compare/v0.6.0...v0.6.0) (2024-08-14)

### ⚠ BREAKING CHANGES

- **minotaur:** ServerActor and ConnActor now use a more streamlined typing system that may affect consumers of the old API.
- **vivid:** ActorId generation and parsing logic has been modified to support cluster identifiers. This may affect existing systems that rely on the previous format.
- **vivid:** GetActorIdByActorRef function has been removed. Update yourcode to use the new Id() method from the ActorRef interface.
- **chrono:** This modifies the internal scheduler logic, which might affect existing clients relying on the previous behavior.
- **vivid:** Mailbox now requires an additional parameter for Enqueue method to specify if the message should be delivered instantly. This may affect the clients relying on the previous signature of the Enqueue method.

### Features

- **actor-context:** add discardOld parameter to Become method ([0bf8601](https://github.com/kercylan98/minotaur/commit/0bf86015bd26429d2178b8ac3295b9a5e3d333b3))
- **actor:** implement idle timeout and improve actor termination ([e306a53](https://github.com/kercylan98/minotaur/commit/e306a5319604677772f81082abb0aa08d4af9286))
- collection 包新增 ConvertSliceToBatches、ConvertMapKeysToBatches、ConvertMapValuesToBatches 函数，用于将切片或 map 转换为按特定数量分批的批次切片 ([9dba7ff](https://github.com/kercylan98/minotaur/commit/9dba7ffe19f0b5502e06d3cafcd1602736e6648e))
- collection 包新增 Equel 命名前缀的用于比较切片和 map 元素是否相同的函数，新增 Loop 命名前缀的用于遍历切片和 map 元素的函数 ([756f823](https://github.com/kercylan98/minotaur/commit/756f823ca409477891f7368c5cc33bd1a06174af))
- collection 包新增 MergeSlice 函数，用于将多个同类型对象合并为一个切片 ([4799a8c](https://github.com/kercylan98/minotaur/commit/4799a8cb73d2bccb2b8a1fd0b33bc2a9746a3f89))
- compress 新增 tar 和 zip 解压缩函数 ([6bd987f](https://github.com/kercylan98/minotaur/commit/6bd987fce5d78cbd1c208e098a066dedca0b3fe5))
- **examples:** add websocket echo server implementation ([671ed64](https://github.com/kercylan98/minotaur/commit/671ed649550c16a1557053ce3c87b2a776584486))
- geometry 包新增 SimpleCircle 结构体，用于表示仅由圆心及半径组成的圆形，包含投影、距离等常用函数。优化 geometry 中的计算函数，所有计算入参均会转换为 float64 运算，输出时转换回原有的泛型类型 ([6846c9d](https://github.com/kercylan98/minotaur/commit/6846c9dfc70b8eb6b326529908ef18f29e4a2a30))
- huge 包 NewInt 函数支持 string、bool、float 类型 ([c4605cc](https://github.com/kercylan98/minotaur/commit/c4605cc4c30e4eeee29662265dfa852d58a96549))
- huge 包新增 Float 实现 ([af189ab](https://github.com/kercylan98/minotaur/commit/af189ab26b1ebf876e44ba508dc8d8600c2ec96f))
- **login:** implement account module with basic authentication ([929d706](https://github.com/kercylan98/minotaur/commit/929d7066ad47c0fb039e78bbc6882a8725eddf35))
- **minotaur:** add hooks for HTTP server customization ([b6334f6](https://github.com/kercylan98/minotaur/commit/b6334f66dbd51f4a7c8c28ac98d8140691f138db))
- **minotaur:** add hooks for HTTP server launch ([beb79bf](https://github.com/kercylan98/minotaur/commit/beb79bfb4364520022419f3a39881d6fd3fb36ac))
- **minotaur:** add pprof support and improve application functions ([875f881](https://github.com/kercylan98/minotaur/commit/875f8819a7479f005cec0d0a9b42d1cf00836e6b))
- modular 包新增 Block 接口，当模块化服务实现 modular.Service 后可选择的实现 Block 接口，该接口将适用于具有阻塞等待需求的服务，例如网络服务器。 ([3549fcc](https://github.com/kercylan98/minotaur/commit/3549fcca11691299e311928fb79ee15863a276cf))
- modular 包新增 dimension 概念，适用于根据特定宿主进行模块化，例如房间与房间之间的各组件相互隔离 ([1402b85](https://github.com/kercylan98/minotaur/commit/1402b854c617af0217f25ff49099b650959b7d3f))
- reflects 包新增 FuncWrapper 相关函数 ([5a898f5](https://github.com/kercylan98/minotaur/commit/5a898f58dc37027b3ba1b5deac6f573bb8d3c52f))
- server 包增加消息统计相关函数 ([05aeed0](https://github.com/kercylan98/minotaur/commit/05aeed05a1533e87237aefca0890533134eda4d2))
- server 包新增 service 模式的加载函数 server.BindService ([bdf4a23](https://github.com/kercylan98/minotaur/commit/bdf4a237df5e055851f06002e98d8727c5357e2c))
- server 包新增 WithDisableAutomaticReleaseShunt 可选项，可禁止分流渠道自动释放。增加 ReleaseShunt、HasShunt、GetShuntNum 等函数。优化系统分流渠道将不再能够被释放 ([d9ef347](https://github.com/kercylan98/minotaur/commit/d9ef3474a721ef7d98e0bfd9509378d25d18ef69))
- server 包新增 WithWebsocketConnInitializer 函数，支持对 websocket 连接打开后进行初始化设置 ([7ee4b89](https://github.com/kercylan98/minotaur/commit/7ee4b893cdc3e2bea8c2be39e254e4f3b13c5695))
- server 包新增 WithWebsocketUpgrade 函数，支持自定义 websocket.Upgrader ([e960d07](https://github.com/kercylan98/minotaur/commit/e960d07f49adb83359f92543053b2efe1e35d182))
- server 支持通过 WithLowMessageDuration、WithAsyncLowMessageDuration 函数设置慢消息阈值 ([4e1d075](https://github.com/kercylan98/minotaur/commit/4e1d075a059363ab39853e6b92f20668c11d0b74))
- server.MultipleServer 支持绑定 Service ([6b24b7c](https://github.com/kercylan98/minotaur/commit/6b24b7c5760705138a8818870de8afa692895fdd))
- server.Server 在执行 Shutdown 时将会等待所有消息分发器被释放 ([4f2850b](https://github.com/kercylan98/minotaur/commit/4f2850b355f788199d6270e9c4862188340ff797))
- server.Server.LoadData 函数支持加载 any 类型的数据 ([ebe7a70](https://github.com/kercylan98/minotaur/commit/ebe7a7049692e8aa4cf2e8cae9b1e5bfdd2836e4))
- sher 包增加部分转换和去重相关函数 ([2ff360c](https://github.com/kercylan98/minotaur/commit/2ff360c48c17573573b8ae41856fadf7e8784c3f))
- sher 包新增 FindInSlice 和 FindInSliceByBinary 函数 ([96953d7](https://github.com/kercylan98/minotaur/commit/96953d74e224c92916df4c6743122371895a626e))
- sher 包新增 map 相关映射操作 ([7086281](https://github.com/kercylan98/minotaur/commit/708628139985109a4fa192319c3eed6c33270623))
- sher 包新增将任一切片转换为 []any 的函数 ([bb06cbf](https://github.com/kercylan98/minotaur/commit/bb06cbfeb0418fa462743c3e6b2e7833ade2cbff))
- space.RoomController 支持设置房主 ([a269845](https://github.com/kercylan98/minotaur/commit/a269845dbbdc3827bf8626d038f4a9d9b87f786e))
- super 包新增 RecoverTransform 函数，用于将 recover() 结果转化为 error ([7efe88a](https://github.com/kercylan98/minotaur/commit/7efe88a0f4a8918541087d95a5d94a49a974727f))
- super 包新增 StopWatch 和 StopWatchAndPrintln 函数，用于追踪函数运行时间 ([7fa0e68](https://github.com/kercylan98/minotaur/commit/7fa0e6863613bbd137be98fa5f4d57345622e0c2))
- super 包新增 TryReadChannel、TryReadChannelByHandler 函数用于对 channel 尝试写入 ([959abff](https://github.com/kercylan98/minotaur/commit/959abff85f4cbf0c4c81b586317e39bf3dee3a80))
- super 包新增 TryWriteChannel 函数，支持尝试性的对 channel 进行写入 ([5b53e8a](https://github.com/kercylan98/minotaur/commit/5b53e8a2ac697bd3503740c8e564c8b485b6c664))
- super 包新增 TryWriteChannelByHandler 函数，支持尝试写入 channel，如果 channel 无法写入则执行 handler ([efbde3e](https://github.com/kercylan98/minotaur/commit/efbde3e3f80973800854b3e9c1f3bb27f8004b38))
- super 包新增 WaitGroup 结构，用法同 sync.WaitGroup，包含一个额外的 Exec 函数，用于便捷的执行异步函数。移除 stack.go 相关的无用代码 ([c98d15b](https://github.com/kercylan98/minotaur/commit/c98d15b0f242cff42f8114f03105b147a5a563c6))
- timer 包新增 GetCurrWeekDate 和 GetLastWeekDate 函数 ([ad4777a](https://github.com/kercylan98/minotaur/commit/ad4777a379753750fac9d98b7b4f0b80ca688c39))
- times 包新增 GetWeekdayDateRelativeToNowWithOffset 及 GetWeekdayTimeRelativeToNowWithOffset 函数，用于取代 GetCurrWeekDate 和 GetLastWeekDate 函数 ([92d6c56](https://github.com/kercylan98/minotaur/commit/92d6c5680d1a97540b5c00fe7643fa657e7c20f7))
- **transport:** implement kcp core functionality ([b47311b](https://github.com/kercylan98/minotaur/commit/b47311b7c9a79471aafa152df946d7e37d39bf0b))
- **vivid:** add cluster support and refactor actor system ([7d6abf8](https://github.com/kercylan98/minotaur/commit/7d6abf8ccfd8ced1713832bd912706a92329b643))
- **vivid:** add stop on parent restart feature for actors ([bfed382](https://github.com/kercylan98/minotaur/commit/bfed3827df1aa06508f66024561cc4fbf548dbfb))
- **vivid:** add typed actor support and improve ask pattern ([4d12837](https://github.com/kercylan98/minotaur/commit/4d12837fbb236b7c3f47d4ef5027267c03de62b4))
- 优化 log 包，支持动态修改日志级别 ([3e41068](https://github.com/kercylan98/minotaur/commit/3e4106861967f0b8f7a57f5365c135fdd323f63e))
- 优化项目文档 ([7001e3d](https://github.com/kercylan98/minotaur/commit/7001e3dbab84555f6d3ff35e6a1b95bd51efd801))
- 修复 HTTP 服务器慢消息空指针问题 ([31c68e4](https://github.com/kercylan98/minotaur/commit/31c68e42b758492f008822674e37434b3b5a8ecb))
- 修复 HTTP 服务器满消息空指针问题 ([68bc005](https://github.com/kercylan98/minotaur/commit/68bc005fe1fb997a0d50d94d8b1a653072473512))
- 完善 stream 包对于 []string 的操作 ([a2695f4](https://github.com/kercylan98/minotaur/commit/a2695f4fcf2a266d3fc535d67d05d07259168d2f))
- 支持向 server.Server 绑定一些数据 ([acc4684](https://github.com/kercylan98/minotaur/commit/acc468492fc76faa69d493c83098bfecbb1e720d))
- 新增 charproc 包处理字符、文本操作 ([f0f5f8a](https://github.com/kercylan98/minotaur/commit/f0f5f8a39636da358e866137d9308c6be0c80420))
- 新增 chrono 包，用于替代原本的 timer 及 times 包 ([e608e92](https://github.com/kercylan98/minotaur/commit/e608e9257ef2f3031319f586fcb2738c65214fb1))
- 新增 mask 包，增加 DynamicMask 高性能可变长度掩码实现 ([0878d1a](https://github.com/kercylan98/minotaur/commit/0878d1acbb36aab39e4a347e6f93d5c8fab2a48b))
- 新增 modular 包，用于实现模块化项目 ([c95b206](https://github.com/kercylan98/minotaur/commit/c95b206592b3be31c1fe5ece891434d03e968c73))
- 新增 utils/sher 包，包含了对 slice 及 hash 常用的操作函数。用于未来对 utils/slice 和 utils/hash 包进行替换 ([515cbc6](https://github.com/kercylan98/minotaur/commit/515cbc66ebe609aa3757ea0c89c8ba11a465e74c))
- 新版 server 包 HTTP 基础实现 ([b2c0bb0](https://github.com/kercylan98/minotaur/commit/b2c0bb0da3dd87520fa5fcf574d88c47f5a26a4a))
- 新版 server 包 HTTP 基础实现 ([37f35aa](https://github.com/kercylan98/minotaur/commit/37f35aa602e7172a5719ec35f17e99744be9c483))
- 新版 server 包 websocket 基础实现、actor 模型实现 ([92c4280](https://github.com/kercylan98/minotaur/commit/92c42800f13391940b8fc7c36eb0fb3b99f066ae))
- 新版 server 包 websocket 基础实现、actor 模型实现 ([ef1bb32](https://github.com/kercylan98/minotaur/commit/ef1bb321d7b38b3353ed9095c87cff9228f2dbfc))

### Bug Fixes

- lockstep 定时器导致空指针问题处理 ([ceffa2e](https://github.com/kercylan98/minotaur/commit/ceffa2e46fcf7aa246345af5a12e0c95fbaa50ab))
- **readme:** 更新 go get 命令以引用最新 release 版本 ([3012317](https://github.com/kercylan98/minotaur/commit/301231716e57d65b62c4086b3d62323aa0b04d74))
- 修复 dispatcher.Dispatcher 在消息归零的时候使用协程运行处理函数可能导致不可知问题的情况，修复消息消费时获取生产者可能已经被释放的问题。修复在无消息时候设置消息完成处理函数不会触发一次的问题 ([7528dc4](https://github.com/kercylan98/minotaur/commit/7528dc4a1b616e35d96dedf0e6afc5330af897c0))
- 修复 geometry 包 SimpleCircle.Projection 函数不正确的问题。优化部分注释及添加部分函数 ([f7c3701](https://github.com/kercylan98/minotaur/commit/f7c37016cef946c9e8e3a4366139c43fee5f8eb8))
- 修复 room_options.go 中空指针检查方式错误的问题 ([556d1cd](https://github.com/kercylan98/minotaur/commit/556d1cdc020c1ab1e7cd2e7bee76366311415e9a))
- 修复 server 中分流定时消息无法命中分流渠道的问题 ([de43f53](https://github.com/kercylan98/minotaur/commit/de43f531317c664a31d85dc14264b42652aadde5))
- 修复 server 使用 WebSocket 模式下，路由不支持 :1234/ws 的格式问题 ([f8e368a](https://github.com/kercylan98/minotaur/commit/f8e368a8caa93be1bebfadef6f65649b31eb3640))
- 修复 server 包 None 网络类型启动阻塞的问题。增加传入不支持网络类型将导致 panic 的特性。优化 WebSocket 服务器将不再使用 http.DefaultMuxServer，转而使用 http.NewServeMux ([1645ae4](https://github.com/kercylan98/minotaur/commit/1645ae47df879067ba286affec39e5bed168fa02))
- 修复 server 包 WebSocket 路由兼容性问题 ([590d0a1](https://github.com/kercylan98/minotaur/commit/590d0a1887412d62831a2cbbe5e699ac3fd19a6c))
- 修复 server 包异步分流消息的回调函数在取消分流渠道绑定后会在系统分流渠道执行的问题 ([e760ef2](https://github.com/kercylan98/minotaur/commit/e760ef2a0f6b76cbfe9129ddce19d9004d150866))
- 修复 server 包死锁检测中 Message 读写的竞态问题 ([b81f972](https://github.com/kercylan98/minotaur/commit/b81f972fdadb5fb1e5d13667f558ef4c58788036))
- 修复 server 包连接断开消息分发器阻塞的问题，优化等待消息时打印的日志频率 ([af23744](https://github.com/kercylan98/minotaur/commit/af237448d7b7019adcb5bfc8d6efa135f597c372))
- 修复 server 包部分问题，修复 log 包在 init 函数调用可能产生的空指针问题 ([3402c83](https://github.com/kercylan98/minotaur/commit/3402c83fd44e6c71db4401ef49667cada293c9c9))
- 修复 server.hub 广播时未解锁的问题，优化处理逻辑 ([80f38ff](https://github.com/kercylan98/minotaur/commit/80f38ffe9c5a603e432ee00692c9b9bc35ac65c7))
- 修复 server.LoadData 函数签名的错误 ([e585e12](https://github.com/kercylan98/minotaur/commit/e585e12a7243a37cbb335a72500f425d0de188bf))
- 修复 server.Service 初始化的 for 循环指针问题 ([b633f1a](https://github.com/kercylan98/minotaur/commit/b633f1af9fba3c075a25530543576e50520307ce))
- 修复 server.WithDispatcherBufferSize 过小的情况下，在消息中发布新消息导致永久阻塞的问题 ([b39625c](https://github.com/kercylan98/minotaur/commit/b39625c0cb5e1e8c04d10c263b0a9d75c6de6d40))
- 修复 space.RoomManager.AssumeControl 函数编译错误的问题 ([3f099e6](https://github.com/kercylan98/minotaur/commit/3f099e6f8e0c3a4540d8d88856745d71b1cb28b8))
- 修复 super 包 JSON 解析部分零值不正确的问题 ([36de593](https://github.com/kercylan98/minotaur/commit/36de5934ce1591fb6347d8b34f6550e2fe4811fb))
- 修复 timer.Ticker 并发问题 ([d1d5bd4](https://github.com/kercylan98/minotaur/commit/d1d5bd40d488bc1c7b2c4c1e0cbaa3b7190c87f4))
- 修复 timer.Ticker 死锁 ([612c41f](https://github.com/kercylan98/minotaur/commit/612c41ffd8858ff5f8f792316fc4a51a795df2a7))
- 修复 times.GetCurrWeekDate 和 times.GetLastWeekDate 在 week 参数与今日 week 相同的情况下，会多获取到一周的问题 ([902dada](https://github.com/kercylan98/minotaur/commit/902dadad5abffe215ababb5ce6d310141708f6fc))
- 修复循环依赖问题 ([6d8258b](https://github.com/kercylan98/minotaur/commit/6d8258b153fb7d3354d910d0d31e5cdf790364a0))
- 修复配置导出 go 代码文件时，引用包错误的问题 ([790e317](https://github.com/kercylan98/minotaur/commit/790e31764f10ee06916e222d896ac9050bc32faa))
- 修改 ShowServersInfo 函数可见性，修复服务器地址打印的指针问题 ([26aa2d9](https://github.com/kercylan98/minotaur/commit/26aa2d9ff8c064105f4e8f44bad5da9046d76718))
- 移除 modular 包的自动注入，优化 modular.Service 接口说明 ([d531939](https://github.com/kercylan98/minotaur/commit/d531939903e14fdc191263433214515fc12d8a93))
- 移除 modular.go 中不必要的代码，修复 timer.Ticker 释放后 handler 指针污染的问题 ([17cdad2](https://github.com/kercylan98/minotaur/commit/17cdad2c6e00b2ddb405ff8ec23344be60757640))

### Performance Improvements

- client 包由无界缓冲区调整为基于 chan 实现的缓冲区，新增 RunByBufferSize 函数支持以指定缓冲区大小运行 ([bdbcc1b](https://github.com/kercylan98/minotaur/commit/bdbcc1bb358deacbc2250a46b22ad758c7067f9b))
- 优化 server.Server 连接管理机制，优化 GetOnlineCount、GetOnlineBotCount 性能 ([5e5fe8a](https://github.com/kercylan98/minotaur/commit/5e5fe8acca8b2ef7a302997b0211a3103415bdf9))
- 去除 buffer.Unbounded 不必要的构造函数和 nil 字段 ([7111350](https://github.com/kercylan98/minotaur/commit/711135002217c9ccb16f944dfbf26a9b5f934c0d))
- 更改 server 和 conn 的消息实现为 channel ([d27fa7c](https://github.com/kercylan98/minotaur/commit/d27fa7c246319d2d4119c892b1808643db20a5e7))

### Miscellaneous Chores

- release 0.6.0 ([f761390](https://github.com/kercylan98/minotaur/commit/f761390c0c6a6f638980732f20af6e03b63c1a11))

### Code Refactoring

- **chrono:** update scheduler and task management logic ([4732b99](https://github.com/kercylan98/minotaur/commit/4732b9972719bd2ce9b62715eafa63c199e0d1d8))
- **minotaur:** simplify actor typing and improve network handling ([d542b36](https://github.com/kercylan98/minotaur/commit/d542b3669f3cfb4b113fa027925fbc3efd1f398f))
- **vivid:** optimize actor reference handling and mod status management ([1220b60](https://github.com/kercylan98/minotaur/commit/1220b601bc1675f09db5d0f7fa2de15f0426d4a2))
- **vivid:** optimize message dispatching for instant delivery ([cf23e79](https://github.com/kercylan98/minotaur/commit/cf23e7926adabd23bb6a0158259a717addd4cb5e))
