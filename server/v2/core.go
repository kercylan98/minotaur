package server

type Core interface {
	connectionManager
}

type connectionManager interface {
	Event() chan<- any
}
