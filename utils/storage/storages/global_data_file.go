package storages

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/file"
	"path/filepath"
)

const (
	// GlobalDataFileDefaultSuffix 是 GlobalDataFileStorage 的文件默认后缀
	GlobalDataFileDefaultSuffix = "stock"
)

// NewGlobalDataFileStorage 创建一个 GlobalDataFileStorage，dir 是文件存储目录，generate 是生成数据的函数，options 是可选参数
//   - 生成的文件名为 ${name}.${suffix}，可以通过 WithGlobalDataFileStorageSuffix 修改后缀
//   - 默认使用 JSON 格式存储，可以通过 WithGlobalDataFileStorageEncoder 和 WithGlobalDataFileStorageDecoder 修改编码和解码函数
//   - 内置了 JSON 编码和解码函数，可以通过 FileStorageJSONEncoder 和 FileStorageJSONDecoder 获取
func NewGlobalDataFileStorage[T any](dir string, generate func(name string) T, options ...GlobalDataFileStorageOption[T]) *GlobalDataFileStorage[T] {
	abs, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}
	storage := &GlobalDataFileStorage[T]{
		dir:      abs,
		suffix:   GlobalDataFileDefaultSuffix,
		generate: generate,
		encoder:  FileStorageJSONEncoder[T](),
		decoder:  FileStorageJSONDecoder[T](),
	}
	for _, option := range options {
		option(storage)
	}
	return storage
}

// GlobalDataFileStorage 用于存储全局数据的文件存储器
type GlobalDataFileStorage[T any] struct {
	dir      string
	suffix   string
	generate func(name string) T
	encoder  FileStorageEncoder[T]
	decoder  FileStorageDecoder[T]
}

// Load 从文件中加载数据，如果文件不存在则使用 generate 生成数据
func (slf *GlobalDataFileStorage[T]) Load(name string) T {
	bytes, err := file.ReadOnce(filepath.Join(slf.dir, fmt.Sprintf("%s.%s", name, slf.suffix)))
	if err != nil {
		return slf.generate(name)
	}
	var data = slf.generate(name)
	_ = slf.decoder(bytes, data)
	return data
}

// Save 将数据保存到文件中
func (slf *GlobalDataFileStorage[T]) Save(name string, data T) error {
	bytes, err := slf.encoder(data)
	if err != nil {
		return err
	}
	return file.WriterFile(filepath.Join(slf.dir, fmt.Sprintf("%s.%s", name, slf.suffix)), bytes)
}
