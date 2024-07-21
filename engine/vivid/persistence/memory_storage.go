package persistence

import "github.com/puzpuzpuz/xsync/v3"

var (
	memoryStorageRecords = xsync.NewMapOf[Name, *memoryStorageRecord]()
)

func NewMemoryStorage() *MemoryStorage {
	return new(MemoryStorage)
}

type MemoryStorage struct{}

type memoryStorageRecord struct {
	snapshot Snapshot
	events   []Event
}

func (m *MemoryStorage) Save(name Name, snapshot Snapshot, events []Event) error {
	memoryStorageRecords.Store(name, &memoryStorageRecord{
		snapshot: snapshot,
		events:   events,
	})
	return nil
}

func (m *MemoryStorage) Load(name Name) (snapshot Snapshot, events []Event, err error) {
	if record, ok := memoryStorageRecords.Load(name); ok {
		return record.snapshot, record.events, nil
	}
	return nil, nil, ErrorPersistenceNotHasRecord
}

func (m *MemoryStorage) Clear(name Name) error {
	memoryStorageRecords.Delete(name)
	return nil
}
