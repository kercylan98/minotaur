package typemarshalers

import (
	"bytes"
	"fmt"
	"github.com/kercylan98/minotaur/engine/datasheet"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"go/format"
)

func NewGo(configurator ...GoConfigurator) *Go {
	var config = newGoConfiguration()
	for _, conf := range configurator {
		conf.Configure(config)
	}
	return &Go{
		config: config,
	}
}

type Go struct {
	config   *GoConfiguration
	builder  *bytes.Buffer
	currName string
}

func (g *Go) Marshal(t datasheet.Type) (string, error) {
	if g.builder == nil {
		g.builder = new(bytes.Buffer)
	}
	defer g.builder.Reset()
	if err := g.marshal(t); err != nil {
		return "", err
	}
	code := g.builder.Bytes()
	if g.config.typePrefix {
		code = append([]byte("type "+t.(*datasheet.Struct).StructName), code...)
	}

	v, err := format.Source(code)
	if err != nil {
		return string(code), err
	}

	return string(v), nil
}

func (g *Go) marshal(t datasheet.Type) error {
	switch value := t.(type) {
	case *datasheet.BasicType:
		g.builder.WriteString(value.TypeName)
	case *datasheet.Slice:
		g.builder.WriteString("[]")
		if err := g.marshal(value.ValueType); err != nil {
			return err
		}
	case *datasheet.Array:
		g.builder.WriteString("[")
		g.builder.WriteString(convert.IntToString(value.Length))
		g.builder.WriteString("]")
		if err := g.marshal(value.ValueType); err != nil {
			return err
		}
	case *datasheet.Map:
		g.builder.WriteString("map[")
		if err := g.marshal(value.KeyType); err != nil {
			return err
		}
		g.builder.WriteString("]")
		if err := g.marshal(value.ValueType); err != nil {
			return err
		}
	case *datasheet.Struct:
		if !value.Anonymity {
			g.builder.WriteString(value.StructName)
		} else {
			g.builder.WriteString("struct {")
			for _, field := range value.Fields {
				g.currName = field.FieldName
				g.builder.WriteString(field.FieldName)
				g.builder.WriteByte(' ')
				switch fieldType := field.FieldType.(type) {
				case *datasheet.Struct:
					if !fieldType.Anonymity {
						g.builder.WriteString(fieldType.StructName)
					} else {
						if err := g.marshal(fieldType); err != nil {
							return err
						}
					}
				default:
					if err := g.marshal(fieldType); err != nil {
						return err
					}
				}
				g.builder.WriteByte(';')
			}
			g.builder.WriteByte('}')
		}

	default:
		return fmt.Errorf("unknown type, %v", t)
	}

	return nil
}
