package storages

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/file"
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/storage"
	"os"
	"path/filepath"
	"strings"
)

const (
	// IndexDataFileDefaultSuffix 是 IndexDataFileStorage 的文件默认后缀
	IndexDataFileDefaultSuffix = "stock"

	indexNameFormat     = "%s.%v.%s"
	indexNameFormatTemp = "%s.%v.%s.temp"
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

func (slf *IndexDataFileStorage[I, T]) Load(name string, index I) (T, error) {
	bytes, err := file.ReadOnce(filepath.Join(slf.dir, fmt.Sprintf(indexNameFormat, name, index, slf.suffix)))
	if err != nil {
		return slf.generate(name, index), nil
	}
	var data = slf.generate(name, index)
	return data, slf.decoder(bytes, data)
}

func (slf *IndexDataFileStorage[I, T]) LoadAll(name string) (map[I]T, error) {
	var result = make(map[I]T)
	files, err := os.ReadDir(slf.dir)
	if err != nil {
		return result, err
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
	return result, err
}

func (slf *IndexDataFileStorage[I, T]) Save(name string, index I, data T) error {
	bytes, err := slf.encoder(data)
	if err != nil {
		return err
	}
	return file.WriterFile(filepath.Join(slf.dir, fmt.Sprintf(indexNameFormat, name, index, slf.suffix)), bytes)
}

func (slf *IndexDataFileStorage[I, T]) SaveAll(name string, data map[I]T) error {
	var temps = make([]string, 0, len(data))
	defer func() {
		for _, temp := range temps {
			_ = os.Remove(temp)
		}
	}()
	for index, data := range data {
		bytes, err := slf.encoder(data)
		if err != nil {
			return err
		}
		path := filepath.Join(slf.dir, fmt.Sprintf(indexNameFormatTemp, name, index, slf.suffix))
		temps = append(temps, path)
		if err = file.WriterFile(path, bytes); err != nil {
			return err
		}
	}
	for _, temp := range temps {
		if err := os.Rename(temp, strings.TrimSuffix(temp, ".temp")); err != nil {
			return err
		}
	}
	return nil
}

func (slf *IndexDataFileStorage[I, T]) Delete(name string, index I) error {
	return os.Remove(filepath.Join(slf.dir, fmt.Sprintf(indexNameFormat, name, index, slf.suffix)))
}

func (slf *IndexDataFileStorage[I, T]) DeleteAll(name string) error {
	files, err := os.ReadDir(slf.dir)
	if err != nil {
		return err
	}
	for _, entry := range files {
		if entry.IsDir() || !strings.HasPrefix(entry.Name(), name) || !strings.HasSuffix(entry.Name(), slf.suffix) {
			continue
		}
		if err := os.Remove(filepath.Join(slf.dir, entry.Name())); err != nil {
			return err
		}
	}
	return nil
}
