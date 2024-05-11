package generator_test

import (
	"github.com/kercylan98/minotaur/configuration/generator"
	"github.com/kercylan98/minotaur/configuration/raw"
	"testing"
)

func TestNewGolangSingleFile(t *testing.T) {
	t.Run("TestNewGolangSingleFile", func(t *testing.T) {
		if generator.NewGolangSingleFile("", "") == nil {
			t.Fail()
		}
	})
}

func TestGolangSingleFile_Generate(t *testing.T) {
	t.Run("TestGolangSingleFile_Generate", func(t *testing.T) {
		g := generator.NewGolangSingleFile("test", "")
		if g == nil {
			t.Fail()
		}

		if err := g.Generate(raw.Table{}); err != nil {
			t.Fail()
		}
	})

	t.Run("TestGolangSingleFile_Generate_EmptyPackage", func(t *testing.T) {
		g := generator.NewGolangSingleFile("", "")
		if g == nil {
			t.Fail()
		}

		if err := g.Generate(raw.Table{}); err == nil {
			t.Fail()
		}
	})
}
