package storehouse

import "fmt"

const (
	String FieldType = iota + 1
	Int
	Float
)

type FieldType uint

type Field struct {
	Name     string    // 字段名
	Type     FieldType // 字段类型
	Size     *uint     // 字段长度
	Default  *string   // 默认值
	Nullable *bool     // 是否可以为空
}

func (f *Field) parseDefault() {
	switch f.Type {
	case String:
		if f.Nullable == nil {

		}
	case Int:
		f.Default = "0"
	case Float:
		f.Default = "0.0"
	}
}
func (f *Field) toSqlPart() (sql string, err error) {
	switch f.Type {
	case String:
		sql = fmt.Sprintf("`%s` VARCHAR(%d) NOT NULL DEFAULT '%s'", f.Name, f.Size, f.Default)
	case Int:
		sql = fmt.Sprintf("`%s` INT NOT NULL DEFAULT '%s'", f.Name, f.Default)
	case Float:
		sql = fmt.Sprintf("`%s` FLOAT NOT NULL DEFAULT '%s'", f.Name, f.Default)
	default:
		return "", errors.New("不支持的字段类型")
	}
}
