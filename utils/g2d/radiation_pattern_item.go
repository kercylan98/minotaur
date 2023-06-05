package g2d

type RadiationPatternItem[Type comparable] interface {
	GetGuid() int64
	GetType() Type
}
