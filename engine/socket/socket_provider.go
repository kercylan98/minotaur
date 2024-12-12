package socket

type Provider interface {
	Provide() Actor
}

type FunctionalProvider func() Actor

func (f FunctionalProvider) Provide() Actor {
	return f()
}
