package astgo

import (
	"go/ast"
	"strings"
)

func newComment(name string, cg *ast.CommentGroup) *Comment {
	c := &Comment{}
	if cg == nil {
		return c
	}
	for i, comment := range cg.List {
		c.Comments = append(c.Comments, comment.Text)
		cc := strings.TrimPrefix(strings.Replace(comment.Text, "// ", "//", 1), "//")
		if i == 0 {
			tsc := strings.TrimSpace(cc)
			if strings.HasPrefix(tsc, name) {
				s := strings.TrimSpace(strings.TrimPrefix(tsc, name))
				cc = s
			}
		}
		c.Clear = append(c.Clear, cc)
	}
	return c
}

type Comment struct {
	Comments []string
	Clear    []string
}
