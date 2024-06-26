// Code generated by minotaur. DO NOT EDIT.

package config

import (
	"sync/atomic"
)

type ConfigSign = string // 配置签名

var (
	IndexConfigSign ConfigSign = "IndexConfig"
	EasyConfigSign  ConfigSign = "EasyConfig"

	signedConfigGetters = map[ConfigSign]any{
		IndexConfigSign: GetIndexConfig,
		EasyConfigSign:  GetEasyConfig,
	}

	signedConfigSetters = map[ConfigSign]any{
		IndexConfigSign: SetIndexConfig,
		EasyConfigSign:  SetEasyConfig,
	}

	signedConfigLoaders = map[ConfigSign]any{
		IndexConfigSign: LoadIndexConfig,
		EasyConfigSign:  LoadEasyConfig,
	}
)

// GetConfigBySign 根据签名获取配置
func GetConfigBySign[C any](sign ConfigSign) C {
	return signedConfigGetters[sign].(func() C)()
}

// SetConfigBySign 根据签名设置配置
func SetConfigBySign[C any](sign ConfigSign, config C) {
	signedConfigSetters[sign].(func(C))(config)
}

// LoadConfigBySign 根据签名加载配置
func LoadConfigBySign(sign ConfigSign) {
	signedConfigLoaders[sign].(func())()
}

// GetConfigSigns 获取所有配置签名
func GetConfigSigns() []ConfigSign {
	return []ConfigSign{
		IndexConfigSign,
		EasyConfigSign,
	}
}

var (
	loadedIndexConfig atomic.Pointer[map[int]map[string]*IndexConfig]
	readyIndexConfig  atomic.Pointer[map[int]map[string]*IndexConfig]
	loadedEasyConfig  atomic.Pointer[*EasyConfig]
	readyEasyConfig   atomic.Pointer[*EasyConfig]
)

// GetIndexConfig 获取有索引, 该函数将返回已加载的配置
func GetIndexConfig() map[int]map[string]*IndexConfig {
	return *loadedIndexConfig.Load()
}

// SetIndexConfig 设置有索引, 该函数将待加载的配置进行存储，不影响已加载的配置
func SetIndexConfig(config map[int]map[string]*IndexConfig) {
	readyIndexConfig.Store(&config)
}

// LoadIndexConfig 加载有索引, 该函数将待加载的配置替换已加载的配置
func LoadIndexConfig() {
	loadedIndexConfig.Store(readyIndexConfig.Swap(nil))
}

// GetEasyConfig 获取无索引, 该函数将返回已加载的配置
func GetEasyConfig() *EasyConfig {
	return *loadedEasyConfig.Load()
}

// SetEasyConfig 设置无索引, 该函数将待加载的配置进行存储，不影响已加载的配置
func SetEasyConfig(config *EasyConfig) {
	readyEasyConfig.Store(&config)
}

// LoadEasyConfig 加载无索引, 该函数将待加载的配置替换已加载的配置
func LoadEasyConfig() {
	loadedEasyConfig.Store(readyEasyConfig.Swap(nil))
}

// IndexConfig 有索引
type IndexConfig struct {
	Id    int                 `json:"id,omitempty"`
	Count string              `json:"count,omitempty"`
	Award []string            `json:"award,omitempty"`
	Info  *IndexConfigInfo    `json:"info,omitempty"`
	Other []*IndexConfigOther `json:"other,omitempty"`
}

// EasyConfig 无索引
type EasyConfig struct {
	Id    int                `json:"id,omitempty"`
	Count string             `json:"count,omitempty"`
	Award []string           `json:"award,omitempty"`
	Info  *EasyConfigInfo    `json:"info,omitempty"`
	Other []*EasyConfigOther `json:"other,omitempty"`
}

// IndexConfigInfoInfoExp 信息
type IndexConfigInfoInfoExp struct {
	Mux   int `json:"mux,omitempty"`
	Count int `json:"count,omitempty"`
}

// IndexConfigInfoInfo 信息
type IndexConfigInfoInfo struct {
	Lv  int                     `json:"lv,omitempty"`
	Exp *IndexConfigInfoInfoExp `json:"exp,omitempty"`
}

// IndexConfigInfo 信息
type IndexConfigInfo struct {
	Id   int                  `json:"id,omitempty"`
	Name string               `json:"name,omitempty"`
	Info *IndexConfigInfoInfo `json:"info,omitempty"`
}

// IndexConfigOther 信息2
type IndexConfigOther struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// EasyConfigInfoInfoExp 信息
type EasyConfigInfoInfoExp struct {
	Mux   int `json:"mux,omitempty"`
	Count int `json:"count,omitempty"`
}

// EasyConfigInfoInfo 信息
type EasyConfigInfoInfo struct {
	Lv  int                    `json:"lv,omitempty"`
	Exp *EasyConfigInfoInfoExp `json:"exp,omitempty"`
}

// EasyConfigInfo 信息
type EasyConfigInfo struct {
	Id   int                 `json:"id,omitempty"`
	Name string              `json:"name,omitempty"`
	Info *EasyConfigInfoInfo `json:"info,omitempty"`
}

// EasyConfigOther 信息2
type EasyConfigOther struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
