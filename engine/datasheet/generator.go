package datasheet

type CodeGenerator interface {
	Generate(set *Set) error
}

type DataGenerator interface {
	Generate(data map[string]any) error
}

func NewGenerator(datasheetFiles ...string) (*Generator, error) {
	set, err := LoadDatasheets(datasheetFiles...)
	if err != nil {
		return nil, err
	}
	return &Generator{
		set: set,
	}, nil
}

type Generator struct {
	set *Set
}

func (g *Generator) GenerateCode(codeGenerator CodeGenerator) error {
	return codeGenerator.Generate(g.set)
}

func (g *Generator) GenerateData(dataGenerator DataGenerator) error {
	data, err := g.set.LoadData()
	if err != nil {
		return err
	}
	return dataGenerator.Generate(data)
}
