package cmd

import (
	"errors"
	"github.com/kercylan98/minotaur/planner/pce"
	"github.com/kercylan98/minotaur/planner/pce/cs"
	"github.com/kercylan98/minotaur/planner/pce/tmpls"
	"github.com/kercylan98/minotaur/utils/file"
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/str"
	"github.com/spf13/cobra"
	"github.com/tealeg/xlsx"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func init() {
	var filePath, outPath, exclude string

	exportGo := &cobra.Command{
		Use:   "go",
		Short: "Export go language configuration code | 导出 go 语言配置代码",
		RunE: func(cmd *cobra.Command, args []string) error {

			isDir, err := file.IsDir(outPath)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					isDir = filepath.Ext(outPath) == ""
				} else {
					return err
				}
			}
			if isDir {
				_ = os.MkdirAll(outPath, os.ModePerm)
				outPath = filepath.Join(outPath, "config.go")
			} else {
				_ = os.MkdirAll(filepath.Dir(outPath), os.ModePerm)
			}

			fpd, err := file.IsDir(filePath)
			if err != nil {
				return err
			}

			var xlsxFiles []string
			if fpd {
				files, err := os.ReadDir(filePath)
				if err != nil {
					return err
				}
				for _, f := range files {
					if f.IsDir() || !strings.HasSuffix(f.Name(), ".xlsx") || strings.HasPrefix(f.Name(), "~") {
						continue
					}
					xlsxFiles = append(xlsxFiles, filepath.Join(filePath, f.Name()))
				}
			} else {
				xlsxFiles = append(xlsxFiles, filePath)
			}

			var golang []*pce.TmplStruct
			var exporter = pce.NewExporter()
			loader := pce.NewLoader(pce.GetFields())

			excludes := hash.ToMapBool(str.SplitTrimSpace(exclude, ","))
			for _, xlsxFile := range xlsxFiles {
				xf, err := xlsx.OpenFile(xlsxFile)
				if err != nil {
					return err
				}

				for _, sheet := range xf.Sheets {
					cx := cs.NewXlsx(sheet, cs.XlsxExportTypeServer)
					if strings.HasPrefix(cx.GetDisplayName(), "#") || strings.HasPrefix(cx.GetConfigName(), "#") || excludes[cx.GetConfigName()] || excludes[cx.GetDisplayName()] {
						continue
					}
					golang = append(golang, loader.LoadStruct(cx))
				}
			}

			if raw, err := exporter.ExportStruct(tmpls.NewGolang(filepath.Base(filepath.Dir(outPath))), golang...); err != nil {
				return err
			} else {
				if err := file.WriterFile(outPath, raw); err != nil {
					return err
				}
			}

			_ = exec.Command("gofmt", "-w", outPath).Run()
			return nil
		},
	}

	exportGo.Flags().StringVarP(&filePath, "xlsx", "f", "", "xlsx file path or directory path | xlsx 文件路径或所在目录路径")
	exportGo.Flags().StringVarP(&outPath, "output", "o", "", "output path | 输出的 go 文件路径")
	exportGo.Flags().StringVarP(&exclude, "exclude", "e", "", "excluded configuration names or display names (comma separated) | 排除的配置名或显示名（英文逗号分隔）")
	if err := exportGo.MarkFlagRequired("xlsx"); err != nil {
		panic(err)
	}
	if err := exportGo.MarkFlagRequired("output"); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(exportGo)
}
