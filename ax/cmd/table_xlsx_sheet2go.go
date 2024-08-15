package cmd

import (
	"fmt"
	"github.com/kercylan98/minotaur/ax/cmd/table"
	"github.com/kercylan98/minotaur/ax/cmd/table/fieldparser"
	"github.com/kercylan98/minotaur/ax/cmd/table/lua2jsonparser"
	"github.com/kercylan98/minotaur/ax/cmd/table/type2go"
	"github.com/kercylan98/minotaur/ax/cmd/table/xlsxsheet"
	"github.com/kercylan98/minotaur/toolkit/fileproc"
	"github.com/spf13/cobra"
	"github.com/tealeg/xlsx"
)

var (
	sheet2GoFilepath  string
	sheet2GoSheetName string
	sheet2GoOutput    string
	sheet2GoPackage   string
)

var tableXlsxSheet2GoCmd = &cobra.Command{
	Use:   "sheet2go",
	Short: "Convert xlsx sheet to go configuration code",
	RunE: func(cmd *cobra.Command, args []string) error {
		defer func() {
			fmt.Println("Convert xlsx sheet to go configuration code to: " + sheet2GoOutput)
		}()
		xlsxFile, _ := xlsx.OpenFile(sheet2GoFilepath)
		t1 := xlsxsheet.NewTable(xlsxFile.Sheet[sheet2GoSheetName])
		r := table.GenerateConfigs([]table.Table{t1}, fieldparser.New(), type2go.New(sheet2GoPackage), lua2jsonparser.New())
		code := r.GenerateCode()
		return fileproc.WriteToFile(sheet2GoOutput, code)
	},
}

func init() {
	tableXlsxCmd.AddCommand(tableXlsxSheet2GoCmd)

	tableXlsxSheet2GoCmd.Flags().StringVarP(&sheet2GoFilepath, "filepath", "p", "", "xlsx filepath")
	tableXlsxSheet2GoCmd.Flags().StringVarP(&sheet2GoSheetName, "sheetName", "s", "", "xlsx sheet name")
	tableXlsxSheet2GoCmd.Flags().StringVarP(&sheet2GoOutput, "output", "o", "", "output filepath")
	tableXlsxSheet2GoCmd.Flags().StringVarP(&sheet2GoPackage, "package", "n", "", "package name")

	checkError(tableXlsxSheet2GoCmd.MarkFlagRequired("filepath"))
	checkError(tableXlsxSheet2GoCmd.MarkFlagRequired("sheetName"))
	checkError(tableXlsxSheet2GoCmd.MarkFlagRequired("output"))
	checkError(tableXlsxSheet2GoCmd.MarkFlagRequired("package"))
}
