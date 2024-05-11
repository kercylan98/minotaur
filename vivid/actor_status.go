package vivid

type actorStatus = int32

const (
	actorStatusNone actorStatus = iota
	actorStatusRunning
	actorStatusStopping
	actorStatusStopped
)
