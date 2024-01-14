# Cmd

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/cmd)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)




## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[Execute](#Execute)|将所有子命令添加到根命令并适当设置标志。这是由 main.main() 调用的。 rootCmd 只需要发生一次


***
## 详情信息
#### func Execute()
<span id="Execute"></span>
> 将所有子命令添加到根命令并适当设置标志。这是由 main.main() 调用的。 rootCmd 只需要发生一次

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestExecute(t *testing.T) {
	var filePath, outPath, exclude, exportType, prefix string
	exportType = "s"
	filePath = `.\游戏配置.xlsx`
	filePath = `../xlsx_template.xlsx`
	outPath = `.`
	isDir, err := file.IsDir(outPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			isDir = filepath.Ext(outPath) == ""
		} else {
			panic(err)
		}
	}
	if !isDir {
		panic(errors.New("output must be a directory path"))
	}
	_ = os.MkdirAll(outPath, os.ModePerm)
	fpd, err := file.IsDir(filePath)
	if err != nil {
		panic(err)
	}
	var xlsxFiles []string
	if fpd {
		files, err := os.ReadDir(filePath)
		if err != nil {
			panic(err)
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
			panic(err)
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
				panic(err)
			} else {
				var jsonPath string
				if len(prefix) == 0 {
					jsonPath = filepath.Join(outPath, fmt.Sprintf("%s.json", cx.GetConfigName()))
				} else {
					jsonPath = filepath.Join(outPath, fmt.Sprintf("%s.%s.json", prefix, cx.GetConfigName()))
				}
				if err := file.WriterFile(jsonPath, raw); err != nil {
					panic(err)
				}
			}
		}
	}
}

```


</details>


***
