package effect

type BuffId = uint32

func newBuff(id BuffId) *buff {
	return &buff{
		id: id,
	}
}

type buff struct {
	id BuffId
}
