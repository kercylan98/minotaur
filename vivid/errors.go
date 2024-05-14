package vivid

import "errors"

var (
	ErrActorIdInvalid            = errors.New("localActor id invalid")
	ErrActorBehaviorInvalid      = errors.New("localActor behavior invalid")
	ErrActorNotImplementActorRef = errors.New("localActor not implement ActorRef")
	ErrActorOnPreStartFailed     = errors.New("localActor OnPreStart failed")
	ErrActorAlreadyExists        = errors.New("localActor already exists")
	ErrActorNotFound             = errors.New("localActor not found")
)
