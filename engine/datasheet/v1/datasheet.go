package datasheetv1

import (
	"errors"
	"github.com/kercylan98/minotaur/toolkit"
	"strings"
)

type DataSheetRawCells []*DataSheetRawCell

func (cs DataSheetRawCells) Get(cell int32) *DataSheetRawCell {
	if cell < 0 || cell >= int32(len(cs)) {
		return nil
	}
	return cs[cell]
}

func (cs DataSheetRawCells) First() *DataSheetRawCell {
	return cs.Get(0)
}

func (r *DataSheetRaw) GetWithPos(pos *Pos) *DataSheetRawCell {
	if pos == nil {
		return nil
	}
	if pos.Row < 0 || pos.Row >= int32(len(r.Rows)) {
		return nil
	}
	row := r.Rows[pos.Row]
	if pos.Cell < 0 || pos.Cell >= int32(len(row.Cells)) {
		return nil
	}
	return row.Cells[pos.Cell]
}

func (r *DataSheetRaw) Get(row int32, cell ...int32) DataSheetRawCells {
	if row < 0 || row >= int32(len(r.Rows)) {
		return nil
	}

	if len(cell) == 0 {
		return r.Rows[row].Cells
	}

	var result = make(DataSheetRawCells, len(cell))
	for i, c := range cell {
		if c < 0 || c >= int32(len(r.Rows[row].Cells)) {
			return nil
		}
		result[i] = r.Rows[row].Cells[c]
	}
	return result
}
func (r *DataSheet) HasError() bool {
	return r.Error != ""
}

func (p *Pos) Add(pos *Pos) {
	p.Row += pos.Row
	p.Cell += pos.Cell
}

func ParseStructInfo(typ string) (*DataSheetStruct, error) {
	dataSheetTableStruct := &DataSheetStruct{
		IsOptional:    false,
		Type:          0,
		Name:          "",
		Description:   "",
		TypeInfoOneof: nil,
	}

	if strings.HasPrefix(typ, "*") {
		dataSheetTableStruct.IsOptional = true
		typ = strings.TrimPrefix(typ, "*")
	}

	switch strings.ToLower(typ) {
	case "boolean":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_BOOLEAN
	case "bool":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_BOOL
	case "int":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_INT
	case "int8":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_INT8
	case "int16":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_INT16
	case "int32":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_INT32
	case "int64":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_INT64
	case "uint":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_UINT
	case "uint8":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_UINT8
	case "uint16":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_UINT16
	case "uint32":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_UINT32
	case "uint64":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_UINT64
	case "float":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_FLOAT
	case "double":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_DOUBLE
	case "byte":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_Byte
	case "time":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_TIME
	case "float32":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_FLOAT32
	case "float64":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_FLOAT64
	case "short":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_SHORT
	case "long":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_LONG
	case "string":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_STRING
	case "bigint":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_BIGINT
	case "bigfloat":
		dataSheetTableStruct.Type = DataSheetStructType_DATA_SHEET_FIELD_TYPE_BIGFLOAT
	default:
		typeInfo, err := newParser(typ).parse()
		if err != nil {
			return nil, err
		}

		return parseOther(typeInfo)
	}

	return dataSheetTableStruct, nil
}

func parseOther(typeInfo any) (*DataSheetStruct, error) {
	switch value := typeInfo.(type) {
	case *Type:
		switch value.Name {
		case "struct", "slice", "array", "map":
			return parseOther(value.Value)
		default:
			return ParseStructInfo(value.Name)
		}
	case *Struct:
		oneof := &DataSheetStruct_StructInfo{
			StructInfo: &DataSheetStructInfo{},
		}
		for fieldName, fieldType := range value.Fields {
			var fieldInfo *DataSheetStruct
			var err error
			switch fieldType.Name {
			case "struct", "slice", "array", "map":
				fieldInfo, err = parseOther(fieldType.Value)
			default:
				fieldInfo, err = ParseStructInfo(fieldType.Name)
			}
			if err != nil {
				return nil, err
			}
			fieldInfo.FieldName = fieldName
			oneof.StructInfo.Fields = append(oneof.StructInfo.Fields, fieldInfo)
		}
		return &DataSheetStruct{
			IsOptional:    false,
			Type:          DataSheetStructType_DATA_SHEET_FIELD_TYPE_STRUCT,
			TypeInfoOneof: oneof,
		}, nil
	case *Array:
		var elemInfo *DataSheetStruct
		var err error
		switch value.Elem.Name {
		case "struct", "slice", "array", "map":
			elemInfo, err = parseOther(value.Elem.Value)
		default:
			elemInfo, err = ParseStructInfo(value.Elem.Name)
		}
		if err != nil {
			return nil, err
		}
		return &DataSheetStruct{
			IsOptional: false,
			Type:       toolkit.If(value.Size != -1, DataSheetStructType_DATA_SHEET_FIELD_TYPE_ARRAY, DataSheetStructType_DATA_SHEET_FIELD_TYPE_SLICE),
			TypeInfoOneof: &DataSheetStruct_ArrayInfo{
				ArrayInfo: &DataSheetArrayInfo{
					Length:     int32(value.Size),
					StructInfo: elemInfo,
				},
			},
		}, nil
	case *Map:
		keyElem, err := parseOther(value.Key)
		if err != nil {
			return nil, err
		}
		valElem, err := parseOther(value.Value)
		if err != nil {
			return nil, err
		}
		return &DataSheetStruct{
			IsOptional: false,
			Type:       DataSheetStructType_DATA_SHEET_FIELD_TYPE_MAP,
			TypeInfoOneof: &DataSheetStruct_MapInfo{
				MapInfo: &DataSheetMapInfo{
					KeyStructInfo:   keyElem,
					ValueStructInfo: valElem,
				},
			},
		}, nil
	default:
		return nil, errors.New("unknown type")
	}

}
