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
	"sync"
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
	b.newLine(fmt.Sprintf(`[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)`))
	b.newLine(fmt.Sprintf(`![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)`))
	if len(b.p.FileComments().Clear) != 0 {
		b.newLine().newLine(b.p.FileComments().Clear...).newLine()
	} else {
		b.newLine().newLine("暂无介绍...").newLine()
	}
	b.newLine()
}

func (b *Builder) genMenus() {
	var genTitleOnce sync.Once
	var genTitle = func() {
		genTitleOnce.Do(func() {
			b.title(2, "目录导航")
			b.newLine("列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️")
			b.detailsStart("展开 / 折叠目录导航")
		})
	}

	packageFunction := b.p.PackageFunc()
	var structList []*astgo.Struct
	for _, f := range b.p.Files {
		if strings.HasSuffix(f.FilePath, "_test.go") {
			continue
		}
		for _, structInfo := range f.Structs {
			if structInfo.Test || structInfo.Internal {
				continue
			}
			structList = append(structList, structInfo)
		}
	}

	if len(packageFunction) > 0 {
		var pfGenOnce sync.Once
		var pfGen = func() {
			pfGenOnce.Do(func() {
				genTitle()
				b.quote("包级函数定义").newLine()
				b.tableCel("函数名称", "描述")
			})
		}
		for _, function := range packageFunction {
			if function.Test || function.Internal {
				continue
			}
			pfGen()
			b.tableRow(
				fmt.Sprintf("[%s](#%s)", function.Name, function.Name),
				collection.FindFirstOrDefaultInSlice(function.Comments.Clear, "暂无描述..."),
			)
		}
		b.newLine().newLine()
	}

	if len(structList) > 0 {
		var structGenOnce sync.Once
		var structGen = func() {
			structGenOnce.Do(func() {
				genTitle()
				b.quote("类型定义").newLine()
				b.tableCel("类型", "名称", "描述")
			})
		}
		for _, structInfo := range structList {
			if structInfo.Test || structInfo.Internal {
				continue
			}
			structGen()
			b.tableRow(
				super.If(structInfo.Interface, "`INTERFACE`", "`STRUCT`"),
				fmt.Sprintf("[%s](#%s)", structInfo.Name, strings.ToLower(structInfo.Name)),
				collection.FindFirstOrDefaultInSlice(structInfo.Comments.Clear, "暂无描述..."),
			)
		}
	}
	b.detailsEnd()
	b.newLine("***")
}

func (b *Builder) genStructs() {
	var titleOnce sync.Once
	var titleBuild = func() {
		titleOnce.Do(func() {
			b.title(2, "详情信息")
		})
	}

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
		titleBuild()
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
		b.newLine()
		if example := b.p.GetExampleTest(function); example != nil {
			b.newLine("示例代码：", "```go\n", example.Code(), "```\n")
		}
		if unitTest := b.p.GetUnitTest(function); unitTest != nil {
			b.detailsStart("查看 / 收起单元测试")
			b.newLine("```go\n", unitTest.Code(), "```\n")
			b.detailsEnd()
		}
		if benchmarkTest := b.p.GetBenchmarkTest(function); benchmarkTest != nil {
			b.detailsStart("查看 / 收起基准测试")
			b.newLine("```go\n", benchmarkTest.Code(), "```\n")
			b.detailsEnd()
		}
		b.newLine("***")
	}

	for _, f := range b.p.Files {
		for _, structInfo := range f.Structs {
			if structInfo.Internal || structInfo.Test {
				continue
			}
			titleBuild()
			b.title(3, fmt.Sprintf("%s `%s`", structInfo.Name, super.If(structInfo.Interface, "INTERFACE", "STRUCT")))
			b.newLine(structInfo.Comments.Clear...)
			b.newLine("```go")
			structDefine := fmt.Sprintf("type %s %s",
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
					if structInfo.Type == nil {
						var head = "struct"
						if structInfo.Interface {
							head = "interface"
						}
						sb.WriteString(head + " {")
						if len(structInfo.Fields) > 0 {
							sb.WriteString("\n")
						}
						if structInfo.Interface {
							for _, field := range structInfo.Fields {
								sb.WriteString(fmt.Sprintf("\t%s %s\n", field.Name, strings.TrimPrefix(field.Type.Sign, "func")))
							}
						} else {
							for _, field := range structInfo.Fields {
								sb.WriteString(fmt.Sprintf("\t%s %s\n", field.Name, field.Type.Sign))
							}
						}
						sb.WriteString("}")
					} else {
						sb.WriteString(structInfo.Type.Sign)
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
				b.quote()
				for _, comment := range function.Comments.Clear {
					b.quote(comment)
				}
				if function.Name == "Write" {
					fmt.Println()
				}
				if example := b.p.GetExampleTest(function); example != nil {
					b.newLine("示例代码：", "```go\n", example.Code(), "```\n")
				}
				if unitTest := b.p.GetUnitTest(function); unitTest != nil {
					b.detailsStart("查看 / 收起单元测试")
					b.newLine("```go\n", unitTest.Code(), "```\n")
					b.detailsEnd()
				}
				if benchmarkTest := b.p.GetBenchmarkTest(function); benchmarkTest != nil {
					b.detailsStart("查看 / 收起基准测试")
					b.newLine("```go\n", benchmarkTest.Code(), "```\n")
					b.detailsEnd()
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
	return b.newLine("<details>", "<summary>"+title+"</summary>").newLine().newLine()
}

func (b *Builder) detailsEnd() *Builder {
	return b.newLine().newLine("</details>").newLine().newLine()
}
