package storages_test

import (
	"github.com/kercylan98/minotaur/utils/storage"
	"github.com/kercylan98/minotaur/utils/storage/storages"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type IndexData[I string] struct {
	ID    I
	Value int
}

func (slf *IndexData[I]) GetIndex() I {
	return slf.ID
}

func TestIndexDataFileStorage_Save(t *testing.T) {
	Convey("TestIndexDataFileStorage_Save", t, func() {
		data, err := storage.NewIndexData[string, *IndexData[string]]("index_data_file_test", storages.NewIndexDataFileStorage[string, *IndexData[string]]("./example-data", func(name string, index string) *IndexData[string] {
			return &IndexData[string]{ID: index}
		}, func(name string) *IndexData[string] {
			return new(IndexData[string])
		}))
		So(err, ShouldBeNil)
		data.Handle("INDEX_001", func(name string, index string, data *IndexData[string]) {
			data.Value = 10
		})
		if err := data.SaveData("INDEX_001"); err != nil {
			t.Fatal(err)
		}
		So(data.GetData("INDEX_001").Value, ShouldEqual, 10)
	})
}

func TestIndexDataFileStorage_Load(t *testing.T) {
	Convey("TestIndexDataFileStorage_Load", t, func() {
		data, err := storage.NewIndexData[string, *IndexData[string]]("index_data_file_test", storages.NewIndexDataFileStorage[string, *IndexData[string]]("./example-data", func(name string, index string) *IndexData[string] {
			return &IndexData[string]{ID: index}
		}, func(name string) *IndexData[string] {
			return new(IndexData[string])
		}))
		So(err, ShouldBeNil)
		So(data.GetData("INDEX_001").Value, ShouldEqual, 10)
	})
}
