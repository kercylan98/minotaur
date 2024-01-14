package astgo

import (
	"fmt"
	"go/ast"
	"strings"
)

func newType(expr ast.Expr) *Type {
	var typ = &Type{
		expr: expr,
		Name: newName(expr),
	}
	var str strings.Builder
	switch e := typ.expr.(type) {
	case *ast.KeyValueExpr:
		k, v := newType(e.Key), newType(e.Value)
		str.WriteString(fmt.Sprintf("%s : %s", k.Sign, v.Sign))
	case *ast.ArrayType:
		isSlice := e.Len == nil
		str.WriteByte('[')
		if !isSlice {
			length := newType(e.Len)
			str.WriteString(length.Sign)
		}
		str.WriteByte(']')
		element := newType(e.Elt)
		str.WriteString(element.Sign)
	case *ast.StructType:
		str.WriteString("struct {")
		if e.Fields != nil && len(e.Fields.List) > 0 {
			str.WriteString("\n")
			for _, field := range e.Fields.List {
				f := newField(field)[0]
				str.WriteString(fmt.Sprintf("%s %s\n", f.Name, f.Type.Sign))
			}
		}
		str.WriteString("}")
	case *ast.FuncType:
		var handler = func(fls *ast.FieldList) string {
			var s string
			if fls != nil {
				var brackets bool
				var params []string
				for _, field := range fls.List {
					f := newField(field)[0]
					if !f.Anonymous {
						brackets = true
					}
					params = append(params, fmt.Sprintf("%s %s", f.Name, f.Type.Sign))
				}
				s = strings.Join(params, ", ")
				if brackets {
					s = "(" + s + ")"
				}
			}
			return s
		}
		str.WriteString(strings.TrimSpace(fmt.Sprintf("func %s %s", func() string {
			f := handler(e.Params)
			if len(strings.TrimSpace(f)) == 0 {
				return "()"
			}
			if !strings.HasPrefix(f, "(") {
				return "(" + f + ")"
			}
			return f
		}(), handler(e.Results))))
	case *ast.InterfaceType:
		str.WriteString("interface {")
		if e.Methods != nil && len(e.Methods.List) > 0 {
			str.WriteString("\n")
			for _, field := range e.Methods.List {
				f := newField(field)[0]
				str.WriteString(fmt.Sprintf("%s\n", f.Type.Sign))
			}
		}
		str.WriteString("}")
	case *ast.MapType:
		k, v := newType(e.Key), newType(e.Value)
		str.WriteString(fmt.Sprintf("map[%s]%s", k.Sign, v.Sign))
	case *ast.ChanType:
		t := newType(e.Value)
		str.WriteString(fmt.Sprintf("chan %s", t.Sign))
	case *ast.Ident:
		str.WriteString(e.Name)
	case *ast.Ellipsis:
		element := newType(e.Elt)
		str.WriteString(fmt.Sprintf("...%s", element.Sign))
	case *ast.BasicLit:
		str.WriteString(e.Value)
	//case *ast.FuncLit:
	//case *ast.CompositeLit:
	//case *ast.ParenExpr:
	case *ast.SelectorExpr:
		t := newType(e.X)
		str.WriteString(fmt.Sprintf("%s.%s", t.Sign, e.Sel.Name))
	case *ast.IndexExpr:
		t := newType(e.X)
		generic := newType(e.Index)
		str.WriteString(fmt.Sprintf("%s[%s]", t.Sign, generic.Sign))
	case *ast.IndexListExpr:
		self := newType(e.X)
		var genericStr []string
		for _, index := range e.Indices {
			g := newType(index)
			genericStr = append(genericStr, g.Sign)
		}
		str.WriteString(fmt.Sprintf("%s[%s]", self.Sign, strings.Join(genericStr, ", ")))
	case *ast.SliceExpr:
	case *ast.TypeAssertExpr:
	case *ast.CallExpr:
	case *ast.StarExpr:
		typ.IsPointer = true
		t := newType(e.X)
		str.WriteString(fmt.Sprintf("*%s", t.Sign))
	case *ast.UnaryExpr:
	case *ast.BinaryExpr:
	}
	typ.Sign = str.String()
	return typ
}

type Type struct {
	expr      ast.Expr
	Sign      string // 类型签名
	IsPointer bool   // 指针类型
	Name      string // 类型名称
}
