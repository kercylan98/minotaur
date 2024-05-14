package vivid

type actorTrie struct {
	root *trieActorNode // 根节点
}

func (t *actorTrie) init() *actorTrie {
	t.root = &trieActorNode{
		children: make(map[rune]*trieActorNode),
	}
	return t
}

// insert 插入一个 localActor
func (t *actorTrie) insert(name string, actor *localActor) {
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
}

// find 查找一个 localActor
func (t *actorTrie) find(name string) *localActor {
	node := t.root
	for _, ch := range name {
		node.lock.RLock()
		nextNode, exists := node.children[ch]
		node.lock.RUnlock()
		if !exists {
			return nil
		}
		node = nextNode
	}
	node.lock.RLock()
	defer node.lock.RUnlock()
	if node.isEnd {
		return node.actor
	}
	return nil
}

// has 检查是否存在一个 localActor
func (t *actorTrie) has(name string) bool {
	return t.find(name) != nil
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
