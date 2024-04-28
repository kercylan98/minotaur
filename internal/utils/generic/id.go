package generic

type IdR[ID comparable] interface {
	GetId() ID
}

type IDR[ID comparable] interface {
	GetID() ID
}

type IdW[ID comparable] interface {
	SetId(id ID)
}

type IDW[ID comparable] interface {
	SetID(id ID)
}

type IdR2W[ID comparable] interface {
	IdR[ID]
	IdW[ID]
}

type IDR2W[ID comparable] interface {
	IDR[ID]
	IDW[ID]
}
