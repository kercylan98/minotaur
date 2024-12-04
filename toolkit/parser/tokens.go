package parser

type Tokens[T any] struct {
	tokenized   *Tokenized[T]
	handler     Handler[T]
	pos         int
	undo        int
	symbolTable map[Symbol]bool
}

// Consume 消费一个 Token
func (t *Tokens[T]) Consume() Token {
	if t.pos >= len(t.tokenized.tokens) {
		return ""
	}
	token := t.tokenized.tokens[t.pos]
	t.pos++
	t.undo++
	return token
}

// Undo 回退一个 Token，当没有发生消费时，什么也不会发生
func (t *Tokens[T]) Undo() {
	if t.undo > 0 {
		t.undo--
		t.pos--
	}
}

// Peek 预览下一个 Token
func (t *Tokens[T]) Peek() Token {
	if t.pos >= len(t.tokenized.tokens) {
		return ""
	}
	return t.tokenized.tokens[t.pos]
}

// PeekOffset 预览指定偏移的 Token
func (t *Tokens[T]) PeekOffset(offset int) Token {
	offset = t.pos + offset
	if offset >= len(t.tokenized.tokens) || offset < 0 {
		return ""
	}
	return t.tokenized.tokens[offset]
}

// Reset 重置所有消费
func (t *Tokens[T]) Reset() {
	t.pos = 0
}

// Handle 以当前状态递归处理
func (t *Tokens[T]) Handle() (T, error) {
	return t.handler.Handle(t)
}

// IsSymbol 判断是否为符号
func (t *Tokens[T]) IsSymbol(token Token) bool {
	runeToken := []rune(token)
	if len(runeToken) != 1 {
		return false
	}
	return t.symbolTable[Symbol(runeToken[0])]
}
