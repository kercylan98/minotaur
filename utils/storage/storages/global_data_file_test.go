package storages_test

import (
	"github.com/kercylan98/minotaur/utils/storage"
	"github.com/kercylan98/minotaur/utils/storage/storages"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

type GlobalData struct {
	CreateAt   time.Time
	TotalCount int
}

func TestGlobalDataFileStorage_Save(t *testing.T) {
	Convey("TestGlobalDataFileStorage_Save", t, func() {
		data, err := storage.NewGlobalData[*GlobalData]("global_data_file_test", storages.NewGlobalDataFileStorage[*GlobalData]("./example-data", func(name string) *GlobalData {
			return &GlobalData{
				CreateAt: time.Now(),
			}
		}))
		So(err, ShouldBeNil)
		data.Handle(func(name string, data *GlobalData) {
			data.TotalCount = 10
		})
		So(data.SaveData(), ShouldBeNil)
		So(data.GetData().TotalCount, ShouldEqual, 10)
	})
}

func TestGlobalDataFileStorage_Load(t *testing.T) {
	Convey("TestGlobalDataFileStorage_Load", t, func() {
		data, err := storage.NewGlobalData[*GlobalData]("global_data_file_test", storages.NewGlobalDataFileStorage[*GlobalData]("./example-data", func(name string) *GlobalData {
			return &GlobalData{
				CreateAt: time.Now(),
			}
		}))
		So(err, ShouldBeNil)
		So(data.GetData().TotalCount, ShouldEqual, 10)
	})
}
