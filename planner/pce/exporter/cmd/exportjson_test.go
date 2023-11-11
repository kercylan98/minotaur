package cmd_test

import (
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/planner/pce"
	"github.com/kercylan98/minotaur/planner/pce/cs"
	"github.com/kercylan98/minotaur/planner/pce/tmpls"
	"github.com/kercylan98/minotaur/utils/file"
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/str"
	"github.com/tealeg/xlsx"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExecute(t *testing.T) {
	var filePath, outPath, exclude, exportType, prefix string

	exportType = "s"
	filePath = `.\游戏配置.xlsx`
	filePath = `../xlsx_template.xlsx`
	outPath = `.`

	isDir, err := file.IsDir(outPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			isDir = filepath.Ext(outPath) == ""
		} else {
			panic(err)
		}
	}
	if !isDir {
		panic(errors.New("output must be a directory path"))
	}
	_ = os.MkdirAll(outPath, os.ModePerm)

	fpd, err := file.IsDir(filePath)
	if err != nil {
		panic(err)
	}

	var xlsxFiles []string
	if fpd {
		files, err := os.ReadDir(filePath)
		if err != nil {
			panic(err)
		}
		for _, f := range files {
			if f.IsDir() || !strings.HasSuffix(f.Name(), ".xlsx") || strings.HasPrefix(f.Name(), "~") {
				continue
			}
			xlsxFiles = append(xlsxFiles, filepath.Join(filePath, f.Name()))
		}
	} else {
		xlsxFiles = append(xlsxFiles, filePath)
	}

	var exporter = pce.NewExporter()
	loader := pce.NewLoader(pce.GetFields())

	excludes := hash.ToMapBool(str.SplitTrimSpace(exclude, ","))
	for _, xlsxFile := range xlsxFiles {
		xf, err := xlsx.OpenFile(xlsxFile)
		if err != nil {
			panic(err)
		}

		for _, sheet := range xf.Sheets {
			var cx *cs.Xlsx
			switch strings.TrimSpace(strings.ToLower(exportType)) {
			case "c":
				cx = cs.NewXlsx(sheet, cs.XlsxExportTypeClient)
			case "s":
				cx = cs.NewXlsx(sheet, cs.XlsxExportTypeServer)
			}
			if strings.HasPrefix(cx.GetDisplayName(), "#") || strings.HasPrefix(cx.GetConfigName(), "#") || excludes[cx.GetConfigName()] || excludes[cx.GetDisplayName()] {
				continue
			}

			if raw, err := exporter.ExportData(tmpls.NewJSON(), loader.LoadData(cx)); err != nil {
				panic(err)
			} else {
				var jsonPath string
				if len(prefix) == 0 {
					jsonPath = filepath.Join(outPath, fmt.Sprintf("%s.json", cx.GetConfigName()))
				} else {
					jsonPath = filepath.Join(outPath, fmt.Sprintf("%s.%s.json", prefix, cx.GetConfigName()))
				}
				if err := file.WriterFile(jsonPath, raw); err != nil {
					panic(err)
				}
			}
		}
	}
}
