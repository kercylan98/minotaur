package cluster

import (
	"bytes"
	"encoding/gob"
	"github.com/kercylan98/minotaur/core/vivid"
	"sync"
)

type local struct {
	mu *sync.RWMutex

	Kinds map[string]map[vivid.Kind]struct{} // NodeName => Kinds
}

func (l *local) bytes(notLock ...bool) ([]byte, error) {
	if len(notLock) > 0 && notLock[0] {
		l.mu.Lock()
	}
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(l)
	if len(notLock) > 0 && notLock[0] {
		l.mu.Unlock()
	}
	return buf.Bytes(), err
}

func (l *local) merge(other *local) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Merge Kinds
	if l.Kinds == nil {
		l.Kinds = make(map[string]map[vivid.Kind]struct{})
	}
	for nodeName := range l.Kinds {
		if _, ok := other.Kinds[nodeName]; !ok {
			delete(l.Kinds, nodeName)
		}
	}
	for nodeName, kinds := range other.Kinds {
		l.Kinds[nodeName] = kinds
	}
}
