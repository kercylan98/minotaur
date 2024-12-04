package parser

type Tokens struct {
	tokenized *Tokenized
	handler   Handler
	pos       int
	undo      int
}

// Consume 消费一个 Token
func (t *Tokens) Consume() Token {
	if t.pos >= len(t.tokenized.tokens) {
		return ""
	}
	token := t.tokenized.tokens[t.pos]
	t.pos++
	t.undo++
	return token
}

// Undo 回退一个 Token，当没有发生消费时，什么也不会发生
func (t *Tokens) Undo() {
	if t.undo > 0 {
		t.undo--
		t.pos--
	}
}

// Peek 预览下一个 Token
func (t *Tokens) Peek() Token {
	if t.pos >= len(t.tokenized.tokens) {
		return ""
	}
	return t.tokenized.tokens[t.pos]
}

// Reset 重置所有消费
func (t *Tokens) Reset() {
	t.pos = 0
}

// Handle 以当前状态递归处理
func (t *Tokens) Handle() error {
	return t.handler.Handle(t)
}
