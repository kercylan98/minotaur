package vivid

type Actor interface {
	OnReceive(ctx MessageContext)
}
