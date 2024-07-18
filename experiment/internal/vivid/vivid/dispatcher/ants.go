package dispatcher

import "github.com/panjf2000/ants/v2"

var _ Dispatcher = (*Ants)(nil)

// NewAnts 创建一个使用 ants.Pool 作为分发实现的实例
func NewAnts(size int, options ...ants.Option) (*Ants, error) {
	p, err := ants.NewPool(size, options...)
	if err != nil {
		return nil, err
	}
	return &Ants{pool: p}, nil
}

// Ants 是使用 ants.Pool 作为分发器的实现，当 ants.Pool 无法使用时会退化为 goroutine 处理
type Ants struct {
	pool *ants.Pool
}

func (a *Ants) Dispatch(f func()) {
	if err := a.pool.Submit(f); err != nil {
		go f()
	}
}
