package collection

import "github.com/kercylan98/minotaur/utils/generic"

type ComparisonHandler[V any] func(source, target V) bool
type OrderedValueGetter[V any, N generic.Ordered] func(v V) N
