package analyzer

import (
	"fmt"
	"github.com/kercylan98/minotaur/exporter/configuration"
	"github.com/kercylan98/minotaur/exporter/configuration/golang"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

type Default struct {
}

func (slf *Default) Analyze(filePath string) (map[string]configuration.Configuration, error) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	var configs = make(map[string]configuration.Configuration)

	for _, sheetName := range file.GetSheetList() {

		rs, err := file.GetRows(sheetName)
		if err != nil {
			log.Error("Analyze", zap.Error(err))
			continue
		}
		var rows = make(map[int][]string)
		for ri, row := range rs {
			r := make([]string, 0, len(row))
			for _, col := range row {
				r = append(r, col)
			}
			rows[ri] = r
		}

		if len(rows[0]) < 2 || len(rows[1]) < 2 {
			continue
		}

		var (
			name          = strings.TrimSpace(rows[0][1])
			indexCountStr = strings.TrimSpace(rows[1][1])
			indexCount    int
		)

		if name == "" || indexCountStr == "" {
			continue
		} else {
			indexCount, err = strconv.Atoi(indexCountStr)
			if err != nil {
				log.Error("Analyze", zap.Error(err))
				continue
			}
		}

		if len(rows[3]) < 2 || len(rows[4]) < 2 || len(rows[5]) < 2 || len(rows[6]) < 2 {
			continue
		}

		config := golang.NewConfiguration(name)

		for i := range rows[3] {
			if i == 0 {
				continue
			}
			var field = golang.NewField(i, rows[4][i], golang.GetFieldType(rows[5][i]), i-1 < indexCount)
			config.AddField(field)
		}

		fmt.Println(config.GetStruct())
	}

	return configs, nil
}
