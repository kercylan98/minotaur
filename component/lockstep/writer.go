package lockstep

type Writer[ID comparable, FrameCommand any] interface {
	GetID() ID
	Healthy() bool
	Marshal(frames map[uint32]Frame[FrameCommand]) []byte
	Write(data []byte) error
}
