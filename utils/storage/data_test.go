package storage_test

import (
	"encoding/json"
	"fmt"
	"github.com/kercylan98/minotaur/utils/storage"
	"testing"
)

type Player struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Power int64  `json:"power"`
}

func TestData_Struct(t *testing.T) {
	player := storage.NewSet[string, *Player](new(Player),
		func(data *Player) string {
			return data.ID
		}, storage.WithIndex[string, string, *Player]("id", func(data *Player) string {
			return data.ID
		}), storage.WithIndex[string, string, *Player]("name", func(data *Player) string {
			return data.Name
		}),
	)

	p := player.New()

	p.Handle(func(data *Player) {
		data.ID = "1"
		data.Name = "kercylan"
		data.Power = 100
	})

	str, err := player.Struct(p)
	if err != nil {
		panic(err)
	}
	bytes, err := json.Marshal(str)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
