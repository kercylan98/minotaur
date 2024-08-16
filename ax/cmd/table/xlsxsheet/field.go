package xlsxsheet

import (
	"encoding/json"
	"fmt"
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

//goland:noinspection t
func (f *field) Query(pos int) (val map[string]any, skip, end bool) {
	var jsonStr = "{}"
	if !f.idxTable {
		if pos > 0 {
			return nil, false, true
		}
		desc := strings.TrimSpace(f.sheet.Rows[f.sheetIndex].Cells[0].String())
		if strings.HasPrefix(desc, "#") {
			skip = true
			return
		}

		var raw = f.sheet.Rows[f.sheetIndex].Cells[4].String()
		// string 特殊检查
		switch strings.ToLower(strings.TrimSpace(f.typ)) {
		case "string":
			if !strings.HasPrefix(raw, "\"") && !strings.HasSuffix(raw, "\"") {
				raw = strconv.Quote(raw)
			}
		}

		var err error
		jsonStr, err = sjson.SetRaw(jsonStr, f.GetName(), raw)
		if err != nil {
			panic(err)
		}

		val = make(map[string]any)
		if err = json.Unmarshal([]byte(jsonStr), &val); err != nil {
			panic(err)
		}
		return
	} else {
		row := pos + 7
		if row >= len(f.sheet.Rows) {
			end = true
			return
		}

		desc := strings.TrimSpace(f.sheet.Rows[row].Cells[0].String())
		if strings.HasPrefix(desc, "#") {
			skip = true
			return
		}

		var raw = f.sheet.Rows[row].Cells[f.sheetIndex].String()
		if f.index > 0 && raw == "" {
			end = true
			return
		}

		// string 特殊检查
		switch strings.ToLower(strings.TrimSpace(f.typ)) {
		case "string":
			if !strings.HasPrefix(raw, "\"") && !strings.HasSuffix(raw, "\"") {
				raw = strconv.Quote(raw)
			}
		}
		var err error
		jsonStr, err = sjson.SetRaw(jsonStr, f.GetName(), raw)
		if err != nil {
			panic(err)
		}

		val = make(map[string]any)
		if err = json.Unmarshal([]byte(jsonStr), &val); err != nil {
			fmt.Println(jsonStr)
			panic(err)
		}
		return
	}
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
