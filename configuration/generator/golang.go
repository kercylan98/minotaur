package generator

import (
	"github.com/kercylan98/minotaur/configuration"
	"github.com/kercylan98/minotaur/configuration/raw"
	"go/format"
	"os"
	"strings"
)

// NewGolangSingleFile 创建一个Golang代码生成器
func NewGolangSingleFile(packageName string, writePath string) *GolangSingleFile {
	return &GolangSingleFile{
		packageName: packageName,
		writePath:   writePath,
		builder:     new(configuration.Builder),
		basicTypes: map[string]string{
			"int":     "int",
			"int8":    "int8",
			"int16":   "int16",
			"int32":   "int32",
			"int64":   "int64",
			"uint":    "uint",
			"uint8":   "uint8",
			"uint16":  "uint16",
			"uint32":  "uint32",
			"uint64":  "uint64",
			"float32": "float32",
			"float64": "float64",
			"string":  "string",
			"bool":    "bool",
		},
	}
}

// GolangSingleFile Golang 单文件代码生成器，所有生成的代码均被写入到一个文件中
type GolangSingleFile struct {
	builder    *configuration.Builder
	basicTypes map[string]string
	table      raw.Table

	packageName string
	writePath   string
}

func (g *GolangSingleFile) Generate(table raw.Table) (err error) {
	g.table = table

	g.generatePacket()
	g.generateImports()
	g.generateSignature()
	g.generateGlobalVars()
	g.generateBasicFunc()
	g.generateConfigStruct()
	g.generateFieldStruct()

	if err = g.fmt(); err != nil {
		return err
	}

	var writer = os.Stdout
	if g.writePath != "" {
		writer, err = os.OpenFile(g.writePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}
	}

	_, err = writer.Write([]byte(g.builder.String()))
	return err
}

func (g *GolangSingleFile) generatePacket() {
	g.builder.WriteString("// Code generated by minotaur. DO NOT EDIT.\n\n")
	g.builder.Fprintf("package %s\n\n", g.packageName)
}

func (g *GolangSingleFile) generateImports() {
	var imports = []string{
		"sync/atomic",
	}
	g.builder.WriteString("import (\n")
	for _, s := range imports {
		g.builder.Fprintf("\"%s\"\n", s)
	}
	g.builder.WriteString(")\n")
}

func (g *GolangSingleFile) generateSignature() {
	g.builder.WriteString("type ConfigSign = string // 配置签名\n\n")

	// 生成签名和配置的映射
	g.builder.WriteString("var (\n")
	for configName, config := range g.table.GetConfigs() {
		g.builder.Fprintf("%sSign ConfigSign = \"%s\"\n", g.formatName(config.GetName()), configName)
	}
	g.builder.WriteBytes('\n')

	// 生成签名和配置三个函数的映射
	g.builder.Fprintf("signedConfigGetters = map[ConfigSign]any {\n")
	for configName := range g.table.GetConfigs() {
		g.builder.Fprintf("%sSign: GetProcess%s,\n", g.formatName(configName), g.formatName(configName))
	}
	g.builder.WriteString("}\n\n")
	g.builder.Fprintf("signedConfigSetters = map[ConfigSign]any {\n")
	for configName := range g.table.GetConfigs() {
		g.builder.Fprintf("%sSign: Set%s,\n", g.formatName(configName), g.formatName(configName))
	}
	g.builder.WriteString("}\n\n")
	g.builder.Fprintf("signedConfigLoaders = map[ConfigSign]any {\n")
	for configName := range g.table.GetConfigs() {
		g.builder.Fprintf("%sSign: Load%s,\n", g.formatName(configName), g.formatName(configName))
	}
	g.builder.WriteString("}\n\n")
	g.builder.WriteString(")\n\n")

	// 根据签名获取三个函数
	g.builder.Fprintf("// GetConfigBySign 根据签名获取配置\n")
	g.builder.WriteString("func GetConfigBySign[C any](sign ConfigSign) C {\n")
	g.builder.WriteString("return signedConfigGetters[sign].(func() C)()\n")
	g.builder.WriteString("}\n\n")

	g.builder.Fprintf("// SetConfigBySign 根据签名设置配置\n")
	g.builder.WriteString("func SetConfigBySign[C any](sign ConfigSign, config C) {\n")
	g.builder.WriteString("signedConfigSetters[sign].(func(C))(config)\n")
	g.builder.WriteString("}\n\n")

	g.builder.Fprintf("// LoadConfigBySign 根据签名加载配置\n")
	g.builder.WriteString("func LoadConfigBySign(sign ConfigSign) {\n")
	g.builder.WriteString("signedConfigLoaders[sign].(func())()\n")
	g.builder.WriteString("}\n\n")

	g.builder.Fprintf("// GetConfigSigns 获取所有配置签名\n")
	g.builder.WriteString("func GetConfigSigns() []ConfigSign {\n")
	g.builder.WriteString("return []ConfigSign{\n")
	for configName := range g.table.GetConfigs() {
		g.builder.Fprintf("%sSign,\n", g.formatName(configName))
	}
	g.builder.WriteString("}\n")
	g.builder.WriteString("}\n\n")
}

func (g *GolangSingleFile) generateGlobalVars() {
	builder := configuration.NewBuilder()
	g.builder.WriteString("var (\n")
	for configName, config := range g.table.GetConfigs() {
		for _, field := range config.GetFieldsWithSort() {
			if !field.IsKey() {
				break
			}

			builder.Fprintf("map[%s]", g.formatType(configName, field.GetType()))
		}
		builder.Fprintf("*%s", g.formatName(config.GetName()))
		variableType := builder.String()

		// 生成已加载和待加载的配置变量
		g.builder.Fprintf("loaded%s atomic.Pointer[%s]\n", g.formatName(config.GetName()), variableType)
		g.builder.Fprintf("ready%s atomic.Pointer[%s]\n", g.formatName(config.GetName()), variableType)

		builder.Reset()
	}
	g.builder.WriteString(")\n")
}

func (g *GolangSingleFile) generateBasicFunc() {
	builder := configuration.NewBuilder()
	for configName, config := range g.table.GetConfigs() {
		for _, field := range config.GetFieldsWithSort() {
			if !field.IsKey() {
				break
			}

			builder.Fprintf("map[%s]", g.formatType(configName, field.GetType()))
		}
		builder.Fprintf("*%s", g.formatName(config.GetName()))
		variableType := builder.String()

		// getter
		g.builder.Fprintf("// GetProcess%s 获取%s, 该函数将返回已加载的配置\n", g.formatName(config.GetName()), config.GetDescription())
		g.builder.Fprintf("func GetProcess%s() %s {\n", g.formatName(config.GetName()), variableType)
		g.builder.Fprintf("return *loaded%s.Load()\n", g.formatName(config.GetName()))
		g.builder.WriteString("}\n\n")

		// setter
		g.builder.Fprintf("// Set%s 设置%s, 该函数将待加载的配置进行存储，不影响已加载的配置\n", g.formatName(config.GetName()), config.GetDescription())
		g.builder.Fprintf("func Set%s(config %s) {\n", g.formatName(config.GetName()), variableType)
		g.builder.Fprintf("ready%s.Store(&config)\n", g.formatName(config.GetName()))
		g.builder.WriteString("}\n\n")

		// load
		g.builder.Fprintf("// Load%s 加载%s, 该函数将待加载的配置替换已加载的配置\n", g.formatName(config.GetName()), config.GetDescription())
		g.builder.Fprintf("func Load%s() {\n", g.formatName(config.GetName()))
		g.builder.Fprintf("loaded%s.Store(ready%s.Swap(nil))\n", g.formatName(config.GetName()), g.formatName(config.GetName()))
		g.builder.WriteString("}\n\n")

		builder.Reset()
	}
}

func (g *GolangSingleFile) generateConfigStruct() {
	for configName, config := range g.table.GetConfigs() {
		g.builder.Fprintf("// %s %s\n", g.formatName(config.GetName()), config.GetDescription())
		g.builder.Fprintf("type %s struct {\n", g.formatName(config.GetName()))
		for _, field := range config.GetFields() {
			g.builder.Fprintf("%s %s `json:\"%s,omitempty\"`\n", g.formatName(field.GetName()), g.formatType(configName, field.GetType()), field.GetName())
		}
		g.builder.Fprintf("}\n")
	}
}

func (g *GolangSingleFile) generateFieldStruct() {
	for configName, config := range g.table.GetConfigs() {
		for fieldName, structures := range config.GetFieldStructures() {
			for _, structure := range structures {
				g.builder.Fprintf("// %s%s %s\n", g.formatName(config.GetName()), g.formatName(structure.GetName()), config.GetField(fieldName).GetDescription())
				g.builder.Fprintf("type %s%s struct {\n", g.formatName(config.GetName()), g.formatName(structure.GetName()))
				for _, field := range structure.GetFields() {
					g.builder.Fprintf("%s %s `json:\"%s,omitempty\"`\n", g.formatName(field.GetName()), g.formatType(configName, field.GetType()), field.GetName())
				}
				g.builder.Fprintf("}\n")
			}
		}
	}
}

func (g *GolangSingleFile) formatName(str string) string {
	if g.basicTypes[str] != "" {
		return str
	}
	var camelStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if vv[i] == '_' {
			i++
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				camelStr += string(vv[i])
			} else {
				return str
			}
		} else {
			camelStr += string(vv[i])
		}
	}
	//  首字母大写
	if camelStr[0] >= 97 && camelStr[0] <= 122 {
		vv := []rune(camelStr)
		vv[0] -= 32
		camelStr = string(vv)
	}
	return camelStr
}

func (g *GolangSingleFile) formatType(configName, str string) string {
	if t, ok := g.basicTypes[str]; ok {
		return t
	}
	switch {
	case strings.HasPrefix(str, "[]"): // 切片
		return "[]" + g.formatType(configName, str[2:])
	case strings.HasPrefix(str, "["): // 数组
		index := strings.Index(str, "]")
		if index == -1 {
			panic("数组类型错误")
		}
		return str[:index+1] + g.formatType(configName, str[index+1:])
	case strings.HasPrefix(str, "map["): // map
		index := strings.Index(str, "]")
		if index == -1 {
			panic("map类型错误")
		}
		return "map[" + str[4:index+1] + g.formatType(configName, str[index+1:])
	}

	return "*" + g.formatName(configName) + g.formatName(str)
}

func (g *GolangSingleFile) fmt() error {
	code := g.builder.String()
	source, err := format.Source([]byte(code))
	if err != nil {
		return err
	}

	g.builder.Reset()
	g.builder.Write(source)
	return nil
}
