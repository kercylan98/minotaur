package cmd

import (
	"fmt"
	"github.com/kercylan98/minotaur/ax/cmd/table"
	"github.com/kercylan98/minotaur/ax/cmd/table/fieldparser"
	"github.com/kercylan98/minotaur/ax/cmd/table/lua2jsonparser"
	"github.com/kercylan98/minotaur/ax/cmd/table/type2go"
	"github.com/kercylan98/minotaur/ax/cmd/table/xlsxsheet"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"github.com/kercylan98/minotaur/toolkit/fileproc"
	"github.com/spf13/cobra"
	"github.com/tealeg/xlsx"
	"os"
	"path/filepath"
)

var (
	xlsxTableSheet2GoFilepath  string
	xlsxTableSheet2GoSheetName string
	xlsxTableSheet2GoOutput    string
	xlsxTableSheet2GoPackage   string
)

var tableXlsxSheet2GoCmd = &cobra.Command{
	Use:   "sheet2go",
	Short: "Convert an xlsx sheet to Go configuration code.",
	Long:  `Converts a specified sheet from an xlsx file into Go configuration code, allowing easy integration of configuration data into your Go applications.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return onTableXlsxSheet2Go(xlsxTableSheet2GoFilepath, xlsxTableSheet2GoSheetName, xlsxTableSheet2GoOutput, xlsxTableSheet2GoPackage)
	},
}

func init() {
	tableXlsxCmd.AddCommand(tableXlsxSheet2GoCmd)

	tableXlsxSheet2GoCmd.Flags().StringVarP(&xlsxTableSheet2GoFilepath, "filepath", "p", "", "xlsx filepath")
	tableXlsxSheet2GoCmd.Flags().StringVarP(&xlsxTableSheet2GoSheetName, "sheetName", "s", "", "xlsx sheet name")
	tableXlsxSheet2GoCmd.Flags().StringVarP(&xlsxTableSheet2GoOutput, "output", "o", "", "output filepath")
	tableXlsxSheet2GoCmd.Flags().StringVarP(&xlsxTableSheet2GoPackage, "package", "n", "", "package name")

	checkError(tableXlsxSheet2GoCmd.MarkFlagRequired("filepath"))
	checkError(tableXlsxSheet2GoCmd.MarkFlagRequired("sheetName"))
	checkError(tableXlsxSheet2GoCmd.MarkFlagRequired("output"))
	checkError(tableXlsxSheet2GoCmd.MarkFlagRequired("package"))
}

func onTableXlsxSheet2Go(xlsxFilePath, sheetName, output, packageName string) error {
	xlsxFile, err := xlsx.OpenFile(xlsxFilePath)
	if err != nil {
		return err
	}
	tab := xlsxsheet.NewTable(xlsxFile.Sheet[sheetName], 0, false)
	if tab.IsIgnore() {
		return fmt.Errorf("config sheet %s is ignore", sheetName)
	}

	configs := table.GenerateConfigs([]table.Table{tab}, fieldparser.New(), type2go.New(packageName), lua2jsonparser.New())
	code := configs.GenerateCode()

	exist := fileproc.CheckPathExist(output)
	dir := (!exist && filepath.Ext(output) == charproc.None) || fileproc.CheckIsDir(output)

	if dir {
		checkError(os.MkdirAll(output, os.ModePerm))
		output = filepath.Join(output, charproc.Snake(tab.GetName()))
	} else if filepath.Ext(output) != ".go" {
		output += ".go"
	}

	if err := fileproc.WriteToFile(output, code); err != nil {
		return err
	}

	fmt.Println("Convert xlsx sheet to go configuration code to: " + output)
	return nil
}
