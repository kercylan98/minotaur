package astgo

import (
	"go/ast"
	"strings"
)

func newName(expr ast.Expr) string {
	var str strings.Builder
	switch e := expr.(type) {
	//case *ast.KeyValueExpr:
	//case *ast.ArrayType:
	//case *ast.StructType:
	//case *ast.FuncType:
	//case *ast.InterfaceType:
	//case *ast.MapType:
	//case *ast.ChanType:
	case *ast.Ident:
		str.WriteString(e.Name)
	case *ast.Ellipsis:
		str.WriteString(newName(e.Elt))
	//case *ast.BasicLit:
	//case *ast.FuncLit:
	//case *ast.CompositeLit:
	//case *ast.ParenExpr:
	case *ast.SelectorExpr:
		str.WriteString(newName(e.X))
	case *ast.IndexExpr:
		str.WriteString(newName(e.X))
	case *ast.IndexListExpr:
	case *ast.SliceExpr:
	case *ast.TypeAssertExpr:
	case *ast.CallExpr:
	case *ast.StarExpr:
		str.WriteString(newName(e.X))
	case *ast.UnaryExpr:
	case *ast.BinaryExpr:
	}
	return str.String()
}
