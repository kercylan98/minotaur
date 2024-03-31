package reactor

type queueMessage[M any] struct {
	ident *identifiable
	msg   M
}
