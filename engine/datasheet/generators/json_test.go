package generators_test

import (
	"github.com/kercylan98/minotaur/engine/datasheet"
	"github.com/kercylan98/minotaur/engine/datasheet/generators"
	"testing"
)

func TestGenerateJSON(t *testing.T) {
	generator, err := datasheet.NewGenerator("../datasheet_template.xlsx")
	if err != nil {
		panic(err)
	}

	if err = generator.GenerateData(generators.JSON("./test/generate-json.json", datasheet.ServerGroup)); err != nil {
		panic(err)
	}
}

func TestSeparateJSON(t *testing.T) {
	generator, err := datasheet.NewGenerator("../datasheet_template.xlsx")
	if err != nil {
		panic(err)
	}

	if err = generator.GenerateData(generators.SeparateJSON("./test", datasheet.ServerGroup)); err != nil {
		panic(err)
	}
}
