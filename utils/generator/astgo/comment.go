package astgo

import (
	"go/ast"
	"strings"
)

func newComment(cg *ast.CommentGroup) *Comment {
	c := &Comment{}
	if cg == nil {
		return c
	}
	for i, comment := range cg.List {
		c.Comments = append(c.Comments, comment.Text)
		cc := strings.TrimPrefix(strings.Replace(comment.Text, "// ", "//", 1), "//")
		if i == 0 {
			s := strings.SplitN(cc, " ", 2)
			if len(s) == 2 {
				cc = s[1]
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
