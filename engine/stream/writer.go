package stream

import "github.com/kercylan98/minotaur/engine/vivid"

// Writer 是一个流式写入器，它是来自 Stream Actor 的写入器引用，接收 Packet 并将其写入到 Stream 中
type Writer = vivid.ActorRef
