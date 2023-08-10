# Changelog

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
