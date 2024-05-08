package main

import (
	"fmt"
	jsonIter "github.com/json-iterator/go"
	"github.com/kercylan98/minotaur/configuration"
	"github.com/kercylan98/minotaur/configuration/exporter"
	"github.com/kercylan98/minotaur/configuration/generator"
	"github.com/kercylan98/minotaur/configuration/scanner"
	"github.com/kercylan98/minotaur/configuration/test/config"
	"github.com/tealeg/xlsx"
)

func main() {
	xf, err := xlsx.OpenFile(`D:\sources\minotaur\configuration\xlsx_template.xlsx`)
	if err != nil {
		panic(err)
	}

	scanners := make([]configuration.Scanner, 0, len(xf.Sheets))
	for i, sheet := range xf.Sheets {
		if i == 0 {
			continue
		}
		scanners = append(scanners, scanner.NewXlsxSheetScanner(sheet))
	}

	if err = configuration.GenerateCode(generator.NewGolangSingleFile("config", "./test/config/config.go"), scanners...); err != nil {
		panic(err)
	}

	if err = configuration.ExportData(exporter.NewJSON(), scanners...); err != nil {
		panic(err)
	}

	var c = make(map[int]map[string]*config.IndexConfig)
	if err := jsonIter.Unmarshal([]byte(tempJson), &c); err != nil {
		panic(err)
	}

	config.SetIndexConfig(c)
	config.LoadIndexConfig()
	fmt.Println(config.GetIndexConfig()[1]["a"].Count)
}

const tempJson = `
{
  "1": {
    "a": {
      "award": [
        "asd",
        "12"
      ],
      "info": {
        "id": 1,
        "name": "小明",
        "info": {
          "lv": 1,
          "exp": {
            "mux": 10,
            "count": 100
          }
        }
      },
      "other": [
        {
          "id": 1,
          "name": "张飞"
        },
        {
          "id": 2,
          "name": "刘备"
        }
      ],
      "count": "a",
      "id": 1
    },
    "b": {
      "award": [
        "asd",
        "12"
      ],
      "info": {
        "id": 1,
        "name": "小明",
        "info": {
          "lv": 1,
          "exp": {
            "mux": 10,
            "count": 100
          }
        }
      },
      "other": [
        {
          "id": 1,
          "name": "张飞"
        },
        {
          "id": 2,
          "name": "刘备"
        }
      ],
      "id": 1,
      "count": "b"
    }
  },
  "2": {
    "c": {
      "award": [
        "asd",
        "12"
      ],
      "info": {
        "id": 1,
        "name": "小明",
        "info": {
          "exp": {
            "mux": 10,
            "count": 100
          },
          "lv": 1
        }
      },
      "other": [
        {
          "name": "张飞",
          "id": 1
        },
        {
          "id": 2,
          "name": "刘备"
        }
      ],
      "id": 2,
      "count": "c"
    },
    "d": {
      "info": {
        "id": 1,
        "name": "小明",
        "info": {
          "lv": 1,
          "exp": {
            "mux": 10,
            "count": 100
          }
        }
      },
      "other": [
        {
          "id": 1,
          "name": "张飞"
        },
        {
          "id": 2,
          "name": "刘备"
        }
      ],
      "id": 2,
      "count": "d",
      "award": [
        "asd",
        "12"
      ]
    }
  }
}
`
