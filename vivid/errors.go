package vivid

import "errors"

var (
	ErrActorIdInvalid            = errors.New("actor id invalid")
	ErrActorBehaviorInvalid      = errors.New("actor behavior invalid")
	ErrActorNotImplementActorRef = errors.New("actor not implement ActorRef")
	ErrActorAlreadyExists        = errors.New("actor already exists")
	ErrActorNotFound             = errors.New("actor not found")
	ErrReplyTimeout              = errors.New("actor reply timeout")
)
