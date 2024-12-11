package generators

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/kercylan98/minotaur/engine/datasheet"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"go/format"
	"os"
	"strings"
	"text/template"
)

//go:embed go.tmpl
var golangTemplate string

func Go(packageName, filePath string) datasheet.CodeGenerator {
	return &golang{
		PackageName: packageName,
		filepath:    filePath,
	}
}

type golang struct {
	set      *datasheet.Set
	filepath string

	PackageName   string
	Types         []*golangType
	Datasheets    []*golangType
	DatasheetVars []*golangVar
}

func (g *golang) Generate(set *datasheet.Set) error {
	g.set = set

	g.parse(set)

	buf := new(bytes.Buffer)
	tmpl, err := template.New("config").Parse(strings.TrimSpace(golangTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, g); err != nil {
		return err
	}

	v, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(string(buf.Bytes()))
		return err
	}

	if err = os.WriteFile(g.filepath, v, 0644); err != nil {
		panic(err)
	}

	return nil
}

func (g *golang) parse(set *datasheet.Set) {
	for _, prefab := range set.Prefabs {
		gt := &golangType{
			Name:        prefab.Name,
			Description: prefab.Description,
		}
		switch value := prefab.Type.(type) {
		case *datasheet.Struct:
			for _, field := range value.Fields {
				goField := &golangStructField{
					Name: charproc.BigCamel(field.Name),
					Type: g.parseType(field.Type),
				}
				gt.Fields = append(gt.Fields, goField)
			}
		default:
			gt.Type = g.parseType(prefab.Type)
		}
		g.Types = append(g.Types, gt)
	}
	for _, datasheets := range set.Datasheets {
		for _, d := range datasheets {
			gt := &golangType{
				Name:        d.Name,
				Description: d.Description,
			}
			for _, field := range d.Fields {
				if !field.InGroup(datasheet.ServerGroup) {
					continue
				}
				goField := &golangStructField{
					Name:        charproc.BigCamel(field.Name),
					Type:        g.parseType(field.Type),
					Description: field.Description,
				}
				gt.Fields = append(gt.Fields, goField)
			}
			g.Types = append(g.Types, gt)
			g.Datasheets = append(g.Datasheets, gt)
		}
	}

	for _, datasheets := range set.Datasheets {
		for _, d := range datasheets {
			gt := &golangVar{
				Name:        charproc.BigCamel(d.Name),
				HasIndex:    d.HasIndex(),
				Description: d.Description,
			}
			buf := new(bytes.Buffer)
			for _, field := range d.Fields {
				if field.Index > 0 {
					buf.WriteString("map[")
					buf.WriteString(g.parseType(field.Type))
					buf.WriteString("]")
				} else {
					buf.WriteString("*" + charproc.BigCamel(d.Name))
					break
				}
			}
			gt.Type = buf.String()
			buf.Reset()
			g.DatasheetVars = append(g.DatasheetVars, gt)
		}
	}
}

func (g *golang) parseType(t datasheet.Type) string {
	buf := new(bytes.Buffer)
	defer buf.Reset()

	switch value := t.(type) {
	case *datasheet.Prefab:
		buf.WriteString(charproc.BigCamel(value.Name))
	case *datasheet.Struct:
		buf.WriteString("struct {")
		for i, field := range value.Fields {
			buf.WriteString(charproc.BigCamel(field.Name))
			buf.WriteString(" ")
			if field.Optional {
				buf.WriteString("*")
			}
			buf.WriteString(g.parseType(field.Type))
			if i != len(value.Fields)-1 {
				buf.WriteString(";")
			}
		}
		buf.WriteString("}")
	case *datasheet.Array:
		buf.WriteString("[")
		buf.WriteString(convert.IntToString(value.Length))
		buf.WriteString("]")
		if value.Optional {
			buf.WriteString("*")
		}
		buf.WriteString(g.parseType(value.Type))
	case *datasheet.Slice:
		buf.WriteString("[]")
		if value.Optional {
			buf.WriteString("*")
		}
		buf.WriteString(g.parseType(value.Type))
	case *datasheet.Map:
		buf.WriteString("map[")
		buf.WriteString(g.parseType(value.KeyType))
		buf.WriteString("]")
		if value.ValueOptional {
			buf.WriteString("*")
		}
		buf.WriteString(g.parseType(value.ValueType))
	case *datasheet.Basic:
		switch value.Name {
		case datasheet.BasicTypeBigInt:
			buf.WriteString("datasheet.BigInt")
		case datasheet.BasicTypeBigFloat:
			buf.WriteString("datasheet.BigFloat")
		case datasheet.BasicTypeBoolean:
			buf.WriteString("bool")
		case datasheet.BasicTypeFloat:
			buf.WriteString("float32")
		case datasheet.BasicTypeDouble:
			buf.WriteString("float64")
		case datasheet.BasicTypeShort:
			buf.WriteString("int16")
		case datasheet.BasicTypeLong:
			buf.WriteString("int64")
		case datasheet.BasicTypeTime, datasheet.BasicTypeDateTime, datasheet.BasicTypeDateOnly, datasheet.BasicTypeTimeOnly:
			buf.WriteString("time.Time")
		case datasheet.BasicTypeDuration:
			buf.WriteString("time.Duration")
		case datasheet.BasicTypeStr:
			buf.WriteString("string")
		default:
			buf.WriteString(value.Name)
		}
	}
	return buf.String()
}
