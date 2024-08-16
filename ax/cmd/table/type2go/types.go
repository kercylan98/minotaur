package type2go

type configStruct struct {
	ConfigName string
	Name       string
	Desc       string
	HasDesc    bool
	Fields     []*configStructField
}

type configStructField struct {
	Name    string
	Type    string
	HasDesc bool
	Desc    string
}

type configVar struct {
	Name      string
	Desc      string
	HasDesc   bool
	Type      string
	IsMake    bool
	ValueType string
}
