package storages_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/storage/storages"
	"time"
)

func ExampleNewGlobalDataFileStorage() {
	storage := storages.NewGlobalDataFileStorage[*GlobalData](".", func(name string) *GlobalData {
		return &GlobalData{
			CreateAt: time.Now(),
		}
	})

	fmt.Println(storage != nil)
	// Output:
	// true
}
