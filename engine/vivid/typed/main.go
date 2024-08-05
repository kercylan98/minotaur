package main

import (
	"github.com/kercylan98/minotaur/engine/vivid/typed/internal/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generator.GenerateFile(gen, f, generator.GenerateModeVivid)
		}

		return nil
	})
}
