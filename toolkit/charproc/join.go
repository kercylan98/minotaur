package charproc

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"strings"
)

func NumberJoin[T constraints.Number](elems []T, sep string) string {
	return strings.Join(convert.NumbersToStrings(elems...), sep)
}
