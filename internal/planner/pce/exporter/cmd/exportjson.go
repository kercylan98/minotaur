package cmd

import (
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/planner/pce"
	"github.com/kercylan98/minotaur/planner/pce/cs"
	"github.com/kercylan98/minotaur/planner/pce/tmpls"
"github.com/kercylan98/minotaur/toolkit/collection"
"github.com/kercylan98/minotaur/utils/file"
"github.com/kercylan98/minotaur/utils/str"
"github.com/spf13/cobra"
"github.com/tealeg/xlsx"
"os"
"path/filepath"
"strings"
)

func init() {
	var filePath, outPath, exclude, exportType, prefix string

	exportJson := &cobra.Command{
		Use:   "json",
		Short: "Export json configuration data | 导出 json 配置数据",
		RunE: func(cmd *cobra.Command, args []string) error {

			isDir, err := file.IsDir(outPath)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					isDir = filepath.Ext(outPath) == ""
				} else {
					return err
				}
			}
			if !isDir {
				return errors.New("output must be a directory path")
			}
			_ = os.MkdirAll(outPath, os.ModePerm)

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

			var exporter = pce.NewExporter()
			loader := pce.NewLoader(pce.GetFields())

			excludes := collection.ConvertSliceToBoolMap(str.SplitTrimSpace(exclude, ","))
			for _, xlsxFile := range xlsxFiles {
				xf, err := xlsx.OpenFile(xlsxFile)
				if err != nil {
					return err
				}

				for _, sheet := range xf.Sheets {
					var cx *cs.Xlsx
					switch strings.TrimSpace(strings.ToLower(exportType)) {
					case "c":
						cx = cs.NewXlsx(sheet, cs.XlsxExportTypeClient)
					case "s":
						cx = cs.NewXlsx(sheet, cs.XlsxExportTypeServer)
					}
					if strings.HasPrefix(cx.GetDisplayName(), "#") || strings.HasPrefix(cx.GetConfigName(), "#") || excludes[cx.GetConfigName()] || excludes[cx.GetDisplayName()] {
						continue
					}

					if raw, err := exporter.ExportData(tmpls.NewJSON(), loader.LoadData(cx)); err != nil {
						return err
					} else {
						var jsonPath string
						if len(prefix) == 0 {
							jsonPath = filepath.Join(outPath, fmt.Sprintf("%s.json", cx.GetConfigName()))
						} else {
							jsonPath = filepath.Join(outPath, fmt.Sprintf("%s.%s.json", prefix, cx.GetConfigName()))
						}
						if err := file.WriterFile(jsonPath, raw); err != nil {
							return err
						}
					}
				}
			}

			return nil
		},
	}

	exportJson.Flags().StringVarP(&filePath, "xlsx", "f", "", "xlsx file path or directory path | xlsx 文件路径或所在目录路径")
	exportJson.Flags().StringVarP(&outPath, "output", "o", "", "directory path of the output json file | 输出的 json 文件所在目录路径")
	exportJson.Flags().StringVarP(&exportType, "type", "t", "", "export server configuration[s] or client configuration[c] | 导出服务端配置[s]还是客户端配置[c]")
	exportJson.Flags().StringVarP(&prefix, "prefix", "p", "", "export configuration file name prefix | 导出配置文件名前缀")
	exportJson.Flags().StringVarP(&exclude, "exclude", "e", "", "excluded configuration names or display names (comma separated) | 排除的配置名或显示名（英文逗号分隔）")
	if err := exportJson.MarkFlagRequired("xlsx"); err != nil {
		panic(err)
	}
	if err := exportJson.MarkFlagRequired("output"); err != nil {
		panic(err)
	}
	if err := exportJson.MarkFlagRequired("type"); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(exportJson)
}
