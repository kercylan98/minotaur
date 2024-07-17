package dispatcher

type Dispatcher interface {
	Dispatch(f func())
}
