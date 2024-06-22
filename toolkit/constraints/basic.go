package constraints

import (
	"golang.org/x/exp/constraints"
	"unsafe"
)

type Basic interface {
	Ordered
	bool | []byte | rune | byte
	~bool | ~[]byte | ~rune | ~byte
}

type Hash interface {
	constraints.Integer | constraints.Float | constraints.Complex | ~string | uintptr | ~unsafe.Pointer
}
