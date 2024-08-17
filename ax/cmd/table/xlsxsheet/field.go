package xlsxsheet

import (
	"encoding/json"
	"github.com/kercylan98/minotaur/ax/cmd/table/lua2jsonparser"
	"github.com/tealeg/xlsx"
	"github.com/tidwall/sjson"
	"strconv"
	"strings"
)

type field struct {
	table      *Table
	sheet      *xlsx.Sheet
	sheetIndex int // sheet 中的行或列
	index      int
	name       string
	desc       string
	typ        string
	param      string
	idxTable   bool
}

func (f *field) readPos(pos int) (skip, end bool, val string) {
	if !f.idxTable {
		skip = strings.HasPrefix(strings.TrimSpace(f.sheet.Rows[f.sheetIndex].Cells[0].String()), "#")
		val = f.sheet.Rows[f.sheetIndex].Cells[4].String()
	} else {
		skip = strings.HasPrefix(strings.TrimSpace(f.sheet.Rows[pos].Cells[0].String()), "#")
		val = f.sheet.Rows[pos].Cells[f.sheetIndex].String()
		end = f.index > 0 && val == ""
	}
	return
}

func (f *field) rawTypeFormat(raw string) string {
	switch strings.ToLower(strings.TrimSpace(f.typ)) {
	case "string":
		if !strings.HasPrefix(raw, "\"") && !strings.HasSuffix(raw, "\"") {
			raw = strconv.Quote(raw)
		}
	}

	return raw
}

//goland:noinspection t
func (f *field) Query(pos int) (val map[string]any, skip, end bool) {
	// 位置处理
	if !f.idxTable {
		if pos > 0 {
			return nil, false, true
		}
		pos = 0
	} else {
		row := pos + 7
		if row >= len(f.sheet.Rows) {
			end = true
			return
		}
		pos = row
	}

	// 数据读取及格式化
	var raw string
	skip, end, raw = f.readPos(pos)
	if skip || end {
		return
	}

	// 转换数据为实例
	if f.table.lua {
		raw = f.rawTypeFormat(raw)
		raw = lua2jsonparser.New().Parse(raw)
	} else {
		raw = f.rawTypeFormat(raw)
	}

	var err error
	var jsonStr = "{}"
	jsonStr, err = sjson.SetRaw(jsonStr, f.GetName(), raw)
	if err != nil {
		panic(err)
	}

	val = make(map[string]any)
	if err = json.Unmarshal([]byte(jsonStr), &val); err != nil {
		panic(err)
	}

	return
}

func (f *field) IsIgnore() bool {
	var checker = func(v ...string) bool {
		for _, s := range v {
			ts := strings.TrimSpace(s)
			if ts == "" || strings.HasPrefix(ts, "#") {
				return true
			}
		}
		return false
	}

	if checker(f.name, f.desc, f.typ, f.param) {
		return true
	}

	v := strings.ToLower(strings.TrimSpace(f.param))
	switch v {
	case "s", "c", "sc", "cs":
		switch f.table.exportMode {
		case ExportModeCS:
		case ExportModeS:
			if v == "c" {
				return true
			}
		case ExportModeC:
			if v == "s" {
				return true
			}
		}
	default:
		return true
	}

	return false
}

func (f *field) GetIndex() int {
	return f.index
}

func (f *field) GetName() string {
	return f.name
}

func (f *field) GetDesc() string {
	return f.desc
}

func (f *field) GetType() string {
	return f.typ
}

func (f *field) GetParam() string {
	return f.param
}
