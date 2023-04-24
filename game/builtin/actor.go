package builtin

func NewActor(guid int64) *Actor {
	return &Actor{
		guid: guid,
	}
}

type Actor struct {
	guid int64
}

func (slf *Actor) GetGuid() int64 {
	return slf.guid
}
