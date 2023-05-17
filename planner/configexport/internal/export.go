package internal

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/file"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/tealeg/xlsx"
	"go.uber.org/zap"
	"path/filepath"
	"runtime/debug"
	"sync"
)

func ExportJSON(xlsxPath string, output string) {
	xlsxFile, err := xlsx.OpenFile(xlsxPath)
	if err != nil {
		panic(err)
	}
	var errors []func()
	var wait sync.WaitGroup
	for _, sheet := range xlsxFile.Sheets {
		sheet := sheet
		go func() {
			defer func() {
				if err := recover(); err != nil {
					errors = append(errors, func() {
						log.Error("导出失败", zap.String("名称", xlsxPath), zap.String("Sheet", sheet.Name), zap.Any("err", err))
						fmt.Println(debug.Stack())
					})
				}
			}()
			wait.Add(1)
			config := NewConfig(sheet)
			if err := file.WriterFile(filepath.Join(output, fmt.Sprintf("%s.json", config.GetName())), config.GetJSON()); err != nil {
				panic(err)
			}
			wait.Done()
		}()
	}

	wait.Wait()

	for _, f := range errors {
		f()
	}
}
