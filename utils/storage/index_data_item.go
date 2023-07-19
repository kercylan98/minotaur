package storage

import "github.com/kercylan98/minotaur/utils/generic"

type IndexDataItem[I generic.Ordered] interface {
	GetIndex() I
}
