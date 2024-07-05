package transport

type FiberService interface {
	OnInit(kit *FiberKit)
}
