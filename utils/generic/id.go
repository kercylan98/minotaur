package generic

type IdR[ID comparable] interface {
	GetId() ID
}

type IdW[ID comparable] interface {
	SetId(id ID)
}

type IdRW[ID comparable] interface {
	IdR[ID]
	IdW[ID]
}
