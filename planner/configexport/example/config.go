// Code generated by minotaur-config-export. DO NOT EDIT.
package example

import (
	jsonIter "github.com/json-iterator/go"
	"github.com/kercylan98/minotaur/utils/log"
	"go.uber.org/zap"
	"os"
)

var json = jsonIter.ConfigCompatibleWithStandardLibrary
var (
	// IndexConfig 有索引
	IndexConfig  map[int]map[string]*IndexConfigDefine
	_IndexConfig map[int]map[string]*IndexConfigDefine
	// EasyConfig 无索引
	EasyConfig  *EasyConfigDefine
	_EasyConfig *EasyConfigDefine
)

func LoadConfig(handle func(filename string, config any) error) {
	var err error
	_IndexConfig = make(map[int]map[string]*IndexConfigDefine)
	if err = handle("IndexConfig.json", &_IndexConfig); err != nil {
		log.Error("Config", zap.String("Name", "IndexConfig"), zap.Bool("Invalid", true), zap.Error(err))
	}

	_EasyConfig = new(EasyConfigDefine)
	if err = handle("EasyConfig.json", _EasyConfig); err != nil {
		log.Error("Config", zap.String("Name", "EasyConfig"), zap.Bool("Invalid", true), zap.Error(err))
	}

}

func Refresh() {
	IndexConfig = _IndexConfig
	EasyConfig = _EasyConfig
}

func DefaultLoad(filepath string) {
	LoadConfig(func(filename string, config any) error {
		bytes, err := os.ReadFile(filepath)
		if err != nil {
			return err
		}

		return json.Unmarshal(bytes, &config)
	})
}
