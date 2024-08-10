package storages

import (
	"github.com/jmoiron/sqlx"
	"github.com/kercylan98/minotaur/engine/storehouse"
	"reflect"
	"strings"
)

var _ storehouse.Storage = (*MySQL)(nil)

func NewMySQL(username, password, host, database string) (*MySQL, error) {
	db, err := sqlx.Connect("mysql", username+":"+password+"@tcp("+host+")/"+database)
	if err != nil {
		return nil, err
	}
	return &MySQL{db: db}, nil
}

type MySQL struct {
	db *sqlx.DB `m:"filed_name"`
}

// 创建一张数据表，并
func (m *MySQL) CreateTable(name string, structure any) error {
	var builder strings.Builder
	vof := reflect.ValueOf(structure)
	tof := reflect.Indirect(vof)

	for i := 0; i < tof.NumField(); i++ {
		field := tof.Type().Field(i)
		tag := field.Tag.Get("m")
		if tag == "" {
			continue
		}
		switch field.Type.Kind() {
		case reflect.String:
			builder.WriteString(tag + " varchar(255)")
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			builder.WriteString(tag + " bigint")
		case reflect.Float32, reflect.Float64:
			builder.WriteString(tag + " double")
		case reflect.Bool:
			builder.WriteString(tag + " tinyint(1)")
		default:
			return storehouse.ErrUnsupportedType
		}
	}

	return err
}
