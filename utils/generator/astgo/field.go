package astgo

import (
	"fmt"
	"go/ast"
)

func newField(field *ast.Field) []*Field {
	if len(field.Names) == 0 {
		return []*Field{{
			Anonymous: true,
			Type:      newType(field.Type),
			Comments:  newComment(field.Comment),
		}}
	} else {
		var fs []*Field
		for _, name := range field.Names {
			if name.String() == "Format" {
				fmt.Println()
			}
			fs = append(fs, &Field{
				Anonymous: false,
				Name:      name.String(),
				Type:      newType(field.Type),
				Comments:  newComment(field.Comment),
			})
		}
		return fs
	}
}

type Field struct {
	Anonymous bool     // 匿名字段
	Name      string   // 字段名称
	Type      *Type    // 字段类型
	Comments  *Comment // 注释
}
