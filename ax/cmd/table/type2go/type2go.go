package type2go

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/kercylan98/minotaur/ax/cmd/table"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"go/format"
	"strings"
	"text/template"
)

//go:embed template.tmpl
var typedTemplate string

func New(packageName string) table.CodeParser {
	return &parser{PackageName: packageName}
}

type parser struct {
	PackageName string
	Structs     []*configStruct
	Vars        []*configVar
}

func (p *parser) init(configs []*table.Config) {
	for _, c := range configs {
		// var
		typ := "*" + charproc.FirstUpper(c.GetName())
		for _, v := range c.GetIndexes() {
			switch v := v.(type) {
			case *table.StructType:
				panic("StructType can not be index field")
			case *table.ArrayType:
				panic("ArrayType can not be index field")
			case *table.BasicType:
				typ = fmt.Sprintf("map[%s]", v.Type) + typ
			default:
				panic("unknown type")
			}
		}
		p.Vars = append(p.Vars, &configVar{
			Name:    c.GetName(),
			Desc:    c.GetDesc(),
			HasDesc: c.GetDesc() != "",
			Type:    typ,
			IsMake:  len(c.GetIndexes()) > 0,
		})

		// struct
		for _, t := range c.GetTypes() {
			structure := &configStruct{
				ConfigName: c.GetName(),
			}
			switch v := t.(type) {
			case *table.StructType:
				structure.Name = v.Name
				if v.Name == c.GetName() {
					structure.Desc = c.GetDesc()
					structure.HasDesc = structure.Desc != ""
				}
				structure.Fields = p.parseStruct(c, v)
			case *table.ArrayType:
				structure.Name = charproc.FirstUpper(p.parseArray(v))
			case *table.BasicType:
				structure.Name = charproc.FirstUpper(v.Type)
			default:
				panic("unknown type")
			}
			p.Structs = append(p.Structs, structure)
		}
	}
}

func (p *parser) Parse(configs []*table.Config) string {
	p.init(configs)

	buf := new(bytes.Buffer)
	tmpl, err := template.New("config").Parse(strings.TrimSpace(typedTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, p); err != nil {
		panic(err)
	}

	v, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	return string(v)
}

func (p *parser) parseStruct(config *table.Config, structType *table.StructType) []*configStructField {
	var fields = make([]*configStructField, 0, len(structType.Fields))
	for _, field := range structType.Fields {
		name, typ := p.parseField(field)
		var desc string
		if structType.Name == config.GetName() {
			desc = config.GetFieldDesc(name)
		}
		fields = append(fields, &configStructField{
			Name:    charproc.FirstUpper(name),
			Type:    typ,
			Desc:    desc,
			HasDesc: desc != "",
		})
	}
	return fields
}

func (p *parser) parseField(field *table.StructField) (name, typ string) {
	switch v := field.Type.(type) {
	case *table.StructType:
		return field.Name, "*" + v.Name
	case *table.ArrayType:
		return field.Name, fmt.Sprintf("[]%s", p.parseArray(v))
	case *table.BasicType:
		return field.Name, v.Type
	default:
		panic("unknown type")
	}
}

func (p *parser) parseArray(v *table.ArrayType) string {
	switch v := v.ElementType.(type) {
	case *table.StructType:
		return "*" + v.Name
	case *table.ArrayType:
		return p.parseArray(v)
	case *table.BasicType:
		return v.Type
	default:
		panic("unknown type")
	}
}
