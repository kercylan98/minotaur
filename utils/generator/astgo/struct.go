package astgo

import (
	"go/ast"
)

func newStruct(astGen *ast.GenDecl) *Struct {
	astTypeSpec := astGen.Specs[0].(*ast.TypeSpec)
	s := &Struct{
		Name:     astTypeSpec.Name.String(),
		Comments: newComment(astGen.Doc),
	}
	s.Internal = s.Name[0] >= 97 && s.Name[0] <= 122
	if astTypeSpec.TypeParams != nil {
		for _, field := range astTypeSpec.TypeParams.List {
			s.Generic = append(s.Generic, newField(field)...)
		}
	}

	t, ok := astTypeSpec.Type.(*ast.StructType)
	if ok && t.Fields != nil {
		for _, field := range t.Fields.List {
			s.Fields = append(s.Fields, newField(field)...)
		}
	}
	return s
}

type Struct struct {
	Name     string   // 结构体名称
	Internal bool     // 内部结构体
	Comments *Comment // 注释
	Generic  []*Field // 泛型类型
	Fields   []*Field // 结构体字段
	Test     bool     // 是否是测试结构体
}
