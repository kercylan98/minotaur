package datasheet

type Type interface {
	typeof()
}

func (t *BasicType) typeof() {}
func (t *Array) typeof()     {}
func (t *Map) typeof()       {}
func (t *Slice) typeof()     {}
func (t *Struct) typeof()    {}
