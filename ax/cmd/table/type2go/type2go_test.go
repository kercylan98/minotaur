package type2go

import (
	"github.com/kercylan98/minotaur/ax/cmd/table"
	"github.com/kercylan98/minotaur/ax/cmd/table/fieldparser"
	"github.com/kercylan98/minotaur/ax/cmd/table/lua2jsonparser"
	"github.com/kercylan98/minotaur/ax/cmd/table/xlsxsheet"
	"github.com/kercylan98/minotaur/toolkit/fileproc"
	"github.com/tealeg/xlsx"
	"testing"
)

func TestGen(t *testing.T) {
	xlsxFile, _ := xlsx.OpenFile(`..\xlsxsheet\template.xlsx`)
	t1 := xlsxsheet.NewTable(xlsxFile.Sheets[1])
	r := table.GenerateConfigs([]table.Table{t1}, fieldparser.New(), New("type2go"), lua2jsonparser.New())
	code := r.GenerateCode()
	fileproc.WriteToFile("./template_test_config.go", []byte(code))
}
