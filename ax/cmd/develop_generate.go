package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
)

// developGenerateCmd represents the developGen command
var developGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate necessary files for minotaur project development",
	Long:  `Generate necessary files for minotaur project development, these include the built-in ProtoBuffer and more`,
	Run: func(cmd *cobra.Command, args []string) {
		onDevelopGenerateProtobuf()
		onDevelopGenerateProcessIdInject()
	},
}

func init() {
	developCmd.AddCommand(developGenerateCmd)

}

func onDevelopGenerateProtobuf() {
	var scripts = make([]string, 0)
	cobra.CheckErr(filepath.Walk(viper.GetString(configKeyDevelopWorkdir), func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		abs, err := filepath.Abs(path)
		cobra.CheckErr(err)

		switch os := runtime.GOOS; os {
		case "windows":
			if info.Name() == "generate-protobuf.bat" {
				scripts = append(scripts, abs)
			}
		case "linux", "darwin":
			if info.Name() == "generate-protobuf.sh" {
				scripts = append(scripts, abs)
			}
		default:

		}
		return nil
	}))

	for _, batPath := range scripts {
		fmt.Println("generate:", batPath)
		c := exec.Command(batPath)
		c.Dir = filepath.Dir(batPath)
		if err := c.Run(); err != nil {
			fmt.Println(err)
		}
	}
}

func onDevelopGenerateProcessIdInject() {
	fp := filepath.Join(viper.GetString(configKeyDevelopWorkdir), "engine", "prc", "process_id.pb.go")

	// 读取文件内容
	fd, err := os.ReadFile(fp)
	cobra.CheckErr(err)

	// 创建文件集
	fSet := token.NewFileSet()

	// 解析文件
	file, err := parser.ParseFile(fSet, fp, fd, parser.ParseComments)
	cobra.CheckErr(err)

	// 导入
	// Add the imports
	for i := 0; i < len(file.Decls); i++ {
		d := file.Decls[i]

		switch d.(type) {
		case *ast.FuncDecl:
			// No action
		case *ast.GenDecl:
			dd := d.(*ast.GenDecl)

			// IMPORT Declarations
			if dd.Tok == token.IMPORT {
				// Add the new import
				iSpec := &ast.ImportSpec{Path: &ast.BasicLit{Value: strconv.Quote("sync/atomic")}}
				dd.Specs = append(dd.Specs, iSpec)
			}
		}
	}

	// 查找并修改结构体
	found := false
	ast.Inspect(file, func(n ast.Node) bool {
		if ts, ok := n.(*ast.TypeSpec); ok {
			switch v := ts.Type.(type) {
			case *ast.StructType:
				processCacheField := &ast.Field{
					Names: []*ast.Ident{ast.NewIdent("cache")},
					Type:  ast.NewIdent("atomic.Pointer[Process]"),
				}
				v.Fields.List = append(v.Fields.List, processCacheField)

				//redirectAddressField := &ast.Field{
				//	Names: []*ast.Ident{ast.NewIdent("redirect")},
				//	Type:  ast.NewIdent("atomic.Pointer[ProcessId]"),
				//}
				//v.Fields.List = append(v.Fields.List, redirectAddressField)
				found = true
				return false
			}
		}
		return true
	})

	// 查找并删除特定方法
	var newDecls []ast.Decl
	for _, decl := range file.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			if funcDecl.Name.Name == "GetPhysicalAddress" || funcDecl.Name.Name == "GetLogicalAddress" {
				continue // 跳过这个方法，不加入新的声明列表
			}
		}
		newDecls = append(newDecls, decl)
	}
	file.Decls = newDecls

	if !found {
		cobra.CheckErr(fmt.Errorf("struct prc.ProcessId not found"))
	}

	// 将修改后的 AST 写回文件
	var buf bytes.Buffer
	cobra.CheckErr(format.Node(&buf, fSet, file))

	err = os.WriteFile(fp, buf.Bytes(), 0644)
	cobra.CheckErr(err)

	fmt.Println("prc.ProcessId added Process field successfully!")
}
