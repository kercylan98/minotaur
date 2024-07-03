package cluster

import (
	"bytes"
	"encoding/gob"
	"github.com/kercylan98/minotaur/core"
	"sync"
)

func init() {
	gob.Register(new(metadata))
}

func parseMetadata(data []byte) (*metadata, error) {
	var m = new(metadata)
	b := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(b)
	err := decoder.Decode(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

type metadata struct {
	mu *sync.RWMutex

	ActorSystemAddr        string
	ActorSystemPort        uint16
	ActorSystemRootAddress core.Address
}

func (m *metadata) bytes() ([]byte, error) {
	m.mu.Lock()
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(m)
	m.mu.Unlock()
	return buf.Bytes(), err
}
