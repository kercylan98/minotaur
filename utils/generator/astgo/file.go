package astgo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func newFile(owner *Package, filePath string) (*File, error) {
	af, err := parser.ParseFile(token.NewFileSet(), filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	f := &File{
		af:       af,
		owner:    owner,
		FilePath: filePath,
		Comment:  newComment(af.Doc),
	}
	for _, decl := range af.Decls {
		switch typ := decl.(type) {
		case *ast.FuncDecl:
			f.Functions = append(f.Functions, newFunction(typ))
		case *ast.GenDecl:
			_, ok := typ.Specs[0].(*ast.TypeSpec)
			if ok {
				s := newStruct(typ)
				f.Structs = append(f.Structs, s)
				if strings.HasSuffix(filePath, "_test.go") {
					s.Test = true
				}
			}
		}
	}
	return f, nil
}

type File struct {
	af        *ast.File   // 抽象语法树文件
	owner     *Package    // 所属包
	FilePath  string      // 文件路径
	Structs   []*Struct   // 包含的结构体
	Functions []*Function // 包含的函数
	Comment   *Comment    // 文件注释
}

func (f *File) Package() string {
	return f.af.Name.String()
}
