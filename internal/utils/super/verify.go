package super

// Verify 对特定表达式进行校验，当表达式不成立时，将执行 handle
func Verify[V any](handle func(V)) *VerifyHandle[V] {
	return &VerifyHandle[V]{handle: handle}
}

// VerifyHandle 校验句柄
type VerifyHandle[V any] struct {
	handle func(V)
	v      V
	hit    bool
}

// PreCase 先决校验用例，当 expression 成立时，将跳过 caseHandle 的执行，直接执行 handle 并返回 false
//   - 常用于对前置参数的空指针校验，例如当 a 为 nil 时，不执行 a.B()，而是直接返回 false
func (slf *VerifyHandle[V]) PreCase(expression func() bool, value V, caseHandle func(verify *VerifyHandle[V]) bool) bool {
	if expression() {
		slf.handle(value)
		return false
	}
	return caseHandle(slf)
}

// Case 校验用例，当 expression 成立时，将忽略后续 Case，并将在 Do 时执行 handle，返回 false
func (slf *VerifyHandle[V]) Case(expression bool, value V) *VerifyHandle[V] {
	if !slf.hit && expression {
		slf.v = value
		slf.hit = true
	}
	return slf
}

// Do 执行校验，当校验失败时，将执行 handle，并返回 false
func (slf *VerifyHandle[V]) Do() bool {
	if slf.hit {
		slf.handle(slf.v)
	}
	return !slf.hit
}
