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
	"github.com/kercylan98/minotaur/toolkit/collection"
	"github.com/kercylan98/minotaur/toolkit/fileproc"
	"github.com/spf13/cobra"
	"github.com/tealeg/xlsx"
	"os"
	"path/filepath"
	"strings"
)

var (
	tableXlsxXlsx2JSONFilepath   string
	tableXlsxXlsx2JSONOutput     string
	tableXlsxXlsx2JSONExportMode string
)

var tableXlsxXlsx2JSONCmd = &cobra.Command{
	Use:   "xlsx2json",
	Short: "Convert xlsx file sheets to json data file",
	RunE: func(cmd *cobra.Command, args []string) error {
		return onTableXlsxXlsx2JSON(tableXlsxXlsx2JSONFilepath, tableXlsxXlsx2JSONOutput)
	},
}

func init() {
	tableXlsxCmd.AddCommand(tableXlsxXlsx2JSONCmd)

	tableXlsxXlsx2JSONCmd.Flags().StringVarP(&tableXlsxXlsx2JSONFilepath, "filepath", "p", "", "xlsx filepath")
	tableXlsxXlsx2JSONCmd.Flags().StringVarP(&tableXlsxXlsx2JSONOutput, "output-dir", "o", "", "output dir")
	tableXlsxXlsx2JSONCmd.Flags().StringVarP(&tableXlsxXlsx2JSONExportMode, "export-mode", "m", "sc", "export only the fields contained in the parameters(sc/s/c)")

	checkError(tableXlsxXlsx2JSONCmd.MarkFlagRequired("filepath"))
	checkError(tableXlsxXlsx2JSONCmd.MarkFlagRequired("output-dir"))
}

//goland:noinspection t
func onTableXlsxXlsx2JSON(xlsxFilepath string, output string) error {
	xlsxFile, err := xlsx.OpenFile(xlsxFilepath)
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

	// 整理配置
	var tables = make(map[string]table.Table)
	for _, sheet := range xlsxFile.Sheets {
		tab := xlsxsheet.NewTable(sheet, mode)
		if tab.IsIgnore() {
			fmt.Println("Ignore sheet: " + sheet.Name)
			continue
		}
		tables[sheet.Name] = tab
	}

	if len(tables) == 0 {
		return fmt.Errorf("no configuration found")
	}

	exist := fileproc.CheckPathExist(output)
	dir := (!exist && filepath.Ext(output) == charproc.None) || fileproc.CheckIsDir(output)
	if dir {
		checkError(os.MkdirAll(output, os.ModePerm))
	} else {
		panic("output path is not a directory")
	}

	configs := table.GenerateConfigs(collection.ConvertMapValuesToSlice(tables), fieldparser.New(), type2go.New("none"), lua2jsonparser.New())
	for name, data := range configs.GenerateData() {
		bytes, err := json.MarshalIndent(data, "", "  ")
		checkError(err)
		checkError(fileproc.WriteToFile(filepath.Join(output, name+".json"), bytes))
		fmt.Println("Convert xlsx sheet to json data file to: " + filepath.Join(output, name+".json"))
	}
	return nil
}
