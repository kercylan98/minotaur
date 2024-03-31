package actor

import "github.com/kercylan98/minotaur/server/internal/v2/dispatcher"

type Actor[M any] struct {
	*dispatcher.Dispatcher[M]
}
