package cmd

import (
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
)

var (
	tableXlsxXlsx2GoFilepath string
	tableXlsxXlsx2GoOutput   string
	tableXlsxXlsx2GoPackage  string
)

var tableXlsxXlsx2GoCmd = &cobra.Command{
	Use:   "xlsx2go",
	Short: "Convert xlsx file sheets to Go configuration code.",
	Long:  `Converts all sheets from a specified xlsx file into Go configuration code, streamlining the integration of configuration data into Go projects.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return onTableXlsxXlsx2Go(tableXlsxXlsx2GoFilepath, tableXlsxXlsx2GoOutput, tableXlsxXlsx2GoPackage)
	},
}

func init() {
	tableXlsxCmd.AddCommand(tableXlsxXlsx2GoCmd)

	tableXlsxXlsx2GoCmd.Flags().StringVarP(&tableXlsxXlsx2GoFilepath, "filepath", "p", "", "xlsx filepath")
	tableXlsxXlsx2GoCmd.Flags().StringVarP(&tableXlsxXlsx2GoOutput, "output", "o", "", "output")
	tableXlsxXlsx2GoCmd.Flags().StringVarP(&tableXlsxXlsx2GoPackage, "package", "n", "", "package name")

	checkError(tableXlsxXlsx2GoCmd.MarkFlagRequired("filepath"))
	checkError(tableXlsxXlsx2GoCmd.MarkFlagRequired("output"))
	checkError(tableXlsxXlsx2GoCmd.MarkFlagRequired("package"))
}

//goland:noinspection t
func onTableXlsxXlsx2Go(xlsxFilepath string, output string, goPackage string) error {
	xlsxFile, err := xlsx.OpenFile(xlsxFilepath)
	if err != nil {
		return err
	}

	// 整理配置
	var tables = make(map[string]table.Table)
	for _, sheet := range xlsxFile.Sheets {
		tab := xlsxsheet.NewTable(sheet, 0, false)
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
		output = filepath.Join(output, filepath.Base(xlsxFilepath)+".go")
	} else if filepath.Ext(output) != ".go" {
		output += ".go"
	}

	configs := table.GenerateConfigs(collection.ConvertMapValuesToSlice(tables), fieldparser.New(), type2go.New(goPackage), lua2jsonparser.New())
	code := configs.GenerateCode()

	if err = fileproc.WriteToFile(output, code); err != nil {
		return err
	}

	fmt.Println("Convert xlsx sheets to go configuration code to: " + output)

	return nil
}
