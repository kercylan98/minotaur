package astgo

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"strings"
)

func newFunction(astFunc *ast.FuncDecl) *Function {
	f := &Function{
		decl:     astFunc,
		Name:     astFunc.Name.String(),
		Comments: newComment(astFunc.Doc),
	}
	f.IsTest = strings.HasPrefix(f.Name, "Test")
	f.IsBenchmark = strings.HasPrefix(f.Name, "Benchmark")
	f.IsExample = strings.HasPrefix(f.Name, "Example")
	f.Test = f.IsTest || f.IsBenchmark || f.IsExample
	f.Internal = f.Name[0] >= 97 && f.Name[0] <= 122
	if astFunc.Type.TypeParams != nil {
		for _, field := range astFunc.Type.TypeParams.List {
			f.Generic = append(f.Generic, newField(field)...)
		}
	}
	if astFunc.Type.Params != nil {
		for _, field := range astFunc.Type.Params.List {
			f.Params = append(f.Params, newField(field)...)
		}
	}
	if astFunc.Type.Results != nil {
		for _, field := range astFunc.Type.Results.List {
			f.Results = append(f.Results, newField(field)...)
		}
	}
	if astFunc.Recv != nil {
		f.Struct = newField(astFunc.Recv.List[0])[0]
	}
	return f
}

type Function struct {
	decl        *ast.FuncDecl
	Name        string   // 函数名
	Internal    bool     // 内部函数
	Generic     []*Field // 泛型定义
	Params      []*Field // 参数字段
	Results     []*Field // 返回值字段
	Comments    *Comment // 注释
	Struct      *Field   // 结构体函数对应的结构体字段
	IsExample   bool     // 是否为测试用例
	IsTest      bool     // 是否为单元测试
	IsBenchmark bool     // 是否为基准测试
	Test        bool     // 是否为测试函数
}

func (f *Function) Code() string {
	if f.Test {
		f.decl.Doc = nil
	}
	var bb bytes.Buffer
	if err := format.Node(&bb, token.NewFileSet(), f.decl); err != nil {
		panic(err)
	}
	return bb.String() + "\n"
}
