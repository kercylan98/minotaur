package vivid

import "sync"

type trieActorNode struct {
	children map[rune]*trieActorNode // 子节点
	isEnd    bool                    // 是否是一个单词的结束
	actor    *actorCore              // Actor
	lock     sync.RWMutex            // 读写锁
}
