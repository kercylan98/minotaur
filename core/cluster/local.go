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

func (l *local) merge(other *local) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Merge Kinds
	if l.Kinds == nil {
		l.Kinds = make(map[string]map[vivid.Kind]struct{})
	}

	for nodeName, kinds := range other.Kinds {
		localOtherKinds, exists := l.Kinds[nodeName]
		if !exists {
			localOtherKinds = make(map[vivid.Kind]struct{})
			l.Kinds[nodeName] = localOtherKinds
		}
		for kind := range kinds {
			localOtherKinds[kind] = struct{}{}
		}
	}
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
