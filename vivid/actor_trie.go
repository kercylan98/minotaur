package vivid

import "sync"

type actorTrie struct {
	root   *trieActorNode         // 根节点
	fast   map[ActorId]*actorCore // 快速查找
	fastRW sync.RWMutex           // 快速查找读写锁
}

func (t *actorTrie) init() *actorTrie {
	t.fast = make(map[ActorId]*actorCore)
	t.root = &trieActorNode{
		children: make(map[rune]*trieActorNode),
	}
	return t
}

// insert 插入一个 localActorRef
func (t *actorTrie) insert(actorId ActorId, actor *actorCore) {
	name := actorId.Name()
	node := t.root
	for _, ch := range name {
		node.lock.Lock()
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = &trieActorNode{
				children: make(map[rune]*trieActorNode),
			}
		}
		nextNode := node.children[ch]
		node.lock.Unlock()
		node = nextNode
	}
	node.lock.Lock()
	node.isEnd = true
	node.actor = actor
	node.lock.Unlock()

	t.fastRW.Lock()
	t.fast[actorId] = actor
	t.fastRW.Unlock()
}

// find 查找一个 localActor
func (t *actorTrie) find(id ActorId) *actorCore {
	t.fastRW.RLock()
	actor, _ := t.fast[id]
	t.fastRW.RUnlock()
	return actor
}

// has 检查是否存在一个 localActorRef
func (t *actorTrie) has(id ActorId) bool {
	t.fastRW.RLock()
	_, exists := t.fast[id]
	t.fastRW.RUnlock()
	return exists
}

// startsWith 搜索一个前缀
func (t *actorTrie) startsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		node.lock.RLock()
		nextNode, exists := node.children[ch]
		node.lock.RUnlock()
		if !exists {
			return false
		}
		node = nextNode
	}
	return true
}

// remove 删除一个单词
func (t *actorTrie) remove(word string) {
	var remove func(node *trieActorNode, i int)
	remove = func(node *trieActorNode, i int) {
		if i == len(word) {
			node.lock.Lock()
			node.isEnd = false
			node.lock.Unlock()
			return
		}
		ch := rune(word[i])
		node.lock.RLock()
		nextNode, exists := node.children[ch]
		node.lock.RUnlock()
		if !exists {
			return
		}
		remove(nextNode, i+1)
		node.lock.Lock()
		if len(node.children[ch].children) == 0 && !node.children[ch].isEnd {
			delete(node.children, ch)
		}
		node.lock.Unlock()
	}
	remove(t.root, 0)
}
