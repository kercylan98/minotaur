package storages

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/file"
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/storage"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	// IndexDataFileDefaultSuffix 是 IndexDataFileStorage 的文件默认后缀
	IndexDataFileDefaultSuffix = "stock"

	indexNameFormat = "%s.%v.%s"
)

// NewIndexDataFileStorage 创建索引数据文件存储器
func NewIndexDataFileStorage[I generic.Ordered, T storage.IndexDataItem[I]](dir string, generate func(name string, index I) T, generateZero func(name string) T, options ...IndexDataFileStorageOption[I, T]) *IndexDataFileStorage[I, T] {
	s := &IndexDataFileStorage[I, T]{
		dir:          dir,
		suffix:       IndexDataFileDefaultSuffix,
		generate:     generate,
		generateZero: generateZero,
		encoder:      FileStorageJSONEncoder[T](),
		decoder:      FileStorageJSONDecoder[T](),
	}
	for _, option := range options {
		option(s)
	}
	return s
}

// IndexDataFileStorage 索引数据文件存储器
type IndexDataFileStorage[I generic.Ordered, T storage.IndexDataItem[I]] struct {
	dir          string
	suffix       string
	generate     func(name string, index I) T
	generateZero func(name string) T
	encoder      FileStorageEncoder[T]
	decoder      FileStorageDecoder[T]
}

func (slf *IndexDataFileStorage[I, T]) Load(name string, index I) T {
	bytes, err := file.ReadOnce(filepath.Join(slf.dir, fmt.Sprintf(indexNameFormat, name, index, slf.suffix)))
	if err != nil {
		return slf.generate(name, index)
	}
	var data = slf.generate(name, index)
	_ = slf.decoder(bytes, data)
	return data
}

func (slf *IndexDataFileStorage[I, T]) LoadAll(name string) map[I]T {
	var result = make(map[I]T)
	files, err := os.ReadDir(slf.dir)
	if err != nil {
		return result
	}
	for _, entry := range files {
		if entry.IsDir() || !strings.HasPrefix(entry.Name(), name) || !strings.HasSuffix(entry.Name(), slf.suffix) {
			continue
		}
		bytes, err := file.ReadOnce(filepath.Join(slf.dir, entry.Name()))
		if err != nil {
			continue
		}
		data := slf.generateZero(name)
		if err := slf.decoder(bytes, data); err == nil {
			result[data.GetIndex()] = data
		}
	}
	return result
}

func (slf *IndexDataFileStorage[I, T]) Save(name string, index I, data T) error {
	bytes, err := slf.encoder(data)
	if err != nil {
		return err
	}
	return file.WriterFile(filepath.Join(slf.dir, fmt.Sprintf(indexNameFormat, name, index, slf.suffix)), bytes)
}

func (slf *IndexDataFileStorage[I, T]) SaveAll(name string, data map[I]T, errHandle func(err error) bool, retryInterval time.Duration) {
	for index, data := range data {
		for {
			if err := slf.Save(name, index, data); err != nil {
				if !errHandle(err) {
					time.Sleep(retryInterval)
					continue
				}
				break
			}
			break
		}
	}
}

func (slf *IndexDataFileStorage[I, T]) Delete(name string, index I) {
	_ = os.Remove(filepath.Join(slf.dir, fmt.Sprintf(indexNameFormat, name, index, slf.suffix)))
}

func (slf *IndexDataFileStorage[I, T]) DeleteAll(name string) {
	files, err := os.ReadDir(slf.dir)
	if err != nil {
		return
	}
	for _, entry := range files {
		if entry.IsDir() || !strings.HasPrefix(entry.Name(), name) || !strings.HasSuffix(entry.Name(), slf.suffix) {
			continue
		}
		_ = os.Remove(filepath.Join(slf.dir, entry.Name()))
	}
}
