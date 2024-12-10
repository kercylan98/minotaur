package generators

type golangType struct {
	Name        string
	Description string
	Type        string // 与 Fields 互斥
	Fields      []*golangStructField
}

func (g *golangType) IsStruct() bool {
	return len(g.Fields) > 0
}

func (g *golangType) HasDescription() bool {
	return g.Description != ""
}

type golangStructField struct {
	Name        string
	Type        string
	Description string
}

func (g *golangStructField) HasDescription() bool {
	return g.Description != ""
}

type golangVar struct {
	Name        string
	Type        string
	HasIndex    bool
	Description string
}

func (g *golangVar) HasDescription() bool {
	return g.Description != ""
}
