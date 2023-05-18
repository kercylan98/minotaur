package xlsxtool

import (
	"github.com/kercylan98/minotaur/utils/g2d/matrix"
	"github.com/tealeg/xlsx"
)

// GetSheetMatrix 将sheet转换为二维矩阵
func GetSheetMatrix(sheet *xlsx.Sheet) *matrix.Matrix[*xlsx.Cell] {
	m := matrix.NewMatrix[*xlsx.Cell](sheet.MaxCol, sheet.MaxRow)
	for y := 0; y < sheet.MaxRow; y++ {
		for x := 0; x < sheet.MaxCol; x++ {
			if x >= len(sheet.Rows[y].Cells) {
				continue
			}
			m.Set(x, y, sheet.Rows[y].Cells[x])
		}
	}
	return m
}
