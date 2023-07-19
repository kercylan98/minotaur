package storages_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/storage/storages"
)

func ExampleNewIndexDataFileStorage() {
	storage := storages.NewIndexDataFileStorage[string, *IndexData[string]]("./example-data", func(name string, index string) *IndexData[string] {
		return &IndexData[string]{ID: index}
	}, func(name string) *IndexData[string] {
		return new(IndexData[string])
	})

	fmt.Println(storage != nil)
	// Output:
	// true
}
