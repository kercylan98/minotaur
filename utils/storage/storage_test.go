package storage_test

import (
	"github.com/kercylan98/minotaur/utils/storage"
	"testing"
)

var fakeDB = map[any][]byte{}

type Player struct {
	ID   string
	Name string
}

type PlayerWarehouse[K string, D *Player] struct {
}

func (slf *PlayerWarehouse[K, D]) GenerateZero() D {
	return &Player{}
}

func (slf *PlayerWarehouse[K, D]) Init() (map[K][]byte, error) {
	return nil, nil
}

func (slf *PlayerWarehouse[K, D]) Query(key K) (data []byte, err error) {
	return fakeDB[key], nil
}

func (slf *PlayerWarehouse[K, D]) Create(key K, data []byte) error {
	fakeDB[key] = data
	return nil
}

func (slf *PlayerWarehouse[K, D]) Save(key K, data []byte) error {
	fakeDB[key] = data
	return nil
}

func TestStorage(t *testing.T) {
	s := storage.New[string, *Player, *PlayerWarehouse[string, *Player]](new(PlayerWarehouse[string, *Player]))
	_ = s.Create("1", &Player{
		ID:   "1",
		Name: "1",
	})
	player, _ := s.Query("1")
	t.Log(player)
}
