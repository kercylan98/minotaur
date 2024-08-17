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
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var (
	tableXlsxDir2JSONDir        string
	tableXlsxDir2JSONOutput     string
	tableXlsxDir2JSONExportMode string
	tableXlsxDir2JSONLua        bool
)

var tableXlsxDir2JSONCmd = &cobra.Command{
	Use:   "dir2json",
	Short: "Convert xlsx file sheets in a directory to JSON data files.",
	Long:  `Converts all xlsx file sheets within a directory into JSON data files, streamlining data export for use in various environments.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return onTableXlsxDir2JSON(tableXlsxDir2JSONDir, tableXlsxDir2JSONOutput, tableXlsxDir2JSONExportMode, tableXlsxDir2JSONLua)
	},
}

func init() {
	tableXlsxCmd.AddCommand(tableXlsxDir2JSONCmd)

	tableXlsxDir2JSONCmd.Flags().StringVarP(&tableXlsxDir2JSONDir, "dir", "p", "", "xlsx dir")
	tableXlsxDir2JSONCmd.Flags().StringVarP(&tableXlsxDir2JSONOutput, "output-dir", "o", "", "output dir")
	tableXlsxDir2JSONCmd.Flags().StringVarP(&tableXlsxDir2JSONExportMode, "export-mode", "m", "sc", "export only the fields contained in the parameters(sc/s/c)")
	tableXlsxDir2JSONCmd.Flags().BoolVarP(&tableXlsxDir2JSONLua, "lua", "l", false, "the data is described as lua")

	checkError(tableXlsxDir2JSONCmd.MarkFlagRequired("dir"))
	checkError(tableXlsxDir2JSONCmd.MarkFlagRequired("output-dir"))
}

//goland:noinspection t
func onTableXlsxDir2JSON(dir, output, exportMode string, lua bool) error {
	var xlsxFiles []*xlsx.File
	checkError(filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(d.Name()) != ".xlsx" || strings.HasPrefix(d.Name(), "~$") {
			return nil
		}
		xlsxFilepath := filepath.Join(dir, d.Name())
		xlsxFile, err := xlsx.OpenFile(xlsxFilepath)
		if err != nil {
			return err
		}
		xlsxFiles = append(xlsxFiles, xlsxFile)
		return nil
	}))

	var mode xlsxsheet.ExportMode
	switch strings.ToLower(exportMode) {
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
	for _, file := range xlsxFiles {
		for _, sheet := range file.Sheets {
			tab := xlsxsheet.NewTable(sheet, mode, lua)
			if tab.IsIgnore() {
				fmt.Println("Ignore sheet: " + sheet.Name)
				continue
			}
			tables[sheet.Name] = tab
		}
	}

	if len(tables) == 0 {
		return fmt.Errorf("no configuration found")
	}

	exist := fileproc.CheckPathExist(output)
	isDir := (!exist && filepath.Ext(output) == charproc.None) || fileproc.CheckIsDir(output)
	if isDir {
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
