package genreadme

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/collection"
	"github.com/kercylan98/minotaur/utils/file"
	"github.com/kercylan98/minotaur/utils/generator/astgo"
	"github.com/kercylan98/minotaur/utils/str"
	"github.com/kercylan98/minotaur/utils/super"
	"go/format"
	"strings"
)

func New(pkgDirPath string, output string) (*Builder, error) {
	p, err := astgo.NewPackage(pkgDirPath)
	if err != nil {
		return nil, err
	}
	b := &Builder{
		p: p,
		b: new(strings.Builder),
		o: output,
	}
	return b, nil
}

type Builder struct {
	p *astgo.Package
	b *strings.Builder
	o string
}

func (b *Builder) Generate() error {
	//fmt.Println(string(super.MarshalIndentJSON(b.p, "", "  ")))
	b.genHeader()
	b.genMenus()
	b.genStructs()

	return file.WriterFile(b.o, []byte(b.b.String()))
}

func (b *Builder) genHeader() {
	b.title(1, str.FirstUpper(b.p.Name))
	b.newLine()
	b.newLine(b.p.FileComments().Clear...).newLine()
	b.newLine(fmt.Sprintf(`[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/%s)`, b.p.Name))
	b.newLine(fmt.Sprintf(`![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)`))
	b.newLine()
}

func (b *Builder) genMenus() {
	b.title(2, "目录")
	b.newLine("列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️")
	b.detailsStart("展开 / 折叠目录")

	b.quote("包级函数定义").newLine()
	b.tableCel("函数", "描述")
	for _, function := range b.p.PackageFunc() {
		if function.Test || function.Internal {
			continue
		}
		b.tableRow(
			fmt.Sprintf("[%s](#%s)", function.Name, function.Name),
			collection.FindFirstOrDefaultInSlice(function.Comments.Clear, "暂无描述..."),
		)
	}
	b.newLine().newLine()
	b.quote("结构体定义").newLine()
	b.tableCel("结构体", "描述")
	for _, f := range b.p.Files {
		if strings.HasSuffix(f.FilePath, "_test.go") {
			continue
		}
		for _, structInfo := range f.Structs {
			if structInfo.Test || structInfo.Internal {
				continue
			}
			b.tableRow(
				fmt.Sprintf("[%s](#%s)", structInfo.Name, strings.ToLower(structInfo.Name)),
				collection.FindFirstOrDefaultInSlice(structInfo.Comments.Clear, "暂无描述..."),
			)
		}
	}
	b.detailsEnd()
}

func (b *Builder) genStructs() {
	var funcHandler = func(params []*astgo.Field) string {
		var s string
		var brackets bool
		var paramsStr []string
		for _, field := range params {
			if !field.Anonymous {
				brackets = true
			}
			paramsStr = append(paramsStr, fmt.Sprintf("%s %s", field.Name, field.Type.Sign))
		}
		s = strings.Join(paramsStr, ", ")
		if brackets {
			s = "(" + s + ")"
		}
		return s
	}

	for _, function := range b.p.PackageFunc() {
		if function.Internal || function.Test {
			continue
		}
		b.title(4, strings.TrimSpace(fmt.Sprintf("func %s%s %s",
			function.Name,
			func() string {
				f := funcHandler(function.Params)
				if !strings.HasPrefix(f, "(") {
					f = "(" + f + ")"
				}
				return f
			}(),
			funcHandler(function.Results),
		)))
		b.newLine(fmt.Sprintf(`<span id="%s"></span>`, function.Name))
		b.quote()
		for _, comment := range function.Comments.Clear {
			b.quote(comment)
		}
		b.newLine("***")
	}

	for _, f := range b.p.Files {
		for _, structInfo := range f.Structs {
			if structInfo.Internal || structInfo.Test {
				continue
			}
			b.title(3, structInfo.Name)
			b.newLine(structInfo.Comments.Clear...)
			b.newLine("```go")
			structDefine := fmt.Sprintf("type %s struct {%s}",
				func() string {
					var sb strings.Builder
					sb.WriteString(structInfo.Name)
					if len(structInfo.Generic) > 0 {
						sb.WriteString("[")
						var gs []string
						for _, field := range structInfo.Generic {
							gs = append(gs, fmt.Sprintf("%s %s", field.Name, field.Type.Sign))
						}
						sb.WriteString(strings.Join(gs, ", "))
						sb.WriteString("]")
					}
					return sb.String()
				}(),
				func() string {
					var sb strings.Builder
					if len(structInfo.Fields) > 0 {
						sb.WriteString("\n")
					}
					for _, field := range structInfo.Fields {
						sb.WriteString(fmt.Sprintf("\t%s %s\n", field.Name, field.Type.Sign))
					}
					return sb.String()
				}())
			sdb, err := format.Source([]byte(structDefine))
			if err != nil {
				fmt.Println(structDefine)
				panic(err)
			}
			b.newLine(string(sdb))
			b.newLine("```")

			for _, function := range b.p.StructFunc(structInfo.Name) {
				if function.Internal || function.Test {
					continue
				}
				b.title(4, strings.TrimSpace(fmt.Sprintf("func (%s%s) %s%s %s",
					super.If(function.Struct.Type.IsPointer, "*", ""),
					structInfo.Name,
					function.Name,
					func() string {
						f := funcHandler(function.Params)
						if !strings.HasPrefix(f, "(") {
							f = "(" + f + ")"
						}
						return f
					}(),
					funcHandler(function.Results),
				)))
				//b.newLine(fmt.Sprintf(`<span id="%s"></span>`, function.Name))
				b.quote()
				for _, comment := range function.Comments.Clear {
					b.quote(comment)
				}
				b.newLine("***")
			}
		}
	}
}

func (b *Builder) newLine(text ...string) *Builder {
	if len(text) == 0 {
		b.b.WriteString("\n")
		return b
	}
	for _, s := range text {
		b.b.WriteString(s + "\n")
	}
	return b
}

func (b *Builder) title(lv int, str string) *Builder {
	var l string
	for i := 0; i < lv; i++ {
		l += "#"
	}
	b.b.WriteString(l + " " + str)
	b.b.WriteString("\n")
	return b
}

func (b *Builder) quote(str ...string) *Builder {
	for _, s := range str {
		b.b.WriteString("> " + s)
		b.newLine()
	}
	return b
}

func (b *Builder) tableCel(str ...string) *Builder {
	var a string
	var c string
	for _, s := range str {
		a += "|" + s
		c += "|:--"
	}
	b.newLine(a, c)
	return b
}

func (b *Builder) tableRow(str ...string) *Builder {
	var c string
	for _, s := range str {
		c += "|" + s
	}
	return b.newLine(c)
}

func (b *Builder) detailsStart(title string) *Builder {
	return b.newLine("<details>", "<summary>"+title+"</summary").newLine().newLine()
}

func (b *Builder) detailsEnd() *Builder {
	return b.newLine().newLine("</details>").newLine().newLine()
}
