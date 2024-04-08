package queue

func newWrapper[Id, Q comparable, M Message[Q]](queue *Queue[Id, Q, M], message M) wrapper[Id, Q, M] {
	return wrapper[Id, Q, M]{
		message:    message,
		controller: newController[Id, Q, M](queue, message),
	}
}

type wrapper[Id, Q comparable, M Message[Q]] struct {
	message    M
	controller Controller
}
