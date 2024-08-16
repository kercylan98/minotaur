package cmd

import (
	"encoding/json"
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
	"strings"
)

var (
	xlsxTableSheet2JSONFilepath   string
	xlsxTableSheet2JSONSheetName  string
	xlsxTableSheet2JSONOutput     string
	xlsxTableSheet2JSONExportMode string
)

var tableXlsxSheet2JSONCmd = &cobra.Command{
	Use:   "sheet2json",
	Short: "Convert xlsx sheet to json data file",
	RunE: func(cmd *cobra.Command, args []string) error {
		return onTableXlsxSheet2JSON(xlsxTableSheet2JSONFilepath, xlsxTableSheet2JSONSheetName, xlsxTableSheet2JSONOutput)
	},
}

func init() {
	tableXlsxCmd.AddCommand(tableXlsxSheet2JSONCmd)

	tableXlsxSheet2JSONCmd.Flags().StringVarP(&xlsxTableSheet2JSONFilepath, "filepath", "p", "", "xlsx filepath")
	tableXlsxSheet2JSONCmd.Flags().StringVarP(&xlsxTableSheet2JSONSheetName, "sheetName", "s", "", "xlsx sheet name")
	tableXlsxSheet2JSONCmd.Flags().StringVarP(&xlsxTableSheet2JSONOutput, "output-dir", "o", "", "output dir path")
	tableXlsxSheet2JSONCmd.Flags().StringVarP(&xlsxTableSheet2JSONExportMode, "export-mode", "m", "sc", "export only the fields contained in the parameters(sc/s/c)")

	checkError(tableXlsxSheet2JSONCmd.MarkFlagRequired("filepath"))
	checkError(tableXlsxSheet2JSONCmd.MarkFlagRequired("sheetName"))
	checkError(tableXlsxSheet2JSONCmd.MarkFlagRequired("output-dir"))
}

func onTableXlsxSheet2JSON(xlsxFilePath, sheetName, output string) error {
	xlsxFile, err := xlsx.OpenFile(xlsxFilePath)
	if err != nil {
		return err
	}
	var mode xlsxsheet.ExportMode
	switch strings.ToLower(xlsxTableSheet2JSONExportMode) {
	case "c", "cli", "client":
		mode = xlsxsheet.ExportModeC
	case "s", "srv", "server":
		mode = xlsxsheet.ExportModeS
	case "sc", "cs", "cli-srv", "srv-cli":
		mode = xlsxsheet.ExportModeCS
	default:
		checkError("export mode is not support")
	}
	tab := xlsxsheet.NewTable(xlsxFile.Sheet[sheetName], mode)
	if tab.IsIgnore() {
		return fmt.Errorf("config sheet %s is ignore", sheetName)
	}

	exist := fileproc.CheckPathExist(output)
	dir := (!exist && filepath.Ext(output) == charproc.None) || fileproc.CheckIsDir(output)

	if dir {
		checkError(os.MkdirAll(output, os.ModePerm))
	} else {
		panic("output path is not a directory")
	}

	configs := table.GenerateConfigs([]table.Table{tab}, fieldparser.New(), type2go.New("none"), lua2jsonparser.New())
	for name, data := range configs.GenerateData() {
		bytes, err := json.MarshalIndent(data, "", "  ")
		checkError(err)
		checkError(fileproc.WriteToFile(filepath.Join(output, name+".json"), bytes))
		fmt.Println("Convert xlsx sheet to json data file to: " + filepath.Join(output, name+".json"))
	}

	return nil
}
