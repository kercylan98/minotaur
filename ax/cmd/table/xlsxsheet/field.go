package xlsxsheet

import "github.com/kercylan98/minotaur/ax/cmd/table"

type field struct {
	index   int
	name    string
	desc    string
	typ     string
	param   string
	scanner table.FieldDataScanner
}

func (f *field) GetIndex() int {
	return f.index
}

func (f *field) GetName() string {
	return f.name
}

func (f *field) GetDesc() string {
	return f.desc
}

func (f *field) GetType() string {
	return f.typ
}

func (f *field) GetParam() string {
	return f.param
}

func (f *field) GetData() table.FieldDataScanner {
	return f.scanner
}
