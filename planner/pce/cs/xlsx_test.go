package cs_test

import (
	"github.com/kercylan98/minotaur/planner/pce/cs"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tealeg/xlsx"
	"testing"
)

func TestNewIndexXlsxConfig(t *testing.T) {
	Convey("TestNewIndexXlsxConfig", t, func() {
		f, err := xlsx.OpenFile("./xlsx_template.xlsx")
		if err != nil {
			panic(err)
		}
		config := cs.NewXlsx(f.Sheets[1])
		So(config, ShouldNotBeNil)
	})
}

func TestXlsxIndexConfig_GetConfigName(t *testing.T) {
	Convey("TestXlsxIndexConfig_GetConfigName", t, func() {
		f, err := xlsx.OpenFile("./xlsx_template.xlsx")
		if err != nil {
			panic(err)
		}
		config := cs.NewXlsx(f.Sheets[1])
		So(config.GetConfigName(), ShouldEqual, "IndexConfig")
	})
}

func TestXlsxIndexConfig_GetDisplayName(t *testing.T) {
	Convey("TestXlsxIndexConfig_GetDisplayName", t, func() {
		f, err := xlsx.OpenFile("./xlsx_template.xlsx")
		if err != nil {
			panic(err)
		}
		config := cs.NewXlsx(f.Sheets[1])
		So(config.GetDisplayName(), ShouldEqual, "有索引")
	})
}

func TestXlsxIndexConfig_GetDescription(t *testing.T) {
	Convey("TestXlsxIndexConfig_GetDescription", t, func() {
		f, err := xlsx.OpenFile("./xlsx_template.xlsx")
		if err != nil {
			panic(err)
		}
		config := cs.NewXlsx(f.Sheets[1])
		So(config.GetDescription(), ShouldEqual, "暂无描述")
	})
}

func TestXlsxIndexConfig_GetIndexCount(t *testing.T) {
	Convey("TestXlsxIndexConfig_GetIndexCount", t, func() {
		f, err := xlsx.OpenFile("./xlsx_template.xlsx")
		if err != nil {
			panic(err)
		}
		config := cs.NewXlsx(f.Sheets[1])
		So(config.GetIndexCount(), ShouldEqual, 2)
	})
}
