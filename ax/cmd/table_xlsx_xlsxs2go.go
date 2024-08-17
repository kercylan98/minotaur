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
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var (
	tableXlsxDir2GoDir     string
	tableXlsxDir2GoOutput  string
	tableXlsxDir2GoPackage string
)

var tableXlsxDir2GoCmd = &cobra.Command{
	Use:   "dir2go",
	Short: "Convert xlsx file sheets in a directory to Go configuration code.",
	Long:  `Converts all xlsx file sheets within a directory into Go configuration code, automating the process of integrating multiple configuration files into your Go project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return onTableXlsxDir2Go(tableXlsxDir2GoDir, tableXlsxDir2GoOutput, tableXlsxDir2GoPackage)
	},
}

func init() {
	tableXlsxCmd.AddCommand(tableXlsxDir2GoCmd)

	tableXlsxDir2GoCmd.Flags().StringVarP(&tableXlsxDir2GoDir, "dir", "p", "", "xlsx dir")
	tableXlsxDir2GoCmd.Flags().StringVarP(&tableXlsxDir2GoOutput, "output", "o", "", "output")
	tableXlsxDir2GoCmd.Flags().StringVarP(&tableXlsxDir2GoPackage, "package", "n", "", "package name")

	checkError(tableXlsxDir2GoCmd.MarkFlagRequired("dir"))
	checkError(tableXlsxDir2GoCmd.MarkFlagRequired("output"))
	checkError(tableXlsxDir2GoCmd.MarkFlagRequired("package"))
}

//goland:noinspection t
func onTableXlsxDir2Go(dir string, output string, goPackage string) error {
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

	// 整理配置
	var tables = make(map[string]table.Table)
	for _, file := range xlsxFiles {
		for _, sheet := range file.Sheets {
			tab := xlsxsheet.NewTable(sheet, 0, false)
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
		output = filepath.Join(output, filepath.Base(output)+".go")
	} else if filepath.Ext(output) != ".go" {
		output += ".go"
	}

	configs := table.GenerateConfigs(collection.ConvertMapValuesToSlice(tables), fieldparser.New(), type2go.New(goPackage), lua2jsonparser.New())
	code := configs.GenerateCode()

	checkError(fileproc.WriteToFile(output, code))

	fmt.Println("Convert xlsx sheets to go configuration code to: " + output)

	return nil
}
