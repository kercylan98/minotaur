package vivid

import "errors"

var (
	ErrInvalidActorId         = errors.New("invalid actor id")
	ErrDispatcherNotFound     = errors.New("dispatcher not found")
	ErrMailboxFactoryNotFound = errors.New("mailbox factory not found")
	ErrActorAlreadyExists     = errors.New("actor already exists")
	ErrActorDeadOrNotExist    = errors.New("actor dead or not exist")
	ErrMessageReplyTimeout    = errors.New("message reply timeout")
	ErrClusterNotEnabled      = errors.New("cluster not enabled")
)
