package astgo

import (
	"go/ast"
)

func newStruct(astGen *ast.GenDecl) *Struct {
	astTypeSpec := astGen.Specs[0].(*ast.TypeSpec)
	s := &Struct{
		Name:     astTypeSpec.Name.String(),
		Comments: newComment(astTypeSpec.Name.String(), astGen.Doc),
	}
	s.Internal = s.Name[0] >= 97 && s.Name[0] <= 122
	if astTypeSpec.TypeParams != nil {
		for _, field := range astTypeSpec.TypeParams.List {
			s.Generic = append(s.Generic, newField(field)...)
		}
	}

	switch ts := astTypeSpec.Type.(type) {
	case *ast.StructType:
		if ts.Fields != nil {
			for _, field := range ts.Fields.List {
				s.Fields = append(s.Fields, newField(field)...)
			}
		}
	case *ast.InterfaceType:
		s.Interface = true
		if ts.Methods != nil {
			for _, field := range ts.Methods.List {
				s.Fields = append(s.Fields, newField(field)...)
			}
		}
	default:
		s.Type = newType(ts)
	}
	return s
}

type Struct struct {
	Name      string   // 结构体名称
	Internal  bool     // 内部结构体
	Interface bool     // 接口定义
	Comments  *Comment // 注释
	Generic   []*Field // 泛型类型
	Fields    []*Field // 结构体字段
	Type      *Type    // 和结构体字段二选一，处理类似 type a string 的情况
	Test      bool     // 是否是测试结构体
}
