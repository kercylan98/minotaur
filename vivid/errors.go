package vivid

import "errors"

var (
	ErrActorIdInvalid            = errors.New("actor id invalid")
	ErrActorBehaviorInvalid      = errors.New("actor behavior invalid")
	ErrActorNotImplementActorRef = errors.New("actor not implement ActorRef")
	ErrActorAlreadyExists        = errors.New("actor already exists")
	ErrActorNotFound             = errors.New("actor not found")
	ErrReplyTimeout              = errors.New("actor reply timeout")
	ErrActorNotUnique            = errors.New("actor not unique")
	ErrActorTerminated           = errors.New("actor terminated or not exists")
	ErrActorPreStart             = errors.New("actor pre start error")
)
