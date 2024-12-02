package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strconv"
)

func main() {
	fp := "./engine/prc/v1/process_id.pb.go"

	fd, err := os.ReadFile(fp)
	if err != nil {
		panic(err)
	}

	fSet := token.NewFileSet()

	file, err := parser.ParseFile(fSet, fp, fd, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(file.Decls); i++ {
		d := file.Decls[i]

		switch d.(type) {
		case *ast.FuncDecl:
		case *ast.GenDecl:
			dd := d.(*ast.GenDecl)
			if dd.Tok == token.IMPORT {
				iSpec := &ast.ImportSpec{Path: &ast.BasicLit{Value: strconv.Quote("sync/atomic")}}
				dd.Specs = append(dd.Specs, iSpec)
			}
		}
	}

	found := false
	ast.Inspect(file, func(n ast.Node) bool {
		if ts, ok := n.(*ast.TypeSpec); ok {
			switch v := ts.Type.(type) {
			case *ast.StructType:
				processCacheField := &ast.Field{
					Names: []*ast.Ident{ast.NewIdent("Cache")},
					Type:  ast.NewIdent("atomic.Pointer[any]"),
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
		if err != nil {
			panic(err)
		}
	}

	// 将修改后的 AST 写回文件
	var buf bytes.Buffer
	if err = format.Node(&buf, fSet, file); err != nil {
		panic(err)
	}

	err = os.WriteFile(fp, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("prc.ProcessId added Process field successfully!")
}
