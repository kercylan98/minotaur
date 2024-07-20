package behavior

func New[T any]() Behavior[T] {
	return &behavior[T]{}
}

type Behavior[T any] interface {
	Become(performance Performance[T])

	BecomeStacked(performance Performance[T])

	UnBecomeStacked()

	Perform(ctx T)
}

type behavior[T any] struct {
	performances []Performance[T]
}

func (b *behavior[T]) Become(performance Performance[T]) {
	b.performances = b.performances[:0]
	b.performances = append(b.performances, performance)
}

func (b *behavior[T]) BecomeStacked(performance Performance[T]) {
	b.performances = append(b.performances, performance)
}

func (b *behavior[T]) UnBecomeStacked() {
	if len(b.performances) > 0 {
		l := len(b.performances) - 1
		b.performances = b.performances[:l]
	}
}

func (b *behavior[T]) Perform(ctx T) {
	if len(b.performances) > 0 {
		performance := b.performances[len(b.performances)-1]
		performance.Perform(ctx)
	}
}
