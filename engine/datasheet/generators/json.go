package generators

import (
	"bytes"
	"fmt"
	jsonIter "github.com/json-iterator/go"
	"github.com/kercylan98/minotaur/engine/datasheet"
	"os"
	"path/filepath"
)

func JSON(outputFilepath string, groupType datasheet.GroupType) datasheet.DataGenerator {
	return &json{
		groupType:      groupType,
		outputFilepath: outputFilepath,
	}
}

type json struct {
	outputFilepath string
	groupType      datasheet.GroupType
}

func (j *json) Group() datasheet.GroupType {
	return j.groupType
}

func (j *json) Generate(data map[string]any) error {
	abs, err := filepath.Abs(j.outputFilepath)
	if err != nil {
		return err
	}

	buffer := &bytes.Buffer{}
	encoder := jsonIter.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	if err = encoder.Encode(data); err != nil {
		return err
	}

	info, err := os.Stat(abs)
	if info != nil && info.IsDir() {
		return fmt.Errorf("output filepath is a directory")
	}

	err = os.MkdirAll(filepath.Dir(j.outputFilepath), 0755)
	if err != nil {
		return err
	}

	return os.WriteFile(j.outputFilepath, buffer.Bytes(), 0644)
}

func SeparateJSON(outputDir string, groupType datasheet.GroupType) datasheet.DataGenerator {
	return &separateJSON{
		outputDir: outputDir,
		groupType: groupType,
	}
}

type separateJSON struct {
	outputDir string
	groupType datasheet.GroupType
}

func (s *separateJSON) Group() datasheet.GroupType {
	return s.groupType
}

func (s *separateJSON) Generate(data map[string]any) error {
	dirPath, err := filepath.Abs(s.outputDir)
	if err != nil {
		return err
	}

	if info, _ := os.Stat(dirPath); info != nil && !info.IsDir() {
		return fmt.Errorf("output dir is not a directory")
	}

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return err
	}

	var result = make(map[string][]byte)
	for k, v := range data {
		buffer := &bytes.Buffer{}
		encoder := jsonIter.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", "  ")
		if err = encoder.Encode(v); err != nil {
			return err
		}

		result[k] = buffer.Bytes()
	}

	for k, v := range result {
		err = os.WriteFile(filepath.Join(dirPath, k+".json"), v, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
