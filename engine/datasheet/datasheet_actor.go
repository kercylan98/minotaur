package datasheet

import (
	"bytes"
	datasheetv1 "github.com/kercylan98/minotaur/engine/datasheet/v1"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/xuri/excelize/v2"
	"strings"
)

type actor struct {
}

func (a *actor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *datasheetv1.LoadDataSheetFileRequest:
		a.onLoadDatasheetFileRequest(ctx, m)
	case *datasheetv1.GenerateDataSheetRequest:
		a.onGenerateDataSheetRequest(ctx, m)
	}
}

func (a *actor) onLoadDatasheetFileRequest(ctx vivid.ActorContext, m *datasheetv1.LoadDataSheetFileRequest) {
	reader := bytes.NewReader(m.BinaryData)

	dataSheetFile, err := excelize.OpenReader(reader)
	if err != nil {
		ctx.Reply(err)
	}

	defer func() {
		if err = dataSheetFile.Close(); err != nil {
			ctx.System().Logger().Warn("DataSheet", log.String("close", "data sheet binary file"), log.Err(err))
		}
	}()

	var response = &datasheetv1.LoadDataSheetFileResponse{}
	for i, sheetName := range dataSheetFile.GetSheetList() {
		var sheet = &datasheetv1.DataSheetRaw{
			Index: int32(i),
			Name:  sheetName,
		}

		rows, err := dataSheetFile.GetRows(sheetName)
		if err != nil {
			response.LoadDataSheetFailedInfos = append(response.LoadDataSheetFailedInfos, &datasheetv1.LoadDataSheetFailedInfo{
				Name: sheetName,
				Err:  err.Error(),
			})
			continue
		}

		sheet.Rows = make([]*datasheetv1.DataSheetRawRow, len(rows))
		for rowIndex, cells := range rows {
			sheet.Rows[rowIndex] = &datasheetv1.DataSheetRawRow{
				Cells: make([]*datasheetv1.DataSheetRawCell, len(cells)),
			}
			for cellIndex, cell := range cells {
				sheet.Rows[rowIndex].Cells[cellIndex] = &datasheetv1.DataSheetRawCell{
					Value: cell,
				}
			}
		}
		for _, row := range rows {
			for _, cell := range row {
				sheet.Name = cell
				break
			}
		}

		response.DataSheetRaws = append(response.DataSheetRaws, sheet)
	}

	ctx.Reply(response)
}

func (a *actor) onGenerateDataSheetRequest(ctx vivid.ActorContext, m *datasheetv1.GenerateDataSheetRequest) {
	dataSheet := &datasheetv1.DataSheet{}

	a.onGenerateDataSheetName(ctx, m, dataSheet)
	a.onGenerateDataSheetFields(ctx, m, dataSheet)

	ctx.Reply(&datasheetv1.GenerateDataSheetResponse{DataSheet: dataSheet})
}

func (a *actor) onGenerateDataSheetName(ctx vivid.ActorContext, m *datasheetv1.GenerateDataSheetRequest, sheet *datasheetv1.DataSheet) {
	if sheet.HasError() {
		return
	}

	if cell := m.DataSheetRaw.GetWithPos(m.DataSheetNamePos); cell != nil {
		sheet.Name = cell.Value
	} else {
		sheet.Name = m.DataSheetRaw.Name
	}
}

func (a *actor) onGenerateDataSheetFields(ctx vivid.ActorContext, m *datasheetv1.GenerateDataSheetRequest, sheet *datasheetv1.DataSheet) {
	if sheet.HasError() {
		return
	}

	// 字段名称
	nameCurr := m.FieldScannerConfig.NameStartPos
	descriptionCurr := m.FieldScannerConfig.DescriptionStartPos
	typCurr := m.FieldScannerConfig.TypeStartPos
	groupsCurr := m.FieldScannerConfig.GroupsStartPos
	for {
		nameCell := m.DataSheetRaw.GetWithPos(nameCurr)
		descriptionCell := m.DataSheetRaw.GetWithPos(descriptionCurr)
		typCell := m.DataSheetRaw.GetWithPos(typCurr)
		groupsCell := m.DataSheetRaw.GetWithPos(groupsCurr)

		if nameCell == nil || typCell == nil || groupsCell == nil || descriptionCell == nil {
			break
		}

		field := &datasheetv1.DataSheetField{
			Name:        nameCell.Value,
			Description: descriptionCell.Value,
			Groups:      strings.Split(groupsCell.Value, ","),
		}
		structInfo, err := datasheetv1.ParseStructInfo(typCell.Value)
		if err != nil {
			sheet.Error = err.Error()
			return
		}
		field.StructInfo = structInfo

		sheet.Fields = append(sheet.Fields, field)

		nameCurr.Add(m.FieldScannerConfig.NameNextDeltaPos)
		descriptionCurr.Add(m.FieldScannerConfig.DescriptionNextDeltaPos)
		typCurr.Add(m.FieldScannerConfig.TypeNextDeltaPos)
		groupsCurr.Add(m.FieldScannerConfig.GroupsNextDeltaPos)
	}
}
