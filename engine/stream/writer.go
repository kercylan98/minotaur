package stream

import "github.com/kercylan98/minotaur/engine/vivid"

// Writer 是一个流式写入器，它是来自 Stream Actor 的写入器引用，接收 Packet 并将其写入到 Stream 中
//
// Deprecated: 该设计加大了理解成本，且不易于使用，将考虑新的方案用于处理网络连接。至 v0.7.0 版本及以后，stream 包将被移除。
type Writer = vivid.ActorRef
