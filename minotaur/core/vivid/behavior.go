package vivid

var _ Behavior = (*behavior)(nil)

type Behavior interface {
	Become(receive OnReceiveFunc)

	BecomeStacked(receive OnReceiveFunc)

	UnBecomeStacked()

	Receive(context ActorContext)
}

type behavior struct {
	data []OnReceiveFunc
}

func newBehavior() *behavior {
	return &behavior{
		data: make([]OnReceiveFunc, 0, 4), // 预分配一定容量
	}
}

func (b *behavior) Become(receive OnReceiveFunc) {
	b.clear()
	b.push(receive)
}

func (b *behavior) BecomeStacked(receive OnReceiveFunc) {
	b.push(receive)
}

func (b *behavior) UnBecomeStacked() {
	b.pop()
}

func (b *behavior) Receive(context ActorContext) {
	if f, ok := b.peek(); ok {
		f(context)
	}
}

func (b *behavior) clear() {
	b.data = b.data[:0]
}

func (b *behavior) peek() (OnReceiveFunc, bool) {
	if len(b.data) > 0 {
		return b.data[len(b.data)-1], true
	}
	return nil, false
}

func (b *behavior) push(v OnReceiveFunc) {
	b.data = append(b.data, v)
}

func (b *behavior) pop() (OnReceiveFunc, bool) {
	if len(b.data) > 0 {
		l := len(b.data) - 1
		v := b.data[l]
		b.data = b.data[:l]
		return v, true
	}
	return nil, false
}
