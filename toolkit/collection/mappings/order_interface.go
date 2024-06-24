package mappings

import "github.com/kercylan98/minotaur/toolkit/constraints"

type OrderInterface[K constraints.Hash, V any] interface {
	Get(key K) (value V, exists bool)

	Add(key K, value V)

	Set(key K, value V)

	Len() int

	Del(key K)

	Range(handle func(key K, value V) bool)
}
