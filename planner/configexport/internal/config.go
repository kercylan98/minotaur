package internal

import (
	"bytes"
	"fmt"
	jsonIter "github.com/json-iterator/go"
	"github.com/kercylan98/minotaur/utils/g2d/matrix"
	"github.com/kercylan98/minotaur/utils/str"
	"github.com/kercylan98/minotaur/utils/xlsxtool"
	"github.com/tealeg/xlsx"
	"strconv"
	"strings"
	"text/template"
)

// NewConfig 定位为空将读取Sheet名称
//   - 表格中需要严格遵守 描述、名称、类型、导出参数、数据列的顺序
func NewConfig(sheet *xlsx.Sheet) (*Config, error) {
	config := &Config{
		ignore:        "#",
		excludeFields: map[int]bool{0: true},
	}
	if err := config.initField(sheet); err != nil {
		return nil, err
	}
	return config, nil
}

type Config struct {
	DisplayName string
	Name        string
	Describe    string
	ExportParam string
	IndexCount  int
	Fields      []*Field

	matrix        *matrix.Matrix[*xlsx.Cell]
	excludeFields map[int]bool // 排除的字段
	ignore        string
	horizontal    bool
	dataStart     int

	dataServer map[any]any
	dataClient map[any]any
}

func (slf *Config) initField(sheet *xlsx.Sheet) error {
	slf.matrix = xlsxtool.GetSheetMatrix(sheet)
	var displayName *Position
	name, indexCount := NewPosition(1, 0), NewPosition(1, 1)
	if displayName == nil {
		slf.DisplayName = sheet.Name
	} else {
		if value := slf.matrix.Get(displayName.X, displayName.Y); value == nil {
			return ErrReadConfigFailedWithDisplayName
		} else {
			slf.DisplayName = strings.TrimSpace(value.String())
		}
	}

	if name == nil {
		slf.Name = sheet.Name
	} else {
		if value := slf.matrix.Get(name.X, name.Y); value == nil {
			return ErrReadConfigFailedWithName
		} else {
			slf.Name = str.FirstUpper(strings.TrimSpace(value.String()))
		}
	}

	if indexCount == nil {
		slf.IndexCount, _ = strconv.Atoi(sheet.Name)
	} else {
		if value := slf.matrix.Get(indexCount.X, indexCount.Y); value == nil {
			return ErrReadConfigFailedWithIndexCount
		} else {
			indexCount, err := value.Int()
			if err != nil {
				return err
			}
			if indexCount < 0 {
				return ErrReadConfigFailedWithIndexCountLessThanZero
			}
			slf.IndexCount = indexCount
		}
	}

	var (
		describeStart                  int
		horizontal                     bool
		fields                         = make(map[string]bool)
		dx, dy, nx, ny, tx, ty, ex, ey int
	)
	horizontal = slf.IndexCount > 0
	slf.horizontal = horizontal

	if horizontal {
		describeStart = 3
		dy = describeStart
		ny = dy + 1
		ty = ny + 1
		ey = ty + 1
		slf.dataStart = ey + 1
	} else {
		delete(slf.excludeFields, 0)
		describeStart = 4
		dy = describeStart
		ny = dy
		ty = dy
		ey = dy
		nx = dx + 1
		tx = nx + 1
		ex = tx + 1
		slf.dataStart = ey + 1
	}
	var index = slf.IndexCount
	for {
		var (
			describe, fieldName, fieldType, exportParam string
		)

		if value := slf.matrix.Get(dx, dy); value == nil {
			return ErrReadConfigFailedWithFieldPosition
		} else {
			describe = value.String()
		}
		if value := slf.matrix.Get(nx, ny); value == nil {
			return ErrReadConfigFailedWithFieldPosition
		} else {
			fieldName = str.FirstUpper(strings.TrimSpace(value.String()))
		}
		if value := slf.matrix.Get(tx, ty); value == nil {
			return ErrReadConfigFailedWithFieldPosition
		} else {
			fieldType = strings.TrimSpace(value.String())
		}
		if value := slf.matrix.Get(ex, ey); value == nil {
			return ErrReadConfigFailedWithFieldPosition
		} else {
			exportParam = strings.ToLower(strings.TrimSpace(value.String()))
		}
		var field = NewField(slf.Name, fieldName, fieldType)
		field.Describe = describe
		field.ExportParam = exportParam
		switch field.ExportParam {
		case "s", "sc", "cs":
			field.Server = true
		}

		if horizontal {
			dx++
			nx++
			tx++
			ex++
		} else {
			dy++
			ny++
			ty++
			ey++
		}

		field.Ignore = slf.excludeFields[len(slf.Fields)]
		if !field.Ignore {
			if strings.HasPrefix(field.Describe, slf.ignore) {
				field.Ignore = true
			} else if strings.HasPrefix(field.Name, slf.ignore) {
				field.Ignore = true
			} else if strings.HasPrefix(field.Type, slf.ignore) {
				field.Ignore = true
			} else if strings.HasPrefix(field.ExportParam, slf.ignore) {
				field.Ignore = true
			}
		}
		if !field.Ignore {
			switch exportParam {
			case "s", "c", "sc", "cs":
			default:
				return ErrReadConfigFailedWithExportParamException
			}
		}

		if fields[field.Name] && !field.Ignore {
			return ErrReadConfigFailedWithNameDuplicate
		}
		if index > 0 && !field.Ignore {
			if _, exist := basicTypeName[field.Type]; !exist {
				return ErrReadConfigFailedWithIndexTypeException
			}
			index--
		}

		fields[field.Name] = true
		slf.Fields = append(slf.Fields, field)

		if horizontal {
			if dx >= slf.matrix.GetWidth() {
				break
			}
		} else {
			if dy >= slf.matrix.GetHeight() {
				break
			}
		}
	}

	return slf.initData()
}

func (slf *Config) initData() error {
	var x, y int
	if slf.horizontal {
		y = slf.dataStart
	} else {
		x = 4
		y = 4
	}
	var dataSourceServer = make(map[any]any)
	var dataSourceClient = make(map[any]any)
	for {
		var dataServer = dataSourceServer
		var dataClient = dataSourceClient
		var currentIndex = 0

		var lineServer = map[any]any{}
		var lineClient = map[any]any{}
		for i := 0; i < len(slf.Fields); i++ {
			if slf.excludeFields[i] {
				continue
			}

			field := slf.Fields[i]
			var value any
			if slf.horizontal {
				value = getValueWithType(field.SourceType, slf.matrix.Get(x+i, y).String())
			} else {
				value = getValueWithType(field.SourceType, slf.matrix.Get(x, y+i).String())
			}
			switch field.ExportParam {
			case "s":
				lineServer[field.Name] = value
			case "c":
				lineClient[field.Name] = value
			case "sc", "cs":
				lineServer[field.Name] = value
				lineClient[field.Name] = value
			}

			if currentIndex < slf.IndexCount {
				currentIndex++
				m, exist := dataServer[value]
				if !exist {
					if currentIndex == slf.IndexCount {
						dataServer[value] = lineServer
					} else {
						m = map[any]any{}
						dataServer[value] = m
					}
				}
				if currentIndex < slf.IndexCount {
					dataServer = m.(map[any]any)
				}

				m, exist = dataClient[value]
				if !exist {
					if currentIndex == slf.IndexCount {
						dataClient[value] = lineClient
					} else {
						m = map[any]any{}
						dataClient[value] = m
					}
				}
				if currentIndex < slf.IndexCount {
					dataClient = m.(map[any]any)
				}
			}
		}
		if slf.horizontal {
			y++
			if y >= slf.matrix.GetHeight() {
				break
			}
		} else {
			x++
			if x >= slf.matrix.GetWidth() {
				slf.dataServer = lineServer
				slf.dataClient = lineClient
				break
			}
		}
	}
	if slf.horizontal {
		slf.dataServer = dataSourceServer
		slf.dataClient = dataSourceClient
	}
	return nil
}

func (slf *Config) String() string {
	tmpl, err := template.New("struct").Parse(generateConfigTemplate)
	if err != nil {
		return ""
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, slf); err != nil {
		return ""
	}
	var result string
	_ = str.RangeLine(buf.String(), func(index int, line string) error {
		if len(strings.TrimSpace(line)) == 0 {
			return nil
		}
		result += fmt.Sprintf("%s\n", strings.ReplaceAll(line, "\t\t", "\t"))
		if len(strings.TrimSpace(line)) == 1 {
			result += "\n"
		}
		return nil
	})
	return result
}

func (slf *Config) JsonServer() []byte {
	d, _ := jsonIter.MarshalIndent(slf.dataServer, "", "  ")
	return d
}

func (slf *Config) JsonClient() []byte {
	d, _ := jsonIter.MarshalIndent(slf.dataClient, "", "  ")
	return d
}

func (slf *Config) GetVariable() string {
	var result string
	var count int
	if slf.IndexCount > 0 {
		for _, field := range slf.Fields {
			if field.Ignore {
				continue
			}
			result += fmt.Sprintf("map[%s]", field.Type)
			count++
			if count >= slf.IndexCount {
				break
			}
		}
	}
	return fmt.Sprintf("%s*%s", result, slf.Name)
}
